package usecase

import (
	"context"
	"elibrary/domain/models"
	"elibrary/internal/repository"
	"errors"
)

type BookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) *BookUsecase {
	return &BookUsecase{repo: repo}
}

func (u *BookUsecase) UploadBook(ctx context.Context, book models.Book) error {
	if book.Title == "" || book.FilePath == "" {
		return errors.New("missing required fields")
	}
	return u.repo.SaveBook(ctx, book)
}

func (u *BookUsecase) GetBooks(ctx context.Context) ([]models.Book, error) {
	return u.repo.FetchBooks(ctx)
}

func (u *BookUsecase) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	return u.repo.GetBookByID(ctx, id)
}
