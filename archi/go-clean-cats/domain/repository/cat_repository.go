package repository

import "go-clean-cats/domain/entity"

// CatRepository is the interface that any persistence layer must implement.
// This interface is defined in the domain layer (use case) and implemented in the infrastructure layer.
// This is the Dependency Inversion Principle in action:
// - High-level modules (use cases) define the interface
// - Low-level modules (persistence) implement it
// - Both depend on the abstraction, not on each other directly
type CatRepository interface {
	// Create persists a new cat and returns the created cat
	Create(cat *entity.Cat) (*entity.Cat, error)

	// GetByID retrieves a cat by its ID
	GetByID(id string) (*entity.Cat, error)

	// GetAll retrieves all cats
	GetAll() ([]*entity.Cat, error)

	// Update modifies an existing cat
	Update(cat *entity.Cat) (*entity.Cat, error)

	// Delete removes a cat by its ID
	Delete(id string) error
}
