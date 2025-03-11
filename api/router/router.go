package router

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"bitbucket.org/Local/Scouter/pkg/handler"

	"github.com/gorilla/mux"
)

// SetupRouter initializes routes and serves static files for the frontend
func SetupRouter(playerHandler *handler.PlayerHandler, teamHandler *handler.TeamHandler) *mux.Router {
	r := mux.NewRouter()

	// Serve the React frontend (index.html)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving current directory:", err)
		return nil
	}
	parentDir := filepath.Dir(currentDir)
	clientBuildDir := path.Join(parentDir, "frontend", "build")
	clientStaticDir := path.Join(clientBuildDir, "static")

	// Serve index.html for the frontend
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(clientBuildDir, "index.html"))
	})

	// Serve static files (JS, CSS, etc.)
	fs := http.FileServer(http.Dir(clientStaticDir))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Player routes
	r.HandleFunc("/register", playerHandler.HTTPHandlePlayerInsert).Methods("POST")
	r.HandleFunc("/player/{id}/stats", playerHandler.GetPlayerStats).Methods("GET") // Player stats route

	// Team routes
	r.HandleFunc("/teams", teamHandler.GetTeams).Methods("GET")
	r.HandleFunc("/teams", teamHandler.CreateTeams).Methods("POST")

	// You can add CORS, authentication, or any other middleware as needed

	return r
}
