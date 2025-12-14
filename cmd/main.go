// cmd/main.go
package main

import (
	"account-guru/config"
	"account-guru/database"
	"account-guru/internal/routes"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	// Load dynamic config
	if err := config.LoadConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to load config")
		return
	}

	// Initialize DB connection
	if err := database.InitDB(); err != nil {
		log.Fatal().Err(err).Msg("Database init failed")
	}

	// Register routes
	router := routes.InitRoutes()

	// Start server
	port := config.Config.String("port")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Info().Msgf("Server starting on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
