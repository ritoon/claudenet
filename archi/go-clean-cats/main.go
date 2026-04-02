package main

import (
	"fmt"
	"log"
	"net/http"

	"go-clean-cats/adapter/controller"
	"go-clean-cats/adapter/presenter"
	"go-clean-cats/infrastructure/persistence"
	"go-clean-cats/infrastructure/router"
	"go-clean-cats/usecase"
)

func main() {
	// Initialize the dependency chain (dependency injection)
	// This demonstrates how clean architecture allows easy swapping of implementations

	// Infrastructure layer: persistence
	catRepository := persistence.NewMemoryCatRepository()

	// Use case layer
	catUseCase := usecase.NewCatUseCase(catRepository)

	// Adapter layer: presenter and controller
	catPresenter := presenter.NewCatPresenter()
	catController := controller.NewCatController(catUseCase, catPresenter)

	// Infrastructure layer: routing
	handler := router.SetupRouter(catController)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting Cat CRUD API server on %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST   /cats              - Create a new cat")
	fmt.Println("  GET    /cats              - Get all cats")
	fmt.Println("  GET    /cats/{id}         - Get a specific cat")
	fmt.Println("  PUT    /cats/{id}         - Update a cat")
	fmt.Println("  DELETE /cats/{id}         - Delete a cat")
	fmt.Println()

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
