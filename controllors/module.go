package controllors

import (
	"go.uber.org/fx"
	// "apx103.com/super-mid/controllors/bot"
	// "apx103.com/super-mid/controllors/bot/config"
	"apx103.com/super-mid/controllors/common"
	"apx103.com/super-mid/controllors/demo"
)

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(common.Controller)),
		fx.ResultTags(`group:"controller"`),
	)
}

// Add Router Hear

var Module = fx.Module("Controller",
	fx.Provide(
		// fx.Annotate(
		// 	bot.NewEventsMap,
		// 	fx.ParamTags(`group:"botEvents"`),
		// ),
		AsController(demo.NewPingController),
	),
)
