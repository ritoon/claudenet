package main

import (
	"fmt"
	"log"
	"net/http"

	"go-srp-cats/internal/handler"
	"go-srp-cats/internal/id"
	"go-srp-cats/internal/repository"
	"go-srp-cats/internal/routing"
	"go-srp-cats/internal/service"
	"go-srp-cats/internal/storage"
	"go-srp-cats/internal/validation"
)

func main() {
	// Initialize dependencies
	store := storage.NewMemoryStore()
	repo := repository.NewCatRepository(store)
	validator := validation.NewCatValidator()
	idGen := id.NewGenerator()

	// Create service
	svc := service.NewCatService(repo, validator, idGen)

	// Create handlers
	createHandler := handler.NewCreateCatHandler(svc)
	getCatHandler := handler.NewGetCatHandler(svc)
	getAllHandler := handler.NewGetAllCatsHandler(svc)
	updateHandler := handler.NewUpdateCatHandler(svc)
	deleteHandler := handler.NewDeleteCatHandler(svc)

	// Create router
	router := routing.NewRouter(createHandler, getCatHandler, getAllHandler, updateHandler, deleteHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Starting Cat CRUD API server on http://localhost%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  POST   /cats       - Create a new cat")
	fmt.Println("  GET    /cats       - Get all cats")
	fmt.Println("  GET    /cats/{id}  - Get a specific cat")
	fmt.Println("  PUT    /cats/{id}  - Update a cat")
	fmt.Println("  DELETE /cats/{id}  - Delete a cat")

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
