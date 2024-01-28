//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/safciplak/capila/src/database"
	helpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/http/context"
	"github.com/safciplak/capila/src/logger"

	microserviceStarter "github.com/safciplak/microservice-starter/src"
	dummyHandlers "github.com/safciplak/microservice-starter/src/business/dummy/handlers"
	dummyRepositories "github.com/safciplak/microservice-starter/src/business/dummy/repositories"
	dummyServices "github.com/safciplak/microservice-starter/src/business/dummy/services"
	healthHandlers "github.com/safciplak/microservice-starter/src/business/health/handlers"
	healthRepositories "github.com/safciplak/microservice-starter/src/business/health/repositories"
	healthServices "github.com/safciplak/microservice-starter/src/business/health/services"
)

// InitializeService specifies which dependencies need to be initialized for the service to start
func InitializeService() (_ *microserviceStarter.Service) {
	wire.Build(
		microserviceStarter.NewService,
		microserviceStarter.NewRoutes,
		database.NewDatabase,
		context.NewContext,
		helpers.NewEnvironmentHelper,
		logger.NewLogger,
		healthHandlers.NewHealthHandler,
		healthServices.NewHealthService,
		healthRepositories.NewHealthRepository,
		dummyHandlers.NewDummyHandler,
		dummyServices.NewDummyService,
		dummyRepositories.NewDummyRepository,
	)
	return
}
