package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite", "db/linksheet.db")
	if err != nil {
		log.Fatal("[DB] Failed to open database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("[DB] Failed to connect to database:", err)
	}

	log.Println("[DB] Successfully connected to SQLite")
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("[DB] Connection closed")
	}
}
