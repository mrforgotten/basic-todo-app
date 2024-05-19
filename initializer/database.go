package initializer

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func InitDb(opts *pg.Options) *pg.DB {

	var db *pg.DB = pg.Connect(opts)

	err := db.Ping(db.Context())

	if err != nil {
		log.Fatalln("Failed to connect to database", err)
		panic("Failed to connect to database")
	}

	log.Printf("Connected to database")

	return db
}
