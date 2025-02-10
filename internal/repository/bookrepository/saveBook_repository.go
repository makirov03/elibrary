package bookrepository

import (
	"context"
	"elibrary/domain/models"
)

func (r *bookRepo) SaveBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx, "INSERT INTO books (title, author, file_path, uploaded_by) VALUES ($1, $2, $3, $4)",
		book.Title, book.Author, book.FilePath, book.UploadedBy)
	return err
}
