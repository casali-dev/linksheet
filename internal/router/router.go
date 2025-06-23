package router

import (
	"net/http"

	"github.com/casali-dev/linksheet/internal/handlers"
	"github.com/casali-dev/linksheet/internal/middleware"
)

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/db-test", handlers.DBTestHandler)
	mux.HandleFunc("/links/public", handlers.PublicLinksHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)

	mux.Handle("/links", middleware.AuthMiddleware(http.HandlerFunc(handlers.LinkHandler)))

	return Chain(
		mux,
		middleware.RecoverMiddleware,
		middleware.RateLimitMiddleware,
		middleware.LogMiddleware,
		middleware.JSONMiddleware,
	)
}
