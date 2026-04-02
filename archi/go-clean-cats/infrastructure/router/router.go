package router

import (
	"net/http"
	"strings"

	"go-clean-cats/adapter/controller"
)

// SetupRouter configures all HTTP routes for the application
// This is in the infrastructure layer as it deals with the HTTP framework (net/http)
func SetupRouter(catController *controller.CatController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remove trailing slash for consistency
		path := strings.TrimSuffix(r.URL.Path, "/")

		switch {
		// POST /cats - Create a new cat
		case r.Method == http.MethodPost && path == "/cats":
			catController.HandleCreateCat(w, r)

		// GET /cats - List all cats
		case r.Method == http.MethodGet && path == "/cats":
			catController.HandleListCats(w, r)

		// GET /cats/{id} - Get a single cat
		case r.Method == http.MethodGet && strings.HasPrefix(path, "/cats/"):
			catController.HandleGetCat(w, r)

		// PUT /cats/{id} - Update a cat
		case r.Method == http.MethodPut && strings.HasPrefix(path, "/cats/"):
			catController.HandleUpdateCat(w, r)

		// DELETE /cats/{id} - Delete a cat
		case r.Method == http.MethodDelete && strings.HasPrefix(path, "/cats/"):
			catController.HandleDeleteCat(w, r)

		// 404 - Not Found
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})
}
