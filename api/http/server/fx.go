package server

import "go.uber.org/fx"

var Option = fx.Option(
	fx.Provide(NewServer),
)
