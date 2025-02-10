package userhandlers

import (
	"context"
	"elibrary/domain/models"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.usecase.RegisterUser(ctx, user); err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
