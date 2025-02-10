package userusecase

import (
	"context"
	"errors"
)

func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
