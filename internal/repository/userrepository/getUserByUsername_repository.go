package userrepository

import (
	"context"
	"elibrary/domain/models"
)

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, "SELECT id, username, password, role FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
