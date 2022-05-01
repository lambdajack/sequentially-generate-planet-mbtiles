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

	pth.coastlineFolder = filepath.Join(pth.dataFolder, "coastline")
	makeFolder(pth.coastlineFolder)

	pth.landcoverFolder = filepath.Join(pth.dataFolder, "landcover")
	makeFolder(pth.landcoverFolder)

	pth.landCoverUrbanDepth = filepath.Join(pth.landcoverFolder, "ne_10m_urban_areas")
	makeFolder(pth.landCoverUrbanDepth)

	pth.landCoverIceShelvesDepth = filepath.Join(pth.landcoverFolder, "ne_10m_antarctic_ice_shelves_polys")
	makeFolder(pth.landCoverIceShelvesDepth)

	pth.landCoverGlaciatedDepth = filepath.Join(pth.landcoverFolder, "ne_10m_glaciated_areas")
	makeFolder(pth.landCoverGlaciatedDepth)

	pth.pbfFolder = filepath.Join(pth.dataFolder, "pbf")
	makeFolder(pth.pbfFolder)

	pth.pbfSlicesFolder = filepath.Join(pth.pbfFolder, "slices")
	makeFolder(pth.pbfSlicesFolder)

	pth.pbfQuadrantSlicesFolder = filepath.Join(pth.pbfFolder, "quadrants")
	makeFolder(pth.pbfQuadrantSlicesFolder)

	pth.mbtilesFolder = filepath.Join(pth.dataFolder, "mbtiles")
	makeFolder(pth.mbtilesFolder)

	pth.logsFolder = filepath.Join(pth.dataFolder, "logs")
	makeFolder(pth.logsFolder)
}

func makeFolder(fldr string) {
	if err := os.MkdirAll(fldr, os.ModePerm); err != nil {
		log.Printf("Unable to create %s folder", fldr)
		os.Exit(exitPermissions)
	}
}
