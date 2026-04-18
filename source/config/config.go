package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppPort    string
	AppVersion string

	DBHost string
	DBUser string
	DBPass string
	DBPort string
	DBName string

	JWTSecret string
}

func readVersionFile() string {
	content, err := os.ReadFile("version")
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(content))
}

func Load() *Config {
	_ = godotenv.Load()
	version := readVersionFile()

	return &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
		AppVersion: version,

		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
