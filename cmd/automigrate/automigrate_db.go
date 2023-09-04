package main

import (
	"log"

	"azt.com/api-sample/infrastructure/persistence"
	"azt.com/api-sample/model/entity"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	persistence := persistence.Create()
	persistence.AztDB.DB.AutoMigrate(&entity.Contact{})
}
