package osmboundaryextract

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func OsmBoundaryExtract(srcPbfFilePath, srcFileName, dstFilePath, dstFileName, bbox, containerName string) error {
	finalPath := filepath.Clean(dstFilePath + "/" + dstFileName)

	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		cmdString := fmt.Sprintf("docker run --rm -v %s:/pbf -v %s:/out %s osmium extract -b %s --set-bounds /pbf/%s -o /out/%s", srcPbfFilePath, dstFilePath, containerName, bbox, srcFileName, dstFileName)

		err := execute.OutputToConsole(cmdString)
		if err != nil {
			return stderrorhandler.StdErrorHandler(fmt.Sprintf("osmboundaryextract.go | Failed to extract %s from %s", bbox, srcFileName), err)
		}
	} else {
		log.Printf("%s already exists. Skipping...", finalPath)
	}

	return nil
}
