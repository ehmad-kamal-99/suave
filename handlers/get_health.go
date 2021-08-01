package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/ehmad-kamal-99/suave"
	"github.com/ehmad-kamal-99/suave/gen/models"
	"github.com/ehmad-kamal-99/suave/gen/restapi/operations"
)

// NewGetHealth - replaces swagger's GetHealth handler.
func NewGetHealth(rt *runtime.Runtime) operations.GetHealthHandler {
	return &getHealth{rt: rt}
}

type getHealth struct {
	rt *runtime.Runtime
}

// Handle - handles get health request.
func (g *getHealth) Handle(params operations.GetHealthParams) middleware.Responder {
	log().Debugf("getHealth request: %s", params.HTTPRequest.Host)

	health := g.rt.Service().GetSvcHealth()

	return operations.NewGetHealthOK().WithPayload(&models.Health{
		Status:    swag.String(health.State),
		Main:      health.Main,
		DataStore: health.RedisBloom,
	})
}
