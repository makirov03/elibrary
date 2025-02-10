package bookrepository

import (
	"context"
	"elibrary/domain/models"
)

func (r *bookRepo) FetchBooks(ctx context.Context) ([]models.Book, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, author, file_path, uploaded_by FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.FilePath, &book.UploadedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
