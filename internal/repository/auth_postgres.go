package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func newAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) error {

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2)", usersTable)

	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	err := r.db.Get(&user, query, username, password)
	return user, err
}
