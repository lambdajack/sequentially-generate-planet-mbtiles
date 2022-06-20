package extract

import (
	"fmt"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/docker"
)

func Extract(src, dst, bbox string, osmium *docker.Container) (string, error) {
	src, err := filepath.Abs(src)
	if err != nil {
		return "", err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return "", err
	}

	osmium.Volumes = []docker.Volume{
		{
			Container: "/pbf",
			Host:      filepath.Dir(src),
		},
		{
			Container: "/out",
			Host:      filepath.Dir(dst),
		},
	}
	
	// cmdString := fmt.Sprintf("docker run --rm -v %v:/pbf -v %s:/out %s osmium extract -b %s --set-bounds --overwrite /pbf/%s -o /out/%s", filepath.Dir(src), filepath.Dir(dst), containerName, bbox, filepath.Base(src), filepath.Base(dst))
	err = osmium.Execute([]string{"osmium", "extract", "-b", bbox, "--set-bounds", "--overwrite", "/pbf/" + filepath.Base(src), "-o", "/out/" + filepath.Base(dst)})
	if err != nil {
		return "", fmt.Errorf("osmboundaryextract.go | Failed to extract %s from %s \n %v", bbox, filepath.Base(src), err)
	}

	return dst, nil
}
