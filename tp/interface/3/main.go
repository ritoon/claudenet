package main

import (
	"context"
	"encoding/json"
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

type CaracterResponse map[string]any

func (cr *CaracterResponse) Validate() error {
	// Implémenter une validation de la structure du payload
	return nil
}

func handlerCaracters(c *gin.Context) {
	// 1 appeler getCaracter()
	data, err := getCaracter()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 2 initialiser une structure de type any
	// var payload any
	payload := make(CaracterResponse)

	// 3 réaliser un bind avec le package encoding/json

	//if err := c.BindJSON(&payload); err != nil {
	if err := json.Unmarshal(data, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4 Renvoyer le payload
	c.JSON(http.StatusOK, payload)
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
