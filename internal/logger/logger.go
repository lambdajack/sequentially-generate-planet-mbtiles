package logger

import (
	"log"
	"os"
)

func AppendReport(msg string) {
	reportFile := "REPORT.txt"

	f, err := os.OpenFile(reportFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}

	if _, err = f.WriteString(msg); err != nil {
		log.Println(err)
	}
}
