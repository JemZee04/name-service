package repository

import "github.com/jmoiron/sqlx"

type Human interface {
}

type Repository struct {
	Human
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
