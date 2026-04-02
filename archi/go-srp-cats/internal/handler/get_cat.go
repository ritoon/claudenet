package handler

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// GetCatHandler handles GET requests for a specific cat.
// SRP: This handler has ONE reason to change - if the GET/:id endpoint's logic changes.
// It only does: extract ID → call service → encode response.
type GetCatHandler struct {
	service *service.CatService
}

// NewGetCatHandler creates a new get cat handler.
func NewGetCatHandler(svc *service.CatService) *GetCatHandler {
	return &GetCatHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *GetCatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
	cat, err := h.service.GetCatByID(id)
	if err != nil {
		if errors.IsNotFoundError(err) {
			serialization.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to get cat")
		return
	}

	// Encode response
	response := serialization.NewCatResponse(cat)
	serialization.WriteJSON(w, http.StatusOK, response)
}
