// Package api provides the API handlers for the application.
package api

import (
	"net/http"
	"time"

	"github.com/dbubel/intake/v2"
	"github.com/sirupsen/logrus"
)

// APIHandler handles API requests and implements the application's endpoints.
type APIHandler struct {
	log *logrus.Logger
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