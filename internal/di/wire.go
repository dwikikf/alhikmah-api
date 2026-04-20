//go:build wireinject
// +build wireinject

package di

import (
	"github.com/dwikikf/alhikmah-api/internal/infrastucture/database"
	"github.com/dwikikf/alhikmah-api/pkg/validator"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeApp() (*App, error) {
	wire.Build(
		database.Connect,
		wire.Bind(new(database.DBTX), new(*pgxpool.Pool)),
		validator.NewValidator,

		repositorySet,
		usecaseSet,
		handlerSet,

		wire.Struct(new(Handlers), "*"),
		wire.Struct(new(App), "*"),
	)
	return nil, nil
}
