package sequentiallygenerateplanetmbtiles

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/docker"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/system"
)

type containers struct {
	gdal       *docker.Container
	osmium     *docker.Container
	tilemaker  *docker.Container
	tippecanoe *docker.Container
}

var ct containers

func setupContainers(osmiumDf []byte) {
	ct = containers{
		gdal: docker.New(docker.Container{
			Name:       "sequential-gdal",
			Dockerfile: filepath.Join(gh.gdal.Dst, "docker", "alpine-small", "Dockerfile"),
			Context:    gh.gdal.Dst,
		}),
		osmium: docker.New(docker.Container{
			Name:       "sequential-osmium",
			Dockerfile: setOsmiumDockerfile(osmiumDf),
			// Context set to pth.temp/osmium since the docker file needs to pull in two separate repos, both in the same dir
			Context: filepath.Join(pth.temp, "osmium"),
		}),
		tilemaker: docker.New(docker.Container{
			Name:       "sequential-tilemaker",
			Dockerfile: filepath.Join(gh.tilemaker.Dst, "Dockerfile"),
			Context:    gh.tilemaker.Dst,
		}),
		tippecanoe: docker.New(docker.Container{
			Name:       "sequential-tippecanoe",
			Dockerfile: filepath.Join(gh.tippecanoe.Dst, "Dockerfile"),
			Context:    gh.tippecanoe.Dst,
		}),
	}

	err := ct.osmium.Build()
	if err != nil {
		lg.err.Fatalln("failed to build osmium container:", err)
	}

	err = ct.gdal.Build()
	if err != nil {
		lg.err.Fatalln("failed to build gdal container:", err)
	}

	err = ct.tilemaker.Build()
	if err != nil {
		lg.err.Fatalln("failed to build tilemaker container:", err)
	}

	err = ct.tippecanoe.Build()
	if err != nil {
		lg.err.Fatalln("failed to build tippecanoe container:", err)
	}
}

func cleanContainers() {
	err := ct.gdal.Clean()
	if err != nil {
		lg.err.Println("failed to clean gdal container:", err)
	}

	err = ct.osmium.Clean()
	if err != nil {
		lg.err.Println("failed to clean osmium container:", err)
	}

	err = ct.tilemaker.Clean()
	if err != nil {
		lg.err.Println("failed to clean tilemaker container:", err)
	}

	err = ct.tippecanoe.Clean()
	if err != nil {
		lg.err.Println("failed to clean tippecanoe container:", err)
	}
}

func setOsmiumDockerfile(df []byte) string {
	f, err := os.Create(filepath.Join(pth.temp, "Dockerfile"))
	if err != nil {
		lg.err.Fatalln("failed to create Osmium Dockerfile:", err)
	}

	_, err = f.Write(df)
	if err != nil {
		lg.err.Fatalln("failed to write Osmium Dockerfile:", err)
	}

	system.SetUserOwner(f.Name())

	return f.Name()
}
