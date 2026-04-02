package memory

import (
	"fmt"
	"sync"

	"go-hexagonal-cats/internal/core/domain"
)

// CatRepository is an in-memory implementation of the CatRepository outbound port
// It stores cats in a map with thread-safe access using sync.RWMutex
type CatRepository struct {
	mu   sync.RWMutex
	cats map[string]*domain.Cat
}

// NewCatRepository creates a new in-memory cat repository
func NewCatRepository() *CatRepository {
	return &CatRepository{
		cats: make(map[string]*domain.Cat),
	}
}

// Save stores or updates a cat
func (r *CatRepository) Save(cat *domain.Cat) error {
	if cat == nil {
		return fmt.Errorf("cat cannot be nil")
	}
	if cat.ID == "" {
		return fmt.Errorf("cat ID cannot be empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.cats[cat.ID] = cat
	return nil
}

// FindByID retrieves a cat by its ID
func (r *CatRepository) FindByID(id string) (*domain.Cat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cat, exists := r.cats[id]
	if !exists {
		return nil, nil
	}

	// Return a copy to prevent external modifications
	catCopy := *cat
	return &catCopy, nil
}

// FindAll retrieves all cats
func (r *CatRepository) FindAll() ([]*domain.Cat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cats := make([]*domain.Cat, 0, len(r.cats))
	for _, cat := range r.cats {
		catCopy := *cat
		cats = append(cats, &catCopy)
	}

	return cats, nil
}

// Delete removes a cat by its ID
func (r *CatRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.cats[id]; !exists {
		return fmt.Errorf("cat with ID %s not found", id)
	}

	delete(r.cats, id)
	return nil
}

// Exists checks if a cat with the given ID exists
func (r *CatRepository) Exists(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.cats[id]
	return exists
}
