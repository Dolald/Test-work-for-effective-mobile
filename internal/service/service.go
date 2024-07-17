package service

import (
	"test/internal/repository"
)

type Serv interface {
}

type Service struct {
	repository repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: *repository,
	}
}
