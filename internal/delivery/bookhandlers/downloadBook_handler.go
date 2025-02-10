package bookhandlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *BookHandler) DownloadBookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bookusecase ID", http.StatusBadRequest)
		return
	}

	book, err := h.usecase.GetBookByID(ctx, id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, book.FilePath)
}
