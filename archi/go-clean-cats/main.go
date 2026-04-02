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
	dogRepository := persistence.NewMemoryDogRepository()

	// Use case layer
	catUseCase := usecase.NewCatUseCase(catRepository)
	dogUseCase := usecase.NewDogUseCase(dogRepository)

	// Adapter layer: presenter and controller
	catPresenter := presenter.NewCatPresenter()
	catController := controller.NewCatController(catUseCase, catPresenter)

	dogPresenter := presenter.NewDogPresenter()
	dogController := controller.NewDogController(dogUseCase, dogPresenter)

	// Infrastructure layer: routing
	handler := router.SetupRouter(catController, dogController)

	// Start the server
	port := ":8080"
	fmt.Printf("Starting Cat & Dog CRUD API server on %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /health            - Health check")
	fmt.Println("  POST   /cats              - Create a new cat")
	fmt.Println("  GET    /cats              - Get all cats")
	fmt.Println("  GET    /cats/{id}         - Get a specific cat")
	fmt.Println("  PUT    /cats/{id}         - Update a cat")
	fmt.Println("  DELETE /cats/{id}         - Delete a cat")
	fmt.Println("  POST   /dogs              - Create a new dog")
	fmt.Println("  GET    /dogs              - Get all dogs")
	fmt.Println("  GET    /dogs/{id}         - Get a specific dog")
	fmt.Println("  PUT    /dogs/{id}         - Update a dog")
	fmt.Println("  DELETE /dogs/{id}         - Delete a dog")
	fmt.Println()

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
