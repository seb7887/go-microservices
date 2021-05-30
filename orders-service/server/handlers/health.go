package handlers

import (
	"fmt"
	"net/http"
)

// Health is the endpoint for simple health check; it returns 200 OK
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "ok")
}