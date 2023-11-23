// Code generated by goa v3.14.0, DO NOT EDIT.
//
// catalog HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	catalog "github.com/tektoncd/hub/api/gen/catalog"
	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goa "goa.design/goa/v3/pkg"
)

// RefreshResponseBody is the type of the "catalog" service "Refresh" endpoint
// HTTP response body.
type RefreshResponseBody struct {
	// id of the job
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the catalog
	CatalogName string `form:"catalogName" json:"catalogName" xml:"catalogName"`
	// status of the job
	Status string `form:"status" json:"status" xml:"status"`
}

// RefreshAllResponseBody is the type of the "catalog" service "RefreshAll"
// endpoint HTTP response body.
type RefreshAllResponseBody []*JobResponse

// CatalogErrorResponseBody is the type of the "catalog" service "CatalogError"
// endpoint HTTP response body.
type CatalogErrorResponseBody struct {
	// Catalog errors
	Data []*CatalogErrorsResponseBody `form:"data" json:"data" xml:"data"`
}

// RefreshNotFoundResponseBody is the type of the "catalog" service "Refresh"
// endpoint HTTP response body for the "not-found" error.
type RefreshNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshInternalErrorResponseBody is the type of the "catalog" service
// "Refresh" endpoint HTTP response body for the "internal-error" error.
type RefreshInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshAllInternalErrorResponseBody is the type of the "catalog" service
// "RefreshAll" endpoint HTTP response body for the "internal-error" error.
type RefreshAllInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CatalogErrorInternalErrorResponseBody is the type of the "catalog" service
// "CatalogError" endpoint HTTP response body for the "internal-error" error.
type CatalogErrorInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// JobResponse is used to define fields on response body types.
type JobResponse struct {
	// id of the job
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the catalog
	CatalogName string `form:"catalogName" json:"catalogName" xml:"catalogName"`
	// status of the job
	Status string `form:"status" json:"status" xml:"status"`
}

// CatalogErrorsResponseBody is used to define fields on response body types.
type CatalogErrorsResponseBody struct {
	// Catalog Errror type
	Type string `form:"type" json:"type" xml:"type"`
	// Catalog Error message
	Errors []string `form:"errors" json:"errors" xml:"errors"`
}

// NewRefreshResponseBody builds the HTTP response body from the result of the
// "Refresh" endpoint of the "catalog" service.
func NewRefreshResponseBody(res *catalogviews.JobView) *RefreshResponseBody {
	body := &RefreshResponseBody{
		ID:          *res.ID,
		CatalogName: *res.CatalogName,
		Status:      *res.Status,
	}
	return body
}

// NewRefreshAllResponseBody builds the HTTP response body from the result of
// the "RefreshAll" endpoint of the "catalog" service.
func NewRefreshAllResponseBody(res []*catalog.Job) RefreshAllResponseBody {
	body := make([]*JobResponse, len(res))
	for i, val := range res {
		body[i] = marshalCatalogJobToJobResponse(val)
	}
	return body
}

// NewCatalogErrorResponseBody builds the HTTP response body from the result of
// the "CatalogError" endpoint of the "catalog" service.
func NewCatalogErrorResponseBody(res *catalog.CatalogErrorResult) *CatalogErrorResponseBody {
	body := &CatalogErrorResponseBody{}
	if res.Data != nil {
		body.Data = make([]*CatalogErrorsResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalCatalogCatalogErrorsToCatalogErrorsResponseBody(val)
		}
	} else {
		body.Data = []*CatalogErrorsResponseBody{}
	}
	return body
}

// NewRefreshNotFoundResponseBody builds the HTTP response body from the result
// of the "Refresh" endpoint of the "catalog" service.
func NewRefreshNotFoundResponseBody(res *goa.ServiceError) *RefreshNotFoundResponseBody {
	body := &RefreshNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshInternalErrorResponseBody builds the HTTP response body from the
// result of the "Refresh" endpoint of the "catalog" service.
func NewRefreshInternalErrorResponseBody(res *goa.ServiceError) *RefreshInternalErrorResponseBody {
	body := &RefreshInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshAllInternalErrorResponseBody builds the HTTP response body from
// the result of the "RefreshAll" endpoint of the "catalog" service.
func NewRefreshAllInternalErrorResponseBody(res *goa.ServiceError) *RefreshAllInternalErrorResponseBody {
	body := &RefreshAllInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCatalogErrorInternalErrorResponseBody builds the HTTP response body from
// the result of the "CatalogError" endpoint of the "catalog" service.
func NewCatalogErrorInternalErrorResponseBody(res *goa.ServiceError) *CatalogErrorInternalErrorResponseBody {
	body := &CatalogErrorInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshPayload builds a catalog service Refresh endpoint payload.
func NewRefreshPayload(catalogName string, token string) *catalog.RefreshPayload {
	v := &catalog.RefreshPayload{}
	v.CatalogName = catalogName
	v.Token = token

	return v
}

// NewRefreshAllPayload builds a catalog service RefreshAll endpoint payload.
func NewRefreshAllPayload(token string) *catalog.RefreshAllPayload {
	v := &catalog.RefreshAllPayload{}
	v.Token = token

	return v
}

// NewCatalogErrorPayload builds a catalog service CatalogError endpoint
// payload.
func NewCatalogErrorPayload(catalogName string, token string) *catalog.CatalogErrorPayload {
	v := &catalog.CatalogErrorPayload{}
	v.CatalogName = catalogName
	v.Token = token

	return v
}
