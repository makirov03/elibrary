package userhandlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.usecase.Login(ctx, creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}
