package bookrepository

import "github.com/jackc/pgx/v5"

func NewBookRepository(db *pgx.Conn) BookRepository {
	return &bookRepo{db: db}
}
