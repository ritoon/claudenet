package usecase

import (
	"errors"
	"go-clean-cats/domain/entity"
	"go-clean-cats/domain/repository"
)

// DogUseCase handles all business logic related to dogs.
// It depends on the DogRepository interface defined in the domain layer.
type DogUseCase struct {
	dogRepository repository.DogRepository
}

// NewDogUseCase creates a new instance of DogUseCase
func NewDogUseCase(dogRepository repository.DogRepository) *DogUseCase {
	return &DogUseCase{
		dogRepository: dogRepository,
	}
}

// CreateDog creates a new dog
func (uc *DogUseCase) CreateDog(id, name, breed string, age int, color string) (*entity.Dog, error) {
	if name == "" || breed == "" || color == "" {
		return nil, errors.New("name, breed, and color are required")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	dog := entity.NewCat(id, name, breed, age, color)
	return uc.dogRepository.Create(dog)
}

// GetDogByID retrieves a single dog by ID
func (uc *DogUseCase) GetDogByID(id string) (*entity.Dog, error) {
	if id == "" {
		return nil, errors.New("dog id is required")
	}

	return uc.dogRepository.GetByID(id)
}

// ListAllDogs retrieves all dogs
func (uc *DogUseCase) ListAllDogs() ([]*entity.Dog, error) {
	return uc.dogRepository.GetAll()
}

// UpdateDog modifies an existing dog
func (uc *DogUseCase) UpdateDog(id, name, breed string, age int, color string) (*entity.Dog, error) {
	if id == "" {
		return nil, errors.New("dog id is required")
	}

	if name == "" || breed == "" || color == "" {
		return nil, errors.New("name, breed, and color are required")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	dog := entity.NewCat(id, name, breed, age, color)
	return uc.dogRepository.Update(dog)
}

// DeleteDog removes a dog by ID
func (uc *DogUseCase) DeleteDog(id string) error {
	if id == "" {
		return errors.New("dog id is required")
	}

	return uc.dogRepository.Delete(id)
}
