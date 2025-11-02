.PHONY: build install test clean help release

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
	@rm -rf $(BUILD_DIR) $(DIST_DIR)
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
	@echo "  build            - Build the binary"
	@echo "  install          - Install to GOPATH/bin"
	@echo "  install-local    - Install to ~/bin (for Uberspace)"
	@echo "  release          - Build release binaries for all platforms"
	@echo "  test             - Run tests"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  test-integration - Run integration tests"
	@echo "  clean            - Remove build artifacts"
	@echo "  fmt              - Format code"
	@echo "  lint             - Run linter"
	@echo "  deps             - Download and tidy dependencies"
	@echo "  help             - Show this help message"
