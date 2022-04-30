package extractiontree

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Slicer(filePath string) {
	minX, minY, maxX, maxY, _ := getExtent(filePath)

	slicePoint, _ := getSlicePoint(minX, maxX)

	bb := formatBoundingBox(minX, minY, slicePoint, maxY)
	fmt.Println(bb)

	// Run the extraction
	// ...

	// check the two new files created for size
	// ...

	// call Slicer again with the new files if necessary
	// ...

}

func formatBoundingBox(minX, minY, maxX, maxY float64) string {
	return fmt.Sprintf("%f,%f,%f,%f", minX, minY, maxX, maxY)
}

func getExtent(filePath string) (float64, float64, float64, float64, error) {

	command := "ogrinfo -so -al " + filePath
	args := strings.Split(command, " ")
	out, _ := exec.Command(args[0], args[1:]...).Output()

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

	minX, _ := strconv.ParseFloat(coords[0], 64)
	minY, _ := strconv.ParseFloat(coords[1], 64)
	maxX, _ := strconv.ParseFloat(coords[2], 64)
	maxY, _ := strconv.ParseFloat(coords[3], 64)

	fmt.Println(minX)
	fmt.Println(minY)
	fmt.Println(maxX)
	fmt.Println(maxY)

	return minX, minY, maxX, maxY, nil
}

func getSlicePoint(minX, maxX float64) (float64, error) {
	slicePoint := (minX + maxX) / 2
	// 0.01 is added to ensure no data is lost during the slicing process
	return slicePoint + 0.01, nil
}
