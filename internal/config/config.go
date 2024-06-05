package config

import "os"

// Config holds the configuration settings for the application.
type Config struct {
	ServerPort string
}

// LoadConfig loads configuration from environment variables or defaults.
func LoadConfig() *Config {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // default port
	}
	return &Config{
		ServerPort: ":" + port,
	}
}
