package domain

// Cat represents a cat entity in the domain
type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// CreateCatRequest is the input for creating a new cat
type CreateCatRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// UpdateCatRequest is the input for updating a cat
type UpdateCatRequest struct {
	Name  *string `json:"name,omitempty"`
	Breed *string `json:"breed,omitempty"`
	Age   *int    `json:"age,omitempty"`
	Color *string `json:"color,omitempty"`
}
