package bookusecase

import (
	"context"
	"elibrary/domain/models"
)

func (u *BookUsecase) GetBooks(ctx context.Context) ([]models.Book, error) {
	return u.repo.FetchBooks(ctx)
}
