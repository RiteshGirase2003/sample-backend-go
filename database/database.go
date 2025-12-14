package database

import (
	"fmt"

	"account-guru/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	if err := config.LoadConfig(); err != nil {
		log.Error().
			Err(err).
			Msg("database_init_config_failed")
		return err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
		config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().
			Err(err).
			Msg("database_connection_failed")
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	log.Info().
		Msg("database_connected")

	return nil
}
