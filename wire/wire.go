//go:build wireinject
// +build wireinject

package wire

import (
	"basic-rest-api-orm/api/handler"
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/repository"
	authorservice "basic-rest-api-orm/service/author"
	todoservice "basic-rest-api-orm/service/todo"

	"github.com/go-pg/pg/v10"
	"github.com/google/wire"
)

var authorSet = wire.NewSet(
	repository.ProvideAuthorRepository,
	authorservice.ProvideAuthorService,
	handler.ProviderAuthorHandler,
)

var todoSet = wire.NewSet(
	repository.ProvideTodoRepository,
	todoservice.ProvideTodoService,
	handler.ProviderTodoHandler,
)

func InitApi(_ *pg.DB) initializer.Provider {
	wire.Build(
		initializer.InitProvider,
		todoSet,
		authorSet,
	)

	return initializer.Provider{}
}
