package userhandlers

import (
	"elibrary/internal/usecase/userusecase"
)

type UserHandler struct {
	usecase *userusecase.UserUsecase
}
