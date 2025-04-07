// Package api provides the API handlers and route definitions for the go-api-template.
// This is a template repository for building Go API services with a standardized structure.
// Define your API routes in this file and group them logically for better organization.
package api

import (
	"time"

	"github.com/dbubel/intake/v2"
)

var upTime = time.Now()

// GetEndpoints returns all the API endpoints organized by functional area.
// This function aggregates all route groups and should be used in cmd/main.go to register all endpoints.
//
// Template users should:
// 1. Define logical endpoint groups (e.g., healthEndpoints, userEndpoints, authEndpoints)
// 2. Keep related functionality grouped together for better organization and maintenance
// 3. Use consistent API versioning patterns (e.g., /api/v1/resource)
func GetEndpoints(apiHandler *APIHandler, upTime time.Time, buildDate string, buildTag string) intake.Endpoints {
	// Health endpoints
	healthEndpoints := intake.Endpoints{
		intake.GET("/health", apiHandler.Health(upTime, buildDate, buildTag)),
	}

	// Add more endpoint groups here as needed
	// For example:
	// userEndpoints := intake.Endpoints{
	//     intake.POST("/api/v1/users", apiHandler.CreateUser),
	//     intake.GET("/api/v1/users", apiHandler.GetUsers),
	// }

	// Combine all endpoints
	endpoints := healthEndpoints
	// endpoints := append(healthEndpoints, userEndpoints...)

	return endpoints
}
