package main

import (
	"basic-rest-api-orm/config"
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/wire"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var (
	dbOpts   *pg.Options
	dbDriver *pg.DB
)

func main() {
	dbOpts = config.GetDbConfig()

	dbDriver := initializer.InitDb(dbOpts)
	defer dbDriver.Close()

	appOpts := config.GetAppConfig()
	p := wire.InitApi(dbDriver)
	var gin *gin.Engine = p.InitApp()
	log.Printf("Server started at 127.0.0.1:%v\n", appOpts.AppPort)
	gin.Run(fmt.Sprintf("127.0.0.1:%v", appOpts.AppPort))
}
