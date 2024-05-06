package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/configs"
)

func NewDB(config configs.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode))
	if err != nil {
		fmt.Print("213123123123131")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
