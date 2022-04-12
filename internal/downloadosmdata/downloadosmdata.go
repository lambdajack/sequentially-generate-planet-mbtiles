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
	landCoverUrban := downloadInformation{
		url:          "https://naciscdn.org/naturalearth/10m/cultural/ne_10m_urban_areas.zip",
		destFileName: "ne_10m_urban_areas.zip",
		destFolder:   folders.DataFolder,
	}
	landCoverIceShelves := downloadInformation{
		url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_antarctic_ice_shelves_polys.zip",
		destFileName: "ne_10m_antarctic_ice_shelves_polys.zip",
		destFolder:   folders.DataFolder,
	}
	landCoverGlaciated := downloadInformation{
		url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_glaciated_areas.zip",
		destFileName: "ne_10m_glaciated_areas.zip",
		destFolder:   folders.DataFolder,
	}

	downloads := [...]downloadInformation{osmPlanetPbf, waterPolygons, landCoverUrban, landCoverIceShelves, landCoverGlaciated}

	for _, dl := range downloads {
		if _, err := os.Stat(filepath.Clean(dl.destFolder + "/" + dl.destFileName)); os.IsNotExist(err) {
			err := downloadurl.DownloadURL(dl.url, dl.destFileName, dl.destFolder)
			if err != nil {
				stderrorhandler.StdErrorHandler("main.go | Failed downloading required initial data. Unable to proceed", err)
				panic(err)
			}
			fmt.Printf("Download success: %v\n", dl.destFileName)
		} else {
			fmt.Printf("main.go | %v already exists. Skipping download.\n", dl.destFileName)
		}
	}
}
