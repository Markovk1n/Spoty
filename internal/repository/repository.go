package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/models"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
	}
}
