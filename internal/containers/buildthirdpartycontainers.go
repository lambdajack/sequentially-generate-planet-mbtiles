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
type containers struct {
	tilemaker  container
	tippecanoe container
	osmium     container
}

var ct = &containers{
	tilemaker: container{
		name:       "sequential-tilemaker",
		dockerfile: filepath.Clean("third_party/tilemaker/Dockerfile"),
		context:    filepath.Clean("third_party/tilemaker"),
	},
	tippecanoe: container{
		name:       "sequential-tippecanoe",
		dockerfile: filepath.Clean("third_party/tippecanoe/Dockerfile"),
		context:    filepath.Clean("third_party/tippecanoe"),
	},
	osmium: container{
		name:       "sequential-osmium",
		dockerfile: filepath.Clean("build/osmium/Dockerfile"),
		context:    filepath.Clean("third_party"),
	},
}

func BuildAll() error {
	err := buildContainer(ct.tilemaker.name, ct.tilemaker.dockerfile, ct.tilemaker.context)
	if err != nil {
		return err
	}

	err = buildContainer(ct.tippecanoe.name, ct.tippecanoe.dockerfile, ct.tippecanoe.context)
	if err != nil {
		return err
	}

	err = buildContainer(ct.osmium.name, ct.osmium.dockerfile, ct.osmium.context)
	if err != nil {
		return err
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
