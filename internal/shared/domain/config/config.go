package config

import (
	"os"
	"strconv"
)

// Configurations exported
type Config struct {
	Server      ServerConfig
	Environment string
}

// ServerConfig exported
type ServerConfig struct {
	Port      int
	ConnLimit uint8
	Timeout   uint8
}

func LoadConfig() Config {
	var port, connLimit, timeout int

	if os.Getenv("API_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("API_PORT"))
	} else {
		port = 4000
	}

	if os.Getenv("API_CONN_LIMIT") != "" {
		connLimit, _ = strconv.Atoi(os.Getenv("API_CONN_LIMIT"))
	} else {
		connLimit = 5
	}

	if os.Getenv("API_TIMEOUT") != "" {
		timeout, _ = strconv.Atoi(os.Getenv("API_TIMEOUT"))
	} else {
		timeout = 60
	}

	srvConfig := ServerConfig{
		Port:      port,
		ConnLimit: uint8(connLimit),
		Timeout:   uint8(timeout),
	}

	return Config{
		Server:      srvConfig,
		Environment: os.Getenv("APP_ENV"),
	}
}
