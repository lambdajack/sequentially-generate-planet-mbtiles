package generatembtiles

import (
	"fmt"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func GenerateMbTiles(inputFile, outputFile, pbfFolder, mbtilesFolder, coastlineFolder, containerName, configFolder, configFile, processFolder, processFile string) error {

	generateMbtilesCmd := fmt.Sprintf("docker run -v %s:/pbf -v %s:/mbtiles -v %s:/coastline -v %s:/config -v %s:/process %s --input /pbf/%s --output /mbtiles/%s --config /config/%s --process /process/%s", pbfFolder, mbtilesFolder, coastlineFolder, configFolder, processFolder, containerName, inputFile, outputFile, configFile, processFile)

	err := execute.OutputToConsole(generateMbtilesCmd)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("generatembtiles.go | Failed to generate mbtiles for %s", inputFile), err)
	}
	return nil
}
