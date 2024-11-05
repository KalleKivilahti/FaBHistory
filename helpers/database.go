package helpers

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./matchhistory.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS matches (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        player1 TEXT,
        player2 TEXT,
        deck1 TEXT,
        deck2 TEXT,
        winner TEXT,
        turns TEXT,
        date TEXT
    );`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
