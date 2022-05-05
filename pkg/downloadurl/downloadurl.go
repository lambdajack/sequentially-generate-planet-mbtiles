// Downloads a file from the given url and saves it to the given destination.

package downloadurl

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	DownloadAttemptFailed = errors.New("downloadurl.go | Download attempt failed.")
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

func DownloadURL(URL, destFileName, destFolder string) error {
	err := os.MkdirAll(destFolder, os.ModePerm)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create %v folder, in %v", destFolder, destFolder)
	}

	writeFile := filepath.Join(destFolder, "/", destFileName)

	// Setup file to write download to.
	f, err := os.Create(writeFile)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to create %s file, in %s", destFileName, destFolder)
	}

	// Set client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Set request
	r, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to set download request for %v", URL)
	}

	// Attempt download
	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to Download %v", URL)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return fmt.Errorf("bad status code when downloading %v - got %v", URL, resp.StatusCode)
	}

	// Initiate struct implementing writer for progress reporting.
	di := downloadInformation{
		destFileName:  destFileName,
		counter:       0,
		contentLength: resp.ContentLength,
	}

	// Write to file and log progress
	tee := io.TeeReader(resp.Body, &di)
	_, err = io.Copy(f, tee)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to save %s to %s", URL, destFileName)
	}

	return nil
}
