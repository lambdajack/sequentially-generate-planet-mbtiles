package sequentiallygenerateplanetmbtiles

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeFolder(t *testing.T) {
	tmp := t.TempDir()

	folder := filepath.Join(tmp, "test")

	makeFolder(folder)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		t.Errorf("Folder %s does not exist", folder)
	}
}

func TestInitFolderStruc(t *testing.T) {

	initFolderStructure()

	folderMap := map[string]string{
		"outFolder":                filepath.Clean(cfg.OutDir),
		"coastlineFolder":          filepath.Join(pth.dataFolder, "coastline"),
		"landcoverFolder":          filepath.Join(pth.dataFolder, "landcover"),
		"landCoverUrbanDepth":      filepath.Join(pth.landcoverFolder, "ne_10m_urban_areas"),
		"landCoverIceShelvesDepth": filepath.Join(pth.landcoverFolder, "ne_10m_antarctic_ice_shelves_polys"),
		"landCoverGlaciatedDepth":  filepath.Join(pth.landcoverFolder, "ne_10m_glaciated_areas"),
		"pbfFolder":                filepath.Join(pth.dataFolder, "pbf"),
		"pbfSlicesFolder":          filepath.Join(pth.pbfFolder, "slices"),
		"pbfQuadrantSlicesFolder":  filepath.Join(pth.pbfFolder, "quadrants"),
		"mbtilesFolder":            filepath.Join(pth.dataFolder, "mbtiles"),
		"logsFolder":               filepath.Join(pth.dataFolder, "logs"),
	}

	for folder := range folderMap {
		if _, err := os.Stat(folderMap[folder]); os.IsNotExist(err) {
			t.Errorf("Folder %s does not exist", folderMap[folder])
		}
	}
}
