// Package cmd provides the command-line interface for the API server.
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dbubel/go-api-template/api"
	"github.com/dbubel/go-api-template/pkg/middleware"
	"github.com/dbubel/intake/v2"
	"github.com/sirupsen/logrus"
)

var upTime = time.Now()

// Config holds the application configuration
type Config struct {
	// Environment is the current runtime environment (local, staging, production)
	Environment string
	// Port is the port number that the server will listen on
	Port int
	// BuildTag is the version or tag of the build
	BuildTag string
	// BuildDate is the timestamp when the build was created
	BuildDate string
}

// ServeCommand implements the server command for running the API.
type ServeCommand struct {
	// Config holds the application configuration
	Config Config

	// Log is the configured logger instance for application-wide logging
	Log *logrus.Logger
}

// Help returns the command-line usage help text for the serve command.
func (c *ServeCommand) Help() string {
	return `go run main.go serve`
}

// Synopsis returns a brief description of the serve command.
func (c *ServeCommand) Synopsis() string {
	return "Run the API server"
}

// Run executes the serve command, starting the API server.
func (c *ServeCommand) Run(args []string) int {
	app := intake.New()

	// Initialize middleware
	middlewares := middleware.NewMiddleware(c.Log)

	// Add global middleware
	app.AddGlobalMiddleware(middlewares.Recover)
	app.AddGlobalMiddleware(middlewares.Logging)
	app.AddGlobalMiddleware(middlewares.Timeout(5 * time.Second))
	app.AddGlobalMiddleware(middlewares.CORS)

	// Create the API handler
	apiHandler := api.NewAPIHandler(c.Log)

	// Get all endpoints
	endpoints := api.GetEndpoints(apiHandler, upTime, c.Config.BuildDate, c.Config.BuildTag)
	app.AddEndpoints(endpoints)

	// Print registered routes
	routes := app.GetRoutes()
	c.Log.Info("Registered routes:")
	for path, methods := range routes {
		c.Log.Infof("%s: %v", path, methods)
	}

	// Start the HTTP server with configured timeouts and limits
	c.Log.Infof("Starting server on port %d", c.Config.Port)
	app.Run(&http.Server{
		Addr:           fmt.Sprintf(":%d", c.Config.Port),
		Handler:        app.Mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	})

	return 0
}

