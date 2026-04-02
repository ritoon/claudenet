package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-clean-cats/adapter/controller"
)

// SetupRouter configures all HTTP routes for the application
// This is in the infrastructure layer as it deals with the HTTP framework (net/http)
func SetupRouter(catController *controller.CatController, dogController *controller.DogController) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remove trailing slash for consistency
		path := strings.TrimSuffix(r.URL.Path, "/")

		switch {
		// GET /health - Health check
		case r.Method == http.MethodGet && path == "/health":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})

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

		// POST /dogs - Create a new dog
		case r.Method == http.MethodPost && path == "/dogs":
			dogController.HandleCreateDog(w, r)

		// GET /dogs - List all dogs
		case r.Method == http.MethodGet && path == "/dogs":
			dogController.HandleListDogs(w, r)

		// GET /dogs/{id} - Get a single dog
		case r.Method == http.MethodGet && strings.HasPrefix(path, "/dogs/"):
			dogController.HandleGetDog(w, r)

		// PUT /dogs/{id} - Update a dog
		case r.Method == http.MethodPut && strings.HasPrefix(path, "/dogs/"):
			dogController.HandleUpdateDog(w, r)

		// DELETE /dogs/{id} - Delete a dog
		case r.Method == http.MethodDelete && strings.HasPrefix(path, "/dogs/"):
			dogController.HandleDeleteDog(w, r)

		// 404 - Not Found
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})
}
