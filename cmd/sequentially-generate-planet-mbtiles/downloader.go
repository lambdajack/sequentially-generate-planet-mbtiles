package sequentiallygenerateplanetmbtiles

import (
	"os"
	"path/filepath"

	"github.com/lambdajack/lj_go/pkg/lj_http"
)

type downloadInformation struct {
	url, destFileName, destDir string
}

func downloadOsmData() {
	var downloads = []downloadInformation{
		{
			url:          "https://planet.openstreetmap.org/pbf/planet-latest.osm.pbf",
			destFileName: "planet-latest.osm.pbf",
			destDir:      pth.pbfDir,
		},
		{
			url:          "https://osmdata.openstreetmap.de/download/water-polygons-split-4326.zip",
			destFileName: "water-polygons-split-4326.zip",
			destDir:      pth.workingDir,
		},
		{
			url:          "https://naciscdn.org/naturalearth/10m/cultural/ne_10m_urban_areas.zip",
			destFileName: "ne_10m_urban_areas.zip",
			destDir:      pth.workingDir,
		},
		{
			url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_antarctic_ice_shelves_polys.zip",
			destFileName: "ne_10m_antarctic_ice_shelves_polys.zip",
			destDir:      pth.workingDir,
		},
		{
			url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_glaciated_areas.zip",
			destFileName: "ne_10m_glaciated_areas.zip",
			destDir:      pth.workingDir,
		},
	}

	for _, dl := range downloads {
		if _, err := os.Stat(filepath.Join(dl.destDir, dl.destFileName)); os.IsNotExist(err) {

			if dl.destFileName == "planet-latest.osm.pbf" {
				if fl.planetFile != "" {
					lg.rep.Printf("source file provided - skipping download %s", dl.url)
					continue
				}
			}

			err := lj_http.Download(dl.url, dl.destDir, dl.destFileName)
			if err != nil {
				lg.err.Printf("error downloading %s: %s", dl.url, err)
				os.Exit(exitDownloadURL)
			}
			lg.rep.Printf("Download success: %v\n", dl.destFileName)

		} else {
			lg.rep.Printf("%v already exists. Skipping download.\n", dl.destFileName)
		}
	}
}
