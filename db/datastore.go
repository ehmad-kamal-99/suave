package db

// Datastore - interface for service operations
type Datastore interface {
	AddUsername(username string) error
	CheckUsername(username string) (bool, error)

	Close() error
}
