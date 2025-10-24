# Mumuni Backend Makefile

.PHONY: help build run test test-go test-curl clean deps

# Default target
help:
	@echo "Available commands:"
	@echo "  deps      - Install dependencies"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application"
	@echo "  test      - Run API tests (Go script)"
	@echo "  test-curl - Run API tests (curl commands)"
	@echo "  clean     - Clean build artifacts"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# Build the application
build:
	@echo "Building application..."
	go build -o bin/mumuni_backend main.go

# Run the application
run:
	@echo "Starting server..."
	go run main.go

# Run API tests (Go script)
test:
	@echo "Running API tests..."
	go run test_api.go

# Run API tests (curl commands)
test-curl:
	@echo "Running API tests with curl..."
	@if [ -f test_api.sh ]; then bash test_api.sh; else echo "test_api.sh not found"; fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/

# Run with hot reload (requires air)
dev:
	@echo "Starting development server with hot reload..."
	air

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
