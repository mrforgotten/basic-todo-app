package handler

import (
	"basic-rest-api-orm/model"
	todoservice "basic-rest-api-orm/service/todo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService todoservice.TodoService
}

func NewProviderTodoHandler(s todoservice.TodoService) TodoHandler {
	return TodoHandler{
		todoService: s,
	}
}

func (h *TodoHandler) TodoGetAll(gCtx *gin.Context) {
	todos, err := h.todoService.GetAll()
	if err != nil {

		gCtx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	gCtx.JSON(200, gin.H{
		"data": todos,
	})
}

func (h *TodoHandler) TodoGetByID(gCtx *gin.Context) {
	id, err := strconv.Atoi(gCtx.Param("id"))
	if err != nil {
		gCtx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	todo, err := h.todoService.GetByID(id)
	if err != nil {
		gCtx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	gCtx.JSON(200, todo)
}

func (h *TodoHandler) TodoCreate(gCtx *gin.Context) {
	var todo *model.Todo
	if err := gCtx.ShouldBindJSON(&todo); err != nil {
		gCtx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.todoService.Create(todo)
	if err != nil {
		if err.Error() == "duplicate unique value for column name" {
			gCtx.JSON(400, gin.H{
				"error": "Todo name already exist",
			})
			return

		}
		gCtx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	gCtx.JSON(200, gin.H{
		"data": todo,
	})
}

func (h *TodoHandler) TodoUpdate(gCtx *gin.Context) {
	p := gCtx.Param("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var update *model.Todo
	if err := gCtx.ShouldBindJSON(&update); err != nil {
		gCtx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	update.Id = id

	if err := h.todoService.Update(id, update); err != nil {
		gCtx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	gCtx.JSON(200, gin.H{
		"data": update,
	})
}
