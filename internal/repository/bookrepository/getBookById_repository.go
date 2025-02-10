package bookrepository

import (
	"context"
	"elibrary/domain/models"
)

func (r *bookRepo) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	err := r.db.QueryRow(ctx, "SELECT id, title, author, file_path, uploaded_by FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.FilePath, &book.UploadedBy)
	if err != nil {
		return book, err
	}
	return book, nil
}
