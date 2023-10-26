package service

import (
	"name-service/model"
	"name-service/pkg/repository"
)

type Human interface {
	Create(human model.Human) (int, error)
	GetAll(filter model.FilterHuman, pageSize int) ([]model.Human, error)
	Delete(id int) error
	Update(id int, input model.UpdateHumanInput) error
}

type Service struct {
	Human
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Human: NewHumanService(repos.Human),
	}
}
