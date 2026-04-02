package repository

import (
	"context"

	uuidv4 "github.com/google/uuid"

	"go-layered-cats/model"
	"go-layered-cats/storage"
)

// CatRepository handles all database operations for cats
type CatRepository struct {
	store *storage.Store
}

// NewCatRepository creates a new instance of CatRepository
func NewCatRepository(store *storage.Store) *CatRepository {
	return &CatRepository{
		store: store,
	}
}

// Create stores a new cat in the repository
func (r *CatRepository) Create(cat *model.Cat) error {
	ctx := context.Background()
	r.store.Cat.Set(ctx, cat.UUID.String(), cat)
	return nil
}

// GetByID retrieves a cat by its ID
func (r *CatRepository) GetByID(uuid string) (*model.Cat, error) {
	val, exists := r.store.Cat.Get(context.Background(), uuid)
	if !exists {
		return nil, nil
	}
	return val, nil
}

// GetAll retrieves all cats from the repository
func (r *CatRepository) GetAll() ([]*model.Cat, error) {
	allValues := r.store.Cat.GetAll(context.Background())
	return allValues, nil
}

// Update modifies an existing cat in the repository
func (r *CatRepository) Update(uuid string, cat *model.Cat) error {
	if !r.store.Cat.Exists(context.Background(), uuid) {
		return nil
	}
	_uuid, err := uuidv4.Parse(uuid)
	if err != nil {
		return err
	}
	cat.UUID = _uuid
	r.store.Cat.Set(context.Background(), _uuid.String(), cat)
	return nil
}

// Delete removes a cat from the repository
func (r *CatRepository) Delete(id string) error {
	r.store.Cat.Delete(context.Background(), id)
	return nil
}
