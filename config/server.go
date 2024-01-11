package config

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

type serverConfig struct {
	URI  string
	Port string
	Host string
}

func NewServerConfig() serverConfig {
	return serverConfig{
		URI:  ServerHost,
		Port: ServerPort,
		Host: fmt.Sprintf("%s:%s", ServerHost, ServerPort),
	}
}
