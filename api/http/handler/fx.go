package handler

import (
	"github.com/inhibitor1217/go-fx-playground/api/http/handler/echo"

	"go.uber.org/fx"
)

var Option = fx.Provide(
	echo.NewEchoHandler,
	echo.NewServeMux,
)
