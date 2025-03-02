package handler

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/Local/Scouter/pkg/models"
	"bitbucket.org/Local/Scouter/pkg/service"
)

type PlayerHandler struct {
	playerService *service.PlayerService
}

// NewPlayerHandler initializes a PlayerHandler
func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

// HandlePlayerInsert handles inserting players
func (h *PlayerHandler) HandlePlayerInsert(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	// Decode the request body into the player struct
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the CreatePlayers method and check for any errors
	err := h.playerService.CreatePlayers(&player)
	if err != nil {
		// If an error occurs, return an internal server error response
		http.Error(w, "Failed to insert player", http.StatusInternalServerError)
		return
	}

	// Respond with a successful message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Player inserted successfully"))
}
