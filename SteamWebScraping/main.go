package main

import (
	"SteamWebScraping/dataaccess"
	"SteamWebScraping/scraping"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := dataaccess.ConnectToDb()
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		return
	}
	defer db.Close()

	amountGames, err := scraping.ScrapeGamesCount()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	latestCount, err := dataaccess.CompareAndInsert(db, amountGames)
	if err != nil {
		fmt.Printf("Error comparing and inserting: %v\n", err)
	} else if amountGames > latestCount {
		difference := amountGames - latestCount
		if difference == 1 {
			fmt.Printf("User added %d new game", difference)
		} else if difference > 1 {
			fmt.Printf("User added %d new games", difference)
		}
		// Send a message to RabbitMQ ---
	}
}
