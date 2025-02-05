package server

import (
	"log"
	"net/http"
	"time"

	"github.com/francisco-alonso/https-server/pkg/config"
	"github.com/francisco-alonso/https-server/pkg/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	
	 // Health-check endpoint
	 router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	 
	 
	 // Apply middleware for logging
	 router.Use(LoggingMiddleware)
	 
	 return router
}

// StartServer creates and starts the HTTPS server.
func StartServer(cfg *config.Config) {
    router := NewRouter()

    server := &http.Server{
        Handler:      router,
        Addr:         ":" + cfg.Server.Port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Printf("Starting HTTPS server on port %s", cfg.Server.Port)
    if err := server.ListenAndServeTLS(cfg.Server.CertFile, cfg.Server.KeyFile); err != nil {
        log.Fatalf("HTTPS server failed: %v", err)
    }
}