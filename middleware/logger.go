package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[API] ")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s %s", start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Printf("[DONE] %s %s in %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
