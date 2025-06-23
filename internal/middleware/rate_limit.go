package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

const minInterval = 1 * time.Second

const maxGlobalReqPerSec = 50

var (
	lastSeen = make(map[string]time.Time)
	mu       sync.Mutex
)

var (
	globalCount     = 0
	globalLastReset = time.Now()
	globalMu        sync.Mutex
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		globalMu.Lock()
		now := time.Now()

		if now.Sub(globalLastReset) >= time.Second {
			globalCount = 0
			globalLastReset = now
		}

		if globalCount >= maxGlobalReqPerSec {
			globalMu.Unlock()
			http.Error(w, "Too many requests (global limit)", http.StatusTooManyRequests)
			return
		}

		globalCount++
		globalMu.Unlock()

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "invalid IP", http.StatusInternalServerError)
			return
		}

		mu.Lock()
		last, seen := lastSeen[ip]

		if seen && now.Sub(last) < minInterval {
			mu.Unlock()
			http.Error(w, "Too many requests (per IP limit)", http.StatusTooManyRequests)
			return
		}

		lastSeen[ip] = now
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
