package bookhandlers

import (
	"elibrary/domain/models"
	"encoding/json"
	"net/http"
)

func (h *BookHandler) UploadBookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.usecase.UploadBook(ctx, book); err != nil { // ✅ Pass context
		http.Error(w, "Failed to upload bookusecase", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Book uploaded successfully"})
	if err != nil {
		return
	}
}
