package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/ehmad-kamal-99/suave"
	"github.com/ehmad-kamal-99/suave/gen/models"
	"github.com/ehmad-kamal-99/suave/gen/restapi/operations"
)

// NewGetUsername - replaces swagger's GetUsername handler.
func NewGetUsername(rt *runtime.Runtime) operations.GetUsernameHandler {
	return &getUsername{rt: rt}
}

type getUsername struct {
	rt *runtime.Runtime
}

// Handle - handles get username request.
func (g *getUsername) Handle(params operations.GetUsernameParams) middleware.Responder {
	log().Debugf("getUsername request: %s", params.Username)

	exists, err := g.rt.Service().CheckUsername(params.Username)
	if err != nil {
		log().Errorf("err[500]: failed to check username, err: %+v", err)
		return operations.NewGetUsernameInternalServerError()
	}

	return operations.NewGetUsernameOK().
		WithPayload(&models.Availability{
			Exists:   swag.Bool(exists),
			Username: swag.String(params.Username),
		})
}
