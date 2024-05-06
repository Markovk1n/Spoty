package service

import (
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
