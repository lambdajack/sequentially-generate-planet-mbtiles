package extract

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Slicer(src, dstFolder, ogrContainerName string, targetSizeMb int64) {
	log.Printf("Slicing %s", src)
	
	src = filepath.Clean(src)
	dstFolder = filepath.Clean(dstFolder)

	minX, minY, maxX, maxY, _ := getExtent(src, ogrContainerName)

	// Left slice point
	lsp, err := leftSlicePoint(minX, maxX)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to generate left slice point: %v", err)
	}
	lbb := formatBoundingBox(minX, minY, lsp, maxY)

	// Extract left slice
	f, err := ioutil.TempFile(filepath.Dir(src), "*.osm.pbf")
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to create temp file: %v", err)
	}

	lp, err := Extract(src, filepath.Join(filepath.Dir(src), f.Name()), lbb, ogrContainerName)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to extract left slice: %v", err)
	}

	// Check left file size
	lf, err := os.Stat(lp)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to get file info: %v", err)
	}
	if lf.Size() > targetSizeMb*1024*1024 {
		Slicer(lp, dstFolder, ogrContainerName, targetSizeMb)
	} else {
		log.Printf("Slice %s has reached target size. Saving to working folder", lp)
		os.Rename(lp, filepath.Join(dstFolder, filepath.Base(lp)))
	}

	// Right slice point
	rsp, err := rightSlicePoint(minX, maxX)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to generate right slice point: %v", err)
	}
	rbb := formatBoundingBox(rsp, minY, maxX, maxY)

	// Extract right slice
	f, err = ioutil.TempFile(filepath.Dir(src), "*.osm.pbf")
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to create temp file: %v", err)
	}

	rp, err := Extract(src, filepath.Join(filepath.Dir(src), f.Name()), rbb, ogrContainerName)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to extract right slice: %v", err)
	}

	// Check right file size
	rf, err := os.Stat(rp)
	if err != nil {
		log.Fatalf("extract.go | Slicer | Failed to get file info: %v", err)
	}

	if rf.Size() > targetSizeMb*1024*1024 {
		Slicer(rp, dstFolder, ogrContainerName, targetSizeMb)
	} else {
		log.Printf("Slice %s has reached target size. Saving to working folder", rp)
		os.Rename(rp, filepath.Join(dstFolder, filepath.Base(rp)))
	}
}

func formatBoundingBox(minX, minY, maxX, maxY float64) string {
	return fmt.Sprintf("%f,%f,%f,%f", minX, minY, maxX, maxY)
}

func getExtent(filePath, ogrContainerName string) (minX, minY, maxX, maxY float64, err error) {

	cmd := exec.Command("sudo", "docker", "run", "--rm", "--mount", "type=bind,source="+filePath+",target=/data", "containerName", "ogrinfo", "-so", "-al", "/data")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
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

	fmt.Println(minX)
	fmt.Println(minY)
	fmt.Println(maxX)
	fmt.Println(maxY)

	return minX, minY, maxX, maxY, nil
}

// returns the mid point for the box (which should be used to generate the next slice)
func rightSlicePoint(minX, maxX float64) (float64, error) {
	slicePoint := (minX + maxX) / 2
	// 0.01 is taken away to ensure no data is lost during the slicing process
	return slicePoint - 0.01, nil
}

func leftSlicePoint(minX, maxX float64) (float64, error) {
	slicePoint := (minX + maxX) / 2
	// 0.01 is added to ensure no data is lost during the slicing process
	return slicePoint + 0.01, nil
}
