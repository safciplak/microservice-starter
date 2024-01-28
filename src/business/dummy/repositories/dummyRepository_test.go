package dummyRepositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/database"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	"github.com/safciplak/microservice-starter/src/models"
)

// Test Suite which encapsulate the tests for the repository
type TestSuite struct {
	suite.Suite

	ctx   context.Context
	dummy *models.Dummy
	now   time.Time

	readConn  *database.MockInterfacePGDB
	writeConn *database.MockInterfacePGDB
	query     *database.MockInterfaceORMQuery

	repository InterfaceDummyRepository
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	// Mocks used in the test
	test.readConn = new(database.MockInterfacePGDB)
	test.writeConn = new(database.MockInterfacePGDB)
	test.query = new(database.MockInterfaceORMQuery)

	// Often used test objects
	test.dummy = getDummyStruct()
	test.now = time.Now().UTC()
	test.ctx = context.TODO()

	// Object to be tested
	test.repository = NewDummyRepository(
		&database.Connection{
			Read:  test.readConn,
			Write: test.writeConn,
		},
	)
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.readConn.AssertExpectations(test.T())
	test.writeConn.AssertExpectations(test.T())
	test.query.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// geDummyStruct builds a test example of a dummy object
func getDummyStruct() *models.Dummy {
	return &models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
		Name: "Microservice-Dummy",
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	var (
		expectedResult = make([]models.Dummy, 0)
		dummies        = make([]models.Dummy, 0)
	)

	searchParams := &dummyModels.QueryParams{
		Name: "Microservice-Dummy",
	}

	dummy1 := getDummyStruct()
	dummy2 := getDummyStruct()
	dummy2.Name = "Microservice-Dummy2"
	expectedResult = append(expectedResult, *dummy1, *dummy2)

	test.readConn.On("ModelContext", test.ctx, &dummies).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]models.Dummy)
			*arg = expectedResult
		}).
		Return(test.query).
		Once()
	test.query.On("Where", "dummy.name LIKE ?", searchParams.Name+"%").Return(test.query).Once()
	test.query.On("Where", "dummy.isdeleted = false").Return(test.query).Once()
	test.query.On("Select").Return(nil).Once()

	data, err := test.repository.List(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(2, len(data))
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	searchParams := &dummyModels.QueryParams{
		GUID: test.dummy.GUID,
	}

	test.readConn.On("ModelContext", test.ctx, &models.Dummy{}).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Dummy)
			*arg = *test.dummy
		}).
		Return(test.query).
		Once()
	test.query.On("Where", "dummy.guid = ?", searchParams.GUID).Return(test.query).Once()
	test.query.On("Where", "dummy.isdeleted = false").Return(test.query).Once()
	test.query.On("First").Return(nil).Once()

	data, err := test.repository.Read(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(test.dummy.Name, data.Name)
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	test.writeConn.On("ModelContext", test.ctx, test.dummy).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.Dummy)
			*arg = *test.dummy
			arg.CreatedAt = test.now.String()
		}).Return(test.query).Once()
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Insert").Return(nil, nil).Once()

	err := test.repository.Create(test.ctx, test.dummy)

	test.Nil(err)
	test.Equal(test.now.String(), test.dummy.CreatedAt)
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	test.writeConn.On("ModelContext", test.ctx, test.dummy).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*models.Dummy)
		*arg = *test.dummy
		arg.UpdatedAt = test.now.String()
	}).Return(test.query).Once()
	test.query.On("Column", "name").Return(test.query).Once()
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Where", "dummy.guid = ?", test.dummy.GUID).Return(test.query).Once()
	test.query.On("Where", "dummy.isdeleted = false").Return(test.query).Once()
	test.query.On("Update").Return(nil, nil).Once()

	err := test.repository.Update(test.ctx, test.dummy)

	test.Nil(err)
	test.Equal(test.now.String(), test.dummy.UpdatedAt)
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	test.writeConn.On("ModelContext", test.ctx, test.dummy).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*models.Dummy)
		*arg = *test.dummy
		arg.IsDeleted = true
		arg.UpdatedAt = test.now.String()
	}).Return(test.query).Once()
	test.query.On("Column", "isdeleted").Return(test.query).Once()
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Where", "dummy.guid = ?", test.dummy.GUID).Return(test.query).Once()
	test.query.On("Where", "dummy.isdeleted = false").Return(test.query).Once()
	test.query.On("Update").Return(nil, nil).Once()

	err := test.repository.Delete(test.ctx, test.dummy)

	test.Nil(err)
	test.Equal(test.now.String(), test.dummy.UpdatedAt)
}
