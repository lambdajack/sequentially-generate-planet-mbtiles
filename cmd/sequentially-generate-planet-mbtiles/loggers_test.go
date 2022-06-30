package sequentiallygenerateplanetmbtiles

import (
	"bufio"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestInitLoggers(t *testing.T) {
	lg := &loggers{}
	initLoggers()

	if reflect.TypeOf(lg.prog).String() != "*log.Logger" {
		t.Errorf("got %v, want *log.Logger", reflect.TypeOf(lg.prog))
	}

	if reflect.TypeOf(lg.err).String() != "*log.Logger" {
		t.Errorf("got %v, want *log.Logger", reflect.TypeOf(lg.err))
	}

	if reflect.TypeOf(lg.rep).String() != "*log.Logger" {
		t.Errorf("got %v, want *log.Logger", reflect.TypeOf(lg.rep))
	}
}

func TestInitProg(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Clean(tmpDir + "/prog.log")

	lg := initProg(logFile)

	lg.Println("PROGTEST")

	f, err := os.Open(logFile)
	if err != nil {
		t.Error("Failed to open log file for reading")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.HasSuffix(scanner.Text(), "PROGTEST") {
			break
		}
		t.Errorf("got %v, want suffix PROGTEST", scanner.Text())
	}

}

func TestInitErr(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Clean(tmpDir + "/prog.log")

	lg := initErr(logFile)

	lg.Println("ERRTEST")

	f, err := os.Open(logFile)
	if err != nil {
		t.Error("Failed to open log file for reading")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.HasSuffix(scanner.Text(), "ERRTEST") {
			break
		}
		t.Errorf("got %v, want suffix ERRTEST", scanner.Text())
	}

}

func TestInitRep(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Clean(tmpDir + "/prog.log")

	lg := initRep(logFile)

	lg.Println("REPTEST")

	f, err := os.Open(logFile)
	if err != nil {
		t.Error("Failed to open log file for reading")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.HasSuffix(scanner.Text(), "REPTEST") {
			break
		}
		t.Errorf("got %v, want suffix REPTEST", scanner.Text())
	}

}
