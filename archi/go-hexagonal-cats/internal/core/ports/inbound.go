package ports

import "go-hexagonal-cats/internal/core/domain"

// CatService defines the inbound (driving) port for cat operations
// This is what the application can do - the use cases it exposes
type CatService interface {
	// CreateCat creates a new cat and returns the created cat
	CreateCat(name, breed, color string, age int) (*domain.Cat, error)

	// GetCatByID retrieves a cat by its ID
	GetCatByID(id string) (*domain.Cat, error)

	// GetAllCats retrieves all cats
	GetAllCats() ([]*domain.Cat, error)

	// UpdateCat updates an existing cat
	UpdateCat(id string, request domain.UpdateCatRequest) (*domain.Cat, error)

	// DeleteCat deletes a cat by its ID
	DeleteCat(id string) error
}

// DogService defines the inbound (driving) port for dog operations
type DogService interface {
	// CreateDog creates a new dog and returns the created dog
	CreateDog(name, breed, color string, age int) (*domain.Dog, error)

	// GetDogByID retrieves a dog by its ID
	GetDogByID(id string) (*domain.Dog, error)

	// GetAllDogs retrieves all dogs
	GetAllDogs() ([]*domain.Dog, error)

	// UpdateDog updates an existing dog
	UpdateDog(id string, request domain.UpdateDogRequest) (*domain.Dog, error)

	// DeleteDog deletes a dog by its ID
	DeleteDog(id string) error
}
