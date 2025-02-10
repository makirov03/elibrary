package userhandlers

import "elibrary/internal/usecase/userusecase"

func NewUserHandler(usecase *userusecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}
