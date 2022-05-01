package main

import (
	"testing"
)

func TestEmbeddedFiles(t *testing.T) {
	// This may be unnecessary since the tests won't run if the files aren't embedded.
	// However, for the sake of completeness, I'm leaving it in.

	_, err := embeddedFs.ReadFile("third_party/tilemaker/resources/config-openmaptiles.json")
	if err != nil {
		t.Fatalf("Failed to retrieve tilemaker config file from embedded filesystem: %s", err)
	}

	_, err = embeddedFs.ReadFile("third_party/tilemaker/resources/process-openmaptiles.lua")
	if err != nil {
		t.Fatalf("Failed to retrieve tilemaker process file from embedded filesystem: %s", err)
	}
}
