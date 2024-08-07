// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"basic-rest-api-orm/handler"
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/repository"
	"basic-rest-api-orm/service/author"
	"basic-rest-api-orm/service/todo"
	"github.com/go-pg/pg/v10"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitApi(db *pg.DB) initializer.Provider {
	authorRepository := repository.NewProvideAuthorRepository(db)
	authorService := authorservice.NewProvideAuthorService(authorRepository)
	authorHandler := handler.NewProviderAuthorHandler(authorService)
	todoRepository := repository.NewProvideTodoRepository(db)
	todoService := todoservice.NewProvideTodoService(todoRepository)
	todoHandler := handler.NewProviderTodoHandler(todoService)
	provider := initializer.InitProvider(authorHandler, todoHandler)
	return provider
}

// wire.go:

var authorSet = wire.NewSet(repository.NewProvideAuthorRepository, authorservice.NewProvideAuthorService, handler.NewProviderAuthorHandler)

var todoSet = wire.NewSet(repository.NewProvideTodoRepository, todoservice.NewProvideTodoService, handler.NewProviderTodoHandler)
