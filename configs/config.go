package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfigs struct {
	Postgres     Postgres
	ClientID     string
	ClientSecret string
}
type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig() (AppConfigs, error) {
	envFilePath := filepath.Join("configs", ".env")
	err := godotenv.Load(envFilePath)
	if err != nil {
		return AppConfigs{}, err
	}
	database := Postgres{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	CID := os.Getenv("CLIENT_ID")
	CS := os.Getenv("CLIENT_SECRET")

	return AppConfigs{Postgres: database, ClientID: CID, ClientSecret: CS}, nil
}
