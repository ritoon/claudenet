package model

// Cat represents a cat entity in the system
type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}
