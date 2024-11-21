package config

import (
	"os"
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
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}