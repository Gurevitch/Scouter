package models

import (
	"gorm.io/gorm"
)

// Team struct represents a football team
type Team struct {
	gorm.Model
	Name    string   `json:"name"`
	Value   float64  `json:"value"`
	Goals   int64    `json:"goals"`
	Players []Player `json:"players"`
}
