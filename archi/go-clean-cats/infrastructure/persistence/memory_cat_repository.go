package persistence

import (
	"errors"
	"sync"

	"go-clean-cats/domain/entity"
	"go-clean-cats/domain/repository"
)

// MemoryCatRepository is an in-memory implementation of the CatRepository interface.
// This is in the infrastructure layer because it's a detail of how persistence is implemented.
// The use case doesn't know or care that it's using in-memory storage - it just uses the interface.
// This is a concrete example of the Dependency Inversion Principle.
type MemoryCatRepository struct {
	cats map[string]*entity.Cat
	mu   sync.RWMutex
}

// NewMemoryCatRepository creates a new in-memory cat repository
func NewMemoryCatRepository() repository.CatRepository {
	return &MemoryCatRepository{
		cats: make(map[string]*entity.Cat),
	}
}

// Create persists a new cat
func (r *MemoryCatRepository) Create(cat *entity.Cat) (*entity.Cat, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.cats[cat.ID]; exists {
		return nil, errors.New("cat with this ID already exists")
	}

	r.cats[cat.ID] = cat
	return cat, nil
}

// GetByID retrieves a cat by its ID
func (r *MemoryCatRepository) GetByID(id string) (*entity.Cat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cat, exists := r.cats[id]
	if !exists {
		return nil, errors.New("cat not found")
	}

	return cat, nil
}

// GetAll retrieves all cats
func (r *MemoryCatRepository) GetAll() ([]*entity.Cat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cats := make([]*entity.Cat, 0, len(r.cats))
	for _, cat := range r.cats {
		cats = append(cats, cat)
	}

	return cats, nil
}

// Update modifies an existing cat
func (r *MemoryCatRepository) Update(cat *entity.Cat) (*entity.Cat, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.cats[cat.ID]; !exists {
		return nil, errors.New("cat not found")
	}

	r.cats[cat.ID] = cat
	return cat, nil
}

// Delete removes a cat by its ID
func (r *MemoryCatRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.cats[id]; !exists {
		return errors.New("cat not found")
	}

	delete(r.cats, id)
	return nil
}
