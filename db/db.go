package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	database, err := sql.Open("sqlite3", "api.db")
	DB = database
	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {
	createEventTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER
		)
	`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create events table")
	}
}
