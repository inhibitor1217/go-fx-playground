package http

import (
	"github.com/inhibitor1217/go-fx-playground/api/http/server"

	"go.uber.org/fx"
)

var Option = fx.Options(
	server.Option,
)
