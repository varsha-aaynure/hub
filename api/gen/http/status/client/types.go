// Code generated by goa v3.10.1, DO NOT EDIT.
//
// status HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	status "github.com/tektoncd/hub/api/gen/status"
	goa "goa.design/goa/v3/pkg"
)

// StatusResponseBody is the type of the "status" service "Status" endpoint
// HTTP response body.
type StatusResponseBody struct {
	// List of services and their status
	Services []*HubServiceResponseBody `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
}

// HubServiceResponseBody is used to define fields on response body types.
type HubServiceResponseBody struct {
	// Name of the service
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Status of the service
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Details of the error if any
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
}

// NewStatusResultOK builds a "status" service "Status" endpoint result from a
// HTTP "OK" response.
func NewStatusResultOK(body *StatusResponseBody) *status.StatusResult {
	v := &status.StatusResult{}
	if body.Services != nil {
		v.Services = make([]*status.HubService, len(body.Services))
		for i, val := range body.Services {
			v.Services[i] = unmarshalHubServiceResponseBodyToStatusHubService(val)
		}
	}

	return v
}

// ValidateStatusResponseBody runs the validations defined on StatusResponseBody
func ValidateStatusResponseBody(body *StatusResponseBody) (err error) {
	for _, e := range body.Services {
		if e != nil {
			if err2 := ValidateHubServiceResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateHubServiceResponseBody runs the validations defined on
// HubServiceResponseBody
func ValidateHubServiceResponseBody(body *HubServiceResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "ok" || *body.Status == "error") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"ok", "error"}))
		}
	}
	return
}
