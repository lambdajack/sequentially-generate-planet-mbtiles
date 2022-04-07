package logger

import (
	"fmt"
	"log"
	"os"
)

func AppendReport(msg string) {
	reportFile := "REPORT.txt"
	f, err := os.Create(reportFile)
	if err != nil {
		log.Println("Cannot append REPORT.txt as it does not exist")
	}
	defer f.Close()

	if _, err := f.Write([]byte(fmt.Sprintf("%v\n", msg))); err != nil {
		log.Println(err)
	}
}
