package repository

import (
	"database/sql"

	"github.com/forChin/test-crud/go-zero-crud/internal/logic/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) SaveUser(user models.User) error {
	stmt := `INSERT INTO users (name, age) VALUES ($1, $2);`
	_, err := r.db.Exec(stmt, user.Name, user.Age)

	return err
}
