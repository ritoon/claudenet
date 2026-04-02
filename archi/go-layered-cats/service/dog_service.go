package service

import (
	"errors"
	"fmt"
	"go-layered-cats/model"
	"go-layered-cats/repository"
	"strings"

	"github.com/google/uuid"
)

// DogService handles business logic for dog operations
type DogService struct {
	repo *repository.DogRepository
}

// NewDogService creates a new instance of DogService
func NewDogService(repo *repository.DogRepository) *DogService {
	return &DogService{
		repo: repo,
	}
}

// CreateDog validates input and creates a new dog
func (s *DogService) CreateDog(dog *model.Dog) (*model.Dog, error) {
	// Validate input
	if err := validateDog(dog); err != nil {
		return nil, err
	}

	// Generate UUID for the dog
	dog.UUID = uuid.New()

	// Save to repository
	if err := s.repo.Create(dog); err != nil {
		return nil, fmt.Errorf("failed to create dog: %w", err)
	}

	return dog, nil
}

// GetDogByID retrieves a dog by its ID
func (s *DogService) GetDogByID(id string) (*model.Dog, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("dog ID cannot be empty")
	}

	dog, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve dog: %w", err)
	}

	if dog == nil {
		return nil, errors.New("dog not found")
	}

	return dog, nil
}

// GetAllDogs retrieves all dogs
func (s *DogService) GetAllDogs() ([]*model.Dog, error) {
	dogs, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve dogs: %w", err)
	}

	if dogs == nil {
		dogs = make([]*model.Dog, 0)
	}

	return dogs, nil
}

// UpdateDog validates input and updates an existing dog
func (s *DogService) UpdateDog(id string, dog *model.Dog) (*model.Dog, error) {
	// Validate ID
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("dog ID cannot be empty")
	}

	// Check if dog exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to check dog existence: %w", err)
	}

	if existing == nil {
		return nil, errors.New("dog not found")
	}

	// Validate input
	if err := validateDog(dog); err != nil {
		return nil, err
	}

	// Update in repository
	if err := s.repo.Update(id, dog); err != nil {
		return nil, fmt.Errorf("failed to update dog: %w", err)
	}

	return dog, nil
}

// DeleteDog removes a dog by its ID
func (s *DogService) DeleteDog(id string) error {
	// Validate ID
	if strings.TrimSpace(id) == "" {
		return errors.New("dog ID cannot be empty")
	}

	// Check if dog exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to check dog existence: %w", err)
	}

	if existing == nil {
		return errors.New("dog not found")
	}

	// Delete from repository
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete dog: %w", err)
	}

	return nil
}

// validateDog performs business logic validation on dog data
func validateDog(dog *model.Dog) error {
	if dog == nil {
		return errors.New("dog cannot be nil")
	}

	if strings.TrimSpace(dog.Name) == "" {
		return errors.New("dog name cannot be empty")
	}

	if strings.TrimSpace(dog.Breed) == "" {
		return errors.New("dog breed cannot be empty")
	}

	if dog.Age < 0 {
		return errors.New("dog age cannot be negative")
	}

	if dog.Age > 50 {
		return errors.New("dog age cannot exceed 50 years")
	}

	if strings.TrimSpace(dog.Color) == "" {
		return errors.New("dog color cannot be empty")
	}

	return nil
}
