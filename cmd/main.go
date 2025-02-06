package main

import (
	"log"

	"github.com/francisco-alonso/https-server/pkg/config"
	"github.com/francisco-alonso/https-server/pkg/server"
)

func main() {
	// Load configuration from TOML file.
	cfg := config.LoadConfig("configs/config.toml")

	// Log the loaded configuration for debugging.
	log.Printf("Loaded configuration: %+v", cfg)

	// Start the HTTPS server with the configuration.
	server.StartServer(cfg)
}
