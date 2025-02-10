package repository

import (
	"context"
	"elibrary/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
