package model

import (
	"errors"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// Cat represents a cat entity in the system
type Cat struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Breed string    `json:"breed"`
	Age   int       `json:"age"`
	Color string    `json:"color"`
}

func (c *Cat) BeforeCreate(tx *gorm.DB) (err error) {
	c.UUID = uuid.New()

	if !c.IsValid() {
		err = errors.New("can't save invalid data")
	}
	return
}

func (c *Cat) IsValid() bool {
	return c.Name != "" && c.Breed != "" && c.Age > 0 && c.Color != ""
}
