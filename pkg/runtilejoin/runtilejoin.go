package runtilejoin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/execute"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func RunTileJoin(mbtilesPath, mergedPath, containerName string) error {
	fi, err := os.ReadDir(mbtilesPath)
	if err != nil {
		stderrorhandler.StdErrorHandler(fmt.Sprintf("runtilejoin.go | Failed to read directory %s. Unable to proceed", mbtilesPath), err)
		panic(err)
	}

	b := strings.Builder{}
	for _, f := range fi {
		if f.IsDir() == false {
			b.WriteString("/data/" + f.Name() + " ")
		}
	}

	// check to see if planet.mbtiles exists
	if _, err := os.Stat(filepath.FromSlash(mergedPath + "/" + "planet.mbtiles")); os.IsNotExist(err) {
		f, err := os.Create(filepath.FromSlash(mergedPath + "/" + "planet.mbtiles"))
		if err != nil {
			stderrorhandler.StdErrorHandler(fmt.Sprintf("runtilejoin.go | Failed to create file %s. It may already exist.. tile-join process will attempt to proceed, but will likely fail.", mergedPath), err)
		}
		f.Close()
	}

	mergeMbtilesCmd := fmt.Sprintf("sudo docker run --rm -v %v:/data -v %v:/merged %v tile-join --output=/merged/planet.mbtiles %v", mbtilesPath, mergedPath, containerName, b.String())
	log.Println(mergeMbtilesCmd)
	err = execute.OutputToConsole((mergeMbtilesCmd))
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprint("runtilejoin.go | Failed to merge mbtiles"), err)
	}

	return nil
}
