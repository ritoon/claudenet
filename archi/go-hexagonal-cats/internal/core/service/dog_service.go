package service

import (
	"fmt"

	"go-hexagonal-cats/internal/core/domain"
	"go-hexagonal-cats/internal/core/ports"
)

// DogApplicationService implements the DogService inbound port
// It uses the DogRepository outbound port to persist data
type DogApplicationService struct {
	repository ports.DogRepository
}

// NewDogApplicationService creates a new instance of DogApplicationService
// This is where dependency injection happens - we inject the repository dependency
func NewDogApplicationService(repository ports.DogRepository) *DogApplicationService {
	return &DogApplicationService{
		repository: repository,
	}
}

// CreateDog creates a new dog
func (s *DogApplicationService) CreateDog(name, breed, color string, age int) (*domain.Dog, error) {
	// Business logic: validate inputs
	if name == "" {
		return nil, fmt.Errorf("dog name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("dog age cannot be negative")
	}

	// Generate a unique ID
	id, err := generateID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate dog ID: %w", err)
	}

	// Create the dog entity
	dog := &domain.Dog{
		ID:    id,
		Name:  name,
		Breed: breed,
		Color: color,
		Age:   age,
	}

	// Use the repository to persist the dog (driven port)
	if err := s.repository.Save(dog); err != nil {
		return nil, fmt.Errorf("failed to save dog: %w", err)
	}

	return dog, nil
}

// GetDogByID retrieves a dog by its ID
func (s *DogApplicationService) GetDogByID(id string) (*domain.Dog, error) {
	if id == "" {
		return nil, fmt.Errorf("dog ID cannot be empty")
	}

	dog, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find dog: %w", err)
	}
	if dog == nil {
		return nil, fmt.Errorf("dog with ID %s not found", id)
	}

	return dog, nil
}

// GetAllDogs retrieves all dogs
func (s *DogApplicationService) GetAllDogs() ([]*domain.Dog, error) {
	dogs, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve dogs: %w", err)
	}

	return dogs, nil
}

// UpdateDog updates an existing dog
func (s *DogApplicationService) UpdateDog(id string, request domain.UpdateDogRequest) (*domain.Dog, error) {
	if id == "" {
		return nil, fmt.Errorf("dog ID cannot be empty")
	}

	// Retrieve the existing dog
	dog, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find dog: %w", err)
	}
	if dog == nil {
		return nil, fmt.Errorf("dog with ID %s not found", id)
	}

	// Apply updates
	if request.Name != nil {
		if *request.Name == "" {
			return nil, fmt.Errorf("dog name cannot be empty")
		}
		dog.Name = *request.Name
	}
	if request.Breed != nil {
		dog.Breed = *request.Breed
	}
	if request.Color != nil {
		dog.Color = *request.Color
	}
	if request.Age != nil {
		if *request.Age < 0 {
			return nil, fmt.Errorf("dog age cannot be negative")
		}
		dog.Age = *request.Age
	}

	// Persist the updated dog
	if err := s.repository.Save(dog); err != nil {
		return nil, fmt.Errorf("failed to update dog: %w", err)
	}

	return dog, nil
}

// DeleteDog deletes a dog
func (s *DogApplicationService) DeleteDog(id string) error {
	if id == "" {
		return fmt.Errorf("dog ID cannot be empty")
	}

	// Check if dog exists
	if !s.repository.Exists(id) {
		return fmt.Errorf("dog with ID %s not found", id)
	}

	// Delete the dog
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("failed to delete dog: %w", err)
	}

	return nil
}
