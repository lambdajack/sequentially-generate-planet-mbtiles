package containers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

type container struct {
	name       string
	dockerfile string
	context    string
}

type names struct {
	Tilemaker  string
	Tippecanoe string
	Osmium     string
	Gdal       string
}

var ContainerNames = names{
	Tilemaker:  "sequential-tilemaker",
	Tippecanoe: "sequential-tippecanoe",
	Osmium:     "sequential-osmium",
	Gdal:       "sequential-gdal",
}

var ct = []container{
	{
		name:       ContainerNames.Tilemaker,
		dockerfile: filepath.Clean("third_party/tilemaker/Dockerfile"),
		context:    filepath.Clean("third_party/tilemaker"),
	},
	{
		name:       ContainerNames.Tippecanoe,
		dockerfile: filepath.Clean("third_party/tippecanoe/Dockerfile"),
		context:    filepath.Clean("third_party/tippecanoe"),
	},
	{
		name:       ContainerNames.Osmium,
		dockerfile: filepath.Clean("build/osmium/Dockerfile"),
		context:    filepath.Clean("third_party"),
	},
	{
		name:       ContainerNames.Gdal,
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

func CleanAll() error {
	fmt.Println("Attempting to gracefully shutdown containers...")

	var e error

	for _, c := range ct {
		cmd := exec.Command("docker", "ps", "-a", "-q", "--filter", "ancestor="+c.name+":latest")
		out, err := cmd.Output()
		if err != nil {
			e = err
		}
		a := strings.Split(string(out), "\n")
		for _, n := range a {
			if len(n) > 0 {
				fmt.Println("Stopping container: ", n, c.name)
				err := execute.OutputToConsole(fmt.Sprintf("docker stop %s", n))
				if err != nil {
					fmt.Println("Failed to stop container: ", n)
					e = err
				}
			}
		}
	}

	return e
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
