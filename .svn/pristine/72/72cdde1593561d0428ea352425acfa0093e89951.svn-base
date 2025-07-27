# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
```bash
# Run with default config
go run cmd/main.go

# Run with custom config file
go run cmd/main.go --config configs/config.yaml
```

### Building
```bash
# Build binary
go build -o bin/example-api cmd/main.go

# Install dependencies
go mod tidy
```

### Testing
```bash
# Run tests (if any exist)
go test ./...
```

## Architecture Overview

This is a Go REST API project following clean architecture principles with a layered approach:

### Project Structure
- `cmd/main.go` - Application entry point with Cobra CLI framework
- `pkg/config/` - Configuration management using Viper with YAML support
- `pkg/util/` - Shared utilities (structured logging with slog)
- `pkg/api/handler/` - HTTP request handlers (Echo v4 framework)
- `pkg/api/service/` - Business logic layer with in-memory storage
- `pkg/api/vo/` - Value objects (DTOs) for request/response structures

### Key Technologies
- **Echo v4** - Web framework for REST API
- **Cobra** - CLI framework for command-line interface
- **Viper** - Configuration management (YAML, environment variables)
- **slog** - Native Go structured logging

### Data Flow Pattern
Request → Handler → Service → Response
- Handlers validate input and manage HTTP concerns
- Services contain business logic and data operations
- Value objects define request/response structures
- In-memory storage with mutex-based concurrency control

### Configuration
- Default config: `configs/config.yaml`
- Supports environment variable overrides via Viper
- Configuration structure defined in `pkg/config/config.go`

### Logging
- Structured logging using Go's native `slog` package
- Component-based logger instances with context
- Log levels: debug, info, warn, error

### API Endpoints
Base path: `/api/v1/users`
- GET `/` - List all users
- GET `/:id` - Get user by ID
- POST `/` - Create new user
- PUT `/:id` - Update user
- DELETE `/:id` - Delete user

### Error Handling
- Consistent error response format via `vo.ErrorResponse`
- Proper HTTP status codes
- Structured error logging with context