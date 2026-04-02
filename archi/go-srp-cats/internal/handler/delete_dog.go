package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// DeleteDogHandler handles DELETE requests to remove a dog.
// SRP: This handler has ONE reason to change - if the DELETE /:id dog endpoint's logic changes.
type DeleteDogHandler struct {
	service *service.DogService
}

// NewDeleteDogHandler creates a new delete dog handler.
func NewDeleteDogHandler(svc *service.DogService) *DeleteDogHandler {
	return &DeleteDogHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *DeleteDogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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
	err := h.service.DeleteDog(id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to delete dog")
		return
	}

	// Return 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
