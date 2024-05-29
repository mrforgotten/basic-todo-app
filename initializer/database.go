package initializer

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDb(opts *pg.Options) *pg.DB {

	var db *pg.DB = pg.Connect(opts)
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Println("Connection failed:", err)
		panic(err)
	}

	log.Printf("Connected to database")

	return db
}
