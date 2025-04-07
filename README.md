# Go API Template

A template repository for building Go API services with a standardized structure. This template provides a solid foundation for developing RESTful APIs with Go, focusing on modularity, maintainability, and best practices.

## Features

- HTTP server with configurable port
- Centralized route management through a dedicated routes file
- Middleware for logging, CORS, timeouts, and panic recovery
- CLI-based command structure (using mitchellh/cli)
- Environment-based configuration
- Structured logging with logrus
- Health endpoint with build information

## Requirements

- Go 1.20 or higher

## Getting Started

### Installation

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

### Running the server

```bash
go run main.go serve
```

By default, the server runs on port 3000. You can configure the port using the `PORT` environment variable:

```bash
PORT=8080 go run main.go serve
```

## Project Structure

```
├── api/            # API handlers and route definitions
│   ├── handler.go  # HTTP handler implementations
│   └── routes.go   # Centralized route definitions
├── cmd/            # Command-line commands
│   └── main.go     # Server command implementation
├── pkg/            # Shared packages
│   ├── config/     # Application configuration
│   └── middleware/ # HTTP middleware
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
├── main.go         # Application entry point
└── README.md       # Project documentation
```

## API Endpoints

### Health Check

```
GET /health
```

Returns information about the server's health, uptime, and build details.

Example response:

```json
{
  "status": "ok",
  "upTime": "5m32s",
  "buildDate": "2023-04-10T12:00:00Z",
  "buildTag": "dev"
}
```

## Configuration

The application can be configured using environment variables:

- `PORT`: The port on which the server listens (default: 3000)
- `ENVIRONMENT`: The environment in which the application is running (default: "local")

## Extending the Template

This template is designed to be extended with your specific API functionality:

### Adding New Routes

1. Define new handler methods in `api/handler.go` or create domain-specific handler files
2. Register routes in `api/routes.go` by adding them to the appropriate endpoint group
3. Group related endpoints for better organization

Example:

```go
// In api/routes.go
userEndpoints := intake.Endpoints{
    intake.POST("/api/v1/users", apiHandler.CreateUser),
    intake.GET("/api/v1/users", apiHandler.GetUsers),
    intake.GET("/api/v1/users/:id", apiHandler.GetUser),
}

// Combine with other endpoint groups
endpoints := append(healthEndpoints, userEndpoints...)
```

### Adding Configuration

Extend the `Config` struct in `pkg/config/config.go` with your application-specific settings:

```go
type Config struct {
    // Existing fields...
    
    // Database configuration
    Database struct {
        Host     string `envconfig:"DB_HOST" default:"localhost"`
        Port     int    `envconfig:"DB_PORT" default:"5432"`
        Username string `envconfig:"DB_USER" required:"true"`
        Password string `envconfig:"DB_PASS" required:"true"`
        Name     string `envconfig:"DB_NAME" required:"true"`
    }
}
```

### Adding Middleware

Add custom middleware in `pkg/middleware/middleware.go` and register it in `cmd/main.go`:

```go
// Global middleware (applies to all routes)
app.AddGlobalMiddleware(middlewares.YourCustomMiddleware)

// Or specific to route groups
endpoints.Prepend(middlewares.AnotherMiddleware)
```

## License

[Add your license here]