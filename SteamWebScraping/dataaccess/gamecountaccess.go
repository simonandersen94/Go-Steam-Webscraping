package dataaccess

import (
	"database/sql"
	"fmt"
)

func get(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT count FROM steamgamecount ORDER BY id DESC LIMIT 1").Scan(&count)
	if err != nil {
		fmt.Printf("Error retrieving the latest count: %v\n", err)
		return 0, err
	}
	fmt.Printf("Antal: %d\n", count)
	return count, nil
}

func insert(db *sql.DB, count int) error {
	query := "INSERT INTO steamgamecount (count) VALUES (?)"
	_, err := db.Exec(query, count)
	if err != nil {
		fmt.Printf("Error inserting into database: %v\n", err)
		return err
	}
	return nil
}

func CompareAndInsert(db *sql.DB, input int) (int, error) {
	latestCount, err := get(db)
	if err != nil {
		return 0, err
	}
	if input > latestCount {
		err = insert(db, input)
		if err != nil {
			return 0, err
		}
		return latestCount, nil
	} else if input < latestCount {
		fmt.Printf("Input %d er mindre end den seneste værdi %d.\n", input, latestCount)
		err = insert(db, input)
		if err != nil {
			return 0, err
		}
		return latestCount, nil
	} else {
		fmt.Printf("Input %d er ikke større end den seneste værdi %d. Ingen indsættelse.\n", input, latestCount)
		return latestCount, nil
	}
}
