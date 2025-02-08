package main

import (
	"log"

	"github.com/francisco-alonso/https-server/pkg/config"
	"github.com/francisco-alonso/https-server/pkg/mqtt"
	"github.com/francisco-alonso/https-server/pkg/server"
)

func main() {
	// Load configuration from TOML file.
	cfg := config.LoadConfig("configs/config.toml")

	// Log the loaded configuration for debugging.
	log.Printf("Loaded configuration: %+v", cfg)

	mqttAdapter := mqtt.NewAdapter()
	
	// For Kubernetes, use the service DNS (e.g., "tcp://mqtt-service:1883").
	// For local testing, you might use "tcp://localhost:1883".
	if err := mqttAdapter.Connect("tcp://mqtt-service:1883", "https-server-client"); err != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", err)
	}
	
	// Start the HTTPS server with the configuration.
	server.StartServer(cfg)
}
