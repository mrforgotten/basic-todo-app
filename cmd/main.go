package main

import (
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/wire"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var (
	opts     *pg.Options
	dbDriver *pg.DB
)

func main() {

	opts = &pg.Options{
		Addr:     os.Getenv("PG_HOST") + ":" + os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_NAME"),
	}

	db := initializer.InitDb(opts)

	p := wire.InitApi(db)
	var gin *gin.Engine = p.InitApp()

	gin.Run("127.0.0.1:8080")
}
