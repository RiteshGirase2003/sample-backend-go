// internal/database/connection.go
package database

import (
	"account-guru/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	// Dynamically read DB config values
	host := config.Config.String("db.host")
	port := config.Config.Int("db.port")
	user := config.Config.String("db.user")
	password := config.Config.String("db.password")
	dbname := config.Config.String("db.dbname")

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	log.Println("Successfully connected to the database.")
	return nil
}
