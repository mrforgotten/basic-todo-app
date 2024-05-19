//go:build wireinject
// +build wireinject

package wire

import (
	"basic-rest-api-orm/api/handler"
	"basic-rest-api-orm/initializer"
	"basic-rest-api-orm/repository"
	authorservice "basic-rest-api-orm/service/author"

	"github.com/go-pg/pg/v10"
	"github.com/google/wire"
)

func InitApi(_ *pg.DB) initializer.Provider {
	wire.Build(
		initializer.InitProvider,
		handler.ProviderAuthorHandler,
		repository.ProvideAuthorRepository,
		authorservice.ProvideAuthorService,
		handler.ProviderTodoHandler,
	)

	return initializer.Provider{}
}
