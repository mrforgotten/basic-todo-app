package handler

import (
	"basic-rest-api-orm/dto"
	"basic-rest-api-orm/helper"
	"basic-rest-api-orm/model"
	authorservice "basic-rest-api-orm/service/author"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService authorservice.AuthorService
}

func NewProviderAuthorHandler(s authorservice.AuthorService) AuthorHandler {
	return AuthorHandler{
		authorService: s,
	}
}

func (h *AuthorHandler) AuthorGetAll(gCtx *gin.Context) {
	authors, err := h.authorService.GetAll()
	if err != nil {

		gCtx.JSON(500, helper.ToRes(err))
		return
	}
	gCtx.JSON(200, gin.H{
		"data": authors,
	})
}

func (h *AuthorHandler) AuthorGetByID(gCtx *gin.Context) {
	id, err := strconv.Atoi(gCtx.Param("id"))
	if err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}
	author, err := h.authorService.GetByID(id)
	if err != nil {
		gCtx.JSON(500, helper.ToRes(err))
		return
	}
	gCtx.JSON(200, author)
}

func (h *AuthorHandler) AuthorCreate(gCtx *gin.Context) {

	var input *dto.AuthorCreateInput

	if err := gCtx.ShouldBindJSON(&input); err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}

	var data = &model.Author{
		Name: input.Name,
	}

	err := h.authorService.Create(data)
	if err != nil {
		if err.Error() == "duplicate unique value for column name" {

			error := "author name already exist"

			gCtx.JSON(400, helper.ToRes(error))

			return

		}
		gCtx.JSON(500, helper.ToRes(err))
		return
	}
	gCtx.JSON(200, helper.ToRes(err))
}

func (h *AuthorHandler) AuthorUpdate(gCtx *gin.Context) {
	p := gCtx.Param("id")

	id, err := strconv.Atoi(p)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, helper.ToRes(err))
		return
	}

	var input *dto.AuthorUpdateInput
	if err := gCtx.ShouldBindJSON(&input); err != nil {
		gCtx.JSON(400, helper.ToRes(err))
		return
	}

	var data = &model.Author{
		Name: input.Name,
	}

	data.Id = id

	if err := h.authorService.Update(id, data); err != nil {
		gCtx.JSON(500, helper.ToRes(err))
		return
	}

	gCtx.JSON(200, helper.ToRes(data))
}
