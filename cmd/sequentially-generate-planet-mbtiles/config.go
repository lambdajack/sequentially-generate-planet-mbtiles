package sequentiallygenerateplanetmbtiles

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/handlejson"
)

func initConfig() {
	if fl.config == "" {
		setConfigByFlags()
	} else {
		setConfigByJSON()
	}
}

func setConfigByJSON() {
	err := handlejson.DecodeTo(cfg, fl.config)
	if err != nil {
		log.Printf("Unable to decode config: %s", err)
		os.Exit(exitInvalidJSON)
	}
}

func setConfigByFlags() {
	cfg.PlanetFile = fl.planetFile
	cfg.DataDir = fl.dataDir
	cfg.OutDir = fl.outDir
	cfg.IncludeOcean = fl.includeOcean
	cfg.IncludeLanduse = fl.includeLanduse
	cfg.TilemakerConfig = fl.tilemakerConfig
	cfg.TilemakerProcess = fl.tilemakerProcess
	cfg.MaxRamMb = getRam()
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
