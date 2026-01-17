package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort    string
	APIHost    string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading directly from environment variables")
	}
	cfg := &Config{
		APIPort:    getEnv("API_PORT", "8999"),
		APIHost:    getEnv("API_HOST", "127.0.0.1"),
		DBUser:     getEnv(os.Getenv("DB_USER"), "taasadmin"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME", "taas"),
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("CRITICAL: DB_PASSWORD environment variable is not set")
	}
	return cfg, nil
}
