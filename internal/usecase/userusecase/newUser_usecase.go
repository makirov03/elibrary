package userusecase

import (
	"elibrary/internal/repository/userrepository"
)

func NewUserUsecase(repo userrepository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
