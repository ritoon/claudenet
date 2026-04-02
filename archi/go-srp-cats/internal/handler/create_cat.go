package handler

import (
	"net/http"

	"go-srp-cats/internal/errors"
	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// CreateCatHandler handles POST requests to create a new cat.
// SRP: This handler has ONE reason to change - if the CREATE endpoint's logic changes.
// It only does: decode request → call service → encode response.
// Each handler is independent, so changing one endpoint doesn't affect others.
type CreateCatHandler struct {
	service *service.CatService
}

// NewCreateCatHandler creates a new create cat handler.
func NewCreateCatHandler(svc *service.CatService) *CreateCatHandler {
	return &CreateCatHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *CreateCatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Decode request
	req, err := serialization.DecodeCreateCatRequest(r.Body)
	if err != nil {
		serialization.WriteError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Call service
	cat, err := h.service.CreateCat(req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		if errors.IsValidationError(err) {
			serialization.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		serialization.WriteError(w, http.StatusInternalServerError, "failed to create cat")
		return
	}

	// Encode response
	response := serialization.NewCatResponse(cat)
	serialization.WriteJSON(w, http.StatusCreated, response)
}
