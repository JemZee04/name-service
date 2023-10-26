package service

import (
	"name-service/model"
	"name-service/pkg/repository"
)

type AuthService struct {
	repo repository.Human
}

func NewHumanService(repo repository.Human) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(human model.Human) (int, error) {
	return s.repo.Create(human)
}

func (s *AuthService) GetAll(filter model.FilterHuman, pageSize int) ([]model.Human, error) {
	return s.repo.GetAll(filter, pageSize)
}

func (s *AuthService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *AuthService) Update(id int, input model.UpdateHumanInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
