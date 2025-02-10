package bookhandlers

import (
	"elibrary/internal/usecase/bookusecase"
)

type BookHandler struct {
	usecase *bookusecase.BookUsecase
}
