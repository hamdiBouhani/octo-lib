package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {
	// Load .env file if present
	_ = godotenv.Load()

	cfg := Config{
		DBURL:        MustGet("DB_URL"),
		KafkaBrokers: MustGet("KAFKA_BROKERS"),
		HTTPPort:     GetOrDefault("HTTP_PORT", ":8080"),
		LogLevel:     GetOrDefault("LOG_LEVEL", "info"),
	}

	return cfg
}

func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}

func GetOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
