package logger

import (
	"fmt"
	"log"
	"os"
)

func AppendReport(msg string) {
	// Ensure REPORT.txt exists
	reportFile := "mbtiles/merged/REPORT.txt"
	f, err := os.OpenFile(reportFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Cannot append REPORT.txt as it does not exist")
	}
	if _, err := f.Write([]byte(fmt.Sprintf("%v\n", msg))); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
