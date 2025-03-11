package router

import (
	"net/http"

	"bitbucket.org/Local/Scouter/pkg/handler"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

// SetupRouter initializes routes
func SetupRouter(PlayerHandler *handler.PlayerHandler, teamHandler *handler.TeamHandler) http.Handler {
	r := mux.NewRouter()

	// Player routes
	r.HandleFunc("/register", PlayerHandler.HTTPHandlePlayerInsert).Methods("POST")
	r.HandleFunc("/player/{id}/stats", PlayerHandler.GetPlayerStats).Methods("GET") // Player stats route

	// Team routes
	r.HandleFunc("/teams", teamHandler.GetTeams).Methods("GET")
	r.HandleFunc("/teams", teamHandler.CreateTeams).Methods("POST")

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(r) // Wrap the router with CORS handler
}
