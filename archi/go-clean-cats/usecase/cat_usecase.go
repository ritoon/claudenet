package usecase

import (
	"errors"
	"go-clean-cats/domain/entity"
	"go-clean-cats/domain/repository"
)

// CatUseCase handles all business logic related to cats.
// It depends on the CatRepository interface defined in the domain layer.
// This layer contains application-specific business rules.
type CatUseCase struct {
	catRepository repository.CatRepository
}

// NewCatUseCase creates a new instance of CatUseCase
// This demonstrates dependency injection - the repository is injected,
// allowing us to swap implementations without changing the use case
func NewCatUseCase(catRepository repository.CatRepository) *CatUseCase {
	return &CatUseCase{
		catRepository: catRepository,
	}
}

// CreateCat creates a new cat
func (uc *CatUseCase) CreateCat(id, name, breed string, age int, color string) (*entity.Cat, error) {
	if name == "" || breed == "" || color == "" {
		return nil, errors.New("name, breed, and color are required")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	cat := entity.NewCat(id, name, breed, age, color)
	return uc.catRepository.Create(cat)
}

// GetCatByID retrieves a single cat by ID
func (uc *CatUseCase) GetCatByID(id string) (*entity.Cat, error) {
	if id == "" {
		return nil, errors.New("cat id is required")
	}

	return uc.catRepository.GetByID(id)
}

// ListAllCats retrieves all cats
func (uc *CatUseCase) ListAllCats() ([]*entity.Cat, error) {
	return uc.catRepository.GetAll()
}

// UpdateCat modifies an existing cat
func (uc *CatUseCase) UpdateCat(id, name, breed string, age int, color string) (*entity.Cat, error) {
	if id == "" {
		return nil, errors.New("cat id is required")
	}

	if name == "" || breed == "" || color == "" {
		return nil, errors.New("name, breed, and color are required")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	cat := entity.NewCat(id, name, breed, age, color)
	return uc.catRepository.Update(cat)
}

// DeleteCat removes a cat by ID
func (uc *CatUseCase) DeleteCat(id string) error {
	if id == "" {
		return errors.New("cat id is required")
	}

	return uc.catRepository.Delete(id)
}
