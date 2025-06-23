package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/casali-dev/linkhub/internal/db"
	"github.com/casali-dev/linkhub/internal/middleware"
	"github.com/casali-dev/linkhub/internal/repositories"
	"github.com/casali-dev/linkhub/internal/services"
)

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signupResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		middleware.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req signupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middleware.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	repo := repositories.NewUserRepository(db.DB)
	service := services.NewUserService(repo)

	user, err := service.Signup(req.Email, req.Password)
	if err != nil {
		middleware.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := signupResponse{ID: user.ID, Email: user.Email, Role: user.Role}
	middleware.WriteJSON(w, http.StatusCreated, resp, "Signup successful")
}
