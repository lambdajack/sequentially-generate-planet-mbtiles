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

var TilemakerConfigFile = formatFolderString(Pwd + "/" + "third_party" + "/" + "tilemaker" + "/" + "resources" + "/" + "config-openmaptiles.json")
var TilemakerProcessFile = formatFolderString(Pwd + "/" + "third_party" + "/" + "tilemaker" + "/" + "resources" + "/" + "process-openmaptiles.lua")

var DataFolder = formatFolderString(Pwd + "/" + "data")
var CoastlineFolder = formatFolderString(DataFolder + "/" + "coastline")

var LandcoverFolder = formatFolderString(DataFolder + "/" + "landcover")
var LandCoverUrbanDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_urban_areas")
var LandCoverIceShelvesDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_antarctic_ice_shelves_polys")
var LandCoverGlaciatedDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_glaciated_areas")

var PbfFolder = formatFolderString(DataFolder + "/" + "pbf")
var PbfSlicesFolder = formatFolderString(PbfFolder + "/" + "slices")
var PbfQuadrantSlicesFolder = formatFolderString(PbfSlicesFolder + "/" + "quadrants")
var MbtilesFolder = formatFolderString(DataFolder + "/" + "mbtiles")
var MbtilesMergedFolder = formatFolderString(MbtilesFolder + "/" + "merged")

func SetupFolderStructure() {
	if config.Config.DataDir != "" {
		DataFolder = formatFolderString(config.Config.DataDir)
		CoastlineFolder = formatFolderString(DataFolder + "/" + "coastline")
		LandcoverFolder = formatFolderString(DataFolder + "/" + "landcover")
		LandCoverUrbanDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_urban_areas")
		LandCoverIceShelvesDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_antarctic_ice_shelves_polys")
		LandCoverGlaciatedDepth = formatFolderString(LandcoverFolder + "/" + "ne_10m_glaciated_areas")
		PbfFolder = formatFolderString(DataFolder + "/" + "pbf")
		PbfSlicesFolder = formatFolderString(PbfFolder + "/" + "slices")
		PbfQuadrantSlicesFolder = formatFolderString(PbfSlicesFolder + "/" + "quadrants")
		MbtilesFolder = formatFolderString(DataFolder + "/" + "mbtiles")
		MbtilesMergedFolder = formatFolderString(MbtilesFolder + "/" + "merged")
	}

	if config.Config.TilemakerConfig != "" {
		TilemakerConfigFile = formatFolderString(config.Config.TilemakerConfig)
	}

	if config.Config.TilemakerProcess != "" {
		TilemakerProcessFile = formatFolderString(config.Config.TilemakerProcess)
	}

	allFolders := [...]*string{&DataFolder, &CoastlineFolder, &LandcoverFolder, &PbfFolder, &PbfSlicesFolder, &PbfQuadrantSlicesFolder, &MbtilesFolder, &MbtilesMergedFolder, &LandCoverUrbanDepth, &LandCoverIceShelvesDepth, &LandCoverGlaciatedDepth}

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
	return folder
}
