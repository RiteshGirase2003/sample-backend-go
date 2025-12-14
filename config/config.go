package config

import (
	"fmt"

	"github.com/knadh/koanf/v2"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/rs/zerolog/log"
)

var Config = koanf.New(".")

var Database DatabaseConfig

// DatabaseConfig holds the configuration for the database
type DatabaseConfig struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	DBName   string `koanf:"dbname"`
	SSLMode  string `koanf:"sslmode"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

func LoadConfig() error {
	path := "../config/config.dev.yaml"

	if err := Config.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Error().
			Err(err).
			Str("path", path).
			Msg("config_load_failed")
		return fmt.Errorf("config load failed: %w", err)
	}

	if err := Config.Unmarshal("database", &Database); err != nil {
		log.Error().
			Err(err).
			Msg("config_unmarshal_database_failed")
		return fmt.Errorf("config unmarshal database failed: %w", err)
	}

	log.Info().
		Str("path", path).
		Msg("config_loaded")

	return nil
}
