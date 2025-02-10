package repository

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

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role)
	return err
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, "SELECT id, username, password, role FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
