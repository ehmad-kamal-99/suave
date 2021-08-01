// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsernameHandlerFunc turns a function with the right signature into a get username handler
type GetUsernameHandlerFunc func(GetUsernameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsernameHandlerFunc) Handle(params GetUsernameParams) middleware.Responder {
	return fn(params)
}

// GetUsernameHandler interface for that can handle valid get username params
type GetUsernameHandler interface {
	Handle(GetUsernameParams) middleware.Responder
}

// NewGetUsername creates a new http.Handler for the get username operation
func NewGetUsername(ctx *middleware.Context, handler GetUsernameHandler) *GetUsername {
	return &GetUsername{Context: ctx, Handler: handler}
}

/* GetUsername swagger:route GET /username/availability getUsername

check username

*/
type GetUsername struct {
	Context *middleware.Context
	Handler GetUsernameHandler
}

func (o *GetUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUsernameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
