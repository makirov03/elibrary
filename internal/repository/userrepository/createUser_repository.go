package userrepository

import (
	"context"
	"elibrary/domain/models"
)

func (r *userRepo) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role)
	return err
}
