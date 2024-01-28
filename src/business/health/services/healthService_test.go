package healthServices

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	capilaHelpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/http/response"

	healthRepositories "github.com/safciplak/microservice-starter/src/business/health/repositories"
	"github.com/safciplak/microservice-starter/src/models"
)

// setup sets up variables used in multiple tests
func setup() (*healthRepositories.MockInterfaceHealthRepository, *capilaHelpers.MockInterfaceEnvironmentHelper) {
	var (
		repository  = &healthRepositories.MockInterfaceHealthRepository{}
		environment = &capilaHelpers.MockInterfaceEnvironmentHelper{}
	)

	_ = os.Setenv("DEPLOYED_AT", "right now")

	return repository, environment
}

func teardown() {
	_ = os.Unsetenv("DEPLOYED_AT")
}

// TestHealth tests the health method
func TestHealth(t *testing.T) {
	var (
		repository, env = setup()
		migration       = models.SchemaMigrations{
			Version:   1,
			Dirty:     false,
			UpdatedAt: "yesterday",
		}
		service          = NewHealthService(env, repository)
		expectedResponse = response.Response{
			StatusCode:       200,
			Error:            nil,
			ValidationErrors: nil,
			Data: struct {
				Success    bool                    `json:"success"`
				DeployedAt string                  `json:"deployedAt"`
				Migration  models.SchemaMigrations `json:"migration"`
			}{
				true,
				"right now",
				migration,
			},
		}
		ctx = context.Background()
	)

	env.On("Get", "DEPLOYED_AT").Return("right now").Once()
	repository.On("GetLastMigration", ctx).Return(&migration, nil).Once()

	httpResponse := service.Health(ctx)

	// Check if the call has replaced the service
	assert.Equal(t, &expectedResponse, httpResponse)
	teardown()
}
