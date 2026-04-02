package storage

import (
	"sync"
)

// MemoryStore provides thread-safe in-memory key-value storage.
// SRP: Only reason to change - if the storage backend changes (e.g., database, file system, cache).
// This abstraction means switching storage engines only requires changes in this file.
// The generic nature makes it reusable for different types of data.
type MemoryStore struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

// NewMemoryStore creates a new in-memory store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		items: make(map[string]interface{}),
	}
}

// Set stores a value with the given key.
func (s *MemoryStore) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[key] = value
}

// Get retrieves a value by key. Returns the value and a boolean indicating if it exists.
func (s *MemoryStore) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.items[key]
	return val, exists
}

// Delete removes a value by key.
func (s *MemoryStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, key)
}

// GetAll returns a copy of all stored items.
func (s *MemoryStore) GetAll() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Create a copy to avoid external modifications
	result := make(map[string]interface{})
	for k, v := range s.items {
		result[k] = v
	}
	return result
}

// Exists checks if a key exists in the store.
func (s *MemoryStore) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.items[key]
	return exists
}
