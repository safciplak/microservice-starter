package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/safciplak/microservice-starter/src/providers/example"
)

// TestGetProviderExample Tests whether the registered Example provider can be found
func TestGetProviderExample(t *testing.T) {
	providerManager := NewProviders(&example.Repository{})

	provider, expectedError := providerManager.GetProvider("Example")
	assert.Nil(t, expectedError)

	realProvider, ok := provider.(*example.Repository)
	assert.Equal(t, true, ok)
	assert.Equal(t, "Example", realProvider.GetName())
}

// TestGetProviderNotFound Tests whether the error is thrown successfully
func TestGetProviderNotFound(t *testing.T) {
	providerManager := NewProviders(&example.Repository{})

	provider, expectedError := providerManager.GetProvider("")

	assert.Equal(t, provider, nil)
	assert.EqualError(t, expectedError, "unknown provider")
}
