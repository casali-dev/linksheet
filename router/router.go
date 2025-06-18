package router

import (
	"net/http"

	"github.com/casali-dev/linksheet/handlers"
	"github.com/casali-dev/linksheet/middleware"
)

func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/links",   handlers.LinkHandler)
	mux.HandleFunc("/ping",    handlers.PingHandler)
	mux.HandleFunc("/health",  handlers.HealthHandler)
	mux.HandleFunc("/db-test", handlers.DBTestHandler)

	return Chain(mux,
		middleware.RecoverMiddleware,
		middleware.RateLimitMiddleware,
		middleware.LogMiddleware,
		middleware.JSONMiddleware,
	)
}
