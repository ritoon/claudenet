package repository

import (
	"go-layered-cats/model"
	"go-layered-cats/storage"
)

// CatRepository handles all database operations for cats
type CatRepository struct {
	store *storage.MemoryStore
}

// NewCatRepository creates a new instance of CatRepository
func NewCatRepository(store *storage.MemoryStore) *CatRepository {
	return &CatRepository{
		store: store,
	}
}

// Create stores a new cat in the repository
func (r *CatRepository) Create(cat *model.Cat) error {
	r.store.Set(cat.ID, cat)
	return nil
}

// GetByID retrieves a cat by its ID
func (r *CatRepository) GetByID(id string) (*model.Cat, error) {
	val, exists := r.store.Get(id)
	if !exists {
		return nil, nil
	}
	cat := val.(*model.Cat)
	return cat, nil
}

// GetAll retrieves all cats from the repository
func (r *CatRepository) GetAll() ([]*model.Cat, error) {
	allValues := r.store.GetAll()
	cats := make([]*model.Cat, len(allValues))
	for i, val := range allValues {
		cats[i] = val.(*model.Cat)
	}
	return cats, nil
}

// Update modifies an existing cat in the repository
func (r *CatRepository) Update(id string, cat *model.Cat) error {
	if !r.store.Exists(id) {
		return nil
	}
	cat.ID = id
	r.store.Set(id, cat)
	return nil
}

// Delete removes a cat from the repository
func (r *CatRepository) Delete(id string) error {
	r.store.Delete(id)
	return nil
}
