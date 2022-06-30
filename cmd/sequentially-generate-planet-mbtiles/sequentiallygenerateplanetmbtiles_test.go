package sequentiallygenerateplanetmbtiles

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func init() {
	// Ensure that test flags match the ones in EntryPoint()

	flag.BoolVar(&fl.version, "v", false, "")
	flag.BoolVar(&fl.version, "version", false, "")

	flag.BoolVar(&fl.stage, "s", false, "")
	flag.BoolVar(&fl.stage, "stage", false, "")

	flag.StringVar(&fl.config, "c", "", "")
	flag.StringVar(&fl.config, "config", "", "")

	flag.BoolVar(&fl.test, "t", false, "")
	flag.BoolVar(&fl.test, "test", false, "")

	flag.StringVar(&fl.pbfFile, "p", "", "")
	flag.StringVar(&fl.pbfFile, "pbf-file", "", "")

	flag.StringVar(&fl.workingDir, "w", "data", "")
	flag.StringVar(&fl.workingDir, "working-dir", "data", "")

	flag.StringVar(&fl.outDir, "o", "data/out", "")
	flag.StringVar(&fl.outDir, "outdir", "data/out", "")

	flag.BoolVar(&fl.excludeOcean, "eo", false, "")
	flag.BoolVar(&fl.excludeOcean, "exclude-ocean", false, "")

	flag.BoolVar(&fl.excludeLanduse, "el", false, "")
	flag.BoolVar(&fl.excludeLanduse, "exclude-landuse", false, "")

	flag.StringVar(&fl.tilemakerConfig, "tc", "", "")
	flag.StringVar(&fl.tilemakerConfig, "tilemaker-config", "", "")

	flag.StringVar(&fl.tilemakerProcess, "tp", "", "")
	flag.StringVar(&fl.tilemakerProcess, "tilemaker-process", "", "")

	flag.Uint64Var(&fl.maxRamMb, "r", 0, "")
	flag.Uint64Var(&fl.maxRamMb, "ram", 0, "")

	flag.BoolVar(&fl.outAsDir, "od", false, "")
	flag.BoolVar(&fl.outAsDir, "out-as-dir", false, "")

	flag.BoolVar(&fl.skipSlicing, "ss", false, "")
	flag.BoolVar(&fl.skipSlicing, "skip-slicing", false, "")

	flag.BoolVar(&fl.mergeOnly, "mo", false, "")
	flag.BoolVar(&fl.mergeOnly, "merge-only", false, "")

	flag.BoolVar(&fl.skipDownload, "sd", false, "")
	flag.BoolVar(&fl.skipDownload, "skip-download", false, "")

}

func TestEntryPoint(t *testing.T) {
	fmt.Println("IMPLEMENT ENTRY POINT TEST")
}

func TestValidateFlags(t *testing.T) {

	if os.Getenv("TEST_VALIDATE_FLAGS") == "1" {
		os.Args = append(os.Args, "-c", "/into/the/unknown", "-eo")
		flag.Parse()
		validateFlags()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestValidateFlags")
	cmd.Env = append(os.Environ(), "TEST_VALIDATE_FLAGS=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		if e.ExitCode() == exitFlags {
			return
		}
	}

	t.Fatalf("process ran with err %v, want exit status %v", err, exitFlags)
}
