package config

import (
	"os"
)

type Config struct {
	Port    string
	GinMode string
	LogFile string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
}

func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "debug"),
		LogFile: getEnv("LOG_FILE", "app.log"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "admin"),
		DBPass:  getEnv("DB_PASSWORD", ""),
		DBName:  getEnv("DB_NAME", "testdb"),
	}
}

func getEnv(key, defaultVaule string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVaule
}
