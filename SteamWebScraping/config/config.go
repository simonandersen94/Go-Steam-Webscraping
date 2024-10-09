package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ScrapeUrl     string `json:"scrape_url"`
	AllowedDomain string `json:"allowedDomain"`
	DatabaseDSN   string `json:"database_dsn"`
}

func LoadConfig(filePath string) *Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening config file: v%", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: v%", err)
	}

	return &config
}
