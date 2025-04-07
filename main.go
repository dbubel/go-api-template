package main

import (
	"os"
	"time"

	"github.com/dbubel/go-api-template/cmd"
	"github.com/dbubel/go-api-template/pkg/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
)

// Build information variables that are populated during compilation
var (
	// BUILD_TAG contains the version or tag of the build
	BUILD_TAG = "dev"
	// BUILD_DATE contains the timestamp when the build was created
	BUILD_DATE = time.Now().Format(time.RFC3339)
)

// main is the entry point of the application.
// It initializes the logger, loads configuration from environment variables,
// sets up the CLI commands, and runs the selected command.
func main() {
	// Initialize the logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	// Load configuration from environment variables
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.WithError(err).Error("Error parsing config")
		return
	}

	// Add build information to the configuration
	cfg.BuildTag = BUILD_TAG
	cfg.BuildDate = BUILD_DATE

	// Initialize the CLI application
	c := cli.NewCLI("go-api-template", BUILD_TAG)
	c.Args = os.Args[1:] // Use command line arguments excluding the program name

	// Register available commands
	c.Commands = map[string]cli.CommandFactory{
		"serve": func() (cli.Command, error) {
			// Convert to cmd.Config
			cmdConfig := cmd.Config{
				Environment: cfg.Environment,
				Port:        cfg.Port,
				BuildTag:    cfg.BuildTag,
				BuildDate:   cfg.BuildDate,
			}

			return &cmd.ServeCommand{
				Config: cmdConfig,
				Log:    logger,
			}, nil
		},
	}

	// Run the CLI application with the provided arguments
	exitStatus, err := c.Run()
	if err != nil {
		logger.WithError(err).Error("Error running command")
	}

	os.Exit(exitStatus)
}
