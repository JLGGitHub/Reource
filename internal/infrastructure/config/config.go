package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDSN string
	ServerPort  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from environment variables")
	}
	return &Config{
		DatabaseDSN: getEnv("DATABASE_DSN", "root:admin123@tcp(localhost:3306)/resource?parseTime=true"),
		ServerPort:  getEnv("SERVER_PORT", "8081"),
	}
}

// Helper para obtener variables de entorno con un valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
