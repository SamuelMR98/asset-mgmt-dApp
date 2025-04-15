package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            int
	PostgresDSN     string
	RedisAddr       string
	RedisPassword   string
	RedisDB         int
	EthereumNodeURL string
}

// Read .env file and set environment variables
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return &Config{
		Port:            getEnvAsInt("PORT", 8080),
		PostgresDSN:     getEnv("POSTGRES_DSN", "user=postgres password=postgres dbname=asset_mgmt sslmode=disable"),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		RedisDB:         getEnvAsInt("REDIS_DB", 0),
		EthereumNodeURL: getEnv("ETHEREUM_NODE_URL", "http://localhost:8545"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		var value int
		_, err := fmt.Sscanf(valueStr, "%d", &value)
		if err == nil {
			return value
		}
	}
	return defaultVal
}
