package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-clean-cats/adapter/presenter"
	"go-clean-cats/usecase"
)

// DogController handles HTTP requests related to dogs
// It converts HTTP requests to use case calls and HTTP responses
type DogController struct {
	dogUseCase   *usecase.DogUseCase
	dogPresenter *presenter.DogPresenter
}

// NewDogController creates a new dog controller
func NewDogController(dogUseCase *usecase.DogUseCase, dogPresenter *presenter.DogPresenter) *DogController {
	return &DogController{
		dogUseCase:   dogUseCase,
		dogPresenter: dogPresenter,
	}
}

// CreateDogRequest represents the request body for creating a dog
type CreateDogRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// UpdateDogRequest represents the request body for updating a dog
type UpdateDogRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// HandleCreateDog handles POST /dogs requests
func (dc *DogController) HandleCreateDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateDogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dog, err := dc.dogUseCase.CreateDog(req.ID, req.Name, req.Breed, req.Age, req.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dc.dogPresenter.PresentDog(dog))
}

// HandleGetDog handles GET /dogs/:id requests
func (dc *DogController) HandleGetDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Dog ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	dog, err := dc.dogUseCase.GetDogByID(id)
	if err != nil {
		http.Error(w, "Dog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dc.dogPresenter.PresentDog(dog))
}

// HandleListDogs handles GET /dogs requests
func (dc *DogController) HandleListDogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dogs, err := dc.dogUseCase.ListAllDogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dc.dogPresenter.PresentDogs(dogs))
}

// HandleUpdateDog handles PUT /dogs/:id requests
func (dc *DogController) HandleUpdateDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Dog ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	var req UpdateDogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dog, err := dc.dogUseCase.UpdateDog(id, req.Name, req.Breed, req.Age, req.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dc.dogPresenter.PresentDog(dog))
}

// HandleDeleteDog handles DELETE /dogs/:id requests
func (dc *DogController) HandleDeleteDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Dog ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	err := dc.dogUseCase.DeleteDog(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
