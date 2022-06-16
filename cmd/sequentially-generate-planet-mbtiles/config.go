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
	WorkingDir       string `json:"WorkingDir"`
	OutDir           string `json:"outDir"`
	IncludeOcean     bool   `json:"includeOcean"`
	IncludeLanduse   bool   `json:"includeLanduse"`
	TilemakerConfig  string `json:"TilemakerConfig"`
	TilemakerProcess string `json:"TilemakerProcess"`
	MaxRamMb         uint64 `json:"maxRamMb"`
	DiskEfficient    bool   `json:"diskEfficient"`
}

func initConfig() {
	if fl.config == "" {
		setConfigByFlags()
	} else {
		setConfigByJSON()
	}

}

func setConfigByJSON() {
	err := lj_json.DecodeTo(cfg, fl.config, 1000)
	if err != nil {
		log.Printf("Unable to decode config: %s", err)
		os.Exit(exitInvalidJSON)
	}
}

func setConfigByFlags() {
	cfg.PlanetFile = convertAbs(fl.planetFile)
	cfg.WorkingDir = convertAbs(fl.workingDir)
	cfg.OutDir = convertAbs(fl.outDir)
	cfg.IncludeOcean = fl.includeOcean
	cfg.IncludeLanduse = fl.includeLanduse
	cfg.TilemakerConfig = convertAbs(fl.tilemakerConfig)
	cfg.TilemakerProcess = convertAbs(fl.tilemakerProcess)
	cfg.MaxRamMb = getRam()
	cfg.DiskEfficient = fl.diskEfficient
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
