package fx

import (
	"github.com/CTFg/CTFg/actions/fx/configfx"
	"github.com/CTFg/CTFg/actions/fx/loggerfx"
	"go.uber.org/fx"

)

var Module = fx.Options(
	fx.Provide(
		configfx.New,
		loggerfx.New,
	),
)
