# SAPI - Simple API Server

A simple API server with a health endpoint.

## Features

- HTTP server with configurable port
- Health endpoint
- Middleware for logging, CORS, timeouts, and panic recovery
- Custom CLI-based command structure (no external dependencies)

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

## Development

### Project Structure

```
├── api/            # API handlers
├── cmd/            # Command-line commands
├── pkg/            # Shared packages
│   ├── cli/        # Custom CLI implementation with no dependencies
│   └── middleware/ # HTTP middleware
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
├── main.go         # Application entry point
└── README.md       # This file
```