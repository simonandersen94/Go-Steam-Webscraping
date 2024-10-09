package dataaccess

import (
	"SteamWebScraping/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDb() (*sql.DB, error) {
	cfg := config.LoadConfig("config/config.json")

	db, err := sql.Open("mysql", cfg.DatabaseDSN)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
		return nil, err
	}

	return db, nil
}
