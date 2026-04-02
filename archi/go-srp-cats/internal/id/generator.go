package id

import (
	"crypto/rand"
	"fmt"
)

// Generator is responsible for generating unique IDs.
// SRP: Only reason to change - if we want a different ID generation strategy (UUID v4, sequential, etc.).
// This isolation means switching from UUIDs to any other format requires changes ONLY in this file.
type Generator struct{}

// NewGenerator creates a new ID generator.
func NewGenerator() *Generator {
	return &Generator{}
}

// Generate creates a new unique ID using crypto/rand for high security.
// Returns a 16-byte random ID as a hexadecimal string.
func (g *Generator) Generate() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
