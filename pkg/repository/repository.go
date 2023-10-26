package repository

import (
	"github.com/jmoiron/sqlx"
	"name-service/model"
)

type Human interface {
	Create(human model.Human) (int, error)
	GetAll(filter model.FilterHuman, pageSize int) ([]model.Human, error)
	Delete(id int) error
	Update(id int, input model.UpdateHumanInput) error
}

type Repository struct {
	Human
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Human: NewHumanPostgres(db),
	}
}
