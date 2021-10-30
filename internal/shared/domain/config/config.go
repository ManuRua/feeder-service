package config

// Configurations exported
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Environment string
}

// ServerConfig exported
type ServerConfig struct {
	Network   string
	Port      int
	ConnLimit uint8
	Timeout   uint8
}

// DatabaseConfig exported
type DatabaseConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	SSLMode    string
}
