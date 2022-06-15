package containers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type container struct {
	name       string
	dockerfile string
	context    string
}

var ct = []container{
	{
		name:       "sequential-tilemaker",
		dockerfile: filepath.Clean("third_party/tilemaker/Dockerfile"),
		context:    filepath.Clean("third_party/tilemaker"),
	},
	{
		name:       "sequential-tippecanoe",
		dockerfile: filepath.Clean("third_party/tippecanoe/Dockerfile"),
		context:    filepath.Clean("third_party/tippecanoe"),
	},
	{
		name:       "sequential-osmium",
		dockerfile: filepath.Clean("build/osmium/Dockerfile"),
		context:    filepath.Clean("third_party"),
	},
	{
		name:       "sequential-gdal",
		dockerfile: filepath.Clean("third_party/gdal/docker/alpine-small/Dockerfile"),
		context:    filepath.Clean("third_party/gdal"),
	},
}

func BuildAll() error {
	for _, c := range ct {
		err := buildContainer(c.name, c.dockerfile, c.context)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildContainer(name, dockerfile, context string) error {
	cmd := exec.Command("docker", "build", "-t", name, "-f", dockerfile, context)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		err = fmt.Errorf("failed to build container %v, with dockerfile %v, and context %v", name, dockerfile, context)
		return err
	}
	cmd.Wait()
	return nil
}