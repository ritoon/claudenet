package serialization

import (
	"encoding/json"
	"io"
)

// CreateCatRequest is the request DTO for creating a cat.
// SRP: This file has ONE reason to change - if the input format changes.
// Moving from JSON to XML, YAML, or Protocol Buffers only affects this file.
type CreateCatRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// UpdateCatRequest is the request DTO for updating a cat.
type UpdateCatRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// DecodeCreateCatRequest decodes a JSON request body into a CreateCatRequest.
func DecodeCreateCatRequest(body io.ReadCloser) (*CreateCatRequest, error) {
	defer body.Close()
	var req CreateCatRequest
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// DecodeUpdateCatRequest decodes a JSON request body into an UpdateCatRequest.
func DecodeUpdateCatRequest(body io.ReadCloser) (*UpdateCatRequest, error) {
	defer body.Close()
	var req UpdateCatRequest
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// CreateDogRequest is the request DTO for creating a dog.
type CreateDogRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// UpdateDogRequest is the request DTO for updating a dog.
type UpdateDogRequest struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// DecodeCreateDogRequest decodes a JSON request body into a CreateDogRequest.
func DecodeCreateDogRequest(body io.ReadCloser) (*CreateDogRequest, error) {
	defer body.Close()
	var req CreateDogRequest
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// DecodeUpdateDogRequest decodes a JSON request body into an UpdateDogRequest.
func DecodeUpdateDogRequest(body io.ReadCloser) (*UpdateDogRequest, error) {
	defer body.Close()
	var req UpdateDogRequest
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
