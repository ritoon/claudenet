package presenter

import "go-clean-cats/domain/entity"

// DogResponse represents the JSON response structure for a dog.
// This is in the adapter layer because it's specific to HTTP/JSON representation.
type DogResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// DogPresenter converts domain entities to API responses
type DogPresenter struct{}

// NewDogPresenter creates a new presenter instance
func NewDogPresenter() *DogPresenter {
	return &DogPresenter{}
}

// PresentDog converts a Dog entity to a DogResponse
func (p *DogPresenter) PresentDog(dog *entity.Dog) DogResponse {
	return DogResponse{
		ID:    dog.ID,
		Name:  dog.Name,
		Breed: dog.Breed,
		Age:   dog.Age,
		Color: dog.Color,
	}
}

// PresentDogs converts multiple Dog entities to DogResponse list
func (p *DogPresenter) PresentDogs(dogs []*entity.Dog) []DogResponse {
	responses := make([]DogResponse, len(dogs))
	for i, dog := range dogs {
		responses[i] = p.PresentDog(dog)
	}
	return responses
}
