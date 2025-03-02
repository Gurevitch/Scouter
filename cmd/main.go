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
	playerService := service.NewPlayerHandler(playerRepo)

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

	// Start server
	log.Println("🚀 Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
