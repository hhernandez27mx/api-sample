package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AztDbEngine struct {
	DB *gorm.DB
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	return &AztDbEngine{
		DB: db,
	}, nil

}
