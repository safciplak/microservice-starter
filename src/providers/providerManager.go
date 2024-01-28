//go:generate generate-interfaces.sh

// Package providers contains an abstraction layer over external API's so this API can talk to multiple
// external sources over the same interface
package providers

import (
	"errors"

	"github.com/safciplak/microservice-starter/src/providers/example"
)

// ProviderManager contains a map with all the active providers
type ProviderManager struct {
	services map[string]InterfaceProviderService
}

// NewProviders Creates a new ProviderManager and registers the provided repos
func NewProviders(provider *example.Repository) *ProviderManager {
	var (
		providers = new(ProviderManager)
	)

	providers.services = make(map[string]InterfaceProviderService)

	// Register your provider here
	providers.services[provider.GetName()] = provider

	return providers
}

// GetProvider returns the provider by the given name
func (providers *ProviderManager) GetProvider(name string) (InterfaceProviderService, error) {
	if provider, ok := providers.services[name]; ok {
		return provider, nil
	}

	return nil, errors.New("unknown provider")
}
