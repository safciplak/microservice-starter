package dummyHandlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/convert"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	dummyServices "github.com/safciplak/microservice-starter/src/business/dummy/services"
	"github.com/safciplak/microservice-starter/src/models"
)

// Test Suite which encapsulate the tests for the handler!
type TestSuite struct {
	suite.Suite
	ctx      context.Context
	router   *gin.Engine
	recorder *httptest.ResponseRecorder

	handler InterfaceDummyHandler
	service *dummyServices.MockInterfaceDummyService

	dummy *models.Dummy
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)

	test.ctx = context.TODO()
	test.router = gin.New()
	test.recorder = httptest.NewRecorder()
	test.service = new(dummyServices.MockInterfaceDummyService)
	test.handler = NewDummyHandler(test.service)

	test.dummy = &models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
		Name: "Microservice-Dummy",
	}

	// Register the routes just like in the routes.go
	test.router.GET("/dummies", test.handler.List())
	test.router.POST("/dummies", test.handler.Create())
	test.router.GET("/dummies/:guid", test.handler.Read())
	test.router.PUT("/dummies/:guid", test.handler.Update())
	test.router.DELETE("/dummies/:guid", test.handler.Delete())
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.service.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// geDummyStruct builds a test example of a dummy object
func getDummyStruct() models.Dummy {
	return models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dd",
		},
		Name: "Microservice-Dummy",
	}
}

// getDummyRequestBody returns a dummy JSON request
func getDummyRequestBody() string {
	return `
		{
			"guid": "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			"name": "Microservice-Dummy"
		}
	`
}

// getDummyInvalidRequestBody returns a dummy JSON request with invalid validation
func getDummyInvalidRequestBody() string {
	return `
		{
			"guid": "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			"name": "AZ"
		}
	`
}

// getDummyCreateRequestStruct builds a test example of a dummy create object
func getDummyCreateRequestStruct() *dummyModels.CreateRequest {
	return &dummyModels.CreateRequest{
		GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		Name:     "Microservice-Dummy",
		Language: convert.NewString("EN"),
	}
}

// getDummyUpdateRequestStruct builds a test example of a dummy update object
func getDummyUpdateRequestStruct() *dummyModels.UpdateRequest {
	return &dummyModels.UpdateRequest{
		GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		Name:     "Microservice-Dummy",
		Language: convert.NewString("EN"),
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	expectedResult := make([]models.Dummy, 0)
	dummy1 := getDummyStruct()
	dummy2 := getDummyStruct()
	dummy2.Name = "Microservice-Dummy2"
	expectedResult = append(expectedResult, dummy1, dummy2)

	searchParams := dummyModels.ListRequest{
		Name:     "Microservice-Dummy",
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies?name="+searchParams.Name, nil)

	// No routing error should be thrown
	test.Nil(err)

	test.service.On("List", test.ctx, &searchParams).Return(expectedResult, nil).Once()
	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestReadValidationError tests the validation handling of the Read function
func (test *TestSuite) TestListValidationError() {
	searchParams := dummyModels.ListRequest{
		Name:     "AZ",
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies?name="+searchParams.Name, nil)

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String())
}

// TestListServiceError tests the service error handling of the List function
func (test *TestSuite) TestListServiceError() {
	searchParams := dummyModels.ListRequest{
		Name:     "Microservice-Dummy",
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies?name="+searchParams.Name, nil)

	// No routing error should be thrown
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.On("List", test.ctx, &searchParams).Return(nil, expectedError).Once()
	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	searchParams := dummyModels.BaseRequest{
		GUID:     test.dummy.GUID,
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies/"+test.dummy.GUID, nil)

	// No routing error should be thrown
	test.Nil(err)

	test.service.
		On("Read", test.ctx, &searchParams).
		Return(test.dummy, nil).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestReadValidationError tests the validation handling of the Read function
func (test *TestSuite) TestReadValidationError() {
	searchParams := dummyModels.BaseRequest{
		GUID: "WRONG-GUID",
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies/"+searchParams.GUID, nil)

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestReadServiceError tests the service error handling of the Read function
func (test *TestSuite) TestReadServiceError() {
	searchParams := dummyModels.BaseRequest{
		GUID:     test.dummy.GUID,
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/dummies/"+searchParams.GUID, nil)

	// No routing error should be thrown
	test.Nil(err)

	expectedError := errors.New("unknown error")

	test.service.
		On("Read", test.ctx, &searchParams).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	requestBody := getDummyRequestBody()
	dummy := getDummyCreateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/dummies", bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	test.service.
		On("Create", test.ctx, dummy).
		Return(test.dummy, nil).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestCreateValidationError tests the validation handling of the Create function
func (test *TestSuite) TestCreateValidationError() {
	requestBody := getDummyInvalidRequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/dummies", bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String())
}

// TestCreateServiceError tests the service error handling of the Create function
func (test *TestSuite) TestCreateServiceError() {
	requestBody := getDummyRequestBody()
	dummy := getDummyCreateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/dummies", bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.
		On("Create", test.ctx, dummy).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	requestBody := getDummyRequestBody()
	dummy := getDummyUpdateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/dummies/"+test.dummy.GUID, bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	test.service.
		On("Update", test.ctx, dummy).
		Return(test.dummy, nil).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestUpdateValidationError tests the validation handling of the Update function
func (test *TestSuite) TestUpdateValidationError() {
	requestBody := getDummyInvalidRequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/dummies/"+test.dummy.GUID, bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String())
}

// TestUpdateValidationErrorWrongGUID tests the validation handling of the Update function
func (test *TestSuite) TestUpdateValidationErrorWrongGUID() {
	test.dummy.GUID = "WRONG-GUID"
	requestBody := getDummyRequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/dummies/"+test.dummy.GUID, bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestUpdateServiceError tests the service error handling of the Update function
func (test *TestSuite) TestUpdateServiceError() {
	requestBody := getDummyRequestBody()
	dummy := getDummyUpdateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/dummies/"+test.dummy.GUID, bytes.NewBufferString(requestBody))

	// No routing error should be thrown
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.
		On("Update", test.ctx, dummy).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	test.dummy.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/dummies/"+test.dummy.GUID, nil)
	test.Nil(err)

	test.service.
		On("Delete", test.ctx, &dummyModels.BaseRequest{
			GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			Language: convert.NewString("EN"),
		}).
		Return(test.dummy, nil).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestDeleteValidationError tests the validation handling of the Delete function
func (test *TestSuite) TestDeleteValidationError() {
	test.dummy.GUID = "WRONG-GUID"
	test.dummy.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/dummies/"+test.dummy.GUID, nil)

	// No routing error should be thrown
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestDeleteServiceError tests the service error handling of the Delete function
func (test *TestSuite) TestDeleteServiceError() {
	test.dummy.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/dummies/"+test.dummy.GUID, nil)
	test.Nil(err)

	expectedError := errors.New("unknown error")

	test.service.
		On("Delete", test.ctx, &dummyModels.BaseRequest{
			GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			Language: convert.NewString("EN"),
		}).
		Return(test.dummy, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}
