package containers

import (
	"path/filepath"
	"testing"
)

func TestBuildContainer(t *testing.T) {
	var testCt = []container{
		{
			name:       "sequential-tilemaker",
			dockerfile: filepath.Clean("third_party/tilemaker/Dockerfile"),
			context:    filepath.Clean("third_party/tilemaker"),
		},
		{
			name:       "sequential-tippecanoe",
			dockerfile: filepath.Clean("third_party/tippecanoe/Dockerfile"),
			context:    filepath.Clean("third_party/tippecanoe"),
		},
		{
			name:       "sequential-osmium",
			dockerfile: filepath.Clean("build/osmium/Dockerfile"),
			context:    filepath.Clean("third_party"),
		},
		{
			name:       "sequential-gdal",
			dockerfile: filepath.Clean("third_party/gdal/docker/alpine-normal/Dockerfile"),
			context:    filepath.Clean("third_party/gdal"),
		},
	}

	for _, c := range testCt {
		err := buildContainer(c.name, c.dockerfile, c.context)
		if err != nil {
			t.Errorf("Failed to build container %v, with dockerfile %v, and context %v", c.name, c.dockerfile, c.context)
		}
	}
}

// Implement execute test
