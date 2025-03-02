package repository

import (
	"bitbucket.org/Local/Scouter/consts"
	"bitbucket.org/Local/Scouter/pkg/models"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

// NewTeamRepository creates a new instance of TeamRepository
func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

// GetAllTeams retrieves all teams and their players from the database
func (r *TeamRepository) GetAllTeams() ([]models.Team, error) {
	var teams []models.Team

	// Use Preload to eagerly load the associated players for each team
	result := r.db.Preload("Players").Find(&teams)

	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}

// CreateTeam adds a new team
func (r *TeamRepository) CreateTeam(team *models.Team) error {
	return r.db.Create(team).Error
}

// CreateTeamsAndPlayers inserts teams and their associated players
func (r *TeamRepository) CreateTeams() error {
	teams := []models.Team{
		{
			Name:  "Tottenham Hotspur",
			Value: 850.5,
			Players: []models.Player{
				{Name: "Harry Kane", Position: models.Position(consts.Goalkeeper), Wages: 100.0},
				{Name: "Son Heung-min", Position: "Winger", Wages: 85.5, ContractInYears: 3},
				
			},
		},
		{
			Name:  "Manchester United",
			Value: 920.3,
			Players: []models.Player{
				{Name: "Bruno Fernandes", Position: "Midfielder", Wages: 90.0},
				{Name: "Marcus Rashford", Position: "Forward", Wages: 85.0, ContractInYears: 2, Nationality: "England"},
			},
		},
	}

	for _, team := range teams {
		// Check if the team exists based on the Name and create it if it doesn't exist
		var existingTeam models.Team
		err := r.db.Where("name = ?", team.Name).First(&existingTeam).Error
		if err != nil && err.Error() == "record not found" {
			// Team does not exist, create it
			err = r.db.Create(&team).Error
			if err != nil {
				return err
			}
		} else if err == nil {
			// Team exists, update it if necessary (e.g., Value)
			team.ID = existingTeam.ID // Ensure to update the existing team's ID
			err = r.db.Save(&team).Error
			if err != nil {
				return err
			}
		}

		// Insert or update players for this team
		for _, player := range team.Players {
			var existingPlayer models.Player
			err := r.db.Where("name = ?", player.Name).First(&existingPlayer).Error
			if err != nil && err.Error() == "record not found" {
				// Player does not exist, create it
				newPlayer := models.Player{
					Name:            player.Name,
					Position:        player.Position,
					Wages:           player.Wages,
					ContractInYears: player.ContractInYears,
					Nationality:     player.Nationality,
				}
				err = r.db.Create(&newPlayer).Error
				if err != nil {
					return err
				}
			} else if err == nil {
				// Player exists, update the player if necessary
				existingPlayer.Position = player.Position
				existingPlayer.Wages = player.Wages
				existingPlayer.ContractInYears = player.ContractInYears
				existingPlayer.Nationality = player.Nationality
				err = r.db.Save(&existingPlayer).Error
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
