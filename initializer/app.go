package initializer

import (
	"basic-rest-api-orm/api/handler"
	"basic-rest-api-orm/api/route"

	"github.com/gin-gonic/gin"
)

type Provider struct {
	AuthorHandler handler.AuthorHandler

	// Add more dependencies here
	TodoHandler handler.TodoHandler
}

func InitProvider(h1 handler.AuthorHandler, h2 handler.TodoHandler) Provider {
	return Provider{
		AuthorHandler: h1,
		TodoHandler:   h2,
	}
}

func (p Provider) InitApp() *gin.Engine {
	r := gin.Default()

	route.AuthorRoute(r, p.AuthorHandler)
	route.TodoRoute(r, p.TodoHandler)

	return r
}
