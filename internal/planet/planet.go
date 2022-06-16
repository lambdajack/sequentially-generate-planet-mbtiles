package planet

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/containers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func Generate(src, dst string) {
	fi, err := os.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}

	b := strings.Builder{}
	for _, f := range fi {
		if !f.IsDir() {
			b.WriteString("/data/" + f.Name() + " ")
		}
	}

	if _, err := os.Stat(filepath.Join(dst, "planet.mbtiles")); os.IsNotExist(err) {
		f, err := os.Create(filepath.Join(dst, "planet.mbtiles"))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}

	mergeCmd := fmt.Sprintf("docker run --rm -v %v:/data -v %v:/merged %v tile-join --output=/merged/planet.mbtiles %v", src, dst, containers.ContainerNames.Tippecanoe, b.String())
	log.Println(mergeCmd)
	err = execute.OutputToConsole((mergeCmd))
	if err != nil {
		log.Fatalf("Failed to merge mbtiles: %v", err)
	}
}