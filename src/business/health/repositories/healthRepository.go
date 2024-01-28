//go:generate generate-interfaces.sh

package healthRepositories

import (
	"context"

	"github.com/safciplak/capila/src/apm"
	"github.com/safciplak/capila/src/database"

	"github.com/safciplak/microservice-starter/src/models"
)

// HealthRepository is used for future purposes as a receiver
type HealthRepository struct {
	database *database.Connection
}

// NewHealthRepository constructs a new HealthRepository
func NewHealthRepository(db *database.Connection) InterfaceHealthRepository {
	return &HealthRepository{
		database: db,
	}
}

// GetLastMigration gets the last migration
func (repository *HealthRepository) GetLastMigration(ctx context.Context) (*models.SchemaMigrations, error) {
	defer apm.End(apm.Start(ctx, "HealthHandler.Get", "handler"))

	migration := new(models.SchemaMigrations)

	err := repository.database.Read.WithContext(ctx).
		Model(migration).
		Select()

	return migration, err
}
