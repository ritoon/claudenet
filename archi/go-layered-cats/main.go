package main

import (
	"fmt"
	"log"
	"net/http"

	"go-layered-cats/handler"
	"go-layered-cats/repository"
	"go-layered-cats/service"
	"go-layered-cats/storage"
)

func main() {
	// Layer 1: Storage (Data Access)
	// Initialize the in-memory storage
	store := storage.NewMemoryStore()

	// Layer 2: Repository (Data Abstraction)
	// Create repository with storage dependency
	catRepo := repository.NewCatRepository(store)

	// Layer 3: Service (Business Logic)
	// Create service with repository dependency
	catService := service.NewCatService(catRepo)

	// Layer 4: Handler (Presentation/HTTP)
	// Create handler with service dependency
	catHandler := handler.NewCatHandler(catService)

	// Setup HTTP routes
	http.Handle("/cats", catHandler)
	http.Handle("/cats/", catHandler)

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy"}`)
	})

	// Start server
	port := ":8080"
	fmt.Printf("Cat CRUD API server starting on http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST   /cats             - Create a new cat")
	fmt.Println("  GET    /cats             - Get all cats")
	fmt.Println("  GET    /cats/{id}        - Get a specific cat")
	fmt.Println("  PUT    /cats/{id}        - Update a specific cat")
	fmt.Println("  DELETE /cats/{id}        - Delete a specific cat")
	fmt.Println("  GET    /health           - Health check")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
