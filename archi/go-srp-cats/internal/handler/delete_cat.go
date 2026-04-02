package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// DeleteCatHandler handles DELETE requests to remove a cat.
// SRP: This handler has ONE reason to change - if the DELETE /:id endpoint's logic changes.
// It only does: extract ID → call service → return response.
type DeleteCatHandler struct {
	service *service.CatService
}

// NewDeleteCatHandler creates a new delete cat handler.
func NewDeleteCatHandler(svc *service.CatService) *DeleteCatHandler {
	return &DeleteCatHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *DeleteCatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract ID from path: /cats/{id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		serialization.WriteError(w, http.StatusBadRequest, "cat ID is required")
		return
	}
	id := parts[len(parts)-1]

	// Call service
	err := h.service.DeleteCat(id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to delete cat")
		return
	}

	// Return 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
