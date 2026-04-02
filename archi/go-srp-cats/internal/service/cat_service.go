package service

import (
	"go-srp-cats/internal/entity"
	"go-srp-cats/internal/id"
	"go-srp-cats/internal/repository"
	"go-srp-cats/internal/validation"
)

// CatService orchestrates cat-related operations.
// SRP: Only reason to change - if the workflow/orchestration logic changes.
// This service coordinates between validator, ID generator, and repository.
// It contains NO validation logic itself - delegates to validator.
// It contains NO storage logic - delegates to repository.
// This keeps the file focused on business workflow coordination.
type CatService struct {
	repository *repository.CatRepository
	validator  *validation.CatValidator
	idGen      *id.Generator
}

// NewCatService creates a new cat service.
func NewCatService(
	repo *repository.CatRepository,
	validator *validation.CatValidator,
	idGen *id.Generator,
) *CatService {
	return &CatService{
		repository: repo,
		validator:  validator,
		idGen:      idGen,
	}
}

// CreateCat creates a new cat after validation.
func (s *CatService) CreateCat(name, breed, color string, age int) (*entity.Cat, error) {
	// Create entity for validation
	cat := &entity.Cat{
		Name:  name,
		Breed: breed,
		Color: color,
		Age:   age,
	}

	// Validate
	if err := s.validator.ValidateCreate(cat); err != nil {
		return nil, err
	}

	// Generate ID
	generatedID, err := s.idGen.Generate()
	if err != nil {
		return nil, err
	}
	cat.ID = generatedID

	// Save to repository
	if err := s.repository.Save(cat); err != nil {
		return nil, err
	}

	return cat, nil
}

// GetCatByID retrieves a cat by ID.
func (s *CatService) GetCatByID(id string) (*entity.Cat, error) {
	return s.repository.GetByID(id)
}

// GetAllCats retrieves all cats.
func (s *CatService) GetAllCats() ([]*entity.Cat, error) {
	return s.repository.GetAll()
}

// UpdateCat updates an existing cat.
func (s *CatService) UpdateCat(id string, name, breed, color string, age int) (*entity.Cat, error) {
	// Retrieve existing cat
	cat, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	cat.Name = name
	cat.Breed = breed
	cat.Color = color
	cat.Age = age

	// Validate
	if err := s.validator.ValidateUpdate(cat); err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repository.Update(cat); err != nil {
		return nil, err
	}

	return cat, nil
}

// DeleteCat deletes a cat by ID.
func (s *CatService) DeleteCat(id string) error {
	return s.repository.Delete(id)
}
