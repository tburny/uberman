.PHONY: build install test test-short test-properties test-coverage test-coverage-detail test-integration test-config test-appdir test-database test-runtime test-web test-supervisor test-docker test-docker-properties test-docker-integration test-docker-coverage test-docker-build clean help release fmt lint deps install-local

# Variables
BINARY_NAME=uberman
BUILD_DIR=bin
DIST_DIR=dist
GO=go
GOFLAGS=-v
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

# Build the binary
build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/uberman

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

# Run all tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Run tests in short mode (skip integration tests)
test-short:
	@echo "Running unit tests (short mode)..."
	$(GO) test -short -v ./...

# Run only property-based tests
test-properties:
	@echo "Running property-based tests..."
	$(GO) test -v -run Property ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -cover ./...
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with detailed coverage by package
test-coverage-detail:
	@echo "Running tests with detailed coverage..."
	@echo "Package Coverage:"
	@$(GO) test -cover ./internal/config
	@$(GO) test -cover ./internal/appdir
	@$(GO) test -cover ./internal/database
	@$(GO) test -cover ./internal/runtime
	@$(GO) test -cover ./internal/web
	@$(GO) test -cover ./internal/supervisor

# Run integration tests (requires Docker for testcontainers)
test-integration:
	@echo "Running integration tests..."
	@echo "Note: Requires Docker to be running for testcontainers"
	$(GO) test -v -run Integration ./...

# Run specific package tests
test-config:
	@echo "Testing config package..."
	$(GO) test -v ./internal/config

test-appdir:
	@echo "Testing appdir package..."
	$(GO) test -v ./internal/appdir

test-database:
	@echo "Testing database package..."
	$(GO) test -v ./internal/database

test-runtime:
	@echo "Testing runtime package..."
	$(GO) test -v ./internal/runtime

test-web:
	@echo "Testing web package..."
	$(GO) test -v ./internal/web

test-supervisor:
	@echo "Testing supervisor package..."
	$(GO) test -v ./internal/supervisor

# Docker-based testing targets (filesystem isolation)

# Build Docker test image
test-docker-build:
	@echo "Building Docker test image..."
	docker build -f Dockerfile.test -t uberman-test:latest .

# Run all tests in Docker container (isolated filesystem)
test-docker:
	@echo "Running tests in Docker container (isolated filesystem)..."
	@echo "Note: Requires Docker to be running"
	docker compose -f docker-compose.test.yml run --rm test-all

# Run property-based tests in Docker (high-volume filesystem isolation)
test-docker-properties:
	@echo "Running property-based tests in Docker container..."
	@echo "This provides isolation for high-volume filesystem operations"
	docker compose -f docker-compose.test.yml run --rm test-properties

# Run integration tests in Docker
test-docker-integration:
	@echo "Running integration tests in Docker container..."
	docker compose -f docker-compose.test.yml run --rm test-integration

# Run tests with coverage in Docker
test-docker-coverage:
	@echo "Running tests with coverage in Docker container..."
	@mkdir -p coverage
	docker compose -f docker-compose.test.yml run --rm test-coverage
	@echo "Coverage report generated in coverage/ directory"

# Run specific package tests in Docker
test-docker-appdir:
	@echo "Testing appdir package in Docker..."
	docker compose -f docker-compose.test.yml run --rm test-appdir

test-docker-config:
	@echo "Testing config package in Docker..."
	docker compose -f docker-compose.test.yml run --rm test-config

# Clean up Docker test resources
clean-docker:
	@echo "Cleaning Docker test resources..."
	docker compose -f docker-compose.test.yml down --volumes --remove-orphans
	docker rmi -f uberman-test:latest 2>/dev/null || true

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR) $(DIST_DIR)
	@rm -f coverage.out coverage.html
	@rm -rf coverage/
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

# Build release binaries for all platforms
release:
	@echo "Building release binaries..."
	@mkdir -p $(DIST_DIR)
	@echo "Building for Linux amd64..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/uberman
	@echo "Building for Linux arm64..."
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/uberman
	@echo "Building for macOS amd64..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/uberman
	@echo "Building for macOS arm64..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/uberman
	@echo "Building for Windows amd64..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/uberman
	@echo "Building for FreeBSD amd64..."
	GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0 $(GO) build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-freebsd-amd64 ./cmd/uberman
	@echo "Release binaries built in $(DIST_DIR)/"
	@ls -lh $(DIST_DIR)/

# Show help
help:
	@echo "Available targets:"
	@echo "  build                 - Build the binary"
	@echo "  install               - Install to GOPATH/bin"
	@echo "  install-local         - Install to ~/bin (for Uberspace)"
	@echo "  release               - Build release binaries for all platforms"
	@echo ""
	@echo "Testing (Native):"
	@echo "  test                  - Run all tests"
	@echo "  test-short            - Run unit tests only (skip integration)"
	@echo "  test-properties       - Run property-based tests only"
	@echo "  test-coverage         - Run tests with coverage report"
	@echo "  test-coverage-detail  - Show coverage by package"
	@echo "  test-integration      - Run integration tests (requires Docker)"
	@echo "  test-config           - Test config package only"
	@echo "  test-appdir           - Test appdir package only"
	@echo "  test-database         - Test database package only"
	@echo "  test-runtime          - Test runtime package only"
	@echo "  test-web              - Test web package only"
	@echo "  test-supervisor       - Test supervisor package only"
	@echo ""
	@echo "Testing (Docker - Isolated Filesystem):"
	@echo "  test-docker-build         - Build Docker test image"
	@echo "  test-docker               - Run all tests in Docker"
	@echo "  test-docker-properties    - Run property-based tests in Docker (recommended)"
	@echo "  test-docker-integration   - Run integration tests in Docker"
	@echo "  test-docker-coverage      - Run tests with coverage in Docker"
	@echo "  test-docker-appdir        - Test appdir package in Docker"
	@echo "  test-docker-config        - Test config package in Docker"
	@echo "  clean-docker              - Clean Docker test resources"
	@echo ""
	@echo "Development:"
	@echo "  clean                 - Remove build artifacts"
	@echo "  fmt                   - Format code"
	@echo "  lint                  - Run linter"
	@echo "  deps                  - Download and tidy dependencies"
	@echo "  help                  - Show this help message"
