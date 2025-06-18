package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/router"
)

func main() {
	db.Connect()
	defer db.Close()

	db.RunMigrations()

	ln, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}

	log.Println("Servidor rodando na porta 3333")
	err = http.Serve(ln, router.Handler())
	if err != nil {
		log.Printf("Erro no servidor: %v", err)
		os.Exit(1)
	}
}
