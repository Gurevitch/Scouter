package service

import (
	"bitbucket.org/Local/Scouter/pkg/models"
	"bitbucket.org/Local/Scouter/pkg/repository"
)

type PlayerService struct {
	repo *repository.PlayerRepository
}

// NewPlayerService initializes a PlayerHandler
func NewPlayerService(playerService *repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: playerService}
}

// CreatePlayers inserts a player into the database
func (s *PlayerService) CreatePlayers(player *models.Player) error {
	// Ensure you are returning any error that might occur while creating the player
	return s.repo.CreatePlayer(player)
}
func (s *PlayerService) GetPlayerByID(id uint) (*models.Player, error) {
	return s.repo.GetPlayerByID(id)
}
func (s *PlayerService) HandlePlayerInsert(req models.Player) error {
	return s.repo.HandlePlayerInsert(req)
}
func (s *PlayerService) GetPlayerStats(playerID string) (playerStats models.Player, err error) {
	return s.repo.GetPlayerStats(playerID) // Pass pointer to playerStats

}
