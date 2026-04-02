package repository

import (
	"go-srp-cats/internal/entity"
	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/storage"
)

// DogRepository provides data access for Dog entities.
// SRP: Only reason to change - if query logic or data access patterns for dogs change.
type DogRepository struct {
	store *storage.MemoryStore
}

// NewDogRepository creates a new dog repository.
func NewDogRepository(store *storage.MemoryStore) *DogRepository {
	return &DogRepository{
		store: store,
	}
}

// Save stores a dog in the repository.
func (r *DogRepository) Save(dog *entity.Dog) error {
	if r.store.Exists(dog.ID) {
		return &errors.AlreadyExistsError{ID: dog.ID}
	}
	r.store.Set(dog.ID, dog)
	return nil
}

// GetByID retrieves a dog by its ID.
func (r *DogRepository) GetByID(id string) (*entity.Dog, error) {
	val, exists := r.store.Get(id)
	if !exists {
		return nil, &errors.NotFoundError{ID: id}
	}
	dog, ok := val.(*entity.Dog)
	if !ok {
		return nil, &errors.NotFoundError{ID: id}
	}
	return dog, nil
}

// GetAll retrieves all dogs.
func (r *DogRepository) GetAll() ([]*entity.Dog, error) {
	items := r.store.GetAll()
	dogs := make([]*entity.Dog, 0, len(items))

	for _, val := range items {
		dog, ok := val.(*entity.Dog)
		if ok {
			dogs = append(dogs, dog)
		}
	}

	return dogs, nil
}

// Update updates an existing dog.
func (r *DogRepository) Update(dog *entity.Dog) error {
	if !r.store.Exists(dog.ID) {
		return &errors.NotFoundError{ID: dog.ID}
	}
	r.store.Set(dog.ID, dog)
	return nil
}

// Delete removes a dog by its ID.
func (r *DogRepository) Delete(id string) error {
	if !r.store.Exists(id) {
		return &errors.NotFoundError{ID: id}
	}
	r.store.Delete(id)
	return nil
}
