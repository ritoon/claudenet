package main

import (
	"fmt"
	"log"
	"net/http"

	// Inbound adapter (driving the application)
	httpAdapter "go-hexagonal-cats/internal/adapters/inbound/http"
	// Outbound adapter (driven by the application)
	memoryAdapter "go-hexagonal-cats/internal/adapters/outbound/memory"
	// Business logic
	"go-hexagonal-cats/internal/core/service"
)

func main() {
	// Dependency Injection: Wire the application together
	// Start with the outbound adapters (driven side)
	catRepository := memoryAdapter.NewCatRepository()

	// Create the application service (business logic)
	// It depends on the repository port
	catService := service.NewCatApplicationService(catRepository)

	// Create the inbound adapter (driving side)
	// It depends on the service port
	httpHandler := httpAdapter.NewHandler(catService)

	// Set up HTTP routes
	mux := http.NewServeMux()
	httpHandler.RegisterRoutes(mux)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Cat API server starting on http://localhost%s\n", port)
	fmt.Println("API endpoints:")
	fmt.Println("  POST   /api/cats          - Create a new cat")
	fmt.Println("  GET    /api/cats          - Get all cats")
	fmt.Println("  GET    /api/cats/:id      - Get a specific cat")
	fmt.Println("  PUT    /api/cats/:id      - Update a cat")
	fmt.Println("  DELETE /api/cats/:id      - Delete a cat")
	fmt.Println()

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
