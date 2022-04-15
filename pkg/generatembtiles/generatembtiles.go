package generatembtiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func GenerateMbTiles(inputFile, outputFile, pbfFolder, mbtilesFolder, coastlineFolder, containerName, configFile, processFile string) error {

	splitConfigPath := strings.Split(configFile, string(os.PathSeparator))
	splitProcessPath := strings.Split(processFile, string(os.PathSeparator))

	configPath := strings.Join(splitConfigPath[:len(splitConfigPath)-1], string(os.PathSeparator))
	processPath := strings.Join(splitProcessPath[:len(splitProcessPath)-1], string(os.PathSeparator))

	configFile = filepath.Base(configFile)
	processFile = filepath.Base(processFile)

	var generateMbtilesCmd string

	if processPath == configPath {
		generateMbtilesCmd = fmt.Sprintf("docker run --rm -v %s:/pbf -v %s:/mbtiles -v %s:/coastline -v %v:/config %s --input /pbf/%s --output /mbtiles/%s --config /config/%s --process /config/%s", pbfFolder, mbtilesFolder, coastlineFolder, configPath, containerName, inputFile, outputFile, configFile, processFile)
	} else {
		generateMbtilesCmd = fmt.Sprintf("docker run --rm -v %s:/pbf -v %s:/mbtiles -v %s:/coastline -v %v:/config -v %v:/process %s --input /pbf/%s --output /mbtiles/%s --config /config/%s --process /process/%s", pbfFolder, mbtilesFolder, coastlineFolder, configPath, processPath, containerName, inputFile, outputFile, configFile, processFile)
	}

	err := execute.OutputToConsole(generateMbtilesCmd)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("generatembtiles.go | Failed to generate mbtiles for %s", inputFile), err)
	}
	return nil
}
