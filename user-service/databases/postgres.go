package databases

import (
	"database/sql"
	"fmt"
	"log"
	"user-service/configs"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dbConfig, err := configs.LoadPgConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName, dbConfig.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open databases: %v", err)
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to databases: %v", err)
	}

	log.Println("Successfully connected to databases!")
	return db, nil
}
