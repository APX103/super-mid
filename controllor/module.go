package controllor

import (
	"apx103.com/super-mid/controllor/common"
	"apx103.com/super-mid/controllor/demo"
	"apx103.com/super-mid/controllor/task"
	"go.uber.org/fx"
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
		AsController(demo.NewPingController),
		AsController(task.NewTaskController),
	),
)
