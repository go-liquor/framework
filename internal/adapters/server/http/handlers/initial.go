package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-liquor/framework/internal/app/services"
)

type InitialHandler struct {
	svc *services.InitialService
}

func NewInitialHandler(svc *services.InitialService) *InitialHandler {
	return &InitialHandler{
		svc: svc,
	}
}

func (i *InitialHandler) Example(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": i.svc.Initial(c.Request.Context()),
	})
}
