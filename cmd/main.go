package main

import (
	"net/http"

	httpApi "github.com/inhibitor1217/go-fx-playground/api/http"
	log "github.com/inhibitor1217/go-fx-playground/internal/log"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		httpApi.Option,
		log.Option,
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
