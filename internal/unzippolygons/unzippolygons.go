package unzippolygons

import (
	"fmt"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/unzip"
)

type unzipInformation struct {
	srcPath  string
	destPath string
}

func UnzipPolygons() {
	waterPolygons := unzipInformation{srcPath: "water-polygons-split-4326.zip", destPath: folders.CoastlineFolder}
	landCoverUrban := unzipInformation{srcPath: "ne_10m_urban_areas.zip", destPath: folders.LandcoverFolder}
	landCoverIceShelves := unzipInformation{srcPath: "ne_10m_antarctic_ice_shelves_polys.zip", destPath: folders.LandcoverFolder}
	landCoverGlaciated := unzipInformation{srcPath: "ne_10m_glaciated_areas.zip", destPath: folders.LandcoverFolder}

	fileNames := [...]*unzipInformation{&waterPolygons, &landCoverUrban, &landCoverIceShelves, &landCoverGlaciated}

	for i, zipFile := range fileNames {
		fileNames[i].srcPath = filepath.Clean(folders.DataFolder + "/" + zipFile.srcPath)

		err := unzip.Unzip(fileNames[i].srcPath, fileNames[i].destPath)
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("unzippolygons.go | Failed unzipping %v polygons. Unable to proceed...", zipFile.srcPath), err)
			panic(err)
		}
	}

}
