# Testing Cheat Sheet

Quick reference for common testing tasks in Uberman.

## Daily Development (TDD)

```bash
# Fast unit tests (native)
go test -v ./internal/appdir
go test -v ./internal/config

# Watch mode (with entr)
find . -name '*.go' | entr -c go test ./internal/appdir

# Run single test
go test -run TestManager_Create ./internal/appdir

# Run with race detector
go test -race ./...
```

## Before Commit

```bash
# All unit tests
make test-short

# Property-based tests
make test-properties

# Integration tests
make test-integration

# All tests
make test
```

## Docker Testing (Isolated)

```bash
# Build image (first time)
make test-docker-build

# Property tests (recommended)
make test-docker-properties

# All tests
make test-docker

# Integration tests
make test-docker-integration

# Coverage
make test-docker-coverage
```

## Coverage

```bash
# Native
make test-coverage
open coverage.html

# Docker
make test-docker-coverage
open coverage/coverage.html

# Detailed by package
make test-coverage-detail
```

## Specific Packages

```bash
# Native
make test-appdir
make test-config
make test-database

# Docker
make test-docker-appdir
make test-docker-config
```

## Cleanup

```bash
# Clean build artifacts
make clean

# Clean Docker resources
make clean-docker

# Full cleanup
make clean clean-docker
```

## Debugging

```bash
# Verbose output
go test -v ./internal/appdir

# Print test names
go test -v -run . ./internal/appdir | grep "=== RUN"

# Shell in Docker container
docker compose -f docker-compose.test.yml run --rm test-all sh

# Run specific test in Docker
docker compose -f docker-compose.test.yml run --rm test-all \
  go test -v -run TestManager_Create ./internal/appdir
```

## Performance

```bash
# Benchmark tests
go test -bench . ./internal/appdir

# Profile CPU
go test -cpuprofile=cpu.prof ./internal/appdir
go tool pprof cpu.prof

# Profile memory
go test -memprofile=mem.prof ./internal/appdir
go tool pprof mem.prof

# Test count (run N times)
go test -count=10 ./internal/appdir
```

## CI/CD

```bash
# What CI runs (GitHub Actions)
make test-short                    # Unit tests
make test-docker-properties        # Property tests
make test-docker-integration       # Integration tests
make test-coverage                 # Coverage report
```

## Troubleshooting

```bash
# Check Docker
docker ps

# View running tests
docker compose -f docker-compose.test.yml ps

# View logs
docker compose -f docker-compose.test.yml logs

# Force cleanup
make clean-docker
docker system prune -a
```

## Test Patterns

### Unit Test
```go
func TestManager_Create(t *testing.T) {
    tmpDir := testutil.CreateTempDir(t)
    // ... test code
}
```

### Property Test
```go
func TestProperty_DirectoryCreation(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        appName := rapid.StringMatching(`^[a-z]+$`).Draw(t, "appName")
        // ... property checks
    })
}
```

### Integration Test
```go
func TestDatabase_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    // ... integration test
}
```

## Makefile Targets

| Target | Description |
|--------|-------------|
| `test` | Run all tests |
| `test-short` | Unit tests only |
| `test-properties` | Property-based tests |
| `test-integration` | Integration tests |
| `test-coverage` | Coverage report |
| `test-docker` | All tests in Docker |
| `test-docker-properties` | Property tests in Docker |
| `test-docker-coverage` | Coverage in Docker |
| `clean` | Clean build artifacts |
| `clean-docker` | Clean Docker resources |

## Quick Checks

```bash
# Test one package
go test ./internal/appdir

# Test all
go test ./...

# With coverage
go test -cover ./...

# Fail fast
go test -failfast ./...

# Parallel execution
go test -parallel 4 ./...

# Short mode (skip slow tests)
go test -short ./...
```

## Environment Variables

```bash
# Skip integration tests
go test -short ./...

# Enable verbose
VERBOSE=1 go test ./...

# Set Go flags
GOFLAGS="-count=1" go test ./...
```

## Docker Compose Services

| Service | Command |
|---------|---------|
| `test-all` | All tests |
| `test-properties` | Property tests only |
| `test-integration` | Integration tests |
| `test-coverage` | With coverage |
| `test-appdir` | appdir package |
| `test-config` | config package |
