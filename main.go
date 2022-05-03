package main

import (
	"os"

	sequentiallygenerateplanetmbtiles "github.com/lambdajack/sequentially-generate-planet-mbtiles/cmd/sequentially-generate-planet-mbtiles"
)

func main() {
	os.Exit(sequentiallygenerateplanetmbtiles.EntryPoint())
}
