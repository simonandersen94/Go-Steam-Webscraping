package main

import (
	"SteamWebScraping/config"
	"SteamWebScraping/dataaccess"
	"SteamWebScraping/rabbit_MQ"
	"SteamWebScraping/scraping"
	"fmt"
	"time"
)

func main() {
	cfg := config.LoadConfig("config/config.json")
	if cfg == nil {
		fmt.Println("Failed to load configuration")
		return
	}

	rmqService, err := rabbit_MQ.NewRabbitMQService(cfg.RabbitMQUri, cfg.RabbitMQClientProvidedName)
	if err != nil {
		fmt.Printf("Failed to initialize RabbitMQ: %v\n", err)
		return
	}
	defer rmqService.Close()

	db, err := dataaccess.ConnectToDb()
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return
	}
	defer db.Close()

	for {
		amountGames, err := scraping.ScrapeGamesCount(cfg)
		if err != nil {
			fmt.Printf("Error scraping game count: %v\n", err)
			return
		}

		latestCount, err := dataaccess.CompareAndInsert(db, amountGames)
		fmt.Printf("Amount games: %d, Latest count: %d\n", amountGames, latestCount)

		if amountGames > latestCount {
			difference := amountGames - latestCount
			message := fmt.Sprintf("Steamuser %s added %d new game(s)", cfg.SteamID, difference)
			fmt.Println(message)

			err = rabbit_MQ.SendMessage(
				rmqService,
				cfg.RabbitMQExchangeName,
				cfg.RabbitMQRoutingKey,
				message,
			)
			if err != nil {
				fmt.Printf("Error sending message to RabbitMQ: %v\n", err)
			}
		}

		time.Sleep(60 * time.Minute)
	}
}
