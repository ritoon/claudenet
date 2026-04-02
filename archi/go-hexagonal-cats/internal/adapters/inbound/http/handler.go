package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-hexagonal-cats/internal/core/domain"
	"go-hexagonal-cats/internal/core/ports"
)

// Handler is the HTTP adapter that drives the application
// It receives HTTP requests and calls the inbound ports (CatService, DogService)
type Handler struct {
	catService ports.CatService
	dogService ports.DogService
}

// NewHandler creates a new HTTP handler
// This is where we inject the CatService and DogService dependencies
func NewHandler(catService ports.CatService, dogService ports.DogService) *Handler {
	return &Handler{
		catService: catService,
		dogService: dogService,
	}
}

// RegisterRoutes registers all HTTP routes
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/cats", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateCat(w, r)
		case http.MethodGet:
			h.GetAllCats(w, r)
		default:
			h.methodNotAllowed(w)
		}
	})

	mux.HandleFunc("/api/cats/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/cats/")
		if id == "" {
			h.notFound(w)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetCatByID(w, r, id)
		case http.MethodPut:
			h.UpdateCat(w, r, id)
		case http.MethodDelete:
			h.DeleteCat(w, r, id)
		default:
			h.methodNotAllowed(w)
		}
	})

	mux.HandleFunc("/api/dogs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateDog(w, r)
		case http.MethodGet:
			h.GetAllDogs(w, r)
		default:
			h.methodNotAllowed(w)
		}
	})

	mux.HandleFunc("/api/dogs/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/dogs/")
		if id == "" {
			h.notFound(w)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetDogByID(w, r, id)
		case http.MethodPut:
			h.UpdateDog(w, r, id)
		case http.MethodDelete:
			h.DeleteDog(w, r, id)
		default:
			h.methodNotAllowed(w)
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		h.jsonResponse(w, http.StatusOK, map[string]string{"status": "ok"})
	})
}

// CreateCat handles POST /api/cats
func (h *Handler) CreateCat(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateCatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.badRequest(w, "invalid JSON: "+err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		h.badRequest(w, "name is required")
		return
	}

	// Call the inbound port (CatService)
	cat, err := h.catService.CreateCat(req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		h.badRequest(w, err.Error())
		return
	}

	h.jsonResponse(w, http.StatusCreated, cat)
}

// GetCatByID handles GET /api/cats/:id
func (h *Handler) GetCatByID(w http.ResponseWriter, r *http.Request, id string) {
	cat, err := h.catService.GetCatByID(id)
	if err != nil {
		h.notFound(w)
		return
	}

	h.jsonResponse(w, http.StatusOK, cat)
}

// GetAllCats handles GET /api/cats
func (h *Handler) GetAllCats(w http.ResponseWriter, r *http.Request) {
	cats, err := h.catService.GetAllCats()
	if err != nil {
		h.internalError(w, err.Error())
		return
	}

	h.jsonResponse(w, http.StatusOK, cats)
}

// UpdateCat handles PUT /api/cats/:id
func (h *Handler) UpdateCat(w http.ResponseWriter, r *http.Request, id string) {
	var req domain.UpdateCatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.badRequest(w, "invalid JSON: "+err.Error())
		return
	}

	// Call the inbound port (CatService)
	cat, err := h.catService.UpdateCat(id, req)
	if err != nil {
		h.notFound(w)
		return
	}

	h.jsonResponse(w, http.StatusOK, cat)
}

// DeleteCat handles DELETE /api/cats/:id
func (h *Handler) DeleteCat(w http.ResponseWriter, r *http.Request, id string) {
	err := h.catService.DeleteCat(id)
	if err != nil {
		h.notFound(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateDog handles POST /api/dogs
func (h *Handler) CreateDog(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateDogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.badRequest(w, "invalid JSON: "+err.Error())
		return
	}

	// Validate required fields
	if req.Name == "" {
		h.badRequest(w, "name is required")
		return
	}

	// Call the inbound port (DogService)
	dog, err := h.dogService.CreateDog(req.Name, req.Breed, req.Color, req.Age)
	if err != nil {
		h.badRequest(w, err.Error())
		return
	}

	h.jsonResponse(w, http.StatusCreated, dog)
}

// GetDogByID handles GET /api/dogs/:id
func (h *Handler) GetDogByID(w http.ResponseWriter, r *http.Request, id string) {
	dog, err := h.dogService.GetDogByID(id)
	if err != nil {
		h.notFound(w)
		return
	}

	h.jsonResponse(w, http.StatusOK, dog)
}

// GetAllDogs handles GET /api/dogs
func (h *Handler) GetAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := h.dogService.GetAllDogs()
	if err != nil {
		h.internalError(w, err.Error())
		return
	}

	h.jsonResponse(w, http.StatusOK, dogs)
}

// UpdateDog handles PUT /api/dogs/:id
func (h *Handler) UpdateDog(w http.ResponseWriter, r *http.Request, id string) {
	var req domain.UpdateDogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.badRequest(w, "invalid JSON: "+err.Error())
		return
	}

	// Call the inbound port (DogService)
	dog, err := h.dogService.UpdateDog(id, req)
	if err != nil {
		h.notFound(w)
		return
	}

	h.jsonResponse(w, http.StatusOK, dog)
}

// DeleteDog handles DELETE /api/dogs/:id
func (h *Handler) DeleteDog(w http.ResponseWriter, r *http.Request, id string) {
	err := h.dogService.DeleteDog(id)
	if err != nil {
		h.notFound(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper methods for HTTP responses

func (h *Handler) jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) badRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Handler) notFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
}

func (h *Handler) internalError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Handler) methodNotAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
}
