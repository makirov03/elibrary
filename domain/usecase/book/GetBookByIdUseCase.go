package book

import (
	"context"
	"elibrary/domain/models"
)

type GetBookByIdUseCase interface {
	execute(ctx context.Context, id int) (models.Book, error)
}
