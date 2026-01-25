package database

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPostgresDB),
	fx.Provide(NewRedisClient),
)
