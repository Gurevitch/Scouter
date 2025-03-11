package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/Local/Scouter/pkg/models"
	"bitbucket.org/Local/Scouter/pkg/service"
	"github.com/gorilla/mux"
)

type PlayerHandler struct {
	playerService *service.PlayerService
}

// NewPlayerHandler initializes a PlayerHandler
func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}
func (h *PlayerHandler) HTTPHandlePlayerInsert(w http.ResponseWriter, r *http.Request) {
	// Validate Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, `{"error": "Content-Type must be application/json"}`, http.StatusUnsupportedMediaType)
		return
	}

	// Decode JSON request into PlayerRequest
	var req models.Player
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Call the structured function
	if err := h.playerService.HandlePlayerInsert(req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Player inserted successfully"}`))
}

// GetPlayerStats fetches the stats of a player
func (h *PlayerHandler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID := vars["id"]

	// Fetch the stats for the player
	stats, err := h.playerService.GetPlayerStats(playerID)
	if err != nil {
		http.Error(w, "Player stats not found", http.StatusNotFound)
		return
	}

	// Send back the stats in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
