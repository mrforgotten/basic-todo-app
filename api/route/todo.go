package route

import (
	handler "basic-rest-api-orm/api/handler"

	gin "github.com/gin-gonic/gin"
)

func TodoRoute(gr *gin.Engine, h handler.TodoHandler) {
	grg := gr.Group("/todo")
	grg.GET("/", h.TodoGetAll)
}
