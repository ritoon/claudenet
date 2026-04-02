package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	BaseAPI = "https://rickandmortyapi.com/api"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/caracters", handlerCaracters)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func handlerCaracters(c *gin.Context) {
	// 1 appeler getCaracter()

	// 2 initialiser une structure de type any

	// 3 réaliser un bind avec le package encoding/json

	// 4 Renvoyer le payload

}

func getCaracter() ([]byte, error) {
	ctx := context.Background()
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseAPI+"/character", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
