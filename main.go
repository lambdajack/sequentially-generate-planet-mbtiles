package main

import (
	_ "embed"
	"os"

	sequentiallygenerateplanetmbtiles "github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles"
)

//go:embed build/osmium/Dockerfile
var osmiumDockerFile []byte

func main() {
	os.Exit(sequentiallygenerateplanetmbtiles.EntryPoint(osmiumDockerFile))
}
