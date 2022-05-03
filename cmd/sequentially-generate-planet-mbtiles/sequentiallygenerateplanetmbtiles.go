package sequentiallygenerateplanetmbtiles

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type flags struct {
	version          bool
	stage            bool
	config           string
	planetFile       string
	dataDir          string
	outDir           string
	includeOcean     bool
	includeLanduse   bool
	tilemakerConfig  string
	tilemakerProcess string
	maxRamMb         uint64
}

type configuration struct {
	PlanetFile       string `json:"planetFile"`
	DataDir          string `json:"dataDir"`
	OutDir           string `json:"outDir"`
	IncludeOcean     bool   `json:"includeOcean"`
	IncludeLanduse   bool   `json:"includeLanduse"`
	TilemakerConfig  string `json:"TilemakerConfig"`
	TilemakerProcess string `json:"TilemakerProcess"`
	MaxRamMb         uint64 `json:"maxRamMb"`
}

type paths struct {
	dataFolder               string
	outFolder                string
	coastlineFolder          string
	landcoverFolder          string
	landCoverUrbanDepth      string
	landCoverIceShelvesDepth string
	landCoverGlaciatedDepth  string
	pbfFolder                string
	pbfSlicesFolder          string
	pbfQuadrantSlicesFolder  string
	mbtilesFolder            string
	logsFolder               string
}

type loggers struct {
	prog *log.Logger
	err  *log.Logger
	rep  *log.Logger
}

const (
	exitOK          = 0
	exitPermissions = iota + 100
	exitReadInput
	exitFetchURL
	exitFlags
	exitInvalidJSON
)

const sgpmVersion = "3.0.0"

var fl = &flags{}
var cfg = &configuration{}
var pth = &paths{}
var lg = &loggers{}

func init() {
	flag.Usage = func() {
		h := "Sequentially Generate Planet Mbtiles\n"

		h += "\nUsage:\n"
		h += "    sequentially-generate-planet-mbtiles [OPTIONS]\n"

		h += "\nOptions:\n"
		h += "    -h,  --help                 Print this help message\n"
		h += "    -v,  --version              Print version information\n"
		h += "    -s,  --stage                Initialise required containers, folders and logs based on the supplied config file and then exit.\n"
		h += "    -c,  --config               Provide path to a config.json. No configuration is required. If a config.json is provided, all other \"config flags\" are ignored and runtime params are derived solely from the config.json. See documentation for example config.json\n"
		h += "    -p,  --planet-file          Config flag | \"\" | Provide path to your osm.pbf file to be turned into mbtiles. By default planet-latest.osm.pbf will be downloaded directly from OpenStreetMap. If a file is provided, downloading the latest planet osm data from openstreetmap is skipped and the supplied file will be used. You may use this to provide a file other than an entire planet .osm.pbf file (such as a region downloaded from https://download.geofabrik.de).\n"
		h += "    -d,  --datadir              Config flag | data | Provide path to the data directory. This is where files will be downloaded to and files generated as a result of processing osm data will be stored.\n"
		h += "    -o,  --outdir               Config flag | data" + string(os.PathSeparator) + "out | Provide path to output directory for the final planet.mbtiles file.\n"
		h += "    -io, --include-ocean        Config flag | true | Include ocean tiles in final planet.mbtiles\n"
		h += "    -il, --include-landuse      Config flag | true | Include landuse layer in final planet.mbtiles\n"
		h += "    -tc, --tilemaker-config     Config flag | (embedded) | Provide path to tilemaker configuration file. The default configuration is embedded into the release binary. See the default used here: https://github.com/lambdajack/tilemaker/blob/b90347b2a4fd475470b9870b8e44e2829d8e4d6d/resources/config-openmaptiles.json\n"
		h += "    -tp, --tilemaker-process    Config flag | (embedded) | Provide path to tilemaker configuration file. The default process file is embedded into the release binary. See the default used here: https://github.com/lambdajack/tilemaker/blob/b90347b2a4fd475470b9870b8e44e2829d8e4d6d/resources/process-openmaptiles.lua\n"
		h += "    -r,  --ram                  Config flag | (linux: derived from system) (other os: 4096) | Provide the maximum amount of RAM in MB that the process should use. If a linux os is detected, the total system RAM will be detected from /proc/meminfo and a default will be set to a reasonably safe level, maximising the available resourses. This assumes that only a minimal amount of system RAM is currently being used (such as an idle desktop environment (<2G)). If you are having memory problems, consider manually setting this flag to a reduced value. NOTE THIS IS NOT GUARANTEED AND SOME SAFETY MARGIN SHOULD BE ALLOWED\n"

		h += "\nExit Codes:\n"
		h += fmt.Sprintf("    %d\t%s\n", exitOK, "OK")
		h += fmt.Sprintf("    %d\t%s\n", exitPermissions, "Do not have permission")
		h += fmt.Sprintf("    %d\t%s\n", exitReadInput, "Error reading input")
		h += fmt.Sprintf("    %d\t%s\n", exitFetchURL, "Error fetching URL")
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

	flag.StringVar(&fl.dataDir, "d", "data", "")
	flag.StringVar(&fl.dataDir, "datadir", "data", "")

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

	flag.Parse()

	if fl.version {
		fmt.Printf("sequentially-generate-planet-mbtiles version %s\n", sgpmVersion)
		os.Exit(exitOK)
	}

	validateFlags()

	initConfig()

	initFolderStructure()

	initLoggers()

	checkRecursiveClone()

	if fl.stage {
		lg.rep.Println("Stage flag set. Staging completed. Exiting...")
		os.Exit(exitOK)
	}

	// downloadosmdata.DownloadOsmData()
	// unzippolygons.UnzipPolygons()
	// extractiontree.Slicer("./data/pbf/switzerland-latest.osm.pbf")
	// extractquadrants.ExtractQuadrants()
	// extractslices.FromQuadrants()
	// genmbtiles.GenMbtiles()
	// genplanet.GenPlanet()
	return exitOK
}

func checkRecursiveClone() {
	tp := [...]string{"libosmium", "osmium-tool", "tilemaker", "tippecanoe"}

	for _, t := range tp {
		if _, err := os.Stat(filepath.Join("third_party", t)); os.IsNotExist(err) {
			lg.err.Printf("Submodule %v cannot be found. Attempting to fix..", t)
			err := exec.Command("git", "submodule", "update", "--init", "--recursive").Run()
			if err != nil {
				lg.err.Fatal("Failed to recursively clone submodules. Submodules are required to run this programme. Please clone the submodules manually and try again.")
			}
		}
	}
}

func validateFlags() {
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