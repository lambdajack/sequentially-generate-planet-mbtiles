package mbtiles

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/containers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func Generate(src, dst, coastline, landcover, config, process string) {
	src, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}

	dst, err = filepath.Abs(dst)
	if err != nil {
		log.Fatal(err)
	}

	generateMbtilesCmd := fmt.Sprintf("docker run --rm -v %s:/src -v %s:/dst -v %s:/coastline -v %s:/landcover -v %v:/config -v %v:/process %s --input /src/%s --output /dst/%s --config /config/%s --process /process/%s", filepath.Dir(src), filepath.Dir(dst), coastline, landcover, filepath.Dir(config), filepath.Dir(process), containers.ContainerNames.Tilemaker, filepath.Base(src), filepath.Base(dst), filepath.Base(config), filepath.Base(process))

	err = execute.OutputToConsole(generateMbtilesCmd)
	if err != nil {
		log.Fatal(err)
	}
}
