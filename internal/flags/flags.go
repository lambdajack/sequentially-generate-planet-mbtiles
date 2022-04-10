package flags

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/config"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var defaultConfigFile = filepath.FromSlash(folders.ConfigsFolder + "/" + "config.json")

func GetFlags() {
	pathToConfig := flag.String("c", defaultConfigFile, "Relative or absolute path to config.json file")

	flag.Parse()

	_, err := config.LoadConfig(*pathToConfig)
	if err != nil {
		stderrorhandler.StdErrorHandler("flags.go | Failed to load any config.json. Unable to proceed...", err)
		panic(err)
	} else {
		log.Printf("Loaded config: %v", pathToConfig)
	}
}
