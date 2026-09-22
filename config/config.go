package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBURL        string `envconfig:"DB_URL" required:"true"`
	KafkaBrokers string `envconfig:"KAFKA_BROKERS" required:"true"`
	HTTPPort     string `envconfig:"HTTP_PORT" default:":8080"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
}

func Load() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}
