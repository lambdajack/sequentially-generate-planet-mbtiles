// Downloads a file from the given url and saves it to the given destination.

package downloadurl

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/stderrorhandler"
)

type downloadInformation struct {
	destFileName           string
	counter, contentLength int64
}

func (d *downloadInformation) Write(p []byte) (n int, err error) {
	d.counter += int64(len(p))
	if d.contentLength != 0 {
		percentage := d.counter * 100 / d.contentLength
		fmt.Printf("\rDownloading %v of %v bytes (%v%%) of %v", d.counter, d.contentLength, percentage, d.destFileName)
	} else {
		fmt.Printf("\rDownloading %v bytes of %v", d.counter, d.destFileName)
	}
	return int(d.counter), nil
}

func DownloadURL(url, destFileName, destFolder string) error {
	slash := string(os.PathSeparator)

	err := os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("Failed to create %v folder, in %v", destFolder, destFolder), err)
	}

	// Setup file to write download to.
	f, err := os.Create(destFolder + slash + destFileName)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("downloadurl.go | Failed to create %s file, in %s\n", destFileName, destFolder), err)
	}

	// GET request
	r, err := http.Get(url)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("downloadurl.go | Failed to get %s\n", url), err)
	}
	defer r.Body.Close()

	// Initiate struct implementing writer for progress reporting.
	di := downloadInformation{
		destFileName:  destFileName,
		counter:       0,
		contentLength: r.ContentLength,
	}

	// Write to file and log progress
	tee := io.TeeReader(r.Body, &di)
	_, err = io.Copy(f, tee)
	if err != nil {
		return stderrorhandler.StdErrorHandler(fmt.Sprintf("downloadurl.go | Failed to save %s to %s\n", url, destFileName), err)
	}

	return nil
}
