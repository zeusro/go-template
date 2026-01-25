package security

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewOAuth2Provider),
	fx.Provide(NewAuditLogger),
)
