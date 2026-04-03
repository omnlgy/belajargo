package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(fmt.Errorf("failed to open database: %w", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	queryUser := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(queryUser)
	if err != nil {
		panic(fmt.Errorf("failed to create table: %w", err))
	}

	queryEvent := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(queryEvent)
	if err != nil {
		panic(fmt.Errorf("failed to create table: %w", err))
	}

	queryRegistration := `
	CREATE TABLE IF NOT EXISTS registrations (
		id TEXT PRIMARY KEY,
		eventId TEXT NOT NULL,
		userId TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(eventId) REFERENCES events(id),
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(queryRegistration)
	if err != nil {
		panic(fmt.Errorf("failed to create table: %w", err))
	}
}
