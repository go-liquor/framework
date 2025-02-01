package main

import (
	"github.com/go-liquor/framework/internal/adapters/database/migrations"
	"github.com/go-liquor/framework/internal/adapters/database/repository"
	"github.com/go-liquor/framework/internal/adapters/server/http"
	"github.com/go-liquor/framework/internal/app/services"
	"github.com/go-liquor/liquor/sdk/app"
)

func main() {
	app.NewApp(
		http.Server,
		migrations.Migrations,
		app.RegisterService(
			services.NewInitialService,
		),
		app.RegisterRepositories(
			repository.NewInitialRepository,
			//go:liquor:repositories
		),
	)
}
