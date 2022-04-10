package unzipwaterpolygons

import (
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/unzip"
)

func UnzipWaterPolygons() {
	waterPolygonsPath := filepath.FromSlash(folders.DataFolder + "/" + "water-polygons-split-4326.zip")

	err := unzip.Unzip(waterPolygonsPath, folders.CoastlineFolder)
	if err != nil {
		stderrorhandler.StdErrorHandler("unzipwaterpolygons.go | Failed unzipping water polygons. Unable to proceed...", err)
		panic(err)
	}
}
