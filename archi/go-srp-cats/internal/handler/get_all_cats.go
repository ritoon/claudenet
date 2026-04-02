package handler

import (
	"net/http"

	"go-srp-cats/internal/serialization"
	"go-srp-cats/internal/service"
)

// GetAllCatsHandler handles GET requests to retrieve all cats.
// SRP: This handler has ONE reason to change - if the GET /cats endpoint's logic changes.
// It only does: call service → encode response.
type GetAllCatsHandler struct {
	service *service.CatService
}

// NewGetAllCatsHandler creates a new get all cats handler.
func NewGetAllCatsHandler(svc *service.CatService) *GetAllCatsHandler {
	return &GetAllCatsHandler{service: svc}
}

// ServeHTTP handles the HTTP request.
func (h *GetAllCatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serialization.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Call service
	cats, err := h.service.GetAllCats()
	if err != nil {
		serialization.WriteError(w, http.StatusInternalServerError, "failed to retrieve cats")
		return
	}

	// Convert to response DTOs
	responses := make([]*serialization.CatResponse, len(cats))
	for i, cat := range cats {
		responses[i] = serialization.NewCatResponse(cat)
	}

	// Encode response
	serialization.WriteJSON(w, http.StatusOK, responses)
}
