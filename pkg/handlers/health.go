package handlers

import "net/http"

// HealthHandler responds with a simple OK message.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}