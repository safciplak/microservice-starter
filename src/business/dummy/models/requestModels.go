package dummyModels

import (
	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/convert"

	"github.com/safciplak/microservice-starter/src/models"
)

// QueryParams contains the allowed query params per request.
type QueryParams struct {
	GUID     string
	Name     string
	Language string
}

// BaseRequest is the validator for default request param check.
type BaseRequest struct {
	Language *string `binding:"omitempty,min=2,max=5"`
	GUID     string  `uri:"guid" binding:"required,uuid"`
}

// Validate makes sure the correct data is being submitted
func (request *BaseRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindUri(request)
}

// ToQueryParams transforms a request to QueryParams
func (request *BaseRequest) ToQueryParams() *QueryParams {
	return &QueryParams{
		GUID:     request.GUID,
		Language: convert.PointerToString(request.Language),
	}
}

// ToModel transforms a request to a model
func (request *BaseRequest) ToModel(isDeleted bool) *models.Dummy {
	return &models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID:      request.GUID,
			IsDeleted: isDeleted,
		},
	}
}

// ListRequest is the validator for list (getAll) requests.
type ListRequest struct {
	Language *string `binding:"omitempty,min=2,max=5"`
	Name     string  `form:"name" json:"name" binding:"required,min=3,max=255"`
}

// Validate makes sure the correct data is being submitted
func (request *ListRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindQuery(request)
}

// ToQueryParams transforms a request to QueryParams
func (request *ListRequest) ToQueryParams() *QueryParams {
	return &QueryParams{
		Name:     request.Name,
		Language: convert.PointerToString(request.Language),
	}
}

// CreateRequest is the validator for create requests.
type CreateRequest struct {
	Language *string `binding:"omitempty,min=2,max=5"`
	GUID     string  `json:"guid" binding:"omitempty,uuid"`
	Name     string  `json:"name" binding:"required,min=3,max=255"`
}

// Validate makes sure the correct data is being submitted
func (request *CreateRequest) Validate(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(request)
}

// ToModel transforms a request to a model
func (request *CreateRequest) ToModel() *models.Dummy {
	return &models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: request.GUID,
		},
		Name: request.Name,
	}
}

// UpdateRequest is the validator for update requests.
type UpdateRequest struct {
	Language *string `binding:"omitempty,min=2,max=5"`
	GUID     string  `json:"-" binding:"required,uuid"`
	Name     string  `json:"name" binding:"required,min=3,max=255"`
}

// Validate makes sure the correct data is being submitted
func (request *UpdateRequest) Validate(ctx *gin.Context) error {
	request.GUID = ctx.Params.ByName("guid")

	return ctx.ShouldBindJSON(request)
}

// ToModel transforms a request to a model
func (request *UpdateRequest) ToModel() *models.Dummy {
	return &models.Dummy{
		BaseTableModel: models.BaseTableModel{
			GUID: request.GUID,
		},
		Name: request.Name,
	}
}
