package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

func Unzip(zipFilePath, unzipDestFolder string) error {
	zipFilePath = filepath.FromSlash(filepath.FromSlash(zipFilePath))
	unzipDestFolder = filepath.FromSlash(filepath.Clean(unzipDestFolder))
	
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("unzip.go | Failed to open zip file %v", zipFilePath), err)
	}
	defer r.Close()

	// Setup destination dir
	if _, err := os.Stat(unzipDestFolder); os.IsNotExist(err) {
		os.Mkdir(unzipDestFolder, 0755)
		if err != nil {
			return stderrorhandler.StdErrorHandler(fmt.Sprintf("unzip.go | Failed to create destination folder %v", unzipDestFolder), err)
		}
	}
	

	// Iterate through the files in the archive,
	for _, f := range r.File {
		splitName := strings.Split(f.Name, string(os.PathSeparator))
		fmt.Printf("Contents of %s:\n", splitName[len(splitName)-1])
		rc, err := f.Open()
		if err != nil {
			return stderrorhandler.StdErrorHandler(fmt.Sprintf("unzip.go | Failed to open file %v", f.Name), err)
		}
		f, err := os.Create(filepath.FromSlash(unzipDestFolder + "/" + splitName[len(splitName)-1]))
		if err != nil {
			return stderrorhandler.StdErrorHandler(fmt.Sprintf("unzip.go | Failed to open file %v", f.Name()), err)
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			return stderrorhandler.StdErrorHandler(fmt.Sprintf("unzip.go | Failed to copy data into new file %v", f.Name()), err)
		}
	}

	return nil
}
