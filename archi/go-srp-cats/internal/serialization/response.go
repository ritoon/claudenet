package serialization

import (
	"encoding/json"
	"net/http"

	"go-srp-cats/internal/entity"
)

// CatResponse is the response DTO for a Cat.
// SRP: This file has ONE reason to change - if the output format changes.
// Moving from JSON to XML, YAML, or Protocol Buffers only affects this file.
type CatResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// ErrorResponse is the response DTO for errors.
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewCatResponse converts an entity.Cat to a CatResponse.
func NewCatResponse(cat *entity.Cat) *CatResponse {
	return &CatResponse{
		ID:    cat.ID,
		Name:  cat.Name,
		Breed: cat.Breed,
		Color: cat.Color,
		Age:   cat.Age,
	}
}

// WriteJSON writes a response as JSON to the HTTP response writer.
func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// WriteError writes an error response as JSON.
func WriteError(w http.ResponseWriter, statusCode int, message string) error {
	return WriteJSON(w, statusCode, &ErrorResponse{Error: message})
}
