package service

import (
	"bitbucket.org/Local/Scouter/pkg/models"
	"bitbucket.org/Local/Scouter/pkg/repository"
)

type TeamService struct {
	repo *repository.TeamRepository
}

// NewTeamService creates a new TeamService instance
func NewTeamService(repo *repository.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

// CreateTeams calls the repository to insert teams
func (s *TeamService) CreateTeams() error {
	return s.repo.CreateTeams()
}

// GetAllTeams retrieves all teams
func (s *TeamService) GetAllTeams() ([]models.Team, error) {
	return s.repo.GetAllTeams()
}
