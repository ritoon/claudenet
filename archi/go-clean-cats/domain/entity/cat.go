package entity

// Cat represents a cat entity with all its attributes.
// This is the core business rule at the innermost layer of clean architecture.
// It knows nothing about databases, HTTP, or any framework.
type Cat struct {
	ID    string // Unique identifier for the cat
	Name  string // Name of the cat
	Breed string // Breed of the cat
	Age   int    // Age of the cat in years
	Color string // Color/pattern of the cat
}

// NewCat creates a new Cat instance with validation
func NewCat(id, name, breed string, age int, color string) *Cat {
	return &Cat{
		ID:    id,
		Name:  name,
		Breed: breed,
		Age:   age,
		Color: color,
	}
}
