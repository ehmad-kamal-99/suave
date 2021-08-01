package service

import (
	"github.com/ehmad-kamal-99/suave/health"
)

// GetSvcHealth - returns service health status.
func (s *Service) GetSvcHealth() health.Status {
	return s.health.GetSvcHealth()
}
