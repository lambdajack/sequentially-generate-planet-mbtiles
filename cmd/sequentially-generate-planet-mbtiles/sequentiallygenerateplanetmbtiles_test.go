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

	flag.StringVar(&fl.planetFile, "p", "", "")
	flag.StringVar(&fl.planetFile, "planet-file", "", "")

	flag.StringVar(&fl.dataDir, "d", "data", "")
	flag.StringVar(&fl.dataDir, "datadir", "data", "")

	flag.StringVar(&fl.outDir, "o", "data/out", "")
	flag.StringVar(&fl.outDir, "outdir", "data/out", "")

	flag.BoolVar(&fl.includeOcean, "io", true, "")
	flag.BoolVar(&fl.includeOcean, "include-ocean", true, "")

	flag.BoolVar(&fl.includeLanduse, "il", true, "")
	flag.BoolVar(&fl.includeLanduse, "include-landuse", true, "")

	flag.StringVar(&fl.tilemakerConfig, "tc", "", "")
	flag.StringVar(&fl.tilemakerConfig, "tilemaker-config", "", "")

	flag.StringVar(&fl.tilemakerProcess, "tp", "", "")
	flag.StringVar(&fl.tilemakerProcess, "tilemaker-process", "", "")

	flag.Uint64Var(&fl.maxRamMb, "r", 0, "")
	flag.Uint64Var(&fl.maxRamMb, "ram", 0, "")

}

func TestEntryPoint(t *testing.T) {
	fmt.Println("IMPLEMENT ENTRY POINT TEST")
}

func TestValidateFlags(t *testing.T) {

	if os.Getenv("TEST_VALIDATE_FLAGS") == "1" {
		os.Args = append(os.Args, "-c", "/into/the/unknown", "-io")
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

func TestGetEmbeddedFiles(t *testing.T) {
	// Function mirrored in main_test.go.
	// The test relies on a non-global embed.FS variable.
	// While the embed is not needed in the main function,
	// golang only allows embeded files to be within scope of the package.
	// The embed should therefore be in the main package as it is not worth
	// making a mess of the source code by moving the relevant embedded
	// files to within scope.

	// tl:dr - a mirror of this function is tested in main_test.go.
	// It is the developers responsibility to ensure that the mirror is up to date.
}
