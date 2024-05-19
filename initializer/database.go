package initializer

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func InitDb(opts *pg.Options) *pg.DB {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to database")
		panic("Failed to connect to database")
	}

	defer db.Close()

	log.Printf("Connected to database")

	return db
}
