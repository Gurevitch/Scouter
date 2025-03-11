package main

import (
	"log"
	"net/http"

	"bitbucket.org/Local/Scouter/api/router"
	"bitbucket.org/Local/Scouter/pkg/db"
	"bitbucket.org/Local/Scouter/pkg/handler"
	"bitbucket.org/Local/Scouter/pkg/repository"
	"bitbucket.org/Local/Scouter/pkg/service"
)

func main() {
	// Connect to database
	database := db.ConnectDB()

	db.MigrateDB()
	// Initialize repository
	playerRepo := repository.NewPlayerRepository(database)

	// Initialize service
	playerService := service.NewPlayerService(playerRepo)

	// Initialize handler
	playerHandler := handler.NewPlayerHandler(playerService)
	// Initialize repository
	TeamRepo := repository.NewTeamRepository(database)

	// Initialize service
	TeamService := service.NewTeamService(TeamRepo)

	// Initialize handler
	teamHandler := handler.NewTeamHandler(TeamService)

	// Setup router
	r := router.SetupRouter(playerHandler, teamHandler)

	// Serve React app (static files)
	// Ensure the path is correct (adjust based on where your build directory is located)
	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs) // Serve static files from the build directory

	// Start the server on port 8080
	log.Println("🚀 Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
