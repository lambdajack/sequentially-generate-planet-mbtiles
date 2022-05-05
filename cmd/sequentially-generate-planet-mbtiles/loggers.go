package sequentiallygenerateplanetmbtiles

import (
	"io"
	"log"
	"os"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/describeloggers"
)

type loggers struct {
	prog *log.Logger
	err  *log.Logger
	rep  *log.Logger
}

var lg = &loggers{}

func initLoggers() {
	lg.prog = initProg("./data/logs/prog.log")
	lg.err = initErr("./data/logs/err.log")
	lg.rep = initRep("./data/logs/rep.log")
}

func initProg(logPath string) *log.Logger {
	lf, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	writer := io.MultiWriter(os.Stdout, lf)
	lggr := describeloggers.Prog(&writer)
	return lggr
}

func initErr(logPath string) *log.Logger {
	lf, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	writer := io.MultiWriter(os.Stderr, lf)
	lggr := describeloggers.Err(&writer)
	return lggr
}

func initRep(logPath string) *log.Logger {
	lf, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	writer := io.MultiWriter(os.Stdout, lf)
	lggr := describeloggers.Rep(&writer)
	return lggr
}
