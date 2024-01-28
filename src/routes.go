package microserviceStarter

import (
	"github.com/gin-gonic/gin"

	middlewares "github.com/safciplak/capila/src/http/middleware"
	"github.com/safciplak/capila/src/http/web"
	"github.com/safciplak/capila/src/logger"

	dummyHandlers "github.com/safciplak/microservice-starter/src/business/dummy/handlers"
	healthHandlers "github.com/safciplak/microservice-starter/src/business/health/handlers"
)

// Routes contains the logic to route traffic to the microservice
type Routes struct {
	router *gin.Engine
}

// NewRoutes Initializes the router and routes and allows the application to serve the routes
func NewRoutes(
	log *logger.Logger,
	healthHandler healthHandlers.InterfaceHealthHandler,
	dummyHandler dummyHandlers.InterfaceDummyHandler,
) (routes *Routes) {
	routes = new(Routes)
	routes.router = web.Router(log)
	routes.router.Use(middlewares.LanguageMiddleware())

	v1 := routes.router.Group("/v1")
	resource := "/:guid"

	// Health endpoints
	health := v1.Group("/health")
	health.GET("", healthHandler.List())

	// Dummy endpoints
	dummy := v1.Group("/dummies")
	dummy.GET("", dummyHandler.List())
	dummy.POST("", dummyHandler.Create())
	dummy.GET(resource, dummyHandler.Read())
	dummy.PUT(resource, dummyHandler.Update())
	dummy.DELETE(resource, dummyHandler.Delete())

	return
}

// Serve serves the router and the endpoints it encapsulates
func (routes *Routes) Serve() {
	web.Server("8282", routes.router)
}
