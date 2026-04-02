package handler

import (
	"net/http"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// CreateDogHandler handles POST requests to create a new dog.
// SRP: This handler has ONE reason to change - if the CREATE dog endpoint's logic changes.
type CreateDogHandler struct {
	service *service.DogService
}

// NewCreateDogHandler creates a new create dog handler.
func NewCreateDogHandler(svc *service.DogService) *CreateDogHandler {
	return &CreateDogHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *CreateDogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Decode request
	req, err := serialization.DecodeCreateDogRequest(r.Body)
	if err != nil {
		serialization.WriteError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Call service
	dog, err := h.service.CreateDog(req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		if errors.IsValidationError(err) {
			serialization.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to create dog")
		return
	}

	// Encode response
	response := serialization.NewDogResponse(dog)
	serialization.WriteJSON(w, http.StatusCreated, response)
}
