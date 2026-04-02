package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

// parseJSON is a helper function to parse JSON from request body
func parseJSON(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

// writeJSON is a helper function to write JSON response
func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Log error in production, but response is already sent
		_ = err
	}
}

// writeJSONError is a helper function to write JSON error response
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := map[string]string{
		"error": message,
	}
	writeJSON(w, errorResponse, statusCode)
}
