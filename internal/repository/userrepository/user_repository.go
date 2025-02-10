package userrepository

import (
	"context"
	"elibrary/domain/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type userRepo struct {
	db *pgx.Conn
}
