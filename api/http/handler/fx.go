package handler

import (
	"github.com/inhibitor1217/go-fx-playground/internal/handler"

	"go.uber.org/fx"
)

var Option = fx.Provide(
	fx.Annotate(
		NewEchoHandler,
		fx.As(new(handler.Route)),
		fx.ResultTags(`name:"echo"`),
	),
	fx.Annotate(
		NewHelloHandler,
		fx.As(new(handler.Route)),
		fx.ResultTags(`name:"hello"`),
	),
)
