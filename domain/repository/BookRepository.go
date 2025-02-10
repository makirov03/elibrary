package repository

import (
	"context"
	"elibrary/domain/models"
)

type BookRepository interface {
	SaveBook(ctx context.Context, book models.Book) error
	FetchBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id int) (models.Book, error)
}
