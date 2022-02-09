package config

import (
	"os"
	"strings"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	DBName           string `yaml:"dbName"`
	Port             string `yaml:"port"`
}

func (c *Config) SetMongoHost() {
	host := os.Getenv("MONGO_HOST")
	c.ConnectionString = strings.ReplaceAll(c.ConnectionString, "{MONGO_HOST}", host)
}
