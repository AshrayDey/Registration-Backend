package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and creates necessary tables.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to DB: " + err.Error())
	}

	// Set maximum open and idle connections
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create the required table if it doesn't exist
	createTable()
}

// createTable creates the events table if it doesn't already exist.
func createTable() {
	var err error
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT
	);
	`

	_, err = DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user's table: " + err.Error())
	}
}
