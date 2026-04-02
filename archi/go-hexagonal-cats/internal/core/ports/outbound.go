package ports

import "go-hexagonal-cats/internal/core/domain"

// CatRepository defines the outbound (driven) port for cat persistence
// This is what the application needs to accomplish its goals
// The application depends on this interface, but does not depend on its implementation
type CatRepository interface {
	// Save stores a cat or updates it if it already exists
	Save(cat *domain.Cat) error

	// FindByID retrieves a cat by its ID, returns nil if not found
	FindByID(id string) (*domain.Cat, error)

	// FindAll retrieves all cats
	FindAll() ([]*domain.Cat, error)

	// Delete removes a cat by its ID
	Delete(id string) error

	// Exists checks if a cat with given ID exists
	Exists(id string) bool
}
