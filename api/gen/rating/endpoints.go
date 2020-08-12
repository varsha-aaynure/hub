// Code generated by goa v3.2.2, DO NOT EDIT.
//
// rating endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package rating

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "rating" service endpoints.
type Endpoints struct {
	Get    goa.Endpoint
	Update goa.Endpoint
}

// NewEndpoints wraps the methods of the "rating" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Get:    NewGetEndpoint(s, a.JWTAuth),
		Update: NewUpdateEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "rating" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
	e.Update = m(e.Update)
}

// NewGetEndpoint returns an endpoint function that calls the method "Get" of
// service "rating".
func NewGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write"},
			RequiredScopes: []string{"rating:read"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Get(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "Update" of service "rating".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write"},
			RequiredScopes: []string{"rating:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Update(ctx, p)
	}
}