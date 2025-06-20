package handlers

import (
	"net/http"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/middleware"
	"github.com/casali-dev/linksheet/repositories"
	"github.com/casali-dev/linksheet/services"
)

func PublicLinksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		middleware.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	repo := repositories.NewLinkRepository(db.DB)
	service := services.NewLinkService(repo)

	links, err := service.GetPublic()
	if err != nil {
		middleware.WriteError(w, http.StatusInternalServerError, "Failed to fetch public links")
		return
	}

	middleware.WriteJSON(w, http.StatusOK, links, "")
}
