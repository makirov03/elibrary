package bookhandlers

import "elibrary/internal/usecase/bookusecase"

func NewBookHandler(usecase *bookusecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: usecase}
}
