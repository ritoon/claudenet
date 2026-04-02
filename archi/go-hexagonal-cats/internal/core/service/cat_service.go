package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"go-hexagonal-cats/internal/core/domain"
	"go-hexagonal-cats/internal/core/ports"
)

// CatApplicationService implements the CatService inbound port
// It uses the CatRepository outbound port to persist data
type CatApplicationService struct {
	repository ports.CatRepository
}

// NewCatApplicationService creates a new instance of CatApplicationService
// This is where dependency injection happens - we inject the repository dependency
func NewCatApplicationService(repository ports.CatRepository) *CatApplicationService {
	return &CatApplicationService{
		repository: repository,
	}
}

// generateID generates a unique identifier for a cat
func generateID() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// CreateCat creates a new cat
func (s *CatApplicationService) CreateCat(name, breed, color string, age int) (*domain.Cat, error) {
	// Business logic: validate inputs
	if name == "" {
		return nil, fmt.Errorf("cat name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("cat age cannot be negative")
	}

	// Generate a unique ID
	id, err := generateID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate cat ID: %w", err)
	}

	// Create the cat entity
	cat := &domain.Cat{
		ID:    id,
		Name:  name,
		Breed: breed,
		Color: color,
		Age:   age,
	}

	// Use the repository to persist the cat (driven port)
	if err := s.repository.Save(cat); err != nil {
		return nil, fmt.Errorf("failed to save cat: %w", err)
	}

	return cat, nil
}

// GetCatByID retrieves a cat by its ID
func (s *CatApplicationService) GetCatByID(id string) (*domain.Cat, error) {
	if id == "" {
		return nil, fmt.Errorf("cat ID cannot be empty")
	}

	cat, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find cat: %w", err)
	}
	if cat == nil {
		return nil, fmt.Errorf("cat with ID %s not found", id)
	}

	return cat, nil
}

// GetAllCats retrieves all cats
func (s *CatApplicationService) GetAllCats() ([]*domain.Cat, error) {
	cats, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cats: %w", err)
	}

	return cats, nil
}

// UpdateCat updates an existing cat
func (s *CatApplicationService) UpdateCat(id string, request domain.UpdateCatRequest) (*domain.Cat, error) {
	if id == "" {
		return nil, fmt.Errorf("cat ID cannot be empty")
	}

	// Retrieve the existing cat
	cat, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find cat: %w", err)
	}
	if cat == nil {
		return nil, fmt.Errorf("cat with ID %s not found", id)
	}

	// Apply updates
	if request.Name != nil {
		if *request.Name == "" {
			return nil, fmt.Errorf("cat name cannot be empty")
		}
		cat.Name = *request.Name
	}
	if request.Breed != nil {
		cat.Breed = *request.Breed
	}
	if request.Color != nil {
		cat.Color = *request.Color
	}
	if request.Age != nil {
		if *request.Age < 0 {
			return nil, fmt.Errorf("cat age cannot be negative")
		}
		cat.Age = *request.Age
	}

	// Persist the updated cat
	if err := s.repository.Save(cat); err != nil {
		return nil, fmt.Errorf("failed to update cat: %w", err)
	}

	return cat, nil
}

// DeleteCat deletes a cat
func (s *CatApplicationService) DeleteCat(id string) error {
	if id == "" {
		return fmt.Errorf("cat ID cannot be empty")
	}

	// Check if cat exists
	if !s.repository.Exists(id) {
		return fmt.Errorf("cat with ID %s not found", id)
	}

	// Delete the cat
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("failed to delete cat: %w", err)
	}

	return nil
}
