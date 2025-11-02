.PHONY: build install test clean help

# Variables
BINARY_NAME=uberman
BUILD_DIR=bin
GO=go
GOFLAGS=-v

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/uberman/*.go

# Install to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GO) install ./cmd/uberman

# Install to ~/bin (for Uberspace)
install-local: build
	@echo "Installing to ~/bin..."
	@mkdir -p ~/bin
	@cp $(BUILD_DIR)/$(BINARY_NAME) ~/bin/
	@echo "$(BINARY_NAME) installed to ~/bin/$(BINARY_NAME)"

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -cover ./...
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run integration tests (requires Uberspace environment)
test-integration:
	@echo "Running integration tests..."
	$(GO) test -tags=integration -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  build            - Build the binary"
	@echo "  install          - Install to GOPATH/bin"
	@echo "  install-local    - Install to ~/bin (for Uberspace)"
	@echo "  test             - Run tests"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  test-integration - Run integration tests"
	@echo "  clean            - Remove build artifacts"
	@echo "  fmt              - Format code"
	@echo "  lint             - Run linter"
	@echo "  deps             - Download and tidy dependencies"
	@echo "  help             - Show this help message"
