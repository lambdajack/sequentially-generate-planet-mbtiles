package genmbtiles

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/logger"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/generatembtiles"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var mu sync.Mutex

func GenMbtiles() {
	f, err := os.ReadDir(folders.PbfSlicesFolder)
	if err != nil {
		stderrorhandler.StdErrorHandler("genmbtiles.go | Failed to read PBF folder. Unable to proceed...", err)
		panic(err)
	}

	for _, file := range f {
		if file.IsDir() == false {
			outFileName := strings.Split(file.Name(), ".")[0] + ".mbtiles"

			if _, err := os.Stat(filepath.FromSlash(folders.MbtilesFolder + "/" + outFileName)); os.IsNotExist(err) {
				err := generatembtiles.GenerateMbTiles(file.Name(), outFileName, folders.PbfSlicesFolder, folders.MbtilesFolder, folders.CoastlineFolder, buildthirdpartycontainers.ContainerTilemakerName, folders.TilemakerConfigsFolder, "config.json", "process.lua")
				if err != nil {
					stderrorhandler.StdErrorHandler(fmt.Sprintf("genmbtiles.go | Failed to generate mbtiles for %s.", file.Name()), err)
					mu.Lock()
					logger.AppendReport(fmt.Sprintf("GENERATE_MBTILES_FAILED: %s", outFileName))
				} else {
					mu.Lock()
					logger.AppendReport(fmt.Sprintf("GENERATE_MBTILES_SUCCESS: %s", outFileName))
				}
				mu.Unlock()

			} else {
				log.Printf("%s already exists. Skipping...", outFileName)
			}
		}
	}
}
