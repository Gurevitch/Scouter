package models

import (
	"gorm.io/gorm"
)

type Position string

// Player struct represents a football player
type Player struct {
	gorm.Model
	Name            string        `gorm:"not null" json:"name"`
	TeamID          uint          `gorm:"not null" json:"team_id"`
	Team            Team          `gorm:"foreignKey:TeamID" json:"-"` // Ignore in JSON response
	TeamName        string        `gorm:"-" json:"team_name"`         // Virtual field, fetched manually
	Age             int64         `gorm:"not null" json:"age"`
	Position        Position      `gorm:"type:varchar(20);not null" json:"position"`
	Nationality     string        `gorm:"not null" json:"nationality"`
	Wages           float64       `json:"wages"`
	ContractInYears int64         `json:"contract_in_years"`
	SeasonStats     []SeasonStats `gorm:"foreignKey:PlayerID" json:"season_stats"`
}

type SeasonStats struct {
	gorm.Model
	PlayerID   uint   `gorm:"not null" json:"player_id"`
	SeasonYear string `gorm:"not null" json:"season_year"`
	Goals      int64  `gorm:"not null" json:"goals"`
	Assists    int64  `gorm:"not null" json:"assists"`
	Cards      Cards  `gorm:"embedded" json:"cards"`
}

// Ensure Cards struct has default values
type Cards struct {
	YellowCard int64 `gorm:"default:0" json:"yellow_card"`
	RedCard    int64 `gorm:"default:0" json:"red_card"`
}
