package middleware

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)

	res := response{
		Error: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, msg string) {
	w.WriteHeader(status)

	res := response{
		Message: msg,
		Data:    data,
	}

	json.NewEncoder(w).Encode(res)
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
