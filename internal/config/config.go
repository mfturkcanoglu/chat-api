package config

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

type Env string

const (
	EnvFile    Env = "env.json"
	EnvDevFile Env = "env.dev.json"
)

var CurrentEnvFile Env = EnvDevFile

type Config struct {
	Port string `json:"port"`
	log  echo.Logger
}

func NewConfig(log echo.Logger) *Config {
	config := &Config{
		log: log,
	}
	return config.loadConfig()
}

func (c *Config) loadConfig() *Config {
	f, err := os.Open(string(CurrentEnvFile))

	if err != nil {
		log.Fatalf("Could not find env file: %s", string(CurrentEnvFile))
	}

	defer f.Close()

	json.NewDecoder(f).Decode(&c)
	return c
}
