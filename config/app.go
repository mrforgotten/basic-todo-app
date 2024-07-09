package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort uint16
}

func GetAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	var port uint16
	parsed, err := strconv.ParseUint(os.Getenv("APP_PORT"), 10, 16)
	if err != nil {
		port = 3000
		log.Printf("Error during get app port: %v", err)
	} else {
		port = uint16(parsed)
	}

	opts := &AppConfig{
		AppPort: port,
	}

	return opts
}
