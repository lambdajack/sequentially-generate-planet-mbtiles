package describeloggers

import (
	"bufio"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestErr(t *testing.T) {
	tmpDir := t.TempDir()

	f, err := os.Create(tmpDir + "/err.log")
	if err != nil {
		t.Error("Failed to create tmp files for test")
	}

	writer := io.Writer(f)

	progLog := Prog(&writer)

	if reflect.TypeOf(progLog).String() != "*log.Logger" {
		t.Errorf("got %v, want *log.Logger", reflect.TypeOf(progLog))
	}

	progLog.Println("TEST")

	f.Close()

	f, err = os.Open(tmpDir + "/err.log")
	if err != nil {
		t.Error("Failed to open log file for reading")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.HasSuffix(scanner.Text(), "TEST") {
			break
		}
		t.Errorf("got %v, want suffix TEST", scanner.Text())
	}
}
