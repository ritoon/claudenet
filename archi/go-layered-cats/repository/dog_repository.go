package repository

import (
	"context"

	uuidv4 "github.com/google/uuid"

	"go-layered-cats/model"
	"go-layered-cats/storage"
)

// DogRepository handles all database operations for dogs
type DogRepository struct {
	store *storage.Store
}

// NewDogRepository creates a new instance of DogRepository
func NewDogRepository(store *storage.Store) *DogRepository {
	return &DogRepository{
		store: store,
	}
}

// Create stores a new cat in the repository
func (r *DogRepository) Create(cat *model.Cat) error {
	ctx := context.Background()
	r.store.Dog.Set(ctx, cat.UUID.String(), cat)
	return nil
}

// GetByID retrieves a dog by its ID
func (r *DogRepository) GetByID(uuid string) (*model.Cat, error) {
	val, exists := r.store.Dog.Get(context.Background(), uuid)
	if !exists {
		return nil, nil
	}
	return val, nil
}

// GetAll retrieves all dogs from the repository
func (r *DogRepository) GetAll() ([]*model.Cat, error) {
	allValues := r.store.Dog.GetAll(context.Background())
	return allValues, nil
}

// Update modifies an existing dog in the repository
func (r *DogRepository) Update(uuid string, cat *model.Cat) error {
	if !r.store.Dog.Exists(context.Background(), uuid) {
		return nil
	}
	_uuid, err := uuidv4.Parse(uuid)
	if err != nil {
		return err
	}
	cat.UUID = _uuid
	r.store.Dog.Set(context.Background(), _uuid.String(), cat)
	return nil
}

// Delete removes a dog from the repository
func (r *DogRepository) Delete(uuid string) error {
	r.store.Dog.Delete(context.Background(), uuid)
	return nil
}
