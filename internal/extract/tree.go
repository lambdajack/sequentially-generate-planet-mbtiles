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

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/docker"
)

func TreeSlicer(src, dstDir, workingDir string, targetSize uint64, gdal, osmium *docker.Container, elg, plg, rlg *log.Logger) {
	log.Printf("operating on: %s", src)

	src, err := filepath.Abs(src)
	if err != nil {
		elg.Fatal(err)
	}
	dstDir, err = filepath.Abs(dstDir)
	if err != nil {
		elg.Fatal(err)
	}
	workingDir, err = filepath.Abs(workingDir)
	if err != nil {
		elg.Fatal(err)
	}

	minX, minY, maxX, maxY, err := getExtent(src, gdal.Name)
	if err != nil {
		elg.Fatalf("extract.go | Slicer | Failed to get extent: %v", err)
	}

	lsp := leftSlicePoint(minX, maxX)
	lbb := formatBoundingBox(minX, minY, lsp, maxY)
	lp := slice(src, workingDir, lbb, osmium, elg, plg, rlg)

	rsp := rightSlicePoint(minX, maxX)
	rbb := formatBoundingBox(rsp, minY, maxX, maxY)
	rp := slice(src, workingDir, rbb, osmium, elg, plg, rlg)

	if strings.Contains(filepath.Base(src), "-tmp") {
		os.Remove(src)
	} else {
		log.Printf("sparing %s's life", filepath.Base(src))
	}

	if size(lp, targetSize, elg, plg) {
		os.Rename(lp, filepath.Join(dstDir, filepath.Base(lp)))
	} else {
		TreeSlicer(lp, dstDir, workingDir, targetSize, gdal, osmium, elg, plg, rlg)
	}

	if size(rp, targetSize, elg, plg) {
		os.Rename(rp, filepath.Join(dstDir, filepath.Base(rp)))
	} else {
		TreeSlicer(rp, dstDir, workingDir, targetSize, gdal, osmium, elg, plg, rlg)
	}
}

func size(src string, targetMb uint64, elg, plg *log.Logger) bool {
	f, err := os.Stat(src)
	if err != nil {
		elg.Fatalf("slicer failed to get file info: %v\n", err)
	}

	if f.Size() > int64(targetMb*1024*1024) {
		plg.Printf("SLICE: target %s requires further slicing", filepath.Base(src))
		return false
	}

	plg.Printf("SLICE: %s has reached target size; moving to safety.", filepath.Base(src))
	return true
}

func IncompleteProgress(originalSrc, progressDir string, gdal *docker.Container, elg, rlg *log.Logger) string {
	minX, minY, maxX, maxY, err := getExtent(originalSrc, gdal.Name)
	if err != nil {
		elg.Println("failed to get extent for original source; cannot utilise previous progress")
		return ""
	}

	err = filepath.Walk(progressDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, _, mx, _, err := getExtent(path, gdal.Name)
			if err != nil {
				elg.Printf("failed to get extent for %s source; cannot utilise previous progress\n", path)
				return err
			}
			if mx > minX {
				minX = mx
			}
		}
		return nil
	})
	if err != nil {
		return ""
	}

	rlg.Printf("previously incomplete: %f, %f, %f, %f\n", minX, minY, maxX, maxY)

	return formatBoundingBox(minX, minY, maxX, maxY)
}

func slice(src, dst, bb string, osmium *docker.Container, elg, plg, rlg *log.Logger) string {
	f, err := os.CreateTemp(dst, "*-tmp.osm.pbf")
	if err != nil {
		elg.Fatalf("extract.go | Slicer | Failed to create temp file: %v", err)
	}
	defer f.Close()

	rlg.Printf("slicing %s >>> %s (%s)", filepath.Base(src), filepath.Base(f.Name()), bb)
	lp, err := Extract(src, f.Name(), bb, osmium)
	if err != nil {
		elg.Fatalf("extract.go | Slicer | Failed to extract slice: %v", err)
	}
	plg.Printf("SLICE: %s >>> %s (%s) success", filepath.Base(src), filepath.Base(f.Name()), bb)

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
		log.Println("extent not found")
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
