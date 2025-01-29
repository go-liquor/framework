package main

import (
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
)

func main() {
	app.NewApp(
		http.Server,
		app.RegisterService(
			services.NewInitialService,
		),
	)
}
