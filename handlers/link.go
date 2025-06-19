package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/middleware"
	"github.com/casali-dev/linksheet/repositories"
	"github.com/casali-dev/linksheet/services"
)

type createLinkPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewLinkRepository(db.DB)
	service := services.NewLinkService(repo)

	switch r.Method {

	case http.MethodGet:
		links, err := service.GetAll()
		if err != nil {
			middleware.WriteError(w, http.StatusInternalServerError, "Erro ao buscar links")
			return
		}
		middleware.WriteJSON(w, http.StatusOK, links, "")

	case http.MethodPost:
		var payload createLinkPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			middleware.WriteError(w, http.StatusBadRequest, "JSON inválido")
			return
		}

		link, err := service.Create(payload.Name, payload.Description, payload.URL)
		if err != nil {
			middleware.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		middleware.WriteJSON(w, http.StatusCreated, link, "Link criado com sucesso")

	default:
		middleware.WriteError(w, http.StatusMethodNotAllowed, "Método não permitido")
	}
}
