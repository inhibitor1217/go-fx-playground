package main

import (
	"net/http"

	httpApi "github.com/inhibitor1217/go-fx-playground/api/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		httpApi.Option,
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
