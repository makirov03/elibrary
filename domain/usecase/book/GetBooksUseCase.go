package book

import (
	"context"
	"elibrary/domain/models"
)

type GetBooksUseCase interface {
	execute(ctx context.Context) ([]models.Book, error)
}
