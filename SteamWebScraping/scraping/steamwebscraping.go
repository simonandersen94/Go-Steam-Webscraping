package scraping

import (
	"SteamWebScraping/config"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func ScrapeGamesCount(cfg *config.Config) (int, error) {
	var c = colly.NewCollector(colly.AllowedDomains(cfg.AllowedDomain))
	var amountGames int
	var valueGameCountCaptured bool = false

	c.OnHTML("div.value", func(e *colly.HTMLElement) {
		if !valueGameCountCaptured {
			trimmedAmountGames := strings.TrimSpace(e.Text)
			amount, err := strconv.Atoi(trimmedAmountGames)
			if err == nil {
				amountGames = amount
				valueGameCountCaptured = true
			} else {
				fmt.Println("Error parsing game count:", err)
			}
		}
	})

	err := c.Visit(cfg.ScrapeUrl + cfg.SteamID)
	if err != nil {
		fmt.Printf("Error visiting URL: %v\n", err)
		return 0, err
	}
	if !valueGameCountCaptured {
		fmt.Println(err)
		return 0, err
	}

	return amountGames, nil
}
