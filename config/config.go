package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort        string
	AuthPort       string
	AuthDbHost     string
	AuthDbPort     string
	AuthDbUser     string
	AuthDbPassword string
	AuthDbName     string
	AuthDbSSL      string
	SecretKey      string
	RedisPort      string
	RedisName      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := Config{
		ApiPort:        os.Getenv("API_PORT"),
		AuthPort:       os.Getenv("AUTH_PORT"),
		AuthDbHost:     os.Getenv("AUTH_PG_HOST"),
		AuthDbPort:     os.Getenv("AUTH_PG_PORT"),
		AuthDbUser:     os.Getenv("AUTH_PG_USER"),
		AuthDbPassword: os.Getenv("AUTH_PG_PASSWORD"),
		AuthDbName:     os.Getenv("AUTH_PG_DBNAME"),
		AuthDbSSL:      os.Getenv("AUTH_PG_SSLMODE"),
		SecretKey:      os.Getenv("JWT_KEY"),
		RedisPort:      os.Getenv("REDIS_PORT"),
		RedisName:      os.Getenv("REDIS_NAME"),
	}
	return &cfg, nil
}
