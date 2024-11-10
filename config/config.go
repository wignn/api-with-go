package config

import (
	"log"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT reguired"`
	DBHost     string `env:"DB_HOST reguired"`
	DBName     string `env:"DB_NAME reguired"`
	DBUser     string `env:"DB_USER reguired"`
	DBPassword string `env:"DB_PASS reguired"`
	DBSSLmode  string `env:"DB_SSL reguired"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Error parsing .env file")
	}

	return config

}
