package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/clonerepo"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/config"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/downloadsubregion"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/flags"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/logger"
)

var wg sync.WaitGroup

func main() {
		// Set up folder structure and report file
		pbfFolder := "pbf"
		mbtilesFolder := "mbtiles"
		mergedFolder := "mbtiles/merged"
		folders := []string{pbfFolder, mbtilesFolder, mergedFolder}
		for _, folder := range folders {
			if _, err := os.Stat(folder); os.IsNotExist(err) {
				os.MkdirAll(folder, os.ModePerm)
			} else {
				log.Printf("Folder: %v already exists. Skipping creation.\n", folder)
			}
		}
		reportFile := mergedFolder + "/REPORT.txt"
		if _, err := os.Stat(reportFile); os.IsNotExist(err) {
			log.Println("Cannot append REPORT.txt as it does not exist")
		}

			// folder references
	pwd, _ := os.Getwd()
	pbfPath := pwd + "/" + pbfFolder
	mbtilesPath := pwd + "/" + mbtilesFolder


	// Load flags and config.json if supplied; otherwise use defaults.
	pathToConfig := flags.GetFlags()
	config, err := config.LoadConfig(*pathToConfig)
	if err != nil {
		log.Fatalf("Error loading config.json: %v\n", err)
	}
	log.Printf("Config loaded. DOWNLOADING THE FOLLOWING: %v\n", config)
	logger.AppendReport(fmt.Sprintf("------------------------------------------------------\n %v\n", time.Now()))
	logger.AppendReport(fmt.Sprintf("CONFIG LOADED:\n %v\n", config))

	// Clone required repositories
	reposToClone := [2]string{"systemed/tilemaker", "mapbox/tippecanoe"}
	for _, repo := range reposToClone {
		wg.Add(1)
		go func(repo string) {
			err := clonerepo.CloneRepo(repo)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(repo)
	}
	wg.Wait()
	log.Printf("%v successfully cloned.", reposToClone)

	// Build tilemaker inside docker container
	dockerTilemakerName := "sequential-tilemaker"
	execute.OutputToConsole(fmt.Sprintf("sudo docker build -t %v tilemaker", dockerTilemakerName))

	// Build tippecanoe inside docker container
	dockerTippecanoeName := "sequential-tippecanoe"
	execute.OutputToConsole(fmt.Sprintf("sudo docker build -t %v tippecanoe", dockerTippecanoeName))

	// Generate mbtiles for each subregion, downloading pbf files as necessary.
	for _, subRegion := range config.SubRegions {
		// Check to see if subregion already exists, and skip if so.
		splitSubRegion := strings.Split(subRegion, "/")
		finalDestination := pbfFolder + "/" + splitSubRegion[len(splitSubRegion)-1] + ".osm.pbf"
		if _, err := os.Stat(finalDestination); !os.IsNotExist(err) {
			log.Printf("Download: %v already exists. Skipping download.\n", subRegion)
		} else {
			// Download pbf file if required
			ok, err := downloadsubregion.DownloadSubRegion(subRegion, pbfFolder)
			if err != nil {
				log.Printf("There was a problem downloading %v. Moving on to the next one.", subRegion)
				continue
			}
			
			if ok {
				fmt.Printf("%v successfully downloaded.\n", subRegion)
			}
		}

		// Generate mbtiles inside docker container, if they don't already exist
		if _, err := os.Stat(fmt.Sprintf("mbtiles/%v.mbtiles", splitSubRegion[len(splitSubRegion)-1])); os.IsNotExist(err) {
		inputFile := splitSubRegion[len(splitSubRegion)-1] + ".osm.pbf"
		outputFile := splitSubRegion[len(splitSubRegion)-1] + ".mbtiles"
		generateMbtilesCmd := fmt.Sprintf("sudo docker run -v %v:/pbf -v %v:/mbtiles %v --input /pbf/%v --output /mbtiles/%v", pbfPath, mbtilesPath, dockerTilemakerName, inputFile, outputFile)
		execute.OutputToConsole(generateMbtilesCmd)

		} else {
			log.Printf("Mbtiles: mbtiles/%v.mbtiles already exists. Skipping generation.\n", splitSubRegion[len(splitSubRegion)-1])
		}
	}

	// Merge mbtiles into planet.mbtiles
	fi, _ := os.ReadDir("mbtiles")
	b := strings.Builder{}
	for _, f := range fi {
		b.WriteString("/data/" + f.Name() + " ")
	}

	mergeMbtilesCmd := fmt.Sprintf("sudo docker run --rm -v %v:/data %v tile-join --output=/data/merged/planet.mbtiles %v", mbtilesPath, dockerTippecanoeName, b.String())
	log.Println(mergeMbtilesCmd)
	execute.OutputToConsole((mergeMbtilesCmd))
}