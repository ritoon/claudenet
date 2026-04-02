package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// UpdateCatHandler handles PUT requests to update a cat.
// SRP: This handler has ONE reason to change - if the PUT /:id endpoint's logic changes.
// It only does: extract ID → decode request → call service → encode response.
type UpdateCatHandler struct {
	service *service.CatService
}

// NewUpdateCatHandler creates a new update cat handler.
func NewUpdateCatHandler(svc *service.CatService) *UpdateCatHandler {
	return &UpdateCatHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *UpdateCatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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

	// Decode request
	req, err := serialization.DecodeUpdateCatRequest(r.Body)
	if err != nil {
		serialization.WriteError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Call service
	cat, err := h.service.UpdateCat(id, req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.IsValidationError(err) {
			serialization.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to update cat")
		return
	}

	// Encode response
	response := serialization.NewCatResponse(cat)
	serialization.WriteJSON(w, http.StatusOK, response)
}
