package service

import (
	"github.com/ehmad-kamal-99/suave/db"
	"github.com/ehmad-kamal-99/suave/health"
)

// Service holds our datastore and health instance.
type Service struct {
	db     db.Datastore
	health health.Health
}

// NewService returns a new service instance.
func NewService(ds db.Datastore, svcHealth health.Health) *Service {
	return &Service{
		db:     ds,
		health: svcHealth,
	}
}
