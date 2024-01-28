//go:generate generate-interfaces.sh

package healthHandlers

import (
	"github.com/gin-gonic/gin"

	healthServices "github.com/safciplak/microservice-starter/src/business/health/services"
)

// HealthHandler is the health handler receiver
type HealthHandler struct {
	service healthServices.InterfaceHealthService
}

// NewHealthHandler plop
func NewHealthHandler(service healthServices.InterfaceHealthService) InterfaceHealthHandler {
	return &HealthHandler{
		service,
	}
}

// List returns the health response
func (handler HealthHandler) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			requestCtx = ctx.Request.Context()
			response   = handler.service.Health(requestCtx)
		)

		response.ReturnJSON(ctx)
	}
}
