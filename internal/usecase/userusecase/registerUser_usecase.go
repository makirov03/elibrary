package userusecase

import (
	"context"
	"elibrary/domain/models"
	"log"
)

func (u *UserUsecase) RegisterUser(ctx context.Context, user models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error registering user:", err)
		return err
	}

	return nil
}
