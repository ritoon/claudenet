package repository

import "go-clean-cats/domain/entity"

// DogRepository is the interface that any persistence layer must implement for dogs.
// This interface is defined in the domain layer and implemented in the infrastructure layer.
type DogRepository interface {
	// Create persists a new dog and returns the created dog
	Create(dog *entity.Dog) (*entity.Dog, error)

	// GetByID retrieves a dog by its ID
	GetByID(id string) (*entity.Dog, error)

	// GetAll retrieves all dogs
	GetAll() ([]*entity.Dog, error)

	// Update modifies an existing dog
	Update(dog *entity.Dog) (*entity.Dog, error)

	// Delete removes a dog by its ID
	Delete(id string) error
}
