package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func LoadPgConfig() (*PgConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return &PgConfig{}, err
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	return &PgConfig{host, port, user, password, database, sslmode}, nil
}
