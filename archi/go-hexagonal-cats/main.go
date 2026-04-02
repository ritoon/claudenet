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
	dogRepository := memoryAdapter.NewDogRepository()

	// Create the application services (business logic)
	// They depend on the repository ports
	catService := service.NewCatApplicationService(catRepository)
	dogService := service.NewDogApplicationService(dogRepository)

	// Create the inbound adapter (driving side)
	// It depends on the service ports
	httpHandler := httpAdapter.NewHandler(catService, dogService)

	// Set up HTTP routes
	mux := http.NewServeMux()
	httpHandler.RegisterRoutes(mux)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Cat & Dog API server starting on http://localhost%s\n", port)
	fmt.Println("API endpoints:")
	fmt.Println("  POST   /api/cats          - Create a new cat")
	fmt.Println("  GET    /api/cats          - Get all cats")
	fmt.Println("  GET    /api/cats/:id      - Get a specific cat")
	fmt.Println("  PUT    /api/cats/:id      - Update a cat")
	fmt.Println("  DELETE /api/cats/:id      - Delete a cat")
	fmt.Println("  POST   /api/dogs          - Create a new dog")
	fmt.Println("  GET    /api/dogs          - Get all dogs")
	fmt.Println("  GET    /api/dogs/:id      - Get a specific dog")
	fmt.Println("  PUT    /api/dogs/:id      - Update a dog")
	fmt.Println("  DELETE /api/dogs/:id      - Delete a dog")
	fmt.Println("  GET    /health            - Health check")
	fmt.Println()

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
