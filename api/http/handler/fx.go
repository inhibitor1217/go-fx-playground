package handler

import (
	"github.com/inhibitor1217/go-fx-playground/internal/handler"

	"go.uber.org/fx"
)

var Option = fx.Provide(
	handler.AsRoute(NewEchoHandler),
	handler.AsRoute(NewHelloHandler),
)
