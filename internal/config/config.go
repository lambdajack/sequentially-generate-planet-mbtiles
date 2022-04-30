package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	DataDir          string `json:"dataDir"`
	TilemakerConfig  string `json:"TilemakerConfig"`
	TilemakerProcess string `json:"TilemakerProcess"`
	AvailableRam     int    `json:"availableRam"`
}

var Config Configuration

func LoadConfig(pathToConfig string) (*Configuration, error) {
	file, err := os.Open(filepath.Clean(pathToConfig))
	if err != nil {
		log.Fatal("Unable to read config file - Supply a config.json file using the '-c' flag")
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&Config)
	if err != nil {
		log.Fatal("Unable to decode config file - it may be invalid. Supply a config.json file using the '-c' flag")
	}
	return &Config, nil
}
