package storage

import (
	"sync"
)

// MemoryStore provides in-memory data storage using a map
type MemoryStore struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// NewMemoryStore creates a new instance of MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]interface{}),
	}
}

// Set stores a value with the given key
func (m *MemoryStore) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Get retrieves a value by key
func (m *MemoryStore) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

// Delete removes a value by key
func (m *MemoryStore) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// GetAll returns all stored values as a slice of interfaces
func (m *MemoryStore) GetAll() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	values := make([]interface{}, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// Exists checks if a key exists in storage
func (m *MemoryStore) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.data[key]
	return exists
}
