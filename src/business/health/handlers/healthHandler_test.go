package healthHandlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/safciplak/capila/src/http/response"

	healthServices "github.com/safciplak/microservice-starter/src/business/health/services"
)

// setup sets up variables used in multiple tests
func setup() (
	*gin.Engine,
	*httptest.ResponseRecorder,
	context.Context,
	*healthServices.MockInterfaceHealthService) {
	return gin.New(), httptest.NewRecorder(), context.Background(),
		&healthServices.MockInterfaceHealthService{}
}

// TestIndex checks whether the health endpoint works properly
func TestIndex(t *testing.T) {
	var (
		router, recorder, requestContext, service = setup()
		handler                                   = NewHealthHandler(service)
		expectedResponse                          = `{"data":null,"_links":{"self":{"href":"/v1/health"}},"meta":{"query":{}}}`
		httpResponse                              = response.Create()
	)

	service.On("Health", requestContext).Return(httpResponse, nil).Once()

	router.GET("/v1/health", handler.List())

	request, err := http.NewRequestWithContext(requestContext, "GET", "/v1/health", nil)
	router.ServeHTTP(recorder, request)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

// TestIndexError tests the List handler func
func TestIndexError(t *testing.T) {
	var (
		router, recorder, requestContext, service = setup()
		handler                                   = NewHealthHandler(service)
		expectedError                             = errors.New("something went wrong")
		expectedResponse                          = `{"data":null,"_links":{"self":{"href":"/v1/health"}},"meta":{"query":{}}}`
	)

	service.On("Health", requestContext).Return(&response.Response{}, expectedError).Once()

	router.GET("/v1/health", handler.List())

	request, err := http.NewRequestWithContext(requestContext, "GET", "/v1/health", nil)
	assert.Nil(t, err)

	router.ServeHTTP(recorder, request)

	// TODO: when an error is thrown the result code should not be 200
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}
