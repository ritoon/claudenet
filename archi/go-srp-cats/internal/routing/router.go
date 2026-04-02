package routing

import (
	"net/http"
	"strings"

	"go-srp-cats/internal/handler"
)

// Router maps HTTP requests to their corresponding handlers.
// SRP: Only reason to change - if URL routes or HTTP method mappings change.
// Changing route logic doesn't touch individual handler files.
type Router struct {
	createCatHandler   *handler.CreateCatHandler
	getCatHandler      *handler.GetCatHandler
	getAllCatsHandler  *handler.GetAllCatsHandler
	updateCatHandler   *handler.UpdateCatHandler
	deleteCatHandler   *handler.DeleteCatHandler
}

// NewRouter creates a new router with all handlers.
func NewRouter(
	createCat *handler.CreateCatHandler,
	getCat *handler.GetCatHandler,
	getAll *handler.GetAllCatsHandler,
	update *handler.UpdateCatHandler,
	delete *handler.DeleteCatHandler,
) *Router {
	return &Router{
		createCatHandler:   createCat,
		getCatHandler:      getCat,
		getAllCatsHandler:  getAll,
		updateCatHandler:   update,
		deleteCatHandler:   delete,
	}
}

// ServeHTTP routes incoming requests to the appropriate handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimSuffix(req.URL.Path, "/")

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

	// 404 Not Found
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"not found"}`))
}
