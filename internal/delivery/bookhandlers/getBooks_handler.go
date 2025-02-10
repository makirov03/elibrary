package bookhandlers

import (
	"encoding/json"
	"net/http"
)

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	books, err := h.usecase.GetBooks(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}
