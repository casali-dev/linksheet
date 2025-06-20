package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/casali-dev/linksheet/auth"
	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/middleware"
	"github.com/casali-dev/linksheet/repositories"
	"github.com/casali-dev/linksheet/services"
)

type createLinkPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	IsPublic    bool   `json:"isPublic"`
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewLinkRepository(db.DB)
	service := services.NewLinkService(repo)

	switch r.Method {
	case http.MethodGet:
		authInfo, ok := auth.GetAuthInfo(r.Context())

		if !ok {
			middleware.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		links, err := service.GetAllByAuthor(authInfo.AuthorID)

		if err != nil {
			middleware.WriteError(w, http.StatusInternalServerError, "Failed to fetch links")
			return
		}
		middleware.WriteJSON(w, http.StatusOK, links, "")

	case http.MethodPost:
		var payload createLinkPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			middleware.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
			return
		}

		name := strings.TrimSpace(payload.Name)
		desc := strings.TrimSpace(payload.Description)
		url := strings.TrimSpace(payload.URL)

		authInfo, ok := auth.GetAuthInfo(r.Context())
		if !ok {
			middleware.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		authorID := authInfo.AuthorID

		link, err := service.Create(name, desc, url, payload.IsPublic, authorID)
		if err != nil {
			middleware.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		middleware.WriteJSON(w, http.StatusCreated, link, "Link created successfully")

	default:
		middleware.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
