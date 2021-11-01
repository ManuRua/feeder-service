package config

// Configurations exported
type Config struct {
	Server      ServerConfig
	Environment string
}

// ServerConfig exported
type ServerConfig struct {
	Network   string
	Port      int
	ConnLimit uint8
	Timeout   uint8
}
