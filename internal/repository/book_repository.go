package repository

import (
	"context"
	"elibrary/domain/models"
	"github.com/jackc/pgx/v5"
)

type BookRepository interface {
	SaveBook(ctx context.Context, book models.Book) error
	FetchBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id int) (models.Book, error)
}

type bookRepo struct {
	db *pgx.Conn
}

func NewBookRepository(db *pgx.Conn) BookRepository {
	return &bookRepo{db: db}
}

func (r *bookRepo) SaveBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx, "INSERT INTO books (title, author, file_path, uploaded_by) VALUES ($1, $2, $3, $4)",
		book.Title, book.Author, book.FilePath, book.UploadedBy)
	return err
}

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

func (r *bookRepo) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	var book models.Book
	err := r.db.QueryRow(ctx, "SELECT id, title, author, file_path, uploaded_by FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.FilePath, &book.UploadedBy)
	if err != nil {
		return book, err
	}
	return book, nil
}
