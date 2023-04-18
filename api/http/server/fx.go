package server

import "go.uber.org/fx"

var Option = fx.Options(
	fx.Provide(
		NewServer,
		fx.Annotate(
			NewServeMux,
			fx.ParamTags(`name:"echo"`, `name:"hello"`),
		),
	),
)
