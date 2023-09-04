package main

import (
	"log"
	"os"

	"azt.com/api-sample/infrastructure/persistence"
	"azt.com/api-sample/infrastructure/router"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}
func main() {
	persistence := persistence.Create()
	r := router.CreateRoute(persistence)

	port := ":" + os.Getenv("PORT")
	r.Run(port)
}
