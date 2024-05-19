package main

import (
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/wire"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

var (
	opts     *pg.Options
	dbDriver *pg.DB
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	opts = &pg.Options{
		Addr:     os.Getenv("PG_HOST") + ":" + os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			_, err := cn.Exec("SET search_path=public")
			if err != nil {
				return err
			}
			return nil
		},
	}

	dbDriver := initializer.InitDb(opts)
	defer dbDriver.Close()

	p := wire.InitApi(dbDriver)
	var gin *gin.Engine = p.InitApp()
	log.Println("Server started at 127.0.0.1:8080")
	gin.Run("127.0.0.1:8080")
}
