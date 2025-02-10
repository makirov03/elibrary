package delivery

import (
	"elibrary/domain/models"
	"elibrary/internal/usecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
}

func NewBookHandler(usecase *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: usecase}
}

func (h *BookHandler) UploadBookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.usecase.UploadBook(ctx, book); err != nil { // ✅ Pass context
		http.Error(w, "Failed to upload book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Book uploaded successfully"})
	if err != nil {
		return
	}
}

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

func (h *BookHandler) DownloadBookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.usecase.GetBookByID(ctx, id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, book.FilePath)
}
