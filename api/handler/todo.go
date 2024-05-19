package handler

import "github.com/gin-gonic/gin"

type TodoHandler struct{}

func ProviderTodoHandler() TodoHandler {
	return TodoHandler{}
}

func (t TodoHandler) TodoGetAll(gCtx *gin.Context) {
	gCtx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
