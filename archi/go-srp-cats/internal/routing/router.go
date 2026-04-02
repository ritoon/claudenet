package routing

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/handler"
	"go-srp-cats/internal/serialization"
)

// Router maps HTTP requests to their corresponding handlers.
// SRP: Only reason to change - if URL routes or HTTP method mappings change.
// Changing route logic doesn't touch individual handler files.
type Router struct {
	createCatHandler  *handler.CreateCatHandler
	getCatHandler     *handler.GetCatHandler
	getAllCatsHandler  *handler.GetAllCatsHandler
	updateCatHandler  *handler.UpdateCatHandler
	deleteCatHandler  *handler.DeleteCatHandler
	createDogHandler  *handler.CreateDogHandler
	getDogHandler     *handler.GetDogHandler
	getAllDogsHandler  *handler.GetAllDogsHandler
	updateDogHandler  *handler.UpdateDogHandler
	deleteDogHandler  *handler.DeleteDogHandler
}

// NewRouter creates a new router with all handlers.
func NewRouter(
	createCat *handler.CreateCatHandler,
	getCat *handler.GetCatHandler,
	getAllCats *handler.GetAllCatsHandler,
	updateCat *handler.UpdateCatHandler,
	deleteCat *handler.DeleteCatHandler,
	createDog *handler.CreateDogHandler,
	getDog *handler.GetDogHandler,
	getAllDogs *handler.GetAllDogsHandler,
	updateDog *handler.UpdateDogHandler,
	deleteDog *handler.DeleteDogHandler,
) *Router {
	return &Router{
		createCatHandler:  createCat,
		getCatHandler:     getCat,
		getAllCatsHandler:  getAllCats,
		updateCatHandler:  updateCat,
		deleteCatHandler:  deleteCat,
		createDogHandler:  createDog,
		getDogHandler:     getDog,
		getAllDogsHandler:  getAllDogs,
		updateDogHandler:  updateDog,
		deleteDogHandler:  deleteDog,
	}
}

// ServeHTTP routes incoming requests to the appropriate handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimSuffix(req.URL.Path, "/")

	// GET /health - Health check
	if req.Method == http.MethodGet && path == "/health" {
		serialization.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
		return
	}

	// POST /cats - Create a new cat
	if req.Method == http.MethodPost && path == "/cats" {
		r.createCatHandler.ServeHTTP(w, req)
		return
	}

	// GET /cats - Get all cats
	if req.Method == http.MethodGet && path == "/cats" {
		r.getAllCatsHandler.ServeHTTP(w, req)
		return
	}

	// GET /cats/{id} - Get a specific cat
	if req.Method == http.MethodGet && strings.HasPrefix(path, "/cats/") {
		r.getCatHandler.ServeHTTP(w, req)
		return
	}

	// PUT /cats/{id} - Update a cat
	if req.Method == http.MethodPut && strings.HasPrefix(path, "/cats/") {
		r.updateCatHandler.ServeHTTP(w, req)
		return
	}

	// DELETE /cats/{id} - Delete a cat
	if req.Method == http.MethodDelete && strings.HasPrefix(path, "/cats/") {
		r.deleteCatHandler.ServeHTTP(w, req)
		return
	}

	// POST /dogs - Create a new dog
	if req.Method == http.MethodPost && path == "/dogs" {
		r.createDogHandler.ServeHTTP(w, req)
		return
	}

	// GET /dogs - Get all dogs
	if req.Method == http.MethodGet && path == "/dogs" {
		r.getAllDogsHandler.ServeHTTP(w, req)
		return
	}

	// GET /dogs/{id} - Get a specific dog
	if req.Method == http.MethodGet && strings.HasPrefix(path, "/dogs/") {
		r.getDogHandler.ServeHTTP(w, req)
		return
	}

	// PUT /dogs/{id} - Update a dog
	if req.Method == http.MethodPut && strings.HasPrefix(path, "/dogs/") {
		r.updateDogHandler.ServeHTTP(w, req)
		return
	}

	// DELETE /dogs/{id} - Delete a dog
	if req.Method == http.MethodDelete && strings.HasPrefix(path, "/dogs/") {
		r.deleteDogHandler.ServeHTTP(w, req)
		return
	}

	// 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"not found"}`))
}
