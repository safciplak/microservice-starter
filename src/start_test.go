package microserviceStarter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewService should return a valid service
func TestNewService(t *testing.T) {
	result := NewService(nil)
	assert.NotNil(t, result)
}
