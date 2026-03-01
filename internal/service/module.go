package service

import (
	"go.uber.org/fx"
	domainService "github.com/zeusro/go-template/internal/domain/service"
)

var Modules = fx.Options(
	fx.Provide(NewHealthService),
	fx.Provide(NewTranslateService),
	domainService.DomainServiceModule,
	//todo 有新的服务需要添加到这里
)
