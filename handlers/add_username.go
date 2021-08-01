package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/ehmad-kamal-99/suave"
	"github.com/ehmad-kamal-99/suave/gen/restapi/operations"
)

// NewAddUsername - replaces swagger's AddUsername handler.
func NewAddUsername(rt *runtime.Runtime) operations.AddUsernameHandler {
	return &addUsername{rt: rt}
}

type addUsername struct {
	rt *runtime.Runtime
}

// Handle - handles add username request.
func (a *addUsername) Handle(params operations.AddUsernameParams) middleware.Responder {
	log().Debugf("addUsername request: %s", swag.StringValue(params.Username.Username))

	if err := a.rt.Service().AddUsername(swag.StringValue(params.Username.Username)); err != nil {
		log().Errorf("err[500]: failed to add username, err: %+v", err)
		return operations.NewAddUsernameInternalServerError()
	}

	return operations.NewAddUsernameCreated().
		WithPayload(&operations.AddUsernameCreatedBody{
			Username: params.Username.Username,
		})
}
