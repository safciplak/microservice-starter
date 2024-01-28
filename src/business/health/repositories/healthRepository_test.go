package healthRepositories

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/safciplak/capila/src/database"

	"github.com/safciplak/microservice-starter/src/models"
)

// setup sets up variables used in multiple tests
func setup() (
	*database.MockInterfacePGDB,
	*database.MockInterfaceORMQuery,
	*database.Connection) {
	var (
		mockRead   = &database.MockInterfacePGDB{}
		mockQuery  = &database.MockInterfaceORMQuery{}
		connection = &database.Connection{
			Read: mockRead,
		}
	)

	return mockRead, mockQuery, connection
}

// TestGetLastMigration tests the GetLastMigration method
func TestGetLastMigration(t *testing.T) {
	var (
		mockRead, mockQuery, connection = setup()
		repository                      = NewHealthRepository(connection)
		expectedResult                  = models.SchemaMigrations{
			Version:   10,
			Dirty:     true,
			UpdatedAt: "plop",
		}
		ctx = context.Background()
	)

	mockRead.On("WithContext", mock.Anything).Return(mockRead)
	mockRead.On("Model", &models.SchemaMigrations{}).Run(func(args mock.Arguments) {
		arg, ok := args.Get(0).(*models.SchemaMigrations)
		assert.True(t, ok)
		*arg = expectedResult
	}).Return(mockQuery)
	mockQuery.On("Select").Return(nil)

	result, err := repository.GetLastMigration(ctx)
	assert.Nil(t, err)
	assert.Equal(t, &expectedResult, result)
}

// TestGetLastMigrationError tests the GetLastMigration method when an error is encountered
func TestGetLastMigrationError(t *testing.T) {
	var (
		mockRead, mockQuery, connection = setup()
		repository                      = NewHealthRepository(connection)
		expectedResult                  = models.SchemaMigrations{}
		expectedError                   = errors.New("test error")
		ctx                             = context.Background()
	)

	mockRead.On("WithContext", mock.Anything).Return(mockRead)
	mockRead.On("Model", &models.SchemaMigrations{}).Return(mockQuery)
	mockQuery.On("Select").Return(expectedError)

	result, err := repository.GetLastMigration(ctx)

	assert.Equal(t, expectedError, err)
	assert.Equal(t, &expectedResult, result)
}
