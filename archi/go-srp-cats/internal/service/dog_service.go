package service

import (
	"go-srp-cats/internal/entity"
	"go-srp-cats/internal/id"
	"go-srp-cats/internal/repository"
	"go-srp-cats/internal/validation"
)

// DogService orchestrates dog-related operations.
// SRP: Only reason to change - if the workflow/orchestration logic for dogs changes.
type DogService struct {
	repository *repository.DogRepository
	validator  *validation.DogValidator
	idGen      *id.Generator
}

// NewDogService creates a new dog service.
func NewDogService(
	repo *repository.DogRepository,
	validator *validation.DogValidator,
	idGen *id.Generator,
) *DogService {
	return &DogService{
		repository: repo,
		validator:  validator,
		idGen:      idGen,
	}
}

// CreateDog creates a new dog after validation.
func (s *DogService) CreateDog(name, breed, color string, age int) (*entity.Dog, error) {
	// Create entity for validation
	dog := &entity.Dog{
		Name:  name,
		Breed: breed,
		Color: color,
		Age:   age,
	}

	// Validate
	if err := s.validator.ValidateCreate(dog); err != nil {
		return nil, err
	}

	// Generate ID
	generatedID, err := s.idGen.Generate()
	if err != nil {
		return nil, err
	}
	dog.ID = generatedID

	// Save to repository
	if err := s.repository.Save(dog); err != nil {
		return nil, err
	}

	return dog, nil
}

// GetDogByID retrieves a dog by ID.
func (s *DogService) GetDogByID(id string) (*entity.Dog, error) {
	return s.repository.GetByID(id)
}

// GetAllDogs retrieves all dogs.
func (s *DogService) GetAllDogs() ([]*entity.Dog, error) {
	return s.repository.GetAll()
}

// UpdateDog updates an existing dog.
func (s *DogService) UpdateDog(id string, name, breed, color string, age int) (*entity.Dog, error) {
	// Retrieve existing dog
	dog, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	dog.Name = name
	dog.Breed = breed
	dog.Color = color
	dog.Age = age

	// Validate
	if err := s.validator.ValidateUpdate(dog); err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repository.Update(dog); err != nil {
		return nil, err
	}

	return dog, nil
}

// DeleteDog deletes a dog by ID.
func (s *DogService) DeleteDog(id string) error {
	return s.repository.Delete(id)
}
