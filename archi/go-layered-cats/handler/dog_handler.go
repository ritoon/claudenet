package handler

import (
	"net/http"
	"strings"

	"go-layered-cats/model"
	"go-layered-cats/service"
)

// DogHandler handles HTTP requests related to dogs
type DogHandler struct {
	service *service.DogService
}

// NewDogHandler creates a new instance of DogHandler
func NewDogHandler(service *service.DogService) *DogHandler {
	return &DogHandler{
		service: service,
	}
}

// CreateDog handles POST /dogs request
func (h *DogHandler) CreateDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var dog model.Dog
	if err := parseJSON(r.Body, &dog); err != nil {
		writeJSONError(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call service layer
	createdDog, err := h.service.CreateDog(&dog)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	writeJSON(w, createdDog, http.StatusCreated)
}

// GetDog handles GET /dogs/{id} request
func (h *DogHandler) GetDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/dogs/")

	// Call service layer
	dog, err := h.service.GetDogByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJSONError(w, err.Error(), http.StatusNotFound)
		} else {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Return success response
	writeJSON(w, dog, http.StatusOK)
}

// GetAlldogs handles GET /dogs request
func (h *DogHandler) GetAlldogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call service layer
	dogs, err := h.service.GetAllDogs()
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	writeJSON(w, dogs, http.StatusOK)
}

// UpdateDog handles PUT /dogs/{id} request
func (h *DogHandler) UpdateDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/dogs/")

	// Parse request body
	var dog model.Dog
	if err := parseJSON(r.Body, &dog); err != nil {
		writeJSONError(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call service layer
	updatedDog, err := h.service.UpdateDog(id, &dog)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeJSONError(w, err.Error(), http.StatusNotFound)
		} else {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Return success response
	writeJSON(w, updatedDog, http.StatusOK)
}

// DeleteDog handles DELETE /dogs/{id} request
func (h *DogHandler) DeleteDog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	id := strings.TrimPrefix(r.URL.Path, "/dogs/")

	// Call service layer
	if err := h.service.DeleteDog(id); err != nil {
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

// ServeHTTP implements http.Handler interface for routing
func (h *DogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Handle GET requests
	if r.Method == http.MethodGet {
		if path == "/dogs" {
			h.GetAlldogs(w, r)
			return
		} else if strings.HasPrefix(path, "/dogs/") && len(path) > 6 {
			h.GetDog(w, r)
			return
		}
	}

	// Handle POST requests
	if r.Method == http.MethodPost && path == "/dogs" {
		h.CreateDog(w, r)
		return
	}

	// Handle PUT requests
	if r.Method == http.MethodPut && strings.HasPrefix(path, "/dogs/") && len(path) > 6 {
		h.UpdateDog(w, r)
		return
	}

	// Handle DELETE requests
	if r.Method == http.MethodDelete && strings.HasPrefix(path, "/dogs/") && len(path) > 6 {
		h.DeleteDog(w, r)
		return
	}

	// Handle 404
	writeJSONError(w, "Not found", http.StatusNotFound)
}
