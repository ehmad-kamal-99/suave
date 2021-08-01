package runtime

import (
	"github.com/ehmad-kamal-99/suave/db/redisbloom"
	"github.com/ehmad-kamal-99/suave/health"
	"github.com/ehmad-kamal-99/suave/service"
)

// Runtime initializes values for entry point to our application.
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new runtime object.
func NewRuntime() *Runtime {
	return &Runtime{
		svc: service.NewService(redisbloom.NewClient(), health.New()),
	}
}

// Service returns pointer to our service variable.
func (r Runtime) Service() *service.Service {
	return r.svc
}
