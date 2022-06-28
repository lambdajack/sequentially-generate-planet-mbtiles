package sequentiallygenerateplanetmbtiles

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/extract"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/mbtiles"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/planet"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/system"
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

	setTmPaths()

	initLoggers()

	lgInitStr := fmt.Sprintf("sequentially-generate-planet-mbtiles started: %+v\n", cfg)
	cfgCompare(lgInitStr)
	lg.rep.Printf(lgInitStr)

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
			count := 0
			slicingDone := false

			filepath.Walk(pth.pbfSlicesDir, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && !strings.Contains(path, "converted-") {
					count++
				}
				if strings.Contains(path, "converted-") {
					slicingDone = true
				}
				return nil
			})

			if count != 0 {
				lg.rep.Println("previous progress detected; attempting to continue...")
				pbb := extract.IncompleteProgress(cfg.PbfFile, pth.pbfSlicesDir, ct.gdal, lg.err, lg.rep)
				if pbb != "" {
					np, err := extract.Extract(cfg.PbfFile, filepath.Join(pth.pbfDir, "resume.osm.pbf"), pbb, ct.osmium)
					if err != nil || np == "" {
						lg.err.Println("failed to continue previous progress; will attempt to continue from scratch... ", err)
					}
					if np != "" {
						cfg.PbfFile = np
						filepath.Walk(pth.pbfDir, func(path string, info os.FileInfo, err error) error {
							if !info.IsDir() {
								if !strings.Contains(path, "slices") && !strings.Contains(path, "resume") {
									log.Println("removing dirty files: ", path)
									return os.Remove(path)
								}
							}
							return nil
						})
					}
				} else {
					lg.rep.Println("failed to get previous progress; starting from scratch...")
					filepath.Walk(pth.pbfDir, func(path string, info os.FileInfo, err error) error {
						if !info.IsDir() {
							if strings.Contains(path, "resume") || strings.Contains(path, "tmp") {
								log.Println("removing dirty files: ", path)
								return os.Remove(path)
							}
						}
						return nil
					})
				}
			}

			if !slicingDone {
				lg.rep.Println("slice generation started; there may be significant gaps between logs")
				lg.rep.Printf("target file size: %d MB\n", uint64(math.Floor(float64(cfg.MaxRamMb)/15)))
				runtime.GC()
				extract.TreeSlicer(cfg.PbfFile, pth.pbfSlicesDir, pth.pbfDir, uint64(math.Floor(float64(cfg.MaxRamMb)/15)), ct.gdal, ct.osmium, lg.err, lg.prog, lg.rep)
			} else {
				lg.rep.Println("slicing already complete; moving on to tile generation")
			}

			filepath.Walk(pth.pbfSlicesDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					lg.rep.Fatalf(err.Error())
				}
				system.SetUserOwner(path)
				if !info.IsDir() {
					if !strings.Contains(path, "converted-") {
						runtime.GC()
						if !cfg.OutAsDir {
							mbtiles.Generate(path, pth.mbtilesDir, pth.coastlineDir, pth.landcoverDir, cfg.TilemakerConfig, cfg.TilemakerProcess, cfg.OutAsDir, ct.tilemaker, lg.err, lg.prog, lg.rep)
						} else {
							mbtiles.Generate(path, pth.outDir, pth.coastlineDir, pth.landcoverDir, cfg.TilemakerConfig, cfg.TilemakerProcess, cfg.OutAsDir, ct.tilemaker, lg.err, lg.prog, lg.rep)
						}
						os.Rename(path, filepath.Join(filepath.Dir(path), "converted-"+filepath.Base(path)))
					} else {
						lg.rep.Printf("already converted; skipping %s\n", path)
					}
				}
				return nil
			})
		}
	}

	final := pth.outDir

	if !cfg.OutAsDir {
		runtime.GC()
		f := planet.Generate(pth.mbtilesDir, pth.outDir, ct.tippecanoe, lg.err, lg.prog, lg.rep)
		final = f
	}

	if !cfg.OutAsDir && final == pth.outDir {
		lg.rep.Printf("Hmmm - we think you will find success at %s, but we can't quite tell for some reason... Maybe we don't have permission to see?\n", pth.outDir)
	} else {
		lg.rep.Println("success: ", final)
	}

	system.SetUserOwner(final)

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

	
Your carriage awaits you at: ` + out + "\n")

	if !cfg.OutAsDir {
		fmt.Printf("TRY: docker run --rm -it -v %s:/data -p 8080:80 maptiler/tileserver-gl\n\n", filepath.Dir(out))
		fmt.Print("REMEMBER: To view the map with proper styles, you may need to set up a frontend with something like Maplibre or Leaflet.js using the correct style.json, rather than using the tileserver-gl's inbuilt 'Viewer'; although the viewer is great for checking that the mbtiles work and you got the area you were expecting.\n\n")
	}
	fmt.Print("We would love to make this process as easy and reliable as possible for everyone. If you have any feedback, suggestions, or bug reports please come over to https://github.com/lambdajack/sequentially-generate-planet-mbtiles and let us know.\n\n")
}

func cfgCompare(lgInitStr string) {
	f, err := os.Open(filepath.Join(pth.logsDir, "rep.log"))
	if err != nil {
		lg.err.Println("unable to open log file: ", filepath.Join(pth.logsDir, "rep.log"))
		return
	}
	s, err := io.ReadAll(f)
	if err != nil {
		lg.err.Println("unable to read log file: ", filepath.Join(pth.logsDir, "rep.log"))
		return
	}
	idx := strings.LastIndex(string(s), "sequentially-generate-planet-mbtiles started: ")
	if idx > 0 {
		ss := strings.Split(string(s)[idx:], "\n")
		if strings.TrimSpace(ss[0]) != strings.TrimSpace(lgInitStr) {
					lg.rep.Println("resumption of a previously stopped execution with different parameters detected")
					fmt.Println("UNKNOWN THINGS CAN HAPPEN WHEN RESUMING WITH DIFFERENT PARAMETERS. You should only proceed if you are not confident that your changed parameters will have no adverse affect on the end result or are prepared to accept the outcome; otherwise you should clean the working/output directories and start again (you can save and replace downloaded files if requried). DO YOU WISH TO CONTINUE? yes/no")
					r := bufio.NewReader(os.Stdin)
					a, err := r.ReadString('\n')
					if err != nil {
						lg.err.Fatal("unable to read answer to resumption question; aborting; recommend starting fresh")
					}
					if a != "yes\n" && a != "y\n" {
						os.Exit(exitOK)
					}
				}
	}
}
