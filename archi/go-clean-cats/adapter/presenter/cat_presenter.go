package presenter

import "go-clean-cats/domain/entity"

// CatResponse represents the JSON response structure for a cat.
// This is in the adapter layer because it's specific to HTTP/JSON representation.
// The domain entity (Cat) knows nothing about JSON or HTTP concerns.
type CatResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// CatPresenter converts domain entities to API responses
type CatPresenter struct{}

// NewCatPresenter creates a new presenter instance
func NewCatPresenter() *CatPresenter {
	return &CatPresenter{}
}

// PresentCat converts a Cat entity to a CatResponse
func (p *CatPresenter) PresentCat(cat *entity.Cat) CatResponse {
	return CatResponse{
		ID:    cat.ID,
		Name:  cat.Name,
		Breed: cat.Breed,
		Age:   cat.Age,
		Color: cat.Color,
	}
}

// PresentCats converts multiple Cat entities to CatResponse list
func (p *CatPresenter) PresentCats(cats []*entity.Cat) []CatResponse {
	responses := make([]CatResponse, len(cats))
	for i, cat := range cats {
		responses[i] = p.PresentCat(cat)
	}
	return responses
}
