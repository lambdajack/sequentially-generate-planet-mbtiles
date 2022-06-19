package sequentiallygenerateplanetmbtiles

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lambdajack/lj_go/pkg/lj_json"
)

type configuration struct {
	PlanetFile       string `json:"planetFile"`
	WorkingDir       string `json:"workingDir"`
	OutDir           string `json:"outDir"`
	IncludeOcean     bool   `json:"includeOcean"`
	IncludeLanduse   bool   `json:"includeLanduse"`
	TilemakerConfig  string `json:"TilemakerConfig"`
	TilemakerProcess string `json:"TilemakerProcess"`
	MaxRamMb         uint64 `json:"maxRamMb"`
	DiskEfficient    bool   `json:"diskEfficient"`
	OutAsDir         bool   `json:"outAsDir"`
	SkipSlicing      bool   `json:"skipSlicing"`
	MergeOnly        bool   `json:"mergeOnly"`
	SkipDownload	bool  `json:"skipDownload"`
}

func initConfig() {
	if fl.config == "" {
		setConfigByFlags()
	} else {
		setConfigByJSON()
	}

	verifyPaths()
	setAbsolutePaths()
}

func setConfigByJSON() {
	err := lj_json.DecodeTo(cfg, fl.config, 1000)
	if err != nil {
		log.Printf("Unable to decode config: %s", err)
		os.Exit(exitInvalidJSON)
	}
}

func setConfigByFlags() {
	cfg.PlanetFile = fl.planetFile
	cfg.WorkingDir = fl.workingDir
	cfg.OutDir = fl.outDir
	cfg.IncludeOcean = fl.includeOcean
	cfg.IncludeLanduse = fl.includeLanduse
	cfg.TilemakerConfig = fl.tilemakerConfig
	cfg.TilemakerProcess = fl.tilemakerProcess
	cfg.MaxRamMb = getRam()
	cfg.DiskEfficient = fl.diskEfficient
	cfg.OutAsDir = fl.outAsDir
	cfg.SkipSlicing = fl.skipSlicing
	cfg.MergeOnly = fl.mergeOnly
	cfg.SkipDownload = fl.skipDownload
}

func setAbsolutePaths() {

	if cfg.TilemakerProcess == "" || strings.ToLower(cfg.TilemakerProcess) == "tileserver-gl-basic" {
		cfg.TilemakerConfig = "third_party/tilemaker/resources/process-openmaptiles.json"
	}

	if cfg.TilemakerProcess == "sgpm-bright" {
		cfg.TilemakerProcess = "third_party/tilemaker/resources/process-openmaptiles.json"
	}

	cfg.PlanetFile = convertAbs(cfg.PlanetFile)
	cfg.WorkingDir = convertAbs(cfg.WorkingDir)
	cfg.OutDir = convertAbs(cfg.OutDir)
	cfg.TilemakerConfig = convertAbs(cfg.TilemakerConfig)
	cfg.TilemakerProcess = convertAbs(cfg.TilemakerProcess)
}

func verifyPaths() {
	if cfg.PlanetFile != "" {
		if _, err := os.Stat(cfg.PlanetFile); os.IsNotExist(err) {
			log.Fatalf("planet file does not exist: %s", cfg.PlanetFile)
		}
	}

	if cfg.TilemakerConfig != "" {
		if _, err := os.Stat(cfg.TilemakerConfig); os.IsNotExist(err) {
			log.Fatalf("tilemaker config does not exist: %s", cfg.TilemakerConfig)
		}
	}

	if cfg.TilemakerProcess != "" {
		if _, err := os.Stat(cfg.TilemakerProcess); os.IsNotExist(err) {
			log.Fatalf("tilemaker process does not exist: %s", cfg.TilemakerProcess)
		}
	}
}

func getRam() uint64 {

	if fl.maxRamMb != 0 {
		return fl.maxRamMb
	}

	var memTotalKb uint64

	f, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Printf("Unable to open /proc/meminfo: %s", err)
		memTotalKb = 4096 * 1024
	} else {
		defer f.Close()

		scanner := bufio.NewScanner(f)

		memTotalStr := ""

		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "MemTotal:") {
				memTotalStr = regexp.MustCompile(`[0-9]+`).FindString(scanner.Text())
			}
		}

		memTotalKb, err = strconv.ParseUint(memTotalStr, 10, 64)
		if err != nil {
			log.Printf("Unable to parse MemTotal: %s", err)
			memTotalKb = 4096 * 1024
		}
	}

	memTotalMb := memTotalKb / 1024

	totalRamToUse := memTotalMb - 2048

	return totalRamToUse
}
