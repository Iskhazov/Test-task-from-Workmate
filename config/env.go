package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config - Application configuration
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

// Envs — initialized configuration
var Envs = initConfig()

// initConfig - initialize the Config structure
func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "moviehub"),
	}
}

// getEnv returns the value of an environment variable
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}

// GetEnvAsInt tries to read an environment variable as an integer, or returns the default value
func GetEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
