package buildthirdpartycontainers

import (
	"fmt"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

var ContainerTilemakerName = "sequential-tilemaker"
var ContainerTippecanoeName = "sequential-tippecanoe"
var ContainerOsmiumName = "sequential-osmium"

var TilemakerPath = filepath.Clean("./third_party/tilemaker/Dockerfile")
var TippecanoePath = filepath.Clean("./third_party/tippecanoe/Dockerfile")
var OsmiumPath = filepath.Clean("./build/osmium/Dockerfile")

func BuildContainers() {
	execute.OutputToConsole(fmt.Sprintf("docker build -t %s -f %s third_party/tilemaker", ContainerTilemakerName, TilemakerPath))

	execute.OutputToConsole(fmt.Sprintf("docker build -t %s -f %s third_party/tippecanoe", ContainerTippecanoeName, TippecanoePath))

	execute.OutputToConsole(fmt.Sprintf("docker build -t %s -f %s third_party", ContainerOsmiumName, OsmiumPath))
}
