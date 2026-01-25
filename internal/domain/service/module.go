package service

import (
	"go.uber.org/fx"
)

var DomainServiceModule = fx.Options(
	fx.Provide(NewUserService),
)
