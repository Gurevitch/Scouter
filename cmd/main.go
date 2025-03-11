package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/Local/Scouter/api/router"
	"bitbucket.org/Local/Scouter/pkg/db"
	"bitbucket.org/Local/Scouter/pkg/handler"
	"bitbucket.org/Local/Scouter/pkg/repository"
	"bitbucket.org/Local/Scouter/pkg/service"
)

const (
	port = "8080"
)

func main() {
	// Connect to database
	database := db.ConnectDB()

	db.MigrateDB()
	// Initialize repository for player
	playerRepo := repository.NewPlayerRepository(database)

	// Initialize service for player
	playerService := service.NewPlayerService(playerRepo)

	// Initialize handler for player
	playerHandler := handler.NewPlayerHandler(playerService)

	// Initialize repository for team
	teamRepo := repository.NewTeamRepository(database)

	// Initialize service for team
	teamService := service.NewTeamService(teamRepo)

	// Initialize handler for team
	teamHandler := handler.NewTeamHandler(teamService)

	// Initialize router with player and team handlers
	r := router.SetupRouter(playerHandler, teamHandler)

	// Start the server on port 8080
	fmt.Println("🚀 Server is running on port 8080")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
