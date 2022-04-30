package main

import (
	"embed"
	"os"

	sequentiallygenerateplanetmbtiles "github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles"
)

//go:embed third_party/tilemaker/resources/config-openmaptiles.json
//go:embed third_party/tilemaker/resources/process-openmaptiles.lua
var embeddedFs embed.FS

func main() {
	os.Exit(sequentiallygenerateplanetmbtiles.EntryPoint(&embeddedFs))
}
