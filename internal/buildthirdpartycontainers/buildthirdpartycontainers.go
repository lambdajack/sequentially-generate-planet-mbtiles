package buildthirdpartycontainers

import (
	"fmt"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

var ContainerTilemakerName = "sequential-tilemaker"
var ContainerTippecanoeName = "sequential-tippecanoe"
var ContainerOsmiumName = "sequential-osmium"

func BuildContainers() {
	execute.OutputToConsole(fmt.Sprintf("docker build -t %v ./tilemaker", ContainerTilemakerName))

	execute.OutputToConsole(fmt.Sprintf("docker build -t %v ./tippecanoe", ContainerTippecanoeName))

	execute.OutputToConsole(fmt.Sprintf("docker build -t %v build/osmium", ContainerOsmiumName))
}
