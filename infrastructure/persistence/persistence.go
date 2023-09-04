package persistence

import (
	"azt.com/api-sample/infrastructure/persistence/db"
)

type Persistence struct {
	AztDB *db.AztDbEngine
}

func Create() *Persistence {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	return &Persistence{db}
}
