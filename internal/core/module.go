package core

import (
	"go.uber.org/fx"
	"github.com/zeusro/go-template/internal/core/config"
	"github.com/zeusro/go-template/internal/core/database"
	"github.com/zeusro/go-template/internal/core/logprovider"
	"github.com/zeusro/go-template/internal/core/webprovider"
	"github.com/zeusro/go-template/internal/infrastructure/cache"
	"github.com/zeusro/go-template/internal/infrastructure/observability"
	"github.com/zeusro/go-template/internal/infrastructure/security"
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
