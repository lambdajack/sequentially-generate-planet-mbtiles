package folders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/config"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var Pwd, _ = os.Getwd()

var ConfigsFolder = formatFolderString(Pwd + "/" + "configs")

var TilemakerConfigFile = formatFolderString(Pwd + "/" + config.Config.PathToTilemakerConfig)
var TilemakerProcessFile = formatFolderString(Pwd + "/" + config.Config.PathToTilemakerProcess)

var DataFolder = formatFolderString(Pwd + "/" + "data")
var CoastlineFolder = formatFolderString(DataFolder + "/" + "coastline")
var PbfFolder = formatFolderString(DataFolder + "/" + "pbf")
var PbfSlicesFolder = formatFolderString(PbfFolder + "/" + "slices")
var PbfQuadrantSlicesFolder = formatFolderString(PbfSlicesFolder + "/" + "quadrants")
var MbtilesFolder = formatFolderString(DataFolder + "/" + "mbtiles")
var MbtilesMergedFolder = formatFolderString(MbtilesFolder + "/" + "merged")

func SetupFolderStructure() {
	if config.Config.DataDir != "data" {
		DataFolder = formatFolderString(config.Config.DataDir)
		CoastlineFolder = formatFolderString(DataFolder + "/" + "coastline")
		PbfFolder = formatFolderString(DataFolder + "/" + "pbf")
		PbfSlicesFolder = formatFolderString(PbfFolder + "/" + "slices")
		PbfQuadrantSlicesFolder = formatFolderString(PbfSlicesFolder + "/" + "quadrants")
		MbtilesFolder = formatFolderString(DataFolder + "/" + "mbtiles")
		MbtilesMergedFolder = formatFolderString(MbtilesFolder + "/" + "merged")
	}

	if config.Config.PathToTilemakerConfig != "configs/tilemaker/config.json" {
		TilemakerConfigFile = formatFolderString(config.Config.PathToTilemakerConfig)
	}

	if config.Config.PathToTilemakerProcess != "configs/tilemaker/process.lua" {
		TilemakerProcessFile = formatFolderString(config.Config.PathToTilemakerProcess)
	}

	allFolders := [...]*string{&DataFolder, &CoastlineFolder, &PbfFolder, &PbfSlicesFolder, &PbfQuadrantSlicesFolder, &MbtilesFolder, &MbtilesMergedFolder}

	for _, folder := range allFolders {
		err := os.MkdirAll(*folder, os.ModePerm)
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("folders.go | Failed to create %v folder. Unable to proceed. Check permissions etc", *folder), err)
			panic(err)
		}
	}
}

func formatFolderString(folder string) string {
	folder = filepath.Clean(folder)
	folder = filepath.FromSlash(folder)
	return folder
}
