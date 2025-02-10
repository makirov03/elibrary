package usecase

import "elibrary/domain/repository"

type GetBooksUseCaseImpl struct {
	repo repository.BookRepository
}

//func NewUserUsecase(repo repository.BookRepository) usecase.GetBooksUseCase {
//	return &GetBooksUseCaseImpl{repo: usecase.BookRepository}
//}
