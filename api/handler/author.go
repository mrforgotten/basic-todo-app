package handler

import (
	authorservice "basic-rest-api-orm/service/author"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService authorservice.AuthorService
}

func ProviderAuthorHandler(a authorservice.AuthorService) AuthorHandler {
	return AuthorHandler{
		authorService: a,
	}
}

func (a AuthorHandler) AuthorGetAll(gCtx *gin.Context) {
	authors, err := a.authorService.GetAll()
	if err != nil {
		gCtx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	gCtx.JSON(200, gin.H{
		"data": authors,
	})
}
