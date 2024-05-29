package route

import (
	handler "basic-rest-api-orm/api/handler"

	"github.com/gin-gonic/gin"
)

func TodoRoute(gr *gin.Engine, h handler.TodoHandler) *gin.Engine {
	grg := gr.Group("/todo")
	{
		grg.GET("/", h.TodoGetAll)
	}

	return gr
}
