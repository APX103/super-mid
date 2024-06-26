package command

import (
	"apx103.com/super-mid/command/cmd"
	"go.uber.org/fx"
)

var Module = fx.Module("Command",
	fx.Provide(
		fx.Annotate(
			cmd.NewCmdRunnerMap,
			fx.ParamTags(`group:"task_runner"`),
		),
	),
	fx.Provide(),
)
