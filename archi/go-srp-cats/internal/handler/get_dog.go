package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// GetDogHandler handles GET requests for a specific dog.
// SRP: This handler has ONE reason to change - if the GET/:id dog endpoint's logic changes.
type GetDogHandler struct {
	service *service.DogService
}

// NewGetDogHandler creates a new get dog handler.
func NewGetDogHandler(svc *service.DogService) *GetDogHandler {
	return &GetDogHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *GetDogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

	// Call service
	dog, err := h.service.GetDogByID(id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to get dog")
		return
	}

	// Encode response
	response := serialization.NewDogResponse(dog)
	serialization.WriteJSON(w, http.StatusOK, response)
}
