package handler

import (
	"basic-rest-api-orm/dto"
	"basic-rest-api-orm/helper"
	"basic-rest-api-orm/model"
	todoservice "basic-rest-api-orm/service/todo"
	"log"
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
	data, err := h.todoService.GetAll()
	if err != nil {

		gCtx.JSON(500, helper.ToRes(err))
		return
	}

	gCtx.JSON(200, helper.ToRes(data))
}

func (h *TodoHandler) TodoGetByID(gCtx *gin.Context) {
	id, err := strconv.Atoi(gCtx.Param("id"))
	if err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}
	todo, err := h.todoService.GetByID(id)
	if err != nil {
		gCtx.JSON(500, helper.ToRes(err))
		return
	}
	gCtx.JSON(200, todo)
}

func (h *TodoHandler) TodoCreate(gCtx *gin.Context) {
	var input *dto.TodoCreate
	if err := gCtx.ShouldBindJSON(&input); err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}

	authorId, err := strconv.Atoi(gCtx.GetHeader("author_id"))

	if err != nil {
		gCtx.JSON(400, helper.ToRes(err))
	}
	var data = &model.Todo{
		Title:    input.Title,
		AuthorId: authorId,
	}

	err = h.todoService.Create(data)
	if err != nil {
		gCtx.JSON(500, helper.ToRes(err))

		log.Printf("error unable to create: %v", err)

		return
	}
	gCtx.JSON(200, helper.ToRes(data))
}

func (h *TodoHandler) TodoUpdate(gCtx *gin.Context) {
	p := gCtx.Param("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, helper.ToRes(err))
		return
	}

	var data *model.Todo
	if err := gCtx.ShouldBindJSON(&data); err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}

	data.Id = id

	if err := h.todoService.Update(id, data); err != nil {
		gCtx.JSON(500, helper.ToRes(err))
		return
	}

	gCtx.JSON(200, helper.ToRes(data))
}
