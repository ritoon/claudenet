package errors

import "fmt"

// DomainError defines custom error types for the domain.
// SRP: Only reason to change - if the error taxonomy of the domain changes.
// Each error type represents a specific domain error condition.

// ValidationError indicates that validation failed.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return "validation error: " + e.Message
}

// NotFoundError indicates that a resource was not found.
type NotFoundError struct {
	ID string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("cat not found: %s", e.ID)
}

// AlreadyExistsError indicates that a resource already exists.
type AlreadyExistsError struct {
	ID string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("cat already exists: %s", e.ID)
}

// IsValidationError checks if an error is a ValidationError.
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// IsNotFoundError checks if an error is a NotFoundError.
func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// IsAlreadyExistsError checks if an error is an AlreadyExistsError.
func IsAlreadyExistsError(err error) bool {
	_, ok := err.(*AlreadyExistsError)
	return ok
}
