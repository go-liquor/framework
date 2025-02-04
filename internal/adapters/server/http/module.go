package http

import (
	"github.com/go-liquor/framework/internal/adapters/server/http/handlers"
	"github.com/go-liquor/framework/internal/adapters/server/http/routes"
	httpsdk "github.com/go-liquor/liquor-sdk/server/http"
)

var Server = httpsdk.StartModule(&httpsdk.HttpModuleParam{
	Handlers: httpsdk.Handler{
		handlers.NewInitialHandler,
	},
	Routes: httpsdk.Route{
		routes.InitialRoutes,
	},
})
