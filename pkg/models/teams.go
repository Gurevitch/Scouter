package models

import (
	"gorm.io/gorm"
)

// Team struct represents a football team
type Team struct {
	gorm.Model          // This includes ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string   `json:"name"`
	Value      float64  `json:"value"`
	Players    []Player `json:"players"`
}
