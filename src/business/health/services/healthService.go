//go:generate generate-interfaces.sh

package healthServices

import (
	"context"

	"github.com/safciplak/capila/src/apm"
	capilaHelpers "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/http/response"

	healthRepositories "github.com/safciplak/microservice-starter/src/business/health/repositories"
	"github.com/safciplak/microservice-starter/src/models"
)

// HealthService contains the necessary repositories.
type HealthService struct {
	health            healthRepositories.InterfaceHealthRepository
	environmentHelper capilaHelpers.InterfaceEnvironmentHelper
}

// NewHealthService constructs a new health service
func NewHealthService(env capilaHelpers.InterfaceEnvironmentHelper, repository healthRepositories.InterfaceHealthRepository) InterfaceHealthService {
	return &HealthService{
		health:            repository,
		environmentHelper: env,
	}
}

// Health returns the Health of the the microservice
func (healthService *HealthService) Health(ctx context.Context) *response.Response {
	defer apm.End(apm.Start(ctx, "HealthService.Health", "service"))

	healthResponse := response.Create()

	// TODO: VP-4864 create a proper type instead of an anonymous one then reuse it in the test
	data := struct {
		Success    bool                    `json:"success"`
		DeployedAt string                  `json:"deployedAt"`
		Migration  models.SchemaMigrations `json:"migration"`
	}{}

	data.Success = false
	data.DeployedAt = healthService.environmentHelper.Get("DEPLOYED_AT")

	migration, err := healthService.health.GetLastMigration(ctx)
	if err == nil {
		data.Success = true
		data.Migration = *migration
	}

	healthResponse.Data = data

	return healthResponse
}
