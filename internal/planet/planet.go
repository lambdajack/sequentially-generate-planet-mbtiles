package planet

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/docker"
)

func Generate(src, dst string, tippecanoe *docker.Container, elg, plg, rlg *log.Logger) string {
	fi, err := os.ReadDir(src)
	if err != nil {
		elg.Fatal(err)
	}

	var b []string
	for _, f := range fi {
		if !f.IsDir() {
			b = append(b, "/data/"+f.Name())
		}
	}

	if len(b) == 0 {
		elg.Fatalf("cannot find any tiles to merge in %s - have you generated any?", src)
	}

	if _, err := os.Stat(filepath.Join(dst, "planet.mbtiles")); os.IsNotExist(err) {
		f, err := os.Create(filepath.Join(dst, "planet.mbtiles"))
		if err != nil {
			elg.Fatal(err)
		}
		f.Close()
	}

	tippecanoe.Volumes = []docker.Volume{
		{
			Container: "/data",
			Host:      src,
		},
		{
			Container: "/merged",
			Host:      dst,
		},
	}

	rlg.Println("merging: ", strings.ReplaceAll(strings.Join(b, " "), "/data/", "..."))

	err = tippecanoe.Execute(append([]string{"tile-join", "--force", "--output=/merged/planet.mbtiles"}, b...))
	if err != nil {
		elg.Fatalf("failed to merge mbtiles: %v", err)
	}
	plg.Println("MERGING: successfully merged: ", strings.ReplaceAll(strings.Join(b, " "), "/data/", "..."))

	return filepath.Join(dst, "planet.mbtiles")
}
