package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/ehmad-kamal-99/suave"
	"github.com/ehmad-kamal-99/suave/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.UASVCAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewUASVCAPI(spec)

	// user handlers
	handler.AddUsernameHandler = NewAddUsername(rt)
	handler.GetUsernameHandler = NewGetUsername(rt)

	// health handler
	handler.GetHealthHandler = NewGetHealth(rt)

	return handler
}
