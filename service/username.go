package service

// AddUsername - persists username in datastore.
func (s *Service) AddUsername(username string) error {
	return s.db.AddUsername(username)
}

// CheckUsername - checks for username in datastore.
func (s *Service) CheckUsername(username string) (bool, error) {
	return s.db.CheckUsername(username)
}
