package handler

import (
	"basic-rest-api-orm/model"
	todoservice "basic-rest-api-orm/service/todo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService todoservice.TodoService
}

func ProviderTodoHandler(s todoservice.TodoService) TodoHandler {
	return TodoHandler{
		todoService: s,
	}
}

func (h *TodoHandler) TodoGetAll(gCtx *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {
		gCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gCtx.JSON(200, gin.H{
		"data": todos,
	})
}

func (h *TodoHandler) TodoGetByID(gCtx *gin.Context) {
	id, err := strconv.Atoi(gCtx.Param("id"))
	if err != nil {
		gCtx.JSON(400, gin.H{"error": err.Error()})
	}
	todo, err := h.todoService.GetByID(id)
	if err != nil {
		gCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gCtx.JSON(200, todo)
}

func (h *TodoHandler) TodoCreate(gCtx *gin.Context) {
	var todo model.Todo
	if err := gCtx.ShouldBindJSON(&todo); err != nil {
		gCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	data, err := h.todoService.Create(todo)
	if err != nil {
		gCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gCtx.JSON(200, data)
}

func (h *TodoHandler) TodoUpdate(gCtx *gin.Context) {
	var todo model.Todo
	if err := gCtx.ShouldBindJSON(&todo); err != nil {
		gCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	data, err := h.todoService.Update(todo)
	if err != nil {
		gCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	gCtx.JSON(200, data)
}
