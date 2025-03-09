package repository

import (
	"errors"
	"fmt"

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
func (r *PlayerRepository) GetPlayerByID(id uint) (*models.Player, error) {
	var player models.Player

	// Fetch player with team name using JOIN
	err := r.db.Table("players").
		Select("players.*, teams.name as team_name").
		Joins("JOIN teams ON teams.id = players.team_id").
		Where("players.id = ?", id).
		First(&player).Error

	if err != nil {
		return nil, err
	}

	return &player, nil
}
func (r *PlayerRepository) HandlePlayerInsert(req models.Player) error {
	// Validate required fields
	if req.Name == "" || req.TeamID == 0 || req.Age <= 0 || req.Nationality == "" {
		return fmt.Errorf("missing required player fields")
	}

	// Check if the team exists
	var team models.Team
	err := r.db.Where("id = ? or Name  = ?", req.TeamID, req.TeamName).First(&team).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Team not found, create a new one
			newTeam := models.Team{Name: req.TeamName, Value: 0} // Assign a default value
			if createErr := r.db.Create(&newTeam).Error; createErr != nil {
				return fmt.Errorf("failed to create team: %w", createErr)
			}
			// Assign the newly created team's ID to the player
			req.TeamID = newTeam.ID
			team = newTeam // Assign new team to variable
		} else {
			return fmt.Errorf("error finding team: %w", err)
		}
	}

	// Check if the player already exists based on both name and team_id
	var existingPlayer models.Player
	err = r.db.Where("name = ? AND team_id = ?", req.Name, req.TeamID).First(&existingPlayer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Player does not exist, create a new one
			newPlayer := models.Player{
				Name:            req.Name,
				TeamID:          req.TeamID,
				Age:             req.Age,
				Position:        models.Position(req.Position),
				Nationality:     req.Nationality,
				Wages:           req.Wages,
				ContractInYears: req.ContractInYears,
				SeasonStats:     req.SeasonStats, // Directly mapped
			}
			if createErr := r.db.Create(&newPlayer).Error; createErr != nil {
				return fmt.Errorf("error creating player: %w", createErr)
			}
		} else {
			return fmt.Errorf("error checking player: %w", err)
		}
	} else {
		// Player exists, update their details
		existingPlayer.Age = req.Age
		existingPlayer.Position = models.Position(req.Position)
		existingPlayer.Nationality = req.Nationality
		existingPlayer.Wages = req.Wages
		existingPlayer.ContractInYears = req.ContractInYears
		existingPlayer.SeasonStats = req.SeasonStats // Update stats

		// Ensure we're updating the correct player with the correct ID
		if updateErr := r.db.Save(&existingPlayer).Error; updateErr != nil {
			return fmt.Errorf("error updating player: %w", updateErr)
		}
	}

	return nil
}
