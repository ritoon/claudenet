package repository

import (
	"go-srp-cats/internal/entity"
	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/storage"
)

// CatRepository provides data access for Cat entities.
// SRP: Only reason to change - if query logic or data access patterns change.
// This bridges the domain layer with the storage layer.
// Changing how we query or access data only affects this file.
type CatRepository struct {
	store *storage.MemoryStore
}

// NewCatRepository creates a new cat repository.
func NewCatRepository(store *storage.MemoryStore) *CatRepository {
	return &CatRepository{
		store: store,
	}
}

// Save stores a cat in the repository.
func (r *CatRepository) Save(cat *entity.Cat) error {
	if r.store.Exists(cat.ID) {
		return &errors.AlreadyExistsError{ID: cat.ID}
	}
	r.store.Set(cat.ID, cat)
	return nil
}

// GetByID retrieves a cat by its ID.
func (r *CatRepository) GetByID(id string) (*entity.Cat, error) {
	val, exists := r.store.Get(id)
	if !exists {
		return nil, &errors.NotFoundError{ID: id}
	}
	cat, ok := val.(*entity.Cat)
	if !ok {
		return nil, &errors.NotFoundError{ID: id}
	}
	return cat, nil
}

// GetAll retrieves all cats.
func (r *CatRepository) GetAll() ([]*entity.Cat, error) {
	items := r.store.GetAll()
	cats := make([]*entity.Cat, 0, len(items))

	for _, val := range items {
		cat, ok := val.(*entity.Cat)
		if ok {
			cats = append(cats, cat)
		}
	}

	return cats, nil
}

// Update updates an existing cat.
func (r *CatRepository) Update(cat *entity.Cat) error {
	if !r.store.Exists(cat.ID) {
		return &errors.NotFoundError{ID: cat.ID}
	}
	r.store.Set(cat.ID, cat)
	return nil
}

// Delete removes a cat by its ID.
func (r *CatRepository) Delete(id string) error {
	if !r.store.Exists(id) {
		return &errors.NotFoundError{ID: id}
	}
	r.store.Delete(id)
	return nil
}
