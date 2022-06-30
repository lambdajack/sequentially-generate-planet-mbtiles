package sequentiallygenerateplanetmbtiles

import (
	"log"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/system"
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
	mbtilesDir               string
	logsDir                  string
	temp                     string
}

var pth = paths{}

func initDirStructure() {

	tmp, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		log.Printf("Unable to create temp dir: %s", err)
		os.Exit(exitPermissions)
	}

	pth.workingDir = convertAbs(cfg.WorkingDir)
	makeDir(pth.workingDir)

	pth.outDir = convertAbs(cfg.OutDir)
	makeDir(pth.outDir)

	if !cfg.ExcludeOcean {
		pth.coastlineDir = filepath.Join(pth.workingDir, "coastline")
		makeDir(pth.coastlineDir)
	} else {
		pth.coastlineDir = filepath.Join(tmp, "coastline")
	}

	if !cfg.ExcludeLanduse {
		pth.landcoverDir = filepath.Join(pth.workingDir, "landcover")
		makeDir(pth.landcoverDir)
	} else {
		pth.landcoverDir = filepath.Join(tmp, "landcover")
	}

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

	pth.mbtilesDir = filepath.Join(pth.workingDir, "mbtiles")
	makeDir(pth.mbtilesDir)

	pth.logsDir = filepath.Join(pth.workingDir, "logs")
	makeDir(pth.logsDir)

	if system.DockerIsSnap() {
		log.Println("snap version of docker detected; using local tmp folder since snap docker cannot access system /tmp")
		pth.temp = filepath.Join(pth.workingDir, "tmp")
	} else {
		ucd := system.UserCacheDir()
		if ucd != "" {
			pth.temp = filepath.Join(system.UserCacheDir(), "sequentially-generate-planet-mbtiles")
		} else {
			pth.temp = filepath.Join(pth.workingDir, "tmp")
		}
	}
	makeDir(pth.temp)
}

func convertAbs(path string) string {
	if path == "" {
		return ""
	}

	p, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func makeDir(fldr string) {
	if err := os.MkdirAll(fldr, os.ModePerm); err != nil {
		log.Printf("Unable to create %s Dir", fldr)
		os.Exit(exitPermissions)
	}
	system.SetUserOwner(fldr)
}
