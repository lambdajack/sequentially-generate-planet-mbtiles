package extractslices

import (
	"fmt"
	"log"
	"sync"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/buildthirdpartycontainers"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/folders"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/logger"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/osmboundaryextract"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

var wg = &sync.WaitGroup{}
var mu sync.Mutex

var chCount = make(chan int, 2)

func ExtractSlicesFromQuadrants() {
	for i := -180; i < 180; i++ {
		chCount <- 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			k := fmt.Sprintf("%.2f", setK(i))
			j := fmt.Sprintf("%.2f", setJ(i))
			bbox := fmt.Sprintf("%v,-85,%v,85", k, j)
			src := setSrc(i)
			fileName := fmt.Sprintf("%v.osm.pbf", i+180)

			log.Printf("Extracting slice %s / 359 from %s.\n", fileName, src)

			err := osmboundaryextract.OsmBoundaryExtract(folders.PbfQuadrantSlicesFolder, src, folders.PbfSlicesFolder, fileName, bbox, buildthirdpartycontainers.ContainerOsmiumName)
			if err != nil {
				stderrorhandler.StdErrorHandler(fmt.Sprintf("osmboundaryextractslices.go | Failed to extract %s from %s. Skipping and moving onto next one. Data can be retrospecitively filled in manually later.", bbox, src), err)
				mu.Lock()
				logger.AppendReport(fmt.Sprintf("SLICE_EXTRACT_FAILED: %s from %s", bbox, src))
			} else {
				mu.Lock()
				logger.AppendReport(fmt.Sprintf("SLICE_EXTRACT_SUCCESS: %s from %s", bbox, src))
			}
			mu.Unlock()
			<-chCount
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
}

func setK(i int) float32 {
	if i == -180 {
		return float32(i)
	} else {
		return float32(i) - 0.01
	}
}

func setJ(i int) float32 {
	if i == 179 {
		return float32(180)
	} else {
		return float32(i) + 1.01
	}
}

func setSrc(i int) string {
	var src string

	if i < -90 {
		src = fmt.Sprintf("q1.osm.pbf")
	} else if i < 0 {
		src = fmt.Sprintf("q2.osm.pbf")
	} else if i < 90 {
		src = fmt.Sprintf("q3.osm.pbf")
	} else {
		src = fmt.Sprintf("q4.osm.pbf")
	}

	return src
}
