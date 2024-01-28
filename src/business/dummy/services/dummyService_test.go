package dummyServices

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	dummyRepositories "github.com/safciplak/microservice-starter/src/business/dummy/repositories"
	"github.com/safciplak/microservice-starter/src/models"
)

// Test Suite which encapsulate the tests for the test.service
type TestSuite struct {
	suite.Suite
	ctx        context.Context
	repository *dummyRepositories.MockInterfaceDummyRepository
	service    InterfaceDummyService
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	// Mocks used in the test
	test.repository = new(dummyRepositories.MockInterfaceDummyRepository)

	// Often used test objects
	test.ctx = context.TODO()

	// Object to be tested
	test.service = NewDummyService(
		test.repository,
	)
}

// TearDownTest tests whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.repository.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// getDummyStruct builds a test example of a dummy object
func getDummyStruct() models.Dummy {
	return models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
		Name: "Microservice-Dummy",
	}
}

// getDummyCreateRequest builds a test example of a dummy create request
func getDummyCreateRequest() dummyModels.CreateRequest {
	return dummyModels.CreateRequest{
		Name: "Microservice-Dummy",
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
	}
}

// getDummyUpdateRequest builds a test example of a dummy update request
func getDummyUpdateRequest() dummyModels.UpdateRequest {
	return dummyModels.UpdateRequest{
		Name: "Microservice-Dummy2",
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dde",
	}
}

// getDummyBaseRequest builds a test example of a dummy base request
func getDummyBaseRequest() dummyModels.BaseRequest {
	return dummyModels.BaseRequest{
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dde",
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	var (
		expectedResult []models.Dummy
		dummies        []models.Dummy
		err            error
	)

	searchParams := &dummyModels.ListRequest{
		Name: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
	}

	dummy1 := getDummyStruct()
	dummy2 := getDummyStruct()
	expectedResult = append(expectedResult, dummy1, dummy2)

	test.repository.
		On("List", test.ctx, searchParams.ToQueryParams()).
		Return(expectedResult, nil).Once()

	dummies, err = test.service.List(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(2, len(dummies))
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	expectedResult := getDummyStruct()
	searchParams := &dummyModels.BaseRequest{
		GUID: expectedResult.GUID,
	}

	test.repository.
		On("Read", test.ctx, searchParams.ToQueryParams()).
		Return(&expectedResult, nil).Once()

	dummy, err := test.service.Read(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(expectedResult.GUID, dummy.GUID)
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	dummyRequest := getDummyCreateRequest()

	test.repository.
		On("Create", test.ctx, dummyRequest.ToModel()).
		Return(nil).Once()

	dummy, err := test.service.Create(test.ctx, &dummyRequest)

	test.Nil(err)
	test.Equal("Microservice-Dummy", dummy.Name)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43ddf", dummy.GUID)
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	dummyRequest := getDummyUpdateRequest()

	test.repository.
		On("Update", test.ctx, dummyRequest.ToModel()).
		Return(nil).Once()

	dummy, err := test.service.Update(test.ctx, &dummyRequest)

	test.Nil(err)
	test.Equal("Microservice-Dummy2", dummy.Name)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43dde", dummy.GUID)
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	dummyRequest := getDummyBaseRequest()

	test.repository.
		On("Delete", test.ctx, dummyRequest.ToModel(true)).
		Return(nil).Once()

	dummy, err := test.service.Delete(test.ctx, &dummyRequest)

	test.Nil(err)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43dde", dummy.GUID)
}
