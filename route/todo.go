package route

import (
	handler "basic-rest-api-orm/handler"

	"github.com/gin-gonic/gin"
)

func TodoRoute(gr *gin.Engine, h handler.TodoHandler) *gin.Engine {
	grg := gr.Group("/todo")
	{
		grg.GET("/", h.TodoGetAll)
		grg.GET("/:id", h.TodoGetByID)
		grg.POST("/create", h.TodoCreate)
		grg.POST("/update/:id/completed", h.TodoUpdateComplete)
		grg.POST("/update/:id", h.TodoUpdate)
	}

	return gr
}
