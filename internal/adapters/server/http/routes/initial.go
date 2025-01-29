package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liquor/framework/internal/adapters/server/http/handlers"
)

func InitialRoutes(r *gin.Engine, initialHandler *handlers.InitialHandler) {
	group := r.Group("/initial")
	{
		group.GET("/", initialHandler.Example)
	}
}
