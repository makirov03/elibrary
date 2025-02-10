package bookusecase

import (
	"context"
	"elibrary/domain/models"
	"errors"
)

func (u *BookUsecase) UploadBook(ctx context.Context, book models.Book) error {
	if book.Title == "" || book.FilePath == "" {
		return errors.New("missing required fields")
	}
	return u.repo.SaveBook(ctx, book)
}
