package extract

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/containers"
)

func TreeSlicer(src, dstDir, workingDir string, targetSize uint64) {
	log.Printf("Operating on: %s", src)

	src, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}
	dstDir, err = filepath.Abs(dstDir)
	if err != nil {
		log.Fatal(err)
	}
	workingDir, err = filepath.Abs(workingDir)
	if err != nil {
		log.Fatal(err)
	}

	minX, minY, maxX, maxY, err := getExtent(src, containers.ContainerNames.Gdal)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to get extent: %v", err)
	}

	lsp := leftSlicePoint(minX, maxX)
	lbb := formatBoundingBox(minX, minY, lsp, maxY)
	lp := slice(src, workingDir, lbb)

	rsp := rightSlicePoint(minX, maxX)
	rbb := formatBoundingBox(rsp, minY, maxX, maxY)
	rp := slice(src, workingDir, rbb)

	if strings.Contains(filepath.Base(src), "-tmp") {
		os.Remove(src)
	} else {
		log.Printf("Sparing %s's life", filepath.Base(src))
	}

	if size(lp, targetSize) {
		os.Rename(lp, filepath.Join(dstDir, filepath.Base(lp)))
	} else {
		TreeSlicer(lp, dstDir, workingDir, targetSize)
	}

	if size(rp, targetSize) {
		os.Rename(rp, filepath.Join(dstDir, filepath.Base(rp)))
	} else {
		TreeSlicer(rp, dstDir, workingDir, targetSize)
	}
}

func size(src string, targetMb uint64) bool {
	f, err := os.Stat(src)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to get file info: %v", err)
	}

	if f.Size() > int64(targetMb*1024*1024) {
		log.Printf("Target %s requires further slicing", filepath.Base(src))
		return false
	}

	log.Printf("Slice %s has reached target size. Moving to safety.", filepath.Base(src))
	return true
}

func slice(src, dst, bb string) string {
	f, err := os.CreateTemp(dst, "*-tmp.osm.pbf")
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to create temp file: %v", err)
	}
	defer f.Close()

	log.Printf("Slicing: %s >>> %s (%s)", filepath.Base(src), filepath.Base(f.Name()), bb)
	lp, err := Extract(src, f.Name(), bb, containers.ContainerNames.Osmium)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to extract left slice: %v", err)
	}

	return lp
}

func formatBoundingBox(minX, minY, maxX, maxY float64) string {
	return fmt.Sprintf("%f,%f,%f,%f", minX, minY, maxX, maxY)
}

func getExtent(filePath, ogrContainerName string) (minX, minY, maxX, maxY float64, err error) {

	ap, err := filepath.Abs(filePath)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	cmd := exec.Command("docker", "run", "--rm", "--mount", "type=bind,source="+ap+",target=/data", ogrContainerName, "ogrinfo", "-so", "-al", "/data")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error here?")
		return 0, 0, 0, 0, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	var extent string

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Extent") {
			extent = scanner.Text()
			break
		}
	}

	if extent == "" {
		fmt.Println("Extent not found")
		return 0, 0, 0, 0, fmt.Errorf("extent \"Extent\" not found")
	}

	re := regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+")
	coords := re.FindAllString(extent, -1)

	minX, _ = strconv.ParseFloat(coords[0], 64)
	minY, _ = strconv.ParseFloat(coords[1], 64)
	maxX, _ = strconv.ParseFloat(coords[2], 64)
	maxY, _ = strconv.ParseFloat(coords[3], 64)

	return minX, minY, maxX, maxY, nil
}

// returns the mid point for the box (which should be used to generate the next slice)
func rightSlicePoint(minX, maxX float64) float64 {
	slicePoint := (minX + maxX) / 2
	// 0.01 is taken away to ensure no data is lost during the slicing process
	return slicePoint - 0.01
}

func leftSlicePoint(minX, maxX float64) float64 {
	slicePoint := (minX + maxX) / 2
	// 0.01 is added to ensure no data is lost during the slicing process
	return slicePoint + 0.01
}
