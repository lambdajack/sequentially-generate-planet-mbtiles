package buildthirdpartycontainers

import (
	"fmt"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func BuildContainers() (containerNames []string) {
	dockerTilemakerName := "sequential-tilemaker"
	execute.OutputToConsole(fmt.Sprintf("docker build -t %v ./tilemaker", dockerTilemakerName))

	dockerTippecanoeName := "sequential-tippecanoe"
	execute.OutputToConsole(fmt.Sprintf("docker build -t %v ./tippecanoe", dockerTippecanoeName))

	dockerOsmiumName := "sequential-osmium"
	execute.OutputToConsole(fmt.Sprintf("docker build -t %v build/osmium", dockerOsmiumName))

	containerNames = []string{dockerTilemakerName, dockerTippecanoeName, dockerOsmiumName}
	return containerNames
}
