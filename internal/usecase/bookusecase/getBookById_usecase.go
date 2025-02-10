package bookusecase

import (
	"context"
	"elibrary/domain/models"
)

func (u *BookUsecase) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	return u.repo.GetBookByID(ctx, id)
}
