package book

import (
	"context"
	"elibrary/domain/models"
)

type SaveBookUseCase interface {
	execute(ctx context.Context, book models.Book)
}
