package route

import (
	handler "basic-rest-api-orm/api/handler"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(gr *gin.Engine, h handler.AuthorHandler) *gin.Engine {
	grg := gr.Group("/author")
	{
		grg.GET("/", h.AuthorGetAll)
		grg.GET("/:id", h.AuthorGetByID)
		grg.POST("/create", h.AuthorCreate)
	}

	return gr
}
