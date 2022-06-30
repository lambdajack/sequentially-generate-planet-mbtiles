package describeloggers

import (
	"io"
	"log"
)

func Prog(writer *io.Writer) *log.Logger {
	progLog := log.New(*writer, "", log.Ldate|log.Ltime)
	return progLog
}
