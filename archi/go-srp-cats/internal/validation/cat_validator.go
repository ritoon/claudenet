package validation

import (
	"fmt"
	"go-srp-cats/internal/entity"
)

// CatValidator validates Cat business rules.
// SRP: Only reason to change - if business validation rules change.
// Examples: name length requirements, age constraints, breed allowed values, color constraints.
// Moving validation rules is isolated to this file only.
type CatValidator struct{}

// NewCatValidator creates a new validator.
func NewCatValidator() *CatValidator {
	return &CatValidator{}
}

// ValidateCreate validates a Cat for creation (ID should not be set yet).
func (v *CatValidator) ValidateCreate(cat *entity.Cat) error {
	if err := v.validateBasicFields(cat); err != nil {
		return err
	}
	return nil
}

// ValidateUpdate validates a Cat for update.
func (v *CatValidator) ValidateUpdate(cat *entity.Cat) error {
	if cat.ID == "" {
		return fmt.Errorf("cat ID is required for update")
	}
	if err := v.validateBasicFields(cat); err != nil {
		return err
	}
	return nil
}

// validateBasicFields validates common fields used in both create and update.
func (v *CatValidator) validateBasicFields(cat *entity.Cat) error {
	if cat.Name == "" {
		return fmt.Errorf("cat name is required")
	}
	if len(cat.Name) > 100 {
		return fmt.Errorf("cat name must be 100 characters or less")
	}

	if cat.Breed == "" {
		return fmt.Errorf("cat breed is required")
	}
	if len(cat.Breed) > 100 {
		return fmt.Errorf("cat breed must be 100 characters or less")
	}

	if cat.Color == "" {
		return fmt.Errorf("cat color is required")
	}
	if len(cat.Color) > 100 {
		return fmt.Errorf("cat color must be 100 characters or less")
	}

	if cat.Age < 0 || cat.Age > 50 {
		return fmt.Errorf("cat age must be between 0 and 50")
	}

	return nil
}
