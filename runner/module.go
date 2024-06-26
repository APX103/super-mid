package runner

import (
	"apx103.com/super-mid/command/cmd"
	"go.uber.org/fx"
)

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(cmd.TaskRunner)),
		fx.ResultTags(`group:"controller"`),
	)
}
