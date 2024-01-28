//go:generate generate-interfaces.sh

package dummyRepositories

import (
	"context"

	"github.com/safciplak/capila/src/apm"

	"github.com/safciplak/capila/src/database"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	"github.com/safciplak/microservice-starter/src/models"
)

// DummyRepository is used for future purposes as a receiver
type DummyRepository struct {
	database *database.Connection
}

// NewDummyRepository initializes the service. database *database.Connection
func NewDummyRepository(db *database.Connection) InterfaceDummyRepository {
	return &DummyRepository{
		database: db,
	}
}

// List returns a slice of Dummy entities matching the given params
func (repo *DummyRepository) List(ctx context.Context, params *dummyModels.QueryParams) ([]models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyRepository.List", "repository"))

	dummies := make([]models.Dummy, 0)
	err := repo.database.Read.ModelContext(ctx, &dummies).
		Where("dummy.name LIKE ?", params.Name+"%").
		Where("dummy.isdeleted = false").
		Select()

	return dummies, err
}

// Read returns a single Dummy entity matching the params
func (repo *DummyRepository) Read(ctx context.Context, params *dummyModels.QueryParams) (*models.Dummy, error) {
	defer apm.End(apm.Start(ctx, "DummyRepository.Read", "repository"))

	dummy := new(models.Dummy)
	err := repo.database.Read.ModelContext(ctx, dummy).
		Where("dummy.guid = ?", params.GUID).
		Where("dummy.isdeleted = false").
		First()

	return dummy, err
}

// Create creates a new Dummy entity
func (repo *DummyRepository) Create(ctx context.Context, dummy *models.Dummy) error {
	defer apm.End(apm.Start(ctx, "DummyRepository.Create", "repository"))

	_, err := repo.database.Write.ModelContext(ctx, dummy).
		Returning(dummy.BaseTableModel.Returning()).
		Insert()

	return err
}

// Update updates a existing Dummy entity
func (repo *DummyRepository) Update(ctx context.Context, dummy *models.Dummy) error {
	defer apm.End(apm.Start(ctx, "DummyRepository.Update", "repository"))

	_, err := repo.database.Write.ModelContext(ctx, dummy).
		Column("name").
		Returning(dummy.BaseTableModel.Returning()).
		Where("dummy.guid = ?", dummy.GUID).
		Where("dummy.isdeleted = false").
		Update()

	return err
}

// Delete (soft)deletes a existing Dummy entity
func (repo *DummyRepository) Delete(ctx context.Context, dummy *models.Dummy) error {
	defer apm.End(apm.Start(ctx, "DummyRepository.Delete", "repository"))

	dummy.IsDeleted = true

	_, err := repo.database.Write.ModelContext(ctx, dummy).
		Column("isdeleted").
		Returning(dummy.BaseTableModel.Returning()).
		Where("dummy.guid = ?", dummy.GUID).
		Where("dummy.isdeleted = false").
		Update()

	return err
}
