package extract

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
)

type Quadrant struct {
	name, bbox string
}

const maxRoutines = 2

var wg = &sync.WaitGroup{}

func Quadrants(src, dst, containerName string) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	chCount := make(chan int, maxRoutines)

	quadrantsToGenerate := [...]Quadrant{
		{name: "q1.osm.pbf", bbox: "-180,-85,-90.1,85"},
		{name: "q2.osm.pbf", bbox: "-89.9,-85,0.1,85"},
		{name: "q3.osm.pbf", bbox: "-0.1,-85,90.1,85"},
		{name: "q4.osm.pbf", bbox: "89.9,-85,180,85"}}

	for _, quadrant := range quadrantsToGenerate {
		chCount <- 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, quadrant Quadrant) {
			log.Printf("Extracting quadrant %s in parallel. ~1 hour.\n", quadrant.name)
			_, err := Extract(src, dst+quadrant.name, quadrant.bbox, containerName)
			if err != nil {
				fmt.Printf("osmboundaryextractquadrants.go | Failed to extract %s from planet-latest.osm.pbf... unable to proceed", quadrant.bbox)
				panic(err)
			}
			<-chCount
			wg.Done()
		}(wg, quadrant)
	}
	wg.Wait()
}
