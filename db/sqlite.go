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
		log.Fatal("[DB] Erro ao abrir o banco:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("[DB] Erro ao conectar ao banco:", err)
	}

	log.Println("[DB] Conexão com SQLite estabelecida com sucesso")
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("[DB] Conexão encerrada")
	}
}
