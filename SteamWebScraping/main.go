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
	fmt.Println("AmountGames:", amountGames)

	latestCount, err := dataaccess.CompareAndInsert(db, amountGames)
	if err != nil {
		fmt.Printf("Error comparing and inserting: %v\n", err)
	} else {
		fmt.Printf("Comparison and insertion completed successfully. %d > %d\n", amountGames, latestCount)
	}

	// if amountGames > latestCount
	// send message: new game added

	// if amountGames < latestCount
	// send message: a game has been removed

	// if amountGames == latestCount
	// do nothing
}
