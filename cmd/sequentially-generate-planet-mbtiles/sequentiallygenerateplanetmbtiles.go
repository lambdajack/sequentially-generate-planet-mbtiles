package sequentiallygenerateplanetmbtiles

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/extract"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/mbtiles"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/planet"
)

const (
	exitOK          = 0
	exitPermissions = iota + 100
	exitReadInput
	exitDownloadURL
	exitFlags
	exitInvalidJSON
	exitBuildContainers
)

var cfg = &configuration{}

func init() {
	helpMessage()
}

func EntryPoint(df []byte) int {
	initFlags()

	if fl.version {
		fmt.Printf("sequentially-generate-planet-mbtiles version %s\n", sgpmVersion)
		os.Exit(exitOK)
	}

	validateFlags()

	initConfig()

	initDirStructure()

	initLoggers()

	cloneRepos()

	setupContainers(df)

	if fl.stage {
		lg.rep.Println("Stage flag set. Staging completed. Exiting...")
		os.Exit(exitOK)
	}

	if !cfg.MergeOnly {
		downloadOsmData()

		unzipSourceData()

		moveOcean()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			cleanContainers()
			os.Exit(1)
		}()
		defer close(c)

		if !cfg.SkipSlicing {
			lg.rep.Println("slice generation started; there may be significant gaps between logs")
			lg.rep.Printf("target file size: %d MB\n", cfg.MaxRamMb/14)
			extract.TreeSlicer(cfg.PbfFile, pth.pbfSlicesDir, pth.pbfDir, cfg.MaxRamMb/14, ct.gdal, ct.osmium)

			filepath.Walk(pth.pbfSlicesDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Fatalf(err.Error())
				}
				if !info.IsDir() {
					mbtiles.Generate(path, pth.mbtilesDir, pth.coastlineDir, pth.landcoverDir, cfg.TilemakerConfig, cfg.TilemakerProcess, cfg.OutAsDir, ct.tilemaker)
				}
				return nil
			})
		}
	}

	final := pth.outDir

	if !cfg.OutAsDir {
		f := planet.Generate(pth.mbtilesDir, pth.outDir, ct.tippecanoe)
		final = f
	}

	if !cfg.OutAsDir && final == pth.outDir {
		lg.rep.Printf("Hmmm - we think you will find success at %s, but we can't quite tell for some reason... Maybe we don't have permission to see?\n", pth.outDir)
	} else {
		lg.rep.Println("SUCCESS: ", final)
	}

	endMessage(final)

	return exitOK
}

func endMessage(out string) {
	fmt.Println(`
	 __________________________________________________
	|                                                  |
	|                Thank you for using               |
	|     Sequentially Generate Planet Mbtiles!!       |
	|__________________________________________________|

Your carriage awaits you at: ` + out + "\n\n")

	fmt.Printf("TRY: docker run --rm -it -v %s:/data -p 8080:80 maptiler/tileserver-gl\n\n", filepath.Dir(out))
	fmt.Print("REMEMBER: To view the map with proper styles, you may need to set up a frontend with something like Maplibre or Leaflet.js using the correct style.json, rather than using the tileserver-gl's inbuilt 'Viewer'; although the viewer is great for checking that the mbtiles work and you got the area you were expecting.\n\n")

	fmt.Print("We would love to make this process as easy and reliable as possible for everyone. If you have any feedback, suggestions, or bug reports please come over to https://github.com/lambdajack/sequentially-generate-planet-mbtiles and let us know.\n\n")
}
