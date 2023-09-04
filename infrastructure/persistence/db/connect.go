package db

import (
	"database/sql"
	"fmt"
	"os"
)

type AztDbEngine struct {
	DB *sql.DB
}

func NewDB() (*AztDbEngine, error) {

	// Tomamos las variables de entorno para conectar a la DB
	DbHost := os.Getenv("DB_HOST")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &AztDbEngine{
		DB: db,
	}, nil

}
