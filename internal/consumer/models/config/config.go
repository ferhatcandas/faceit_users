package config

import (
	"os"
	"strings"
)

type Config struct {
	RabbitMQUri string `yaml:"rabbitmq"`
}

func (c *Config) SetRabbitHost() {
	host := os.Getenv("RABBIT_HOST")
	c.RabbitMQUri = strings.ReplaceAll(c.RabbitMQUri, "{RABBIT_HOST}", host)
}
