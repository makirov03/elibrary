package bookrepository

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
