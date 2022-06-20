package sequentiallygenerateplanetmbtiles

import (
	"io/fs"
	"os"
	"path/filepath"
)

func moveOcean() {
	if !cfg.ExcludeOcean {
		filepath.Walk(pth.coastlineDir, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				err := os.Rename(path, filepath.Join(pth.coastlineDir, filepath.Base(path)))
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
}
