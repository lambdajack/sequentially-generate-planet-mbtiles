package downloadosmdata

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/downloadurl"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

type downloadInformation struct {
	url, destFileName, destFolder string
}

func DownloadOsmData() {
	osmPlanetPbf := downloadInformation{
		url:          "https://planet.openstreetmap.org/pbf/planet-latest.osm.pbf",
		destFileName: "planet-latest.osm.pbf",
		destFolder:   folders.PbfFolder,
	}
	waterPolygons := downloadInformation{
		url:          "https://osmdata.openstreetmap.de/download/water-polygons-split-4326.zip",
		destFileName: "water-polygons-split-4326.zip",
		destFolder:   folders.DataFolder,
	}
	downloads := [...]downloadInformation{osmPlanetPbf, waterPolygons}

	for _, dl := range downloads {
		if _, err := os.Stat(filepath.FromSlash(dl.destFolder + "/" + dl.destFileName)); os.IsNotExist(err) {
			err := downloadurl.DownloadUrl(dl.url, dl.destFileName, dl.destFolder)
			if err != nil {
				stderrorhandler.StdErrorHandler("main.go | Failed downloading required initial data. Unable to proceed", err)
				panic(err)
			}
		} else {
			fmt.Printf("main.go | %v already exists. Skipping download.\n", dl.destFileName)
		}
	}
}
