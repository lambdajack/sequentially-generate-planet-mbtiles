package folders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var DataFolder = filepath.FromSlash("data")
var CoastlineFolder = filepath.FromSlash(DataFolder + "/" + "coastline")
var PbfFolder = filepath.FromSlash(DataFolder + "/" + "pbf")
var PbfSlicesFolder = filepath.FromSlash(PbfFolder + "/" + "slices")
var PbfQuadrantSlicesFolder = filepath.FromSlash(PbfSlicesFolder + "/" + "quadrants")
var MbtilesFolder = filepath.FromSlash(DataFolder + "/" + "mbtiles")
var MbtilesMergedFolder = filepath.FromSlash(MbtilesFolder + "/" + "merged")

func SetupFolderStructure() {
	allFolders := [...]string{DataFolder, CoastlineFolder, PbfFolder, PbfSlicesFolder, PbfQuadrantSlicesFolder, MbtilesFolder, MbtilesMergedFolder}

	for _, folder := range allFolders {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("folders.go | Failed to create %v folder. Unable to proceed. Check permissions etc", folder), err)
			panic(err)
		}

	}
}
