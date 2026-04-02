package validation

import (
	"fmt"
	"go-srp-cats/internal/entity"
)

// DogValidator validates Dog business rules.
// SRP: Only reason to change - if business validation rules for dogs change.
type DogValidator struct{}

// NewDogValidator creates a new validator.
func NewDogValidator() *DogValidator {
	return &DogValidator{}
}

// ValidateCreate validates a Dog for creation (ID should not be set yet).
func (v *DogValidator) ValidateCreate(dog *entity.Dog) error {
	if err := v.validateBasicFields(dog); err != nil {
		return err
	}
	return nil
}

// ValidateUpdate validates a Dog for update.
func (v *DogValidator) ValidateUpdate(dog *entity.Dog) error {
	if dog.ID == "" {
		return fmt.Errorf("dog ID is required for update")
	}
	if err := v.validateBasicFields(dog); err != nil {
		return err
	}
	return nil
}

// validateBasicFields validates common fields used in both create and update.
func (v *DogValidator) validateBasicFields(dog *entity.Dog) error {
	if dog.Name == "" {
		return fmt.Errorf("dog name is required")
	}
	if len(dog.Name) > 100 {
		return fmt.Errorf("dog name must be 100 characters or less")
	}

	if dog.Breed == "" {
		return fmt.Errorf("dog breed is required")
	}
	if len(dog.Breed) > 100 {
		return fmt.Errorf("dog breed must be 100 characters or less")
	}

	if dog.Color == "" {
		return fmt.Errorf("dog color is required")
	}
	if len(dog.Color) > 100 {
		return fmt.Errorf("dog color must be 100 characters or less")
	}

	if dog.Age < 0 || dog.Age > 50 {
		return fmt.Errorf("dog age must be between 0 and 50")
	}

	return nil
}
