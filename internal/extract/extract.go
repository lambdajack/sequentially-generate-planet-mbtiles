package extract

import (
	"fmt"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func Extract(src, dst, bbox, containerName string) (string, error) {
	src, err := filepath.Abs(src)
	if err != nil {
		return "", err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return "", err
	}

	
		cmdString := fmt.Sprintf("docker run --rm -v %v:/pbf -v %s:/out %s osmium extract -b %s --set-bounds --overwrite /pbf/%s -o /out/%s", filepath.Dir(src), filepath.Dir(dst), containerName, bbox, filepath.Base(src), filepath.Base(dst))

		err = execute.OutputToConsole(cmdString)
		if err != nil {
			return "", fmt.Errorf("osmboundaryextract.go | Failed to extract %s from %s \n %v", bbox, filepath.Base(src), err)
		}
	
	
	return dst, nil
}
