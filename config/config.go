package config

type Config struct {
	DBURL        string `envconfig:"DB_URL" required:"true"`
	KafkaBrokers string `envconfig:"KAFKA_BROKERS" required:"true"`
	HTTPPort     string `envconfig:"HTTP_PORT" default:":8080"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	REDISURL     string `envconfig:"REDIS_URL" default:""`
}
