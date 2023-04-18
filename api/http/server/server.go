package server

import (
	"context"
	"net"
	"net/http"

	"github.com/inhibitor1217/go-fx-playground/internal/handler"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func NewServeMux(routes []handler.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.Pattern(), r)
	}
	return mux
}
