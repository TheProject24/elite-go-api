package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
	DBPassword string
	JWTSecret string
}

func LoadConfig() Config {
	Port := os.Getenv("PORT")
	dbPass := os.Getenv("DB_PASSWORD")
	jwtSecret := os.Getenv("JWT_SECRET")

	if dbPass == "" {
		log.Fatal("CRITICAL ERROR: DB_PASSWORD environmet variable is missing. Shutting down.")
	}

	if jwtSecret == "" {
		log.Fatal("CRITICAL ERROR: JWT_SECRET environment variable is missing. Shutting down")
	}

	if Port == "" {
		Port = "8080"
	}

	return Config{
		Port: Port,
		DBPassword: dbPass,
		JWTSecret: jwtSecret,
	}
}
