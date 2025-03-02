package router

import (
	"bitbucket.org/Local/Scouter/pkg/handler"

	"github.com/gorilla/mux"
)

// SetupRouter initializes routes
func SetupRouter(PlayerHandler *handler.PlayerHandler, teamHandler *handler.TeamHandler) *mux.Router {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/register", PlayerHandler.HandlePlayerInsert).Methods("POST")

	// Team routes
	r.HandleFunc("/teams", teamHandler.GetTeams).Methods("GET")
	r.HandleFunc("/teams", teamHandler.CreateTeams).Methods("POST")

	return r
}
