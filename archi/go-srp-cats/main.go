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
	// Initialize shared dependencies
	idGen := id.NewGenerator()

	// Cat dependencies
	catStore := storage.NewMemoryStore()
	catRepo := repository.NewCatRepository(catStore)
	catValidator := validation.NewCatValidator()
	catSvc := service.NewCatService(catRepo, catValidator, idGen)

	// Cat handlers
	createCatHandler := handler.NewCreateCatHandler(catSvc)
	getCatHandler := handler.NewGetCatHandler(catSvc)
	getAllCatsHandler := handler.NewGetAllCatsHandler(catSvc)
	updateCatHandler := handler.NewUpdateCatHandler(catSvc)
	deleteCatHandler := handler.NewDeleteCatHandler(catSvc)

	// Dog dependencies
	dogStore := storage.NewMemoryStore()
	dogRepo := repository.NewDogRepository(dogStore)
	dogValidator := validation.NewDogValidator()
	dogSvc := service.NewDogService(dogRepo, dogValidator, idGen)

	// Dog handlers
	createDogHandler := handler.NewCreateDogHandler(dogSvc)
	getDogHandler := handler.NewGetDogHandler(dogSvc)
	getAllDogsHandler := handler.NewGetAllDogsHandler(dogSvc)
	updateDogHandler := handler.NewUpdateDogHandler(dogSvc)
	deleteDogHandler := handler.NewDeleteDogHandler(dogSvc)

	// Create router
	router := routing.NewRouter(
		createCatHandler, getCatHandler, getAllCatsHandler, updateCatHandler, deleteCatHandler,
		createDogHandler, getDogHandler, getAllDogsHandler, updateDogHandler, deleteDogHandler,
	)

	// Start server
	port := ":8080"
	fmt.Printf("Starting Cat & Dog CRUD API server on http://localhost%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /health     - Health check")
	fmt.Println("  POST   /cats       - Create a new cat")
	fmt.Println("  GET    /cats       - Get all cats")
	fmt.Println("  GET    /cats/{id}  - Get a specific cat")
	fmt.Println("  PUT    /cats/{id}  - Update a cat")
	fmt.Println("  DELETE /cats/{id}  - Delete a cat")
	fmt.Println("  POST   /dogs       - Create a new dog")
	fmt.Println("  GET    /dogs       - Get all dogs")
	fmt.Println("  GET    /dogs/{id}  - Get a specific dog")
	fmt.Println("  PUT    /dogs/{id}  - Update a dog")
	fmt.Println("  DELETE /dogs/{id}  - Delete a dog")

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
