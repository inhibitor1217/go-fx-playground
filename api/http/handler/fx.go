package handler

import (
	"github.com/inhibitor1217/go-fx-playground/api/http/handler/echo"
	"github.com/inhibitor1217/go-fx-playground/internal/handler"

	"go.uber.org/fx"
)

var Option = fx.Provide(
	fx.Annotate(
		echo.NewEchoHandler,
		fx.As(new(handler.Route)),
	),
)
