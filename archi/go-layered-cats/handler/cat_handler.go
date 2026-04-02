package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"go-layered-cats/model"
	"go-layered-cats/service"
)

// CatHandler handles HTTP requests related to cats
type CatHandler struct {
	service *service.CatService
}

// NewCatHandler creates a new instance of CatHandler
func NewCatHandler(service *service.CatService) *CatHandler {
	return &CatHandler{
		service: service,
	}
}

// CreateCat handles POST /cats request
func (h *CatHandler) CreateCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var cat model.Cat
	if err := parseJSON(r.Body, &cat); err != nil {
		writeJSONError(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call service layer
	createdCat, err := h.service.CreateCat(&cat)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	writeJSON(w, createdCat, http.StatusCreated)
}

// GetCat handles GET /cats/{id} request
func (h *CatHandler) GetCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/cats/")

	// Call service layer
	cat, err := h.service.GetCatByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJSONError(w, err.Error(), http.StatusNotFound)
		} else {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Return success response
	writeJSON(w, cat, http.StatusOK)
}

// GetAllCats handles GET /cats request
func (h *CatHandler) GetAllCats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call service layer
	cats, err := h.service.GetAllCats()
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	writeJSON(w, cats, http.StatusOK)
}

// UpdateCat handles PUT /cats/{id} request
func (h *CatHandler) UpdateCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/cats/")

	// Parse request body
	var cat model.Cat
	if err := parseJSON(r.Body, &cat); err != nil {
		writeJSONError(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call service layer
	updatedCat, err := h.service.UpdateCat(id, &cat)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJSONError(w, err.Error(), http.StatusNotFound)
		} else {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Return success response
	writeJSON(w, updatedCat, http.StatusOK)
}

// DeleteCat handles DELETE /cats/{id} request
func (h *CatHandler) DeleteCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/cats/")

	// Call service layer
	if err := h.service.DeleteCat(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJSONError(w, err.Error(), http.StatusNotFound)
		} else {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Return success response (204 No Content)
	w.WriteHeader(http.StatusNoContent)
}

// parseJSON is a helper function to parse JSON from request body
func parseJSON(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

// writeJSON is a helper function to write JSON response
func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Log error in production, but response is already sent
		_ = err
	}
}

// writeJSONError is a helper function to write JSON error response
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := map[string]string{
		"error": message,
	}
	writeJSON(w, errorResponse, statusCode)
}

// ServeHTTP implements http.Handler interface for routing
func (h *CatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Handle GET requests
	if r.Method == http.MethodGet {
		if path == "/cats" {
			h.GetAllCats(w, r)
			return
		} else if strings.HasPrefix(path, "/cats/") && len(path) > 6 {
			h.GetCat(w, r)
			return
		}
	}

	// Handle POST requests
	if r.Method == http.MethodPost && path == "/cats" {
		h.CreateCat(w, r)
		return
	}

	// Handle PUT requests
	if r.Method == http.MethodPut && strings.HasPrefix(path, "/cats/") && len(path) > 6 {
		h.UpdateCat(w, r)
		return
	}

	// Handle DELETE requests
	if r.Method == http.MethodDelete && strings.HasPrefix(path, "/cats/") && len(path) > 6 {
		h.DeleteCat(w, r)
		return
	}

	// Handle 404
	writeJSONError(w, "Not found", http.StatusNotFound)
}
