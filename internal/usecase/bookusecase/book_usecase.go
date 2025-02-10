package bookusecase

import (
	"elibrary/internal/repository/bookrepository"
)

type BookUsecase struct {
	repo bookrepository.BookRepository
}
