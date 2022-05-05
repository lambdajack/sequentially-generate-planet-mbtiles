package containers

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestBuildContainer(t *testing.T) {
	var testCt = &containers{
		tilemaker: container{
			name:       "sequential-tilemaker",
			dockerfile: filepath.Clean("third_party/tilemaker/Dockerfile"),
			context:    filepath.Clean("third_party/tilemaker"),
		},
		tippecanoe: container{
			name:       "sequential-tippecanoe",
			dockerfile: filepath.Clean("third_party/tippecanoe/Dockerfile"),
			context:    filepath.Clean("third_party/tippecanoe"),
		},
		osmium: container{
			name:       "sequential-osmium",
			dockerfile: filepath.Clean("build/osmium/Dockerfile"),
			context:    filepath.Clean("third_party"),
		},
	}

	if !reflect.DeepEqual(ct, testCt) {
		t.Errorf("\nEXPECTED: %v\n GOT: %v\n", testCt, ct)
	}
}

// Implement execute test
