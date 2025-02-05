package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// ServerConfig holds TLS and port settings.
type ServerConfig struct {
	Port     string `toml:"port"`
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

// Route defines a single HTTP route.
type Route struct {
	Path    string `toml:"path"`
	Method  string `toml:"method"`
	Handler string `toml:"handler"`
}

// Config is the top-level configuration structure.
type Config struct {
	Server ServerConfig `toml:"server"`
	Routes []Route      `toml:"routes"`
}

// LoadConfig loads configuration from the provided TOML file path.
func LoadConfig(path string) *Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	return &config
}
