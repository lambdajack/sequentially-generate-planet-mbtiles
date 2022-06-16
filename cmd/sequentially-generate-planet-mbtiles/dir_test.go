package sequentiallygenerateplanetmbtiles

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeDir(t *testing.T) {
	tmp := t.TempDir()

	Dir := filepath.Join(tmp, "test")

	makeDir(Dir)

	if _, err := os.Stat(Dir); os.IsNotExist(err) {
		t.Errorf("Dir %s does not exist", Dir)
	}
}

func TestInitDirStruc(t *testing.T) {

	initDirStructure()

	DirMap := map[string]string{
		"outDir":                   filepath.Clean(cfg.OutDir),
		"coastlineDir":             filepath.Join(pth.workingDir, "coastline"),
		"landcoverDir":             filepath.Join(pth.workingDir, "landcover"),
		"landCoverUrbanDepth":      filepath.Join(pth.landcoverDir, "ne_10m_urban_areas"),
		"landCoverIceShelvesDepth": filepath.Join(pth.landcoverDir, "ne_10m_antarctic_ice_shelves_polys"),
		"landCoverGlaciatedDepth":  filepath.Join(pth.landcoverDir, "ne_10m_glaciated_areas"),
		"pbfDir":                   filepath.Join(pth.workingDir, "pbf"),
		"pbfSlicesDir":             filepath.Join(pth.pbfDir, "slices"),
		"pbfQuadrantSlicesDir":     filepath.Join(pth.pbfDir, "quadrants"),
		"mbtilesDir":               filepath.Join(pth.workingDir, "mbtiles"),
		"logsDir":                  filepath.Join(pth.workingDir, "logs"),
	}

	for Dir := range DirMap {
		if _, err := os.Stat(DirMap[Dir]); os.IsNotExist(err) {
			t.Errorf("Dir %s does not exist", DirMap[Dir])
		}
	}
}
