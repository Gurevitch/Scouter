package models

// Enum-like constants for predefined positions
type Position string

// Player struct represents a football player
type Player struct {
	ID              uint     `gorm:"primaryKey"`
	Name            string   `gorm:"not null"`
	TeamID          uint     `gorm:"not null"`
	Age             int64    `gorm:"not null"`
	Position        Position `gorm:"type:varchar(20);not null"`
	Nationality     string   `gorm:"not null"`
	Wages           float64  // Allow NULL values
	ContractInYears int64    // Allow NULL values
}
