package describeloggers

import (
	"io"
	"log"
)

func Err(writer *io.Writer) *log.Logger {
	errLog := log.New(*writer, "", log.Ldate|log.Ltime|log.Lshortfile)
	return errLog
}
