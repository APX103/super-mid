package utils

import (
	"apx103.com/super-mid/utils/config"
	"apx103.com/super-mid/utils/mongoc"
	"go.uber.org/fx"
)

var Module = fx.Module("utils",
	fx.Provide(config.NewBaseConfig),
	// db client
	fx.Provide(mongoc.NewMongoClientImpl),
)
