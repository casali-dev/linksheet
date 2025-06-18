package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casali-dev/linksheet/db"
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
			http.Error(w, "Erro ao buscar links", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(links)

	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		newLink := models.NewLink(payload.Name, payload.Description, payload.URL)

		err := repo.Insert(newLink)
		if err != nil {
			http.Error(w, "Erro ao inserir link", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(newLink)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
