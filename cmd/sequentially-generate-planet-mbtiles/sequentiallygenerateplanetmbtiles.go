package sequentiallygenerateplanetmbtiles

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/containers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/extract"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/mbtiles"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/planet"
)

type flags struct {
	version          bool
	stage            bool
	config           string
	planetFile       string
	workingDir       string
	outDir           string
	includeOcean     bool
	includeLanduse   bool
	tilemakerConfig  string
	tilemakerProcess string
	maxRamMb         uint64
	diskEfficient    bool
	outAsDir         bool
	skipSlicing      bool
	mergeOnly        bool
	skipDownload	bool
}

const (
	exitOK          = 0
	exitPermissions = iota + 100
	exitReadInput
	exitDownloadURL
	exitFlags
	exitInvalidJSON
	exitBuildContainers
)

var fl = &flags{}
var cfg = &configuration{}

func init() {
	flag.Usage = func() {
		h := `
Sequentially Generate Planet Mbtiles
____________________________________

Usage:
    sequentially-generate-planet-mbtiles [OPTIONS]

Options:
  -h, --help               Print this help message
  -v, --version            Print version information

  -s, --stage              Initialise required containers, Dirs and logs
                           based on the supplied config file and then exit.

  -c, --config             Provide path to a config.json. No configuration
                           is required. If a config.json is provided, all
                           other "config flags" are ignored and runtime
                           params are derived solely from the config.json.
                           See documentation for example config.json

Config Flags:
  -p, --planet-file        Provide path to your osm.pbf file to be turned 
                           into mbtiles. By default a planet-latest.osm.pbf 
                           will be downloaded directly from OpenStreetMap. If 
                           a file is provided, downloading the latest 
                           planet osm data from openstreetmap is 
                           skipped and the supplied file will be used. 
                           You may use this to provide a file other than 
                           an entire planet .osm.pbf file (such as a region 
                           downloaded from https://download.geofabrik.de).

  -w,  --working-dir       Provide path to the working directory. This is 
                           where files will be downloaded to and files 
                           generated as a result of processing osm data will 
                           be stored. Temporary files will be stored here. 
                           Please ensure your designated working directory 
                           has at least 300 GB of space available.

  -o,  --outdir            Provide path to output directory for the final 
                           planet.mbtiles file.

  -io, --include-ocean     Include ocean tiles in final planet.mbtiles
  -il, --include-landuse   Include landuse layer in final planet.mbtiles
  
  -tc, --tilemaker-config  Provide path to tilemaker configuration file. The 
                           default configuration is embedded into the release 
                           binary. See the default used here: 
                           https://github.com/lambdajack/tilemaker/blob/master/resources/config-openmaptiles.json

  -tp, --tilemaker-process Provide path to tilemaker configuration file. The 
                           default process file is embedded into the release 
                           binary. See the default used here: 
                           https://github.com/lambdajack/tilemaker/blob/master/resources/process-openmaptiles.lua
	
  -r,  --ram               Provide the maximum amount of RAM in MB that the 
                           process should use. If a linux os is detected, 
                           the total system RAM will be detected from 
                           /proc/meminfo and a default will be set to a 
                           reasonably safe level, maximising the available 
                           resources. This assumes that only a minimal amount 
                           of system RAM is currently being used (such as an 
                           idle desktop environment (<2G)). If you are having 
                           memory problems, consider manually setting this flag 
                           to a reduced value. NOTE THIS IS NOT GUARANTEED AND 
                           SOME SAFETY MARGIN SHOULD BE ALLOWED. On non unix 
                           operating systems the default is set to 4096.

  -de, --disk-efficient    Use disk efficient mode. This will skip the 
                           intermediary data slices and proceed straight to the 
                           working slices. Can considerably increase the time 
                           taken, but will save up to approx. 70 GB of disk 
                           space overall. Use only if disk space is a real 
                           consideration.

  -od, --out-as-dir        The final output will be a directory of tiles
                           rather than a single mbtiles file. This will
                           generate hundreds of thousands of files in a
                           predetermined directory structure. More
                           information can ba found about this format here:
                           https://documentation.maptiler.com/hc/en-us/articles/360020886878-Folder-vs-MBTiles-vs-GeoPackage
						   
  -ss, --skip-slicing      Skips the intermediate data processing/slicing
                           and looks for existing files to convert into
                           mbtiles in [workingDir]/pbf/slices. This is 
                           useful if you wish to experiment with different
                           Tilemaker configs/process (for example if you
                           wish to change the zoom levels or style tagging
                           of the final output). Once the existing files
                           have been converted to mbtiles, they will be
                           merged either to a single file, or to a
                           directory, respecting the -od flag.

  -mo, --merge-only        Skips the entire generation process and instead
                           looks for existing mbtiles in [workingDir]/mbtiles
                           and merges them into a single planet.mbtiles file
                           in the [outDir]. This is useful if you already
                           have a tileset you wish to merge.

  -sd, --skip-download     Skips planet download. Must be set in conjunction
                           with -p, -ss or -mo.
`
		h += "\nExit Codes:\n"
		h += fmt.Sprintf("    %d\t%s\n", exitOK, "OK")
		h += fmt.Sprintf("    %d\t%s\n", exitPermissions, "Do not have permission")
		h += fmt.Sprintf("    %d\t%s\n", exitReadInput, "Error reading input")
		h += fmt.Sprintf("    %d\t%s\n", exitDownloadURL, "Error fetching URL")
		h += fmt.Sprintf("    %d\t%s\n", exitFlags, "Error parsing flags")
		h += fmt.Sprintf("    %d\t%s\n", exitInvalidJSON, "Invalid JSON")

		fmt.Fprint(os.Stderr, h)
	}
}

func EntryPoint() int {

	flag.BoolVar(&fl.version, "v", false, "")
	flag.BoolVar(&fl.version, "version", false, "")

	flag.BoolVar(&fl.stage, "s", false, "")
	flag.BoolVar(&fl.stage, "stage", false, "")

	flag.StringVar(&fl.config, "c", "", "")
	flag.StringVar(&fl.config, "config", "", "")

	flag.StringVar(&fl.planetFile, "p", "", "")
	flag.StringVar(&fl.planetFile, "planet-file", "", "")

	flag.StringVar(&fl.workingDir, "d", "data", "")
	flag.StringVar(&fl.workingDir, "working-dir", "data", "")

	flag.StringVar(&fl.outDir, "o", "data/out", "")
	flag.StringVar(&fl.outDir, "outdir", "data/out", "")

	flag.BoolVar(&fl.includeOcean, "io", true, "")
	flag.BoolVar(&fl.includeOcean, "include-ocean", true, "")

	flag.BoolVar(&fl.includeLanduse, "il", true, "")
	flag.BoolVar(&fl.includeLanduse, "include-landuse", true, "")

	flag.StringVar(&fl.tilemakerConfig, "tc", "third_party/tilemaker/resources/config-openmaptiles.json", "")
	flag.StringVar(&fl.tilemakerConfig, "tilemaker-config", "third_party/tilemaker/resources/config-openmaptiles.json", "")

	flag.StringVar(&fl.tilemakerProcess, "tp", "third_party/tilemaker/resources/process-openmaptiles.lua", "")
	flag.StringVar(&fl.tilemakerProcess, "tilemaker-process", "third_party/tilemaker/resources/process-openmaptiles.lua", "")

	flag.Uint64Var(&fl.maxRamMb, "r", 0, "")
	flag.Uint64Var(&fl.maxRamMb, "ram", 0, "")

	flag.BoolVar(&fl.diskEfficient, "de", false, "")
	flag.BoolVar(&fl.diskEfficient, "disk-efficient", false, "")

	flag.BoolVar(&fl.outAsDir, "od", false, "")
	flag.BoolVar(&fl.outAsDir, "out-as-dir", false, "")

	flag.BoolVar(&fl.skipSlicing, "ss", false, "")
	flag.BoolVar(&fl.skipSlicing, "skip-slicing", false, "")

	flag.BoolVar(&fl.mergeOnly, "mo", false, "")
	flag.BoolVar(&fl.mergeOnly, "merge-only", false, "")

	flag.BoolVar(&fl.skipDownload, "sd", false, "")
	flag.BoolVar(&fl.skipDownload, "skip-download", false, "")

	flag.Parse()

	if fl.version {
		fmt.Printf("sequentially-generate-planet-mbtiles version %s\n", sgpmVersion)
		os.Exit(exitOK)
	}

	validateFlags()

	initConfig()

	initDirStructure()

	initLoggers()

	checkRecursiveClone()

	err := containers.BuildAll()
	if err != nil {
		lg.err.Println(err)
		os.Exit(exitBuildContainers)
	}

	if fl.stage {
		lg.rep.Println("Stage flag set. Staging completed. Exiting...")
		os.Exit(exitOK)
	}

	if !cfg.MergeOnly {
		downloadOsmData()

		unzipSourceData()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			err := containers.CleanAll()
			if err != nil {
				lg.err.Println(err)
			}
			os.Exit(1)
		}()
		defer close(c)

		if !cfg.SkipSlicing {
			lg.rep.Println("Starting slice generation. This will take a while and there may be several minutes between progress updates.")

			if !cfg.DiskEfficient {
				if cfg.PlanetFile == "" {
					extract.Quadrants(filepath.Join(pth.pbfDir, "planet-latest.osm.pbf"), pth.pbfQuadrantSlicesDir, containers.ContainerNames.Osmium)
				} else {
					pf, err := filepath.Abs(cfg.PlanetFile)
					if err != nil {
						log.Fatal("failed to locate your planet file: ", cfg.PlanetFile)
					}
					if _, err := os.Stat(cfg.PlanetFile); os.IsNotExist(err) {
						log.Fatal("failed to locate your planet file: ", cfg.PlanetFile)
					}
					extract.Quadrants(pf, pth.pbfQuadrantSlicesDir, containers.ContainerNames.Osmium)
				}
			} else {
				lg.rep.Println("Disk efficient mode enabled. Skipping intermediate quadrant slices.")
			}

			if cfg.DiskEfficient {
				extract.TreeSlicer(cfg.PlanetFile, pth.pbfSlicesDir, pth.pbfDir, 1000)
			} else {
				filepath.Walk(pth.pbfQuadrantSlicesDir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						log.Fatalf(err.Error())
					}
					if !info.IsDir() {
						extract.TreeSlicer(path, pth.pbfSlicesDir, pth.pbfDir, 1000)
					}
					return nil
				})
			}
		}

		filepath.Walk(pth.pbfSlicesDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			if !info.IsDir() {
				mbtiles.Generate(path, pth.mbtilesDir, pth.coastlineDir, pth.landcoverDir, cfg.TilemakerConfig, cfg.TilemakerProcess, cfg.OutAsDir)

			}
			return nil
		})
	}

	final := pth.outDir

	if !cfg.OutAsDir {
		f := planet.Generate(pth.mbtilesDir, pth.outDir)
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

func checkRecursiveClone() {
	tp := [...]string{"libosmium", "osmium-tool", "tilemaker", "tippecanoe", "gdal"}

	for _, t := range tp {
		if _, err := os.Stat(filepath.Join("third_party", t, "README.md")); os.IsNotExist(err) {
			lg.err.Printf("Submodule %v cannot be found. Attempting to fix..", t)
			err := exec.Command("git", "submodule", "update", "--init", "--recursive").Run()
			if err != nil {
				lg.err.Fatal("Failed to recursively clone submodules. Submodules are required to run this programme. Please clone the submodules manually and try again.")
			}
		}
	}
}

func validateFlags() {

	if fl.skipDownload && !fl.skipSlicing && !fl.mergeOnly {
		fmt.Println("-sd must be used with -ss or -mo. See -h for more information.")
		os.Exit(exitFlags)
	}

	configFlag := fl.config

	var defaultConfigFlagValue string

	flag.Visit(func(f *flag.Flag) {
		if f.Name == "config" || f.Name == "c" {
			defaultConfigFlagValue = f.DefValue
		}
	})

	if configFlag != defaultConfigFlagValue {
		invalidFlags := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name != "config" && f.Name != "c" && f.Name != "s" && f.Name != "stage" {
				if f.Value.String() != f.DefValue {
					log.Printf("[WARN] -%s flag was provided but is overwitten by the provided config.json. Please supply either only a config.json file OR configuration flags. See '-h' for more information.", f.Name)
					invalidFlags = true
				}
			}
			if f.Name == "io" || f.Name == "include-ocean" || f.Name == "il" || f.Name == "include-landuse" {
				log.Printf("[WARN] -%s flag was provided but is overwitten by the provided config.json. Please supply either only a config.json file OR configuration flags. See '-h' for more information.", f.Name)
				invalidFlags = true
			}

		})

		if invalidFlags {
			os.Exit(exitFlags)
		}
	}
}

func endMessage(out string) {
	fmt.Println(`
	 __________________________________________________
	|                                                  |
	|                Thank you for using               |
	|     Sequentially Generate Planet Mbtiles!!       |
	|__________________________________________________|

Your carriage awaits you at: ` + out + "\n")

	fmt.Println("We would love to make this process as easy and reliable as possible for everyone. If you have any feedback, suggestions, or bug reports please come over to https://github.com/lambdajack/sequentially-generate-planet-mbtiles and let us know.")
	fmt.Println()

}
