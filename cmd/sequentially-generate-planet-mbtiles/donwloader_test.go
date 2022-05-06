package sequentiallygenerateplanetmbtiles

import (
	"reflect"
	"testing"
)

func TestDownloadOsmData(t *testing.T) {

	testDl := map[string]downloadInformation{
		"osmPlanetPbf": {
			url:          "https://planet.openstreetmap.org/pbf/planet-latest.osm.pbf",
			destFileName: "planet-latest.osm.pbf",
			destFolder:   pth.pbfFolder,
		},
		"waterPolygons": {
			url:          "https://osmdata.openstreetmap.de/download/water-polygons-split-4326.zip",
			destFileName: "water-polygons-split-4326.zip",
			destFolder:   pth.dataFolder,
		},
		"landCoverUrban": {
			url:          "https://naciscdn.org/naturalearth/10m/cultural/ne_10m_urban_areas.zip",
			destFileName: "ne_10m_urban_areas.zip",
			destFolder:   pth.dataFolder,
		},
		"landCoverIceShelves": {
			url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_antarctic_ice_shelves_polys.zip",
			destFileName: "ne_10m_antarctic_ice_shelves_polys.zip",
			destFolder:   pth.dataFolder,
		},
		"landCoverGlaciated": {
			url:          "https://naciscdn.org/naturalearth/10m/physical/ne_10m_glaciated_areas.zip",
			destFileName: "ne_10m_glaciated_areas.zip",
			destFolder:   pth.dataFolder,
		},
	}

	if !reflect.DeepEqual(downloads, testDl) {
		t.Errorf("\nEXPECTED: %v\n GOT: %v\n", testDl, downloads)
	}

	// Set download data for rest of test - geofabrik?
}