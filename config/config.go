package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port            int
	PostgresDSN     string
	RedisAddr       string
	RedisPassword   string
	RedisDB         int
	EthereumNodeURL string
}

func LoadConfig() *Config {
	return &Config{
		Port:            8080,
		PostgresDSN:     getEnv("POSTGRES_DSN", "postgres://user:password@localhost:5432/assetdb?sslmode=disable"),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		RedisDB:         getEnvAsInt("REDIS_DB", 0),
		EthereumNodeURL: getEnv("ETHEREUM_NODE_URL", "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"),
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
