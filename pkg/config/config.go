// Package config provides configuration structures and utilities for the go-api-template.
// This is a template repository for building Go API services with a standardized structure.
//
// The config package handles:
// - Environment-specific settings (local, development, production)
// - Configuration loading from environment variables
// - Connection settings for external services and dependencies
// - Build information for versioning and monitoring
//
// When using this template, extend the Config struct with your application-specific
// configuration needs such as database connections, external API clients, feature flags, etc.
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

// Config is the main configuration structure that contains all application settings.
// This struct is designed to be extended with your specific application configuration needs.
//
// Template users should:
// - Add configuration sections for different components (database, cache, authentication, etc.)
// - Use struct tags for automatic loading from environment variables (via envconfig)
// - Consider using nested structs for complex configuration groups
// - Add validation methods for configuration values when appropriate
//
// Example extensions:
//
//	// Database configuration
//	Database struct {
//	    Host     string `envconfig:"DB_HOST" default:"localhost"`
//	    Port     int    `envconfig:"DB_PORT" default:"5432"`
//	    Username string `envconfig:"DB_USER" required:"true"`
//	    Password string `envconfig:"DB_PASS" required:"true"`
//	    Name     string `envconfig:"DB_NAME" required:"true"`
//	}
type Config struct {
	Environment string `default:"local" envconfig:"ENVIRONMENT"` // Current deployment environment
	Port        int    `default:"3000" envconfig:"PORT"`         // Port the server will listen on
	BuildDate   string // Date when the application was built
	BuildTag    string // Git tag or version identifier for the build

	// Add your application-specific configuration fields below
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

// Dump returns a formatted JSON string representation of the configuration.
// Useful for debugging and logging purposes.
func (c Config) Dump() string {
	s, _ := json.MarshalIndent(c, "", " ")
	return string(s)
}

// IsDevelopment returns true if the environment is set to development.
// Template usage example - add environment-specific helper methods as needed.
func (c Config) IsDevelopment() bool {
	return c.GetEnvironment() == ENV_DEVELOPMENT
}

// IsProduction returns true if the environment is set to production.
// Template usage example - add environment-specific helper methods as needed.
func (c Config) IsProduction() bool {
	return c.GetEnvironment() == ENV_PROD
}

// IsLocal returns true if the environment is set to local development.
// Template usage example - add environment-specific helper methods as needed.
func (c Config) IsLocal() bool {
	return c.GetEnvironment() == ENV_LOCAL
}

