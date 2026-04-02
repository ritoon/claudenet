package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-layered-cats/handler"
	"go-layered-cats/repository"
	"go-layered-cats/service"
	"go-layered-cats/storage"
	"go-layered-cats/storage/postgres"
	"go-layered-cats/storage/sqlite"
)

func main() {
	// Layer 1: Storage (Data Access)
	// Initialize the in-memory storage
	var Env string = os.Getenv("ENV")

	Env = "production" // Set to "production" for SQLite, otherwise it will use PostgreSQL

	var store *storage.Store
	if Env == "production" {
		sqliteStore, err := sqlite.New("cats.db")
		if err != nil {
			log.Fatalf("Failed to initialize SQLite storage: %v", err)
		}
		store = sqliteStore
	} else {
		postgresStore, err := postgres.New("host=localhost user=postgres password=postgres dbname=cats port=5432 sslmode=disable TimeZone=Asia/Shanghai")
		if err != nil {
			log.Fatalf("Failed to initialize PostgreSQL storage: %v", err)
		}
		store = postgresStore
	}

	// Layer 2: Repository (Data Abstraction)
	// Create repository with storage dependency
	catRepo := repository.NewCatRepository(store)
	dogRepo := repository.NewDogRepository(store)

	// Layer 3: Service (Business Logic)
	// Create service with repository dependency
	catService := service.NewCatService(catRepo)
	dogService := service.NewDogService(dogRepo)

	// Layer 4: Handler (Presentation/HTTP)
	// Create handler with service dependency
	catHandler := handler.NewCatHandler(catService)
	dogHandler := handler.NewDogHandler(dogService)

	// Setup HTTP routes
	http.Handle("/cats", catHandler)
	http.Handle("/cats/", catHandler)
	http.Handle("/dogs", dogHandler)
	http.Handle("/dogs/", dogHandler)

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
	fmt.Println("  GET    /dogs             - Get all dogs")
	fmt.Println("  GET    /dogs/{id}        - Get a specific dog")
	fmt.Println("  POST   /dogs             - Create a new dog")
	fmt.Println("  PUT    /dogs/{id}        - Update a specific dog")
	fmt.Println("  DELETE /dogs/{id}        - Delete a specific dog")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
