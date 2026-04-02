package entity

// Cat represents a cat with its attributes.
// SRP: This file has ONE reason to change - if the Cat data structure itself changes (adding/removing fields).
// No validation logic, no JSON handling, no storage logic here - just pure data.
type Cat struct {
	ID    string
	Name  string
	Breed string
	Age   int
	Color string
}
