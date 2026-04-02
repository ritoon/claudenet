package service

import (
	"errors"
	"fmt"
	"go-layered-cats/model"
	"go-layered-cats/repository"
	"strings"

	"github.com/google/uuid"
)

// CatService handles business logic for cat operations
type CatService struct {
	repo *repository.CatRepository
}

// NewCatService creates a new instance of CatService
func NewCatService(repo *repository.CatRepository) *CatService {
	return &CatService{
		repo: repo,
	}
}

// CreateCat validates input and creates a new cat
func (s *CatService) CreateCat(cat *model.Cat) (*model.Cat, error) {
	// Validate input
	if err := validateCat(cat); err != nil {
		return nil, err
	}

	// Generate UUID for the cat
	cat.UUID = uuid.New()

	// Save to repository
	if err := s.repo.Create(cat); err != nil {
		return nil, fmt.Errorf("failed to create cat: %w", err)
	}

	return cat, nil
}

// GetCatByID retrieves a cat by its ID
func (s *CatService) GetCatByID(id string) (*model.Cat, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("cat ID cannot be empty")
	}

	cat, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cat: %w", err)
	}

	if cat == nil {
		return nil, errors.New("cat not found")
	}

	return cat, nil
}

// GetAllCats retrieves all cats
func (s *CatService) GetAllCats() ([]*model.Cat, error) {
	cats, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cats: %w", err)
	}

	if cats == nil {
		cats = make([]*model.Cat, 0)
	}

	return cats, nil
}

// UpdateCat validates input and updates an existing cat
func (s *CatService) UpdateCat(id string, cat *model.Cat) (*model.Cat, error) {
	// Validate ID
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("cat ID cannot be empty")
	}

	// Check if cat exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to check cat existence: %w", err)
	}

	if existing == nil {
		return nil, errors.New("cat not found")
	}

	// Validate input
	if err := validateCat(cat); err != nil {
		return nil, err
	}

	// Update in repository
	if err := s.repo.Update(id, cat); err != nil {
		return nil, fmt.Errorf("failed to update cat: %w", err)
	}

	return cat, nil
}

// DeleteCat removes a cat by its ID
func (s *CatService) DeleteCat(id string) error {
	// Validate ID
	if strings.TrimSpace(id) == "" {
		return errors.New("cat ID cannot be empty")
	}

	// Check if cat exists
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to check cat existence: %w", err)
	}

	if existing == nil {
		return errors.New("cat not found")
	}

	// Delete from repository
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete cat: %w", err)
	}

	return nil
}

// validateCat performs business logic validation on cat data
func validateCat(cat *model.Cat) error {
	if cat == nil {
		return errors.New("cat cannot be nil")
	}

	if strings.TrimSpace(cat.Name) == "" {
		return errors.New("cat name cannot be empty")
	}

	if strings.TrimSpace(cat.Breed) == "" {
		return errors.New("cat breed cannot be empty")
	}

	if cat.Age < 0 {
		return errors.New("cat age cannot be negative")
	}

	if cat.Age > 50 {
		return errors.New("cat age cannot exceed 50 years")
	}

	if strings.TrimSpace(cat.Color) == "" {
		return errors.New("cat color cannot be empty")
	}

	return nil
}
