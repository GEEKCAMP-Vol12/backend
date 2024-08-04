package config

import (
	"os"
)

type Config struct {
    ServerAddress string
    DBHost        string
    DBUser        string
    DBPassword    string
    DBName        string
    DBPort        string
}

func LoadConfig() *Config {
    return &Config{
        ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
        DBHost:        getEnv("DB_HOST", "localhost"),
        DBUser:        getEnv("DB_USER", "postgres"),
        DBPassword:    getEnv("DB_PASSWORD", "password"),
        DBName:        getEnv("DB_NAME", "vol12_server"),
        DBPort:        getEnv("DB_PORT", "5411"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
