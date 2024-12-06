package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func InitConfig() *Config {
	return &Config{
		Token: getToken("TOKEN", ""),
	}
}

func getToken(key string, defaultVal string) string {
	if err := godotenv.Load(); err != nil {
	  log.Fatal("Error loading .env file")
	}

	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultVal
}