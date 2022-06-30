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