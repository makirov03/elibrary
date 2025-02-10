package userrepository

import "github.com/jackc/pgx/v5"

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepo{db: db}
}
