package handler

import (
	"net/http"

	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// GetAllDogsHandler handles GET requests to retrieve all dogs.
// SRP: This handler has ONE reason to change - if the GET /dogs endpoint's logic changes.
type GetAllDogsHandler struct {
	service *service.DogService
}

// NewGetAllDogsHandler creates a new get all dogs handler.
func NewGetAllDogsHandler(svc *service.DogService) *GetAllDogsHandler {
	return &GetAllDogsHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *GetAllDogsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Call service
	dogs, err := h.service.GetAllDogs()
	if err != nil {
		serialization.WriteError(w, http.StatusInternalServerError, "failed to retrieve dogs")
		return
	}

	// Convert to response DTOs
	responses := make([]*serialization.DogResponse, len(dogs))
	for i, dog := range dogs {
		responses[i] = serialization.NewDogResponse(dog)
	}

	// Encode response
	serialization.WriteJSON(w, http.StatusOK, responses)
}
