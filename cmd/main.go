package main

import (
	"basic-rest-api-orm/config"
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/wire"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var (
	opts     *pg.Options
	dbDriver *pg.DB
)

func main() {
	opts = config.GetDbConfig()

	dbDriver := initializer.InitDb(opts)
	defer dbDriver.Close()

	p := wire.InitApi(dbDriver)
	var gin *gin.Engine = p.InitApp()
	log.Println("Server started at 127.0.0.1:8080")
	gin.Run("127.0.0.1:8080")
}
