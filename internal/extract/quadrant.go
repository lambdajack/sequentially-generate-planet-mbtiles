package extract

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"
)

type Quadrant struct {
	name, bbox string
}

const maxRoutines = 2

var wg = &sync.WaitGroup{}

func Quadrants(src, dst, containerName string) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	w := true

	// c := make(chan os.Signal)
    // signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    // go func() {
    //     <-c
	// 	w = false
    //     fmt.Println("\nAttempting to stop containers on interrupt: ", containerName)
	// 	cmd := exec.Command("docker", "ps", "-a", "-q", "--filter", "ancestor="+containerName+":latest")
	// 	out, _ := cmd.Output()
	// 	a := strings.Split(string(out), "\n")
	// 	for _, c := range a {
	// 		execute.OutputToConsole(fmt.Sprintf("docker kill %s", c))
	// 	}
    //     os.Exit(1)
    // }()
	// defer close(c)

	chCount := make(chan int, maxRoutines)

	quadrantsToGenerate := [...]Quadrant{
		{name: "q1.osm.pbf", bbox: "-180,-85,-90.1,85"},
		{name: "q2.osm.pbf", bbox: "-89.9,-85,0.1,85"},
		{name: "q3.osm.pbf", bbox: "-0.1,-85,90.1,85"},
		{name: "q4.osm.pbf", bbox: "89.9,-85,180,85"}}

		
		go func(w *bool) {
			time.Sleep(time.Second)
			for *w {
				fmt.Printf("\rExtracting quadrants... |")
				time.Sleep(time.Second)
				fmt.Printf("\rExtracting quadrants... /")
				time.Sleep(time.Second)
				fmt.Printf("\rExtracting quadrants... -")
				time.Sleep(time.Second)
				fmt.Printf("\rExtracting quadrants... \\")
				time.Sleep(time.Second)
			}
		}(&w)

	for _, quadrant := range quadrantsToGenerate {
		chCount <- 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, quadrant Quadrant) {
			log.Printf("Extracting quadrant %s in parallel. ~1 hour.\n", quadrant.name)
			_, err := Extract(src, filepath.Join(dst, quadrant.name), quadrant.bbox, containerName)
			if err != nil {
				fmt.Printf("osmboundaryextractquadrants.go | Failed to extract %s from planet-latest.osm.pbf... unable to proceed", quadrant.bbox)
				log.Fatal(err)
			}
			<-chCount
			wg.Done()
		}(wg, quadrant)
	}

	wg.Wait()
	w = false
}
