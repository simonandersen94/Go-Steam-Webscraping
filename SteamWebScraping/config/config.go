package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ScrapeUrl                  string `json:"scrape_url"`
	SteamID                    string `json:"steam_id"`
	AllowedDomain              string `json:"allowedDomain"`
	DatabaseDSN                string `json:"database_dsn"`
	RabbitMQUri                string `json:"RabbitMQ_Uri"`
	RabbitMQClientProvidedName string `json:"rabbitMQ_clientProvidedName"`
	RabbitMQExchangeName       string `json:"rabbitMQ_exchangeName"`
	RabbitMQRoutingKey         string `json:"rabbitMQ_routingKey"`
	RabbitMQQueueName          string `json:"rabbitMQ_queueName"`
}

func LoadConfig(filePath string) *Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening config file:", err)
		return nil
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Error decoding config file:", err)
		return nil
	}

	return &config
}
