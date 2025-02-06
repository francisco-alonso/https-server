package server

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs incoming requests and the time taken to process them.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)
		log.Printf("Method: %s, URI: %s, Duration: %v", r.Method, r.RequestURI, duration)
	})
}