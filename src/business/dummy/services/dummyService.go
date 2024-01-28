//go:generate generate-interfaces.sh

package dummyServices

import (
	"context"

	"github.com/safciplak/capila/src/apm"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	dummyRepositories "github.com/safciplak/microservice-starter/src/business/dummy/repositories"
	"github.com/safciplak/microservice-starter/src/models"
)

// DummyService contains the necessary repositories.
type DummyService struct {
	Repository dummyRepositories.InterfaceDummyRepository
}

// NewDummyService instantiates a new DummyService
func NewDummyService(repository dummyRepositories.InterfaceDummyRepository) InterfaceDummyService {
	return &DummyService{
		Repository: repository,
	}
}

// List returns a slice of Dummy entities matching the given params
func (service *DummyService) List(ctx context.Context, params *dummyModels.ListRequest) ([]models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyService.List", "service"))

	queryParams := params.ToQueryParams()

	return service.Repository.List(ctx, queryParams)
}

// Read returns a single Dummy entity matching the params
func (service *DummyService) Read(ctx context.Context, params *dummyModels.BaseRequest) (*models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyService.Read", "service"))

	queryParams := params.ToQueryParams()

	return service.Repository.Read(ctx, queryParams)
}

// Create creates a new Dummy entity
func (service *DummyService) Create(ctx context.Context, params *dummyModels.CreateRequest) (*models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyService.Create", "service"))

	dummy := params.ToModel()
	err := service.Repository.Create(ctx, dummy)

	return dummy, err
}

// Update updates a existing Dummy entity
func (service *DummyService) Update(ctx context.Context, params *dummyModels.UpdateRequest) (*models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyService.Update", "service"))

	dummy := params.ToModel()
	err := service.Repository.Update(ctx, dummy)

	return dummy, err
}

// Delete (soft)deletes a existing Dummy entity
func (service *DummyService) Delete(ctx context.Context, params *dummyModels.BaseRequest) (*models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyService.Delete", "service"))

	dummy := params.ToModel(true)
	err := service.Repository.Delete(ctx, dummy)

	return dummy, err
}
