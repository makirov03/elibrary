-- +goose Up
CREATE TABLE IF NOT EXISTS books (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     author VARCHAR(255) NOT NULL,
                                     file_path TEXT NOT NULL,
                                     uploaded_by VARCHAR(255) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS books;