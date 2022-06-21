package sequentiallygenerateplanetmbtiles

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type flags struct {
	version          bool
	stage            bool
	config           string
	test             bool
	pbfFile          string
	workingDir       string
	outDir           string
	excludeOcean     bool
	excludeLanduse   bool
	tilemakerConfig  string
	tilemakerProcess string
	maxRamMb         uint64
	diskEfficient    bool
	outAsDir         bool
	skipSlicing      bool
	mergeOnly        bool
	skipDownload     bool
}

var fl = &flags{}

func initFlags() {
	flag.BoolVar(&fl.version, "v", false, "")
	flag.BoolVar(&fl.version, "version", false, "")

	flag.BoolVar(&fl.stage, "s", false, "")
	flag.BoolVar(&fl.stage, "stage", false, "")

	flag.StringVar(&fl.config, "c", "", "")
	flag.StringVar(&fl.config, "config", "", "")

	flag.BoolVar(&fl.test, "t", false, "")
	flag.BoolVar(&fl.test, "test", false, "")

	flag.StringVar(&fl.pbfFile, "p", "", "")
	flag.StringVar(&fl.pbfFile, "pbf-file", "", "")

	flag.StringVar(&fl.workingDir, "w", "data", "")
	flag.StringVar(&fl.workingDir, "working-dir", "data", "")

	flag.StringVar(&fl.outDir, "o", "data/out", "")
	flag.StringVar(&fl.outDir, "outdir", "data/out", "")

	flag.BoolVar(&fl.excludeOcean, "eo", false, "")
	flag.BoolVar(&fl.excludeOcean, "exclude-ocean", false, "")

	flag.BoolVar(&fl.excludeLanduse, "el", false, "")
	flag.BoolVar(&fl.excludeLanduse, "exclude-landuse", false, "")

	flag.StringVar(&fl.tilemakerConfig, "tc", "", "")
	flag.StringVar(&fl.tilemakerConfig, "tilemaker-config", "", "")

	flag.StringVar(&fl.tilemakerProcess, "tp", "", "")
	flag.StringVar(&fl.tilemakerProcess, "tilemaker-process", "", "")

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

	if fl.test {
		of := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name != "test" && f.Name != "t" {
				of = true
			}
		})
		if of {
			log.Println("you cannot use the -t or --test flag with any other flags.")
			os.Exit(exitFlags)
		}
	}

	if configFlag != defaultConfigFlagValue {
		invalidFlags := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name != "config" && f.Name != "c" && f.Name != "s" && f.Name != "stage" {
				if f.Value.String() != f.DefValue {
					log.Printf("[WARN] -%s flag was provided but is overwitten by the provided config.json. Please supply either only a config.json file OR configuration flags. See '-h' for more information.", f.Name)
					invalidFlags = true
				}
			}
			if f.Name == "eo" || f.Name == "exclude-ocean" || f.Name == "el" || f.Name == "exclude-landuse" {
				log.Printf("[WARN] -%s flag was provided but is overwitten by the provided config.json. Please supply either only a config.json file OR configuration flags. See '-h' for more information.", f.Name)
				invalidFlags = true
			}
		})

		if invalidFlags {
			os.Exit(exitFlags)
		}
	}
}

func helpMessage() {
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
                           See documentation for example config.json.

  -t, --test               Will run the entire program on a much smaller
                           dataset (morocco-latest.osm.pbf). The program
                           will download the test data and generate a
                           planet.mbtiles from it. This is useful for
                           testing both the output and that your system
                           meets the requirements. You cannot set any
                           other flags in conjunction with this flag.
                           if you wish to run your own custom test then
                           please set a config.json file with your own
                           smaller dataset and other options.
Config Flags:
  -p, --pbf-file           Provide path to your osm.pbf file to be turned 
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

  -eo, --exclude-ocean     Exclude ocean tiles in final planet.mbtiles. This
                           can significantly increase overall speed since
                           there are a lot of ocean tiles which often forms
                           a filesystem io bottleneck when writing them.

  -el, --exclude-landuse   Exclude landuse layer in final planet.mbtiles.
  
  -tc, --tilemaker-config  Provide path to a tilemaker configuration file. The 
                           default configuration is embedded into the release 
                           binary. See the default used here: 
                           https://github.com/lambdajack/tilemaker/blob/master/resources/config-openmaptiles.json

  -tp, --tilemaker-process Provide path to a tilemaker process file OR
                           a special value. Special values are 
                           "tileserver-gl-basic", "sgpm-bright". Setting a
                           special value will use one of the included configs 
                           to match the appropriate target style. See the 
                           default used here: 
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
