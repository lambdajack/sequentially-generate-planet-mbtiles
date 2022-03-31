package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	SubRegions           []string
	KeepDownloadedFiles  bool
	KeepSubRegionMbtiles bool
	TileZoomLevel        int
}

func LoadConfig(pathToConfig string) (*Configuration, error) {
	var Config Configuration
	file, err := os.Open(pathToConfig)
	dec := json.NewDecoder(file)
	err = dec.Decode(&Config)
	if err != nil {
		return nil, err
	}
	return &Config, nil
}
