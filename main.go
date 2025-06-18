package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/router"
)

func main() {
	db.Connect()
	defer db.Close()
	db.RunMigrations()

	server := &http.Server{
		Addr:    ":3333",
		Handler: router.Handler(),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Println("Servidor rodando na porta 3333")
		serverErrors <- server.ListenAndServe()
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Erro no servidor: %v", err)

	case sig := <-sigint:
		log.Printf("Sinal capturado: %v. Encerrando com graceful shutdown...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro ao finalizar servidor: %v", err)
		} else {
			log.Println("Servidor encerrado com sucesso.")
		}
	}
}
