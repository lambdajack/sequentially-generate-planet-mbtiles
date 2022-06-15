package extract

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func Extract(src, dst, bbox, containerName string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		cmdString := fmt.Sprintf("docker run --rm -v %s:/pbf -v %s:/out %s osmium extract -b %s --set-bounds /pbf/%s -o /out/%s", filepath.Dir(src), filepath.Base(src), containerName, bbox, filepath.Base(src), dst)

		err := execute.OutputToConsole(cmdString)
		if err != nil {
			return fmt.Errorf("osmboundaryextract.go | Failed to extract %s from %s \n %v", bbox, filepath.Base(src), err)
		}
	} else {
		log.Printf("%s already exists. Skipping...", dst)
	}

	return nil
}
