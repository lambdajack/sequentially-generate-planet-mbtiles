package folders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var Pwd, _ = os.Getwd()

var ConfigsFolder = filepath.FromSlash(Pwd + "/" + "configs")
var TilemakerConfigsFolder = filepath.FromSlash(ConfigsFolder + "/" + "tilemaker")

var DataFolder = filepath.FromSlash(Pwd + "/" + "data")
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
		fmt.Println(folder)
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("folders.go | Failed to create %v folder. Unable to proceed. Check permissions etc", folder), err)
			panic(err)
		}

	}
}
