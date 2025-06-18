package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/models"
	"github.com/casali-dev/linksheet/repositories"
)

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
		var newLink models.Link
		if err := json.NewDecoder(r.Body).Decode(&newLink); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

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
