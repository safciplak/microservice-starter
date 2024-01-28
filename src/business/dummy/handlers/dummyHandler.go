//go:generate generate-interfaces.sh

package dummyHandlers

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/convert"
	capilaContext "github.com/safciplak/capila/src/http/context"
	"github.com/safciplak/capila/src/http/handlers"

	dummyModels "github.com/safciplak/microservice-starter/src/business/dummy/models"
	dummyServices "github.com/safciplak/microservice-starter/src/business/dummy/services"
)

// DummyHandler is the receiver for the handler functions
type DummyHandler struct {
	service dummyServices.InterfaceDummyService
}

// NewDummyHandler initializes the handler.
func NewDummyHandler(service dummyServices.InterfaceDummyService) InterfaceDummyHandler {
	return &DummyHandler{
		service,
	}
}

// List returns a list of Dummy entities
func (handler *DummyHandler) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dummyModels.ListRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.List(ctx, request)
		})
	}
}

// Read returns a single Dummy entity
func (handler *DummyHandler) Read() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dummyModels.BaseRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Read(ctx, request)
		})
	}
}

// Create creates a new Dummy entity
func (handler *DummyHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dummyModels.CreateRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Create(ctx, request)
		})
	}
}

// Update updates a existing Dummy entity
func (handler *DummyHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dummyModels.UpdateRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Update(ctx, request)
		})
	}
}

// Delete (soft)deletes a existing Dummy entity
func (handler *DummyHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &dummyModels.BaseRequest{
			Language: convert.NewString(strings.ToUpper(capilaContext.GetTwoLetterLanguageCode(ctx.Request.Context()))),
		}

		handlers.GetHandlerFunc(ctx, request, func(ctx context.Context) (interface{}, error) {
			return handler.service.Delete(ctx, request)
		})
	}
}
