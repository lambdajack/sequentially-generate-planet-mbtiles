package sequentiallygenerateplanetmbtiles

import (
	"log"
	"os"
	"path/filepath"
)

func initFolderStructure() {
	pth.dataFolder = filepath.Clean(cfg.DataDir)
	makeFolder(pth.dataFolder)

	pth.outFolder = filepath.Clean(cfg.OutDir)
	makeFolder(pth.outFolder)
	
	pth.coastlineFolder = filepath.Clean(pth.dataFolder + "/coastline")
	makeFolder(pth.coastlineFolder)
	
	pth.landcoverFolder = filepath.Clean(pth.dataFolder + "/landcover")
	makeFolder(pth.landcoverFolder)
	
	pth.landCoverUrbanDepth = filepath.Clean(pth.landcoverFolder + "/ne_10m_urban_areas")
	makeFolder(pth.landCoverUrbanDepth)
	
	pth.landCoverIceShelvesDepth = filepath.Clean(pth.landcoverFolder + "/ne_10m_antarctic_ice_shelves_polys")
	makeFolder(pth.landCoverIceShelvesDepth)
	
	pth.landCoverGlaciatedDepth = filepath.Clean(pth.landcoverFolder + "/ne_10m_glaciated_areas")
	makeFolder(pth.landCoverGlaciatedDepth)
	
	pth.pbfFolder = filepath.Clean(pth.dataFolder + "/pbf")
	makeFolder(pth.pbfFolder)
	
	pth.pbfSlicesFolder = filepath.Clean(pth.pbfFolder + "/slices")
	makeFolder(pth.pbfSlicesFolder)
	
	pth.pbfQuadrantSlicesFolder = filepath.Clean(pth.pbfFolder + "/quadrants")
	makeFolder(pth.pbfQuadrantSlicesFolder)
	
	pth.mbtilesFolder = filepath.Clean(pth.dataFolder + "/mbtiles")
	makeFolder(pth.mbtilesFolder)
	
	pth.logsFolder = filepath.Clean(pth.dataFolder + "/logs")
	makeFolder(pth.logsFolder)
}

func makeFolder(fldr string) {
	if err := os.MkdirAll(fldr, os.ModePerm); err != nil {
		log.Printf("Unable to create %s folder", fldr)
		os.Exit(exitPermissions)
	}
}