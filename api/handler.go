// Package api provides the API handlers for the go-api-template.
// This is a template repository for building Go API services with a standardized structure.
// Customize the handlers in this package to implement your specific API functionality.
package api

import (
	"net/http"
	"time"

	"github.com/dbubel/intake/v2"
	"github.com/sirupsen/logrus"
)

// APIHandler handles API requests and implements the application's endpoints.
// In this template pattern, the APIHandler contains all the handler methods for your API.
// When creating your own API based on this template:
// - Add domain-specific methods to this struct
// - Consider breaking large handler sets into domain-specific handlers if your API grows
// - Use consistent error handling patterns throughout all handlers
type APIHandler struct {
	log *logrus.Logger
	// Add additional dependencies as needed (e.g., database, services, clients)
}

// NewAPIHandler creates a new API handler with the given logger.
func NewAPIHandler(log *logrus.Logger) *APIHandler {
	return &APIHandler{
		log: log,
	}
}

// Health returns a handler function for the health endpoint.
// It provides information about the server's uptime and build details.
func (h *APIHandler) Health(upTime time.Time, buildDate string, buildTag string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		intake.RespondJSON(w, r, http.StatusOK, map[string]any{
			"status":    "ok",
			"upTime":    time.Since(upTime).String(),
			"buildDate": buildDate,
			"buildTag":  buildTag,
		})
	}
}

