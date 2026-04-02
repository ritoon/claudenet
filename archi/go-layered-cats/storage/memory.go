package storage

import (
	"context"
	"go-layered-cats/model"
)

// // MemoryStore provides in-memory data storage using a map
// type MemoryStore struct {
// 	data map[string]interface{}
// 	mu   sync.RWMutex
// }

// // NewMemoryStore creates a new instance of MemoryStore
// func NewMemoryStore() *MemoryStore {
// 	return &MemoryStore{
// 		data: make(map[string]interface{}),
// 	}
// }

type Store struct {
	Cat MemoryStoreCat
	Dog MemoryStoreDog
}

type MemoryStoreCat interface {
	Set(ctx context.Context, key string, value *model.Cat)
	Get(ctx context.Context, key string) (*model.Cat, bool)
	Delete(ctx context.Context, key string)
	GetAll(ctx context.Context) []*model.Cat
	Exists(ctx context.Context, key string) bool
}

type MemoryStoreDog interface {
	Set(ctx context.Context, key string, value *model.Dog)
	Get(ctx context.Context, key string) (*model.Dog, bool)
	Delete(ctx context.Context, key string)
	GetAll(ctx context.Context) []*model.Dog
	Exists(ctx context.Context, key string) bool
}
