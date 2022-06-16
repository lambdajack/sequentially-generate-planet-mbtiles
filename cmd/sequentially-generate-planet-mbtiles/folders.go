package sequentiallygenerateplanetmbtiles

import (
	"log"
	"os"
	"path/filepath"
)

type paths struct {
	workingDir               string
	outDir                   string
	coastlineDir             string
	landcoverDir             string
	landCoverUrbanDepth      string
	landCoverIceShelvesDepth string
	landCoverGlaciatedDepth  string
	pbfDir                   string
	pbfSlicesDir             string
	pbfQuadrantSlicesDir     string
	mbtilesDir               string
	logsDir                  string
}

var pth = paths{}

func initDirStructure() {
	pth.workingDir = filepath.Clean(cfg.WorkingDir)
	makeDir(pth.workingDir)

	pth.outDir = filepath.Clean(cfg.OutDir)
	makeDir(pth.outDir)

	pth.coastlineDir = filepath.Join(pth.workingDir, "coastline")
	makeDir(pth.coastlineDir)

	pth.landcoverDir = filepath.Join(pth.workingDir, "landcover")
	makeDir(pth.landcoverDir)

	pth.landCoverUrbanDepth = filepath.Join(pth.landcoverDir, "ne_10m_urban_areas")
	makeDir(pth.landCoverUrbanDepth)

	pth.landCoverIceShelvesDepth = filepath.Join(pth.landcoverDir, "ne_10m_antarctic_ice_shelves_polys")
	makeDir(pth.landCoverIceShelvesDepth)

	pth.landCoverGlaciatedDepth = filepath.Join(pth.landcoverDir, "ne_10m_glaciated_areas")
	makeDir(pth.landCoverGlaciatedDepth)

	pth.pbfDir = filepath.Join(pth.workingDir, "pbf")
	makeDir(pth.pbfDir)

	pth.pbfSlicesDir = filepath.Join(pth.pbfDir, "slices")
	makeDir(pth.pbfSlicesDir)

	pth.pbfQuadrantSlicesDir = filepath.Join(pth.pbfDir, "quadrants")
	makeDir(pth.pbfQuadrantSlicesDir)

	pth.mbtilesDir = filepath.Join(pth.workingDir, "mbtiles")
	makeDir(pth.mbtilesDir)

	pth.logsDir = filepath.Join(pth.workingDir, "logs")
	makeDir(pth.logsDir)
}

func makeDir(fldr string) {
	if err := os.MkdirAll(fldr, os.ModePerm); err != nil {
		log.Printf("Unable to create %s Dir", fldr)
		os.Exit(exitPermissions)
	}
}
