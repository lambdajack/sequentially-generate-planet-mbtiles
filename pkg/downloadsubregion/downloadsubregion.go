package downloadsubregion

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/logger"
	"github.com/lambdajack/sequentially-generate-planet-mbtiles/pkg/unzip"
)

type DownloadInformation struct {
	SubRegion       string
	ContentLength   int
	TotalDownloaded int
}

func (di *DownloadInformation) Write(p []byte) (n int, err error) {
	di.TotalDownloaded += len(p)
	if di.ContentLength != 0 {
		percentage := di.TotalDownloaded * 100 / di.ContentLength
		fmt.Printf("\rDownloaded %v of %v bytes (%v%%) of %v", di.TotalDownloaded, di.ContentLength, percentage, di.SubRegion)
	} else {
		fmt.Printf("\rDownloaded %v bytes of %v", di.TotalDownloaded, di.SubRegion)
	}
	return di.TotalDownloaded, nil
}

func DownloadSubRegion(subRegion, destFolder string) (ok bool, err error) {
	subRegionUrl := "https://download.geofabrik.de/" + subRegion + "-latest.osm.pbf"

	// HEAD request - find out the content-length (file size in bytes).
	r, err := http.NewRequest("HEAD", subRegionUrl, nil)
	if err != nil {
		return false, err
	}
	rH, err := http.DefaultClient.Do(r)
	if err != nil {
		return false, err
	}
	contentLength, err := strconv.Atoi(rH.Header.Get("Content-Length"))
	if err != nil {
		log.Printf("No content length was provided for the download, therefore progress cannot be displayed. The file will still be downloaded.\n")
	}
	fmt.Printf("Content Length: %T, %v\n", contentLength, contentLength)

	// Setup file to write download to.
	fileName := strings.Split(subRegion, "/")
	f, err := os.Create(destFolder + "/" + fileName[len(fileName)-1] + ".osm.pbf")
	if err != nil {
		log.Printf("Failed to create %v.osm.pbf file, in %v\n", subRegion, destFolder)
		return false, err
	}

	// GET request
	r.Method = "GET"
	rG, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer rG.Body.Close()

	// Initiate struct implementing writer for progress reporting.
	di := &DownloadInformation{
		SubRegion:       subRegion,
		ContentLength:   contentLength,
		TotalDownloaded: 0,
	}

	// Write to file and progress report.
	tee := io.TeeReader(rG.Body, di)
	_, err = io.Copy(f, tee)
	if err != nil {
		log.Printf("There was a problem writing to file: %v.osm.pbf\n", subRegion)
		return false, err
	}
	return true, nil
}

func DownloadOceanPoly(oceanPolyDest string) (ok bool, err error) {
	oceanPolyUrl := "https://osmdata.openstreetmap.de/download/water-polygons-split-4326.zip"

	// HEAD request - find out the content-length (file size in bytes).
	r, err := http.NewRequest("HEAD", oceanPolyUrl, nil)
	if err != nil {
		return false, err
	}
	rH, err := http.DefaultClient.Do(r)
	if err != nil {
		return false, err
	}
	contentLength, err := strconv.Atoi(rH.Header.Get("Content-Length"))
	if err != nil {
		log.Printf("No content length was provided for the download, therefore progress cannot be displayed. The file will still be downloaded.\n")
	}
	fmt.Printf("Content Length: %T, %v\n", contentLength, contentLength)

	// Setup file to write download to.
	oceanPolyFileName := "water-polygons-split-4326.zip"
	f, err := os.Create(oceanPolyFileName)
	if err != nil {
		log.Println("Failed to create ocean poly file... skipping.")
		logger.AppendReport("Failed to create ocean poly file... skipping. Ocean poly's will have to be added manually later if required.")
		return false, err
	}

	// GET request
	r.Method = "GET"
	rG, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer rG.Body.Close()

	// Initiate struct implementing writer for progress reporting.
	di := &DownloadInformation{
		SubRegion:       oceanPolyUrl,
		ContentLength:   contentLength,
		TotalDownloaded: 0,
	}

	// Write to file and progress report.
	tee := io.TeeReader(rG.Body, di)
	_, err = io.Copy(f, tee)
	if err != nil {
		log.Printf("Failed to create ocean poly file... skipping.\n")
		logger.AppendReport("Failed to create ocean poly file... skipping. Ocean poly's will have to be added manually later if required.")
		return false, err
	}

	// Unziping poly into coastline
	unzip.Unzip(oceanPolyFileName, oceanPolyDest)
	if err != nil {
		log.Printf("Failed to unzip ocean poly file... skipping.\n")
		logger.AppendReport("Failed to unzip ocean poly file... skipping. Ocean poly's will have to be added manually later if required.")
		return false, err
	}
	return true, nil
}
