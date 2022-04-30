package describeloggers

import (
	"io"
	"log"
)

func Rep(writer *io.Writer) *log.Logger {
	repLog := log.New(*writer, "", log.Ldate|log.Ltime)
	return repLog
}
