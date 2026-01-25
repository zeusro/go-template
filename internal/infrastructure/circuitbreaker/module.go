package circuitbreaker

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	// Circuit breakers are created on-demand, no global providers needed
)
