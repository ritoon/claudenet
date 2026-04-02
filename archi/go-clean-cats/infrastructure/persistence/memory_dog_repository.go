package persistence

import (
	"errors"
	"sync"

	"go-clean-cats/domain/entity"
	"go-clean-cats/domain/repository"
)

// MemoryDogRepository is an in-memory implementation of the DogRepository interface.
// This is in the infrastructure layer because it's a detail of how persistence is implemented.
type MemoryDogRepository struct {
	dogs map[string]*entity.Dog
	mu   sync.RWMutex
}

// NewMemoryDogRepository creates a new in-memory dog repository
func NewMemoryDogRepository() repository.DogRepository {
	return &MemoryDogRepository{
		dogs: make(map[string]*entity.Dog),
	}
}

// Create persists a new dog
func (r *MemoryDogRepository) Create(dog *entity.Dog) (*entity.Dog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dogs[dog.ID]; exists {
		return nil, errors.New("dog with this ID already exists")
	}

	r.dogs[dog.ID] = dog
	return dog, nil
}

// GetByID retrieves a dog by its ID
func (r *MemoryDogRepository) GetByID(id string) (*entity.Dog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dog, exists := r.dogs[id]
	if !exists {
		return nil, errors.New("dog not found")
	}

	return dog, nil
}

// GetAll retrieves all dogs
func (r *MemoryDogRepository) GetAll() ([]*entity.Dog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dogs := make([]*entity.Dog, 0, len(r.dogs))
	for _, dog := range r.dogs {
		dogs = append(dogs, dog)
	}

	return dogs, nil
}

// Update modifies an existing dog
func (r *MemoryDogRepository) Update(dog *entity.Dog) (*entity.Dog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dogs[dog.ID]; !exists {
		return nil, errors.New("dog not found")
	}

	r.dogs[dog.ID] = dog
	return dog, nil
}

// Delete removes a dog by its ID
func (r *MemoryDogRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dogs[id]; !exists {
		return errors.New("dog not found")
	}

	delete(r.dogs, id)
	return nil
}
