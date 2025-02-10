package bookusecase

import (
	"elibrary/internal/repository/bookrepository"
)

func NewBookUsecase(repo bookrepository.BookRepository) *BookUsecase {
	return &BookUsecase{repo: repo}
}
