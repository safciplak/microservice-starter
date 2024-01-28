package example

import (
	"context"

	"github.com/safciplak/capila/src/apm"
)

// Repository is the repository struct for example
type Repository struct{}

// Initialize makes sure the repository is seeded with the correct settings
func (repository *Repository) Initialize() error {
	defer apm.End(apm.Start(context.Background(), "Repository.Initialize", "repository"))

	return nil
}

// GetName helps to identify the provider
func (repository Repository) GetName() string {
	return "Example"
}
