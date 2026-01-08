package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port           string
	OpenAIKey      string
	InternalToken  string
	MaxWorkers     int
	QueueSize      int
	OpenAIEndpoint string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		OpenAIKey:      getEnv("OPENAI_API_KEY", ""),
		InternalToken:  getEnv("INTERNAL_AUTH_TOKEN", "default-secret-token"),
		MaxWorkers:     getEnvAsInt("MAX_WORKERS", 10),
		QueueSize:      getEnvAsInt("QUEUE_SIZE", 100),
		OpenAIEndpoint: "https://api.openai.com/v1/chat/completions",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, fallback int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}
