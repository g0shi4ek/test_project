package config

import(
	"os"
)

type Config struct{
	auth_port string 
}

func Load() *Config {
	return &Config{
		auth_port: os.Getenv("AUTH_PORT"),
	}
}