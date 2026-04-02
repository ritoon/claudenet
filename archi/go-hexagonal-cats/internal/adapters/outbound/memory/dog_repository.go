package memory

import (
	"fmt"
	"sync"

	"go-hexagonal-cats/internal/core/domain"
)

// DogRepository is an in-memory implementation of the DogRepository outbound port
// It stores dogs in a map with thread-safe access using sync.RWMutex
type DogRepository struct {
	mu   sync.RWMutex
	dogs map[string]*domain.Dog
}

// NewDogRepository creates a new in-memory dog repository
func NewDogRepository() *DogRepository {
	return &DogRepository{
		dogs: make(map[string]*domain.Dog),
	}
}

// Save stores or updates a dog
func (r *DogRepository) Save(dog *domain.Dog) error {
	if dog == nil {
		return fmt.Errorf("dog cannot be nil")
	}
	if dog.ID == "" {
		return fmt.Errorf("dog ID cannot be empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.dogs[dog.ID] = dog
	return nil
}

// FindByID retrieves a dog by its ID
func (r *DogRepository) FindByID(id string) (*domain.Dog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dog, exists := r.dogs[id]
	if !exists {
		return nil, nil
	}

	// Return a copy to prevent external modifications
	dogCopy := *dog
	return &dogCopy, nil
}

// FindAll retrieves all dogs
func (r *DogRepository) FindAll() ([]*domain.Dog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dogs := make([]*domain.Dog, 0, len(r.dogs))
	for _, dog := range r.dogs {
		dogCopy := *dog
		dogs = append(dogs, &dogCopy)
	}

	return dogs, nil
}

// Delete removes a dog by its ID
func (r *DogRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.dogs[id]; !exists {
		return fmt.Errorf("dog with ID %s not found", id)
	}

	delete(r.dogs, id)
	return nil
}

// Exists checks if a dog with the given ID exists
func (r *DogRepository) Exists(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.dogs[id]
	return exists
}
