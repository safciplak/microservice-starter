package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRepository_Initialize tests initialize
func TestRepository_Initialize(t *testing.T) {
	var (
		repository = Repository{}
		err        error
	)

	err = repository.Initialize()

	assert.Nil(t, err)
}

// TestRepository_GetName tests getting the name
func TestRepository_GetName(t *testing.T) {
	repository := Repository{}

	name := repository.GetName()

	assert.Equal(t, `Example`, name)
}
