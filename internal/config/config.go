package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	DataDir                string `json:"dataDir"`
	PathToTilemakerConfig  string `json:"pathToTilemakerConfig"`
	PathToTilemakerProcess string `json:"pathToTilemakerProcess"`
}

var Config Configuration

func LoadConfig(pathToConfig string) (*Configuration, error) {
	file, err := os.Open(pathToConfig)
	dec := json.NewDecoder(file)
	err = dec.Decode(&Config)
	if err != nil {
		log.Fatal("Unable to read config file - config file may be invalid. Unable to proceed...")
	}
	return &Config, nil
}
