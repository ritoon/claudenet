package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-clean-cats/adapter/presenter"
	"go-clean-cats/usecase"
)

// CatController handles HTTP requests related to cats
// It converts HTTP requests to use case calls and HTTP responses
type CatController struct {
	catUseCase  *usecase.CatUseCase
	catPresenter *presenter.CatPresenter
}

// NewCatController creates a new cat controller
func NewCatController(catUseCase *usecase.CatUseCase, catPresenter *presenter.CatPresenter) *CatController {
	return &CatController{
		catUseCase:  catUseCase,
		catPresenter: catPresenter,
	}
}

// CreateCatRequest represents the request body for creating a cat
type CreateCatRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// UpdateCatRequest represents the request body for updating a cat
type UpdateCatRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// HandleCreateCat handles POST /cats requests
func (cc *CatController) HandleCreateCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateCatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cat, err := cc.catUseCase.CreateCat(req.ID, req.Name, req.Breed, req.Age, req.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cc.catPresenter.PresentCat(cat))
}

// HandleGetCat handles GET /cats/:id requests
func (cc *CatController) HandleGetCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Cat ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	cat, err := cc.catUseCase.GetCatByID(id)
	if err != nil {
		http.Error(w, "Cat not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cc.catPresenter.PresentCat(cat))
}

// HandleListCats handles GET /cats requests
func (cc *CatController) HandleListCats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cats, err := cc.catUseCase.ListAllCats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cc.catPresenter.PresentCats(cats))
}

// HandleUpdateCat handles PUT /cats/:id requests
func (cc *CatController) HandleUpdateCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Cat ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	var req UpdateCatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cat, err := cc.catUseCase.UpdateCat(id, req.Name, req.Breed, req.Age, req.Color)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cc.catPresenter.PresentCat(cat))
}

// HandleDeleteCat handles DELETE /cats/:id requests
func (cc *CatController) HandleDeleteCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[len(parts)-1] == "" {
		http.Error(w, "Cat ID is required", http.StatusBadRequest)
		return
	}

	id := parts[len(parts)-1]

	err := cc.catUseCase.DeleteCat(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
