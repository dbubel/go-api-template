// Package config provides configuration structures and utilities for the API.
// It handles environment-specific settings and connections to external services.
package config

import (
	"encoding/json"
)

// ENV represents the deployment environment type
type ENV string

// Environment constants representing different deployment environments
const (
	ENV_LOCAL       ENV = "local"       // Local development environment
	ENV_DEVELOPMENT ENV = "development" // Development/staging environment
	ENV_PROD        ENV = "production"  // Production environment
)

// Config is the main configuration structure that contains all application settings
type Config struct {
	Environment string `default:"local" envconfig:"ENVIRONMENT"` // Current deployment environment
	Port        int    `default:"3000" envconfig:"PORT"`         // Port the server will listen on
	BuildDate   string // Date when the application was built
	BuildTag    string // Git tag or version identifier for the build
}

// GetEnvironment returns the environment as an ENV type
func (c Config) GetEnvironment() ENV {
	switch c.Environment {
	case "local":
		return ENV_LOCAL
	case "development":
		return ENV_DEVELOPMENT
	case "production":
		return ENV_PROD
	default:
		return ENV_LOCAL
	}
}

// Dump returns a formatted JSON string representation of the configuration
// Useful for debugging and logging purposes
func (c Config) Dump() string {
	s, _ := json.MarshalIndent(c, "", " ")
	return string(s)
} 