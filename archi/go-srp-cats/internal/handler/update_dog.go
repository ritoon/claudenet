package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// UpdateDogHandler handles PUT requests to update a dog.
// SRP: This handler has ONE reason to change - if the PUT /:id dog endpoint's logic changes.
type UpdateDogHandler struct {
	service *service.DogService
}

// NewUpdateDogHandler creates a new update dog handler.
func NewUpdateDogHandler(svc *service.DogService) *UpdateDogHandler {
	return &UpdateDogHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *UpdateDogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract ID from path: /dogs/{id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		serialization.WriteError(w, http.StatusBadRequest, "dog ID is required")
		return
	}
	id := parts[len(parts)-1]

	// Decode request
	req, err := serialization.DecodeUpdateDogRequest(r.Body)
	if err != nil {
		serialization.WriteError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Call service
	dog, err := h.service.UpdateDog(id, req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.IsValidationError(err) {
			serialization.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to update dog")
		return
	}

	// Encode response
	response := serialization.NewDogResponse(dog)
	serialization.WriteJSON(w, http.StatusOK, response)
}
