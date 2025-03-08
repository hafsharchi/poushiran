# Poushiran API

A Go-based REST API using Gin framework with Swagger documentation.

## Project Structure

```
.
├── api/
│   ├── handlers/       # HTTP request handlers
│   └── routes/         # Route definitions
├── config/             # Application configuration
├── docs/               # Swagger documentation
├── internal/
│   ├── models/         # Data models
│   └── services/       # Business logic
├── pkg/
│   └── utils/          # Utility functions
├── go.mod              # Go module definition
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- Swag CLI for Swagger generation

## Setup

1. Install dependencies:
   ```
   go mod download
   ```

2. Install Swag CLI (for Swagger docs generation):
   ```
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

3. Generate Swagger documentation:
   ```
   swag init -g main.go -o ./docs
   ```

## Running the Application

```
go run main.go
```

The server will start on http://localhost:8080

## API Documentation

Swagger UI is available at: http://localhost:8080/swagger/index.html

## Available Endpoints

- GET /api/v1/products - Get all products 