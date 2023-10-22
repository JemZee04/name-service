package service

import "name-service/pkg/repository"

type Human interface {
}

type Service struct {
	Human
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
