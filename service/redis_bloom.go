package service

// Close - closes redis-bloom connection.
func (s *Service) Close() error {
	return s.db.Close()
}
