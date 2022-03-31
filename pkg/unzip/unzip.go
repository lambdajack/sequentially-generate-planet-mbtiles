package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Unzip(zipFileName, unzipDest string) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Setup destination dir
	if _, err := os.Stat(unzipDest); os.IsNotExist(err) {
		os.Mkdir(unzipDest, 0755)
	}

	// Iterate through the files in the archive,
	for _, f := range r.File {
		splitName := strings.Split(f.Name, "/")
		fmt.Printf("Contents of %s:\n", splitName[len(splitName)-1])
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(unzipDest + "/" + splitName[len(splitName)-1])
		if err != nil {
			log.Printf("Failed to create files for unzipping: %v\n", err)
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			log.Printf("Failed to populate unzipped files with data: %v\n", err)
		}
	}
}
