package domain

// Dog is a type alias for Cat - they share the same structure
type Dog = Cat

// CreateDogRequest is the input for creating a new dog
type CreateDogRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// UpdateDogRequest is the input for updating a dog
type UpdateDogRequest struct {
	Name  *string `json:"name,omitempty"`
	Breed *string `json:"breed,omitempty"`
	Age   *int    `json:"age,omitempty"`
	Color *string `json:"color,omitempty"`
}
