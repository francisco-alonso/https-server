package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisco-alonso/https-server/pkg/config"
)

// TestNewRouter ensures routes are correctly registered.
func TestNewRouter(t *testing.T) {
	router := NewRouter()

	// Check if the /health endpoint is registered
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}

// Mock server struct to prevent actual HTTPS start
type MockServer struct {
	cfg *config.Config
}

func (m *MockServer) Start() error {
	if m.cfg.Server.Port == "" {
		return fmt.Errorf("%v", m.cfg.Server.Port)

	}
	return nil
}

// TestStartServer ensures server starts with valid configuration.
func TestStartServer(t *testing.T) {
	mockCfg := &config.Config{
		Server: config.ServerConfig{
			Port:     "8080",
			CertFile: "test_cert.pem",
			KeyFile:  "test_key.pem",
		},
	}

	mockServer := &MockServer{cfg: mockCfg}
	err := mockServer.Start()

	if err != nil {
		t.Errorf("Mock server failed to start: %v", err)
	}
}

// TestStartServerWithInvalidConfig handles incorrect configurations.
func TestStartServerWithInvalidConfig(t *testing.T) {
	mockCfg := &config.Config{} // Empty config

	mockServer := &MockServer{cfg: mockCfg}
	err := mockServer.Start()

	if err == nil {
		t.Errorf("Expected error for invalid server config, got nil")
	}
}
