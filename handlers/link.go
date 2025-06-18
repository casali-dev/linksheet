package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/middleware"
	"github.com/casali-dev/linksheet/models"
	"github.com/casali-dev/linksheet/repositories"
)

var payload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewLinkRepository(db.DB)

	switch r.Method {
	case http.MethodGet:
		links, err := repo.GetAll()
		if err != nil {
			middleware.WriteError(w, http.StatusInternalServerError, "Erro ao buscar links")
			return
		}

		middleware.WriteJSON(w, http.StatusOK, links, "")

	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			middleware.WriteError(w, http.StatusBadRequest, "JSON inválido")
			return
		}

		if payload.Name == "" || payload.URL == "" {
			middleware.WriteError(w, http.StatusBadRequest, "Campos 'name' e 'url' são obrigatórios")
			return
		}

		newLink := models.NewLink(payload.Name, payload.Description, payload.URL)

		if err := repo.Insert(newLink); err != nil {
			middleware.WriteError(w, http.StatusInternalServerError, "Erro ao inserir link")
			return
		}

		middleware.WriteJSON(w, http.StatusCreated, newLink, "Link criado com sucesso")

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
