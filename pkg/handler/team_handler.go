package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/Local/Scouter/pkg/service"
)

type TeamHandler struct {
	teamService *service.TeamService
}

// NewTeamHandler initializes a TeamHandler
func NewTeamHandler(teamService *service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

// GetTeams handles GET /teams
func (h *TeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.teamService.GetAllTeams()
	if err != nil {
		http.Error(w, "Failed to fetch teams", http.StatusInternalServerError)
		return
	}

	// Set the response header to indicate the response is in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the teams as JSON response
	if err := json.NewEncoder(w).Encode(teams); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// CreateTeams handles POST /teams
func (h *TeamHandler) CreateTeams(w http.ResponseWriter, r *http.Request) {
	err := h.teamService.CreateTeams()
	if err != nil {
		http.Error(w, "Failed to create teams", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Teams created successfully"))
}
