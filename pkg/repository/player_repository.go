package repository

import (
	"bitbucket.org/Local/Scouter/pkg/models"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

// NewPlayerRepository creates a new instance of TeamRepository
func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// GetAllTeams retrieves all teams from the database
func (r *PlayerRepository) GetAllPlayers() ([]models.Team, error) {
	var teams []models.Team
	result := r.db.Find(&teams)
	return teams, result.Error
}

// CreatePlayer inserts a player into the database
func (r *PlayerRepository) CreatePlayer(player *models.Player) error {
	return r.db.Create(player).Error
}
