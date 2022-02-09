package config

import (
	"os"
	"strings"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	DBName           string `yaml:"dbName"`
	RabbitMQUri      string `yaml:"rabbitmq"`
}

func (c *Config) SetRabbitHost() {
	host := os.Getenv("RABBIT_HOST")
	c.RabbitMQUri = strings.ReplaceAll(c.RabbitMQUri, "{RABBIT_HOST}", host)
}
func (c *Config) SetMongoHost() {
	host := os.Getenv("MONGO_HOST")
	c.ConnectionString = strings.ReplaceAll(c.ConnectionString, "{MONGO_HOST}", host)
}
