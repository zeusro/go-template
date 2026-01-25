package core

import (
	"go.uber.org/fx"
	"zeusro.com/hermes/internal/core/config"
	"zeusro.com/hermes/internal/core/database"
	"zeusro.com/hermes/internal/core/logprovider"
	"zeusro.com/hermes/internal/core/webprovider"
	"zeusro.com/hermes/internal/infrastructure/cache"
	"zeusro.com/hermes/internal/infrastructure/observability"
	"zeusro.com/hermes/internal/infrastructure/security"
)

var CoreModule = fx.Options(
	fx.Provide(config.NewFileConfig),
	fx.Provide(logprovider.GetLogger),
	database.Module,
	cache.Module,
	observability.Module,
	security.Module,
	fx.Provide(webprovider.NewGinEngine),
	fx.Invoke(func(db *database.DB, log logprovider.Logger) {
		if err := database.AutoMigrate(db, log); err != nil {
			log.Panicf("Failed to run migrations: %v", err)
		}
	}),
)
