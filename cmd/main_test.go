package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/francisco-alonso/https-server/pkg/config"
)

// Setup function to create a temporary test config file
func setupTestConfigFile() string {
	testConfigContent := `[server]
		port = "8080"
		cert_file = "test_cert.pem"
		key_file = "test_key.pem"`

	testConfigPath := "configs/config.toml"

	// Create test config directory if it doesn't exist
	os.MkdirAll("configs", os.ModePerm)

	// Write the test config to a file
	err := os.WriteFile(testConfigPath, []byte(testConfigContent), 0644)
	if err != nil {
		log.Fatalf("Failed to create test config file: %v", err)
	}
	return testConfigPath
}

// TestLoadConfig ensures that configuration is properly loaded from the TOML file.
func TestLoadConfig(t *testing.T) {
	testConfigPath := setupTestConfigFile()
	cfg := config.LoadConfig(testConfigPath)

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected Port to be 8080, got %s", cfg.Server.Port)
	}
	if cfg.Server.CertFile != "test_cert.pem" {
		t.Errorf("Expected CertFile to be test_cert.pem, got %s", cfg.Server.CertFile)
	}
	if cfg.Server.KeyFile != "test_key.pem" {
		t.Errorf("Expected KeyFile to be test_key.pem, got %s", cfg.Server.KeyFile)
	}
}

// Mock server struct to avoid starting a real HTTPS server
type MockServer struct{}

func (m *MockServer) Start(cfg config.Config) error {
	if cfg.Server.Port == "" {
		return fmt.Errorf("%v", cfg.Server.Port)
	}
	return nil
}

// TestStartServer ensures that the server receives the correct configuration.
func TestStartServer(t *testing.T) {
	testConfigPath := setupTestConfigFile()
	cfg := config.LoadConfig(testConfigPath)

	mockServer := &MockServer{}
	err := mockServer.Start(*cfg)

	if err != nil {
		t.Errorf("Mock server failed to start: %v", err)
	}
}

// TestStartServerWithInvalidConfig checks handling of invalid configurations.
func TestStartServerWithInvalidConfig(t *testing.T) {
	cfg := config.Config{} // Empty config

	mockServer := &MockServer{}
	err := mockServer.Start(cfg)

	if err == nil {
		t.Errorf("Expected error for invalid server config, got nil")
	}
}

// Cleanup function to remove test artifacts
func cleanupTestConfigFile() {
	os.RemoveAll("configs")
}

// TestMain is the entry point for setup and teardown in testing
func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	// Cleanup test files
	cleanupTestConfigFile()
	// Exit with the test status code
	os.Exit(code)
}
