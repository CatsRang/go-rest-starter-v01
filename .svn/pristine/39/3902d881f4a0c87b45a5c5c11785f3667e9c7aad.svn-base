# Go REST API Example

A simplified example project demonstrating core Go REST API functionality, extracted and refined from the main API Gateway project.

## Features

- **Configuration Management**: Viper-based configuration with YAML support
- **Logging**: Structured logging using Go's native `slog` package
- **CLI Framework**: Cobra for command-line interface
- **REST API**: Echo v4 web framework
- **Clean Architecture**: Handler → Service → VO pattern

## Project Structure

```
_example/
├── cmd/
│   └── main.go                 # Application entry point with Cobra CLI
├── pkg/
│   ├── config/
│   │   └── config.go          # Configuration management with Viper
│   ├── util/
│   │   └── logger.go          # Logging utility with slog
│   └── api/
│       ├── handler/
│       │   └── user_handler.go # HTTP request handlers
│       ├── service/
│       │   └── user_service.go # Business logic layer
│       └── vo/
│           └── user.go        # Value objects (DTOs)
├── configs/
│   └── config.yaml            # Configuration file
├── go.mod                     # Go module definition
└── README.md                  # This file
```

## Quick Start

1. **Install Dependencies**
   ```bash
   cd _example
   go mod tidy
   ```

2. **Run the Application**
   ```bash
   go run cmd/main.go
   # or with custom config
   go run cmd/main.go --config configs/config.yaml
   ```

3. **Build the Application**
   ```bash
   go build -o bin/example-api cmd/main.go
   ./bin/example-api
   ```

## API Endpoints

### Users API (`/api/v1/users`)

- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Example Usage

**Get All Users:**
```bash
curl http://localhost:8080/api/v1/users
```

**Create User:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com"}'
```

**Get User by ID:**
```bash
curl http://localhost:8080/api/v1/users/1
```

**Update User:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.updated@example.com"}'
```

**Delete User:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Configuration

Edit `configs/config.yaml`:

```yaml
server:
  host: "localhost"
  port: 8080

log:
  level: "info"  # debug, info, warn, error
```

## Architecture Overview

### Handler Layer (`pkg/api/handler/`)
- Handles HTTP requests and responses
- Input validation and error handling
- Route registration

### Service Layer (`pkg/api/service/`)
- Business logic implementation
- Data processing and validation
- In-memory storage (for example purposes)

### Value Objects (`pkg/api/vo/`)
- Data transfer objects (DTOs)
- Request/response structures
- Type definitions

### Utilities (`pkg/util/`)
- Shared utilities like logging
- Cross-cutting concerns

### Configuration (`pkg/config/`)
- Application configuration management
- Viper integration for multiple config formats

## Key Technologies

- **Echo v4**: Fast and minimalist web framework
- **Cobra**: CLI framework for command-line applications
- **Viper**: Configuration management with support for multiple formats
- **slog**: Native Go structured logging (Go 1.21+)

## Development Notes

- The example uses in-memory storage for simplicity
- Includes pre-seeded data for testing
- Implements basic CRUD operations
- Follows clean architecture principles
- Demonstrates proper error handling and logging