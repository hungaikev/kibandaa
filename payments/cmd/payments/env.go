package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// LoadDevEnv loads .env file if present
func LoadDevEnv(log *zerolog.Logger) {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Error loading .env file")
		}
	}
}
