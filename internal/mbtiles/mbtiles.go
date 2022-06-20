package mbtiles

import (
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
)

func Generate(src, dst, coastline, landcover, config, process, containerName string, outAsDir bool) {
	src, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		log.Fatal(err)
	}

	if !outAsDir {
		var sb strings.Builder
		r := rand.Perm(10)
		for _, o := range r {
			sb.WriteString(fmt.Sprintf("%d", o))
		}

		dst = filepath.Join(dst, sb.String()+".mbtiles")
	}

	generateMbtilesCmd := fmt.Sprintf("docker run --rm -v %s:/src -v %s:/dst -v %s:/coastline -v %s:/landcover -v %v:/config -v %v:/process %s --input /src/%s --output /dst/%s --config /config/%s --process /process/%s", filepath.Dir(src), filepath.Dir(dst), coastline, landcover, filepath.Dir(config), filepath.Dir(process), containerName, filepath.Base(src), filepath.Base(dst), filepath.Base(config), filepath.Base(process))

	err = execute.OutputToConsole(generateMbtilesCmd)
	if err != nil {
		log.Fatal(err)
	}
}
