package microserviceStarter

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	helpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/logger"

	dummyHandlers "github.com/safciplak/microservice-starter/src/business/dummy/handlers"
	healthHandlers "github.com/safciplak/microservice-starter/src/business/health/handlers"
)

// TestSuite is the suite for testing
type TestSuite struct {
	suite.Suite
	routes *Routes
	health *healthHandlers.MockInterfaceHealthHandler
	dummy  *dummyHandlers.MockInterfaceDummyHandler
}

// SetupTest sets up the test
func (test *TestSuite) SetupTest() {
	test.health = &healthHandlers.MockInterfaceHealthHandler{}
	test.dummy = &dummyHandlers.MockInterfaceDummyHandler{}
}

// TearDownSuite tears down the test
func (test *TestSuite) TearDownTest() {
	test.dummy.AssertExpectations(test.T())
	test.health.AssertExpectations(test.T())
}

// TestTestSuite runs the test suite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestNewRoutes Tests whether the correct handler calls are registered
func (test *TestSuite) TestNewRoutes() {
	log := logger.NewLogger(helpers.NewEnvironmentHelper())
	result := gin.Logger()

	test.health.On("List").Return(result).Once()

	test.dummy.On("List").Return(result).Once()
	test.dummy.On("Read").Return(result).Once()
	test.dummy.On("Create").Return(result).Once()
	test.dummy.On("Update").Return(result).Once()
	test.dummy.On("Delete").Return(result).Once()
	test.routes = NewRoutes(log, test.health, test.dummy)
}
