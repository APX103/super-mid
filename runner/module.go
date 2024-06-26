package runner

import (
	"apx103.com/super-mid/command/cmd"
	"apx103.com/super-mid/runner/jenkins"
	"go.uber.org/fx"
)

func AsRunner(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(cmd.TaskRunner)),
		fx.ResultTags(`group:"task_runner"`),
	)
}

var Module = fx.Module("Command",
	fx.Provide(
		AsRunner(jenkins.NewJenkinsBuildRunner),
	),
)