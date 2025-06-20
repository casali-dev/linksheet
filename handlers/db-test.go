package handlers

import (
	"fmt"
	"net/http"

	"github.com/casali-dev/linksheet/db"
)

func DBTestHandler(w http.ResponseWriter, r *http.Request) {
	conn := db.DB

	var count int
	err := conn.QueryRow("SELECT COUNT(*) FROM links").Scan(&count)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Database access error: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"links": %d}`, count)
}
