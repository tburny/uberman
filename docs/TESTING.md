# Testing Guide

This document describes the testing strategy and tools for Uberman, including containerized testing for filesystem isolation.

## Test Categories

### Unit Tests
Fast, focused tests for individual components. Run without any special requirements.

```bash
# Run all tests
make test

# Run in short mode (skip integration tests)
make test-short

# Test specific package
make test-appdir
make test-config
```

### Property-Based Tests
High-volume tests using `pgregory.net/rapid` that generate many test cases. These tests create 500-600 directory structures per run.

**Native execution:**
```bash
make test-properties
```

**Docker execution (recommended for CI/CD):**
```bash
make test-docker-properties
```

### Integration Tests
Tests requiring external services (MySQL via testcontainers). Always require Docker.

```bash
make test-integration
make test-docker-integration
```

## Containerized Testing

### Why Containerize Tests?

Property-based tests create hundreds of temporary directories. While Go's `t.Cleanup()` handles cleanup, containerization provides:

1. **Complete Isolation**: Tests run in ephemeral containers with tmpfs filesystems
2. **Zero Host Pollution**: All filesystem writes stay in the container
3. **Reproducibility**: Identical environment across developer machines and CI/CD
4. **Resource Management**: Automatic cleanup of all test artifacts

### Docker Test Architecture

```
Dockerfile.test           # Minimal Alpine-based Go test image
docker-compose.test.yml   # Service definitions for different test types
Makefile                  # Convenient targets for both native and Docker tests
```

### Running Containerized Tests

#### Prerequisites
- Docker installed and running
- Docker Compose v2+

#### Quick Start

```bash
# Build the test image (first time only)
make test-docker-build

# Run all tests in Docker
make test-docker

# Run property-based tests (high-volume filesystem isolation)
make test-docker-properties

# Run with coverage
make test-docker-coverage
```

#### Available Docker Test Targets

| Target | Description |
|--------|-------------|
| `test-docker` | Run all tests in isolated container |
| `test-docker-properties` | Run property-based tests only (recommended for heavy filesystem operations) |
| `test-docker-integration` | Run integration tests (with testcontainers support) |
| `test-docker-coverage` | Generate coverage report in `coverage/` directory |
| `test-docker-appdir` | Test appdir package only |
| `test-docker-config` | Test config package only |
| `clean-docker` | Remove Docker test images and containers |

### Manual Docker Execution

If you prefer using Docker Compose directly:

```bash
# Run all tests
docker compose -f docker-compose.test.yml run --rm test-all

# Run property-based tests
docker compose -f docker-compose.test.yml run --rm test-properties

# Run with coverage
docker compose -f docker-compose.test.yml run --rm test-coverage

# Cleanup
docker compose -f docker-compose.test.yml down --volumes
```

## Test-Driven Development Workflow

### Fast TDD Cycle (Native)

For rapid red-green-refactor cycles:

```bash
# Watch and run tests on save (using entr or similar)
find . -name '*.go' | entr -c go test ./internal/appdir

# Or use native test command
go test -v ./internal/appdir
```

### Property Testing Cycle (Docker)

For property-based tests creating many directories:

```bash
# Test locally first
go test -run Property ./internal/appdir

# Verify in isolated environment
make test-docker-properties
```

### Integration Testing Cycle

Always requires Docker for testcontainers:

```bash
make test-integration
```

## Writing Tests

### Standard Unit Test

```go
func TestManager_Create(t *testing.T) {
    tmpDir := testutil.CreateTempDir(t)

    originalHome := os.Getenv("HOME")
    os.Setenv("HOME", tmpDir)
    t.Cleanup(func() {
        os.Setenv("HOME", originalHome)
    })

    manager, err := NewManager("testapp", false, false)
    require.NoError(t, err)

    err = manager.Create()
    require.NoError(t, err)

    assert.True(t, manager.Exists())
}
```

### Property-Based Test

```go
func TestProperty_DirectoryCreation_AllSubdirectoriesExist(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

        _, cleanup := setupTempHome(t)
        defer cleanup()

        manager, err := NewManager(appName, false, false)
        if err != nil {
            t.Fatalf("failed to create manager: %v", err)
        }

        err = manager.Create()
        if err != nil {
            t.Fatalf("failed to create directory structure: %v", err)
        }

        // Property assertions...
    })
}
```

### Integration Test with Testcontainers

```go
func TestMySQLManager_Integration_CreateDatabase(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }

    if _, err := testcontainers.NewDockerClient(); err != nil {
        t.Skip("Docker not available, skipping integration test")
    }

    _, db, _, _ := setupMySQLContainer(t)

    // Test database operations...
}
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test-native:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.24'
      - run: make test-short

  test-docker:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - run: make test-docker-properties
      - run: make test-docker-integration
```

### GitLab CI Example

```yaml
test:native:
  stage: test
  image: golang:1.24
  script:
    - make test-short

test:docker:
  stage: test
  image: docker:latest
  services:
    - docker:dind
  script:
    - make test-docker-properties
    - make test-docker-integration
```

## Test Performance

### Benchmarks

| Test Type | Native | Docker | Notes |
|-----------|--------|--------|-------|
| Unit tests | ~50ms | ~2s | Docker has container startup overhead |
| Property tests | ~200ms | ~3s | Docker adds ~2.8s overhead but provides isolation |
| Integration tests | ~5s | ~6s | Similar performance (both use testcontainers) |

### When to Use Docker Tests

**Use Docker for:**
- CI/CD pipelines
- Property-based tests (high-volume filesystem operations)
- Release verification
- Testing on different architectures

**Use Native for:**
- TDD red-green-refactor cycles
- Quick feedback during development
- Debugging (easier to attach debugger)

## Troubleshooting

### Docker Not Available

```bash
# Check Docker status
docker ps

# Start Docker daemon (Linux)
sudo systemctl start docker

# macOS/Windows: Start Docker Desktop
```

### Tests Failing in Docker but Not Locally

This usually indicates environment differences:

1. **File paths**: Use `filepath.Join()` for cross-platform compatibility
2. **Environment variables**: Docker uses isolated environment
3. **Permissions**: Container runs as root by default

### Slow Docker Test Execution

```bash
# Rebuild image without cache
docker compose -f docker-compose.test.yml build --no-cache

# Prune Docker system
docker system prune -a

# Use tmpfs for better performance (already configured in docker-compose.test.yml)
```

### Container Cleanup Issues

```bash
# Force cleanup
make clean-docker

# Manual cleanup
docker compose -f docker-compose.test.yml down --volumes --remove-orphans
docker ps -a | grep uberman | awk '{print $1}' | xargs docker rm -f
```

## Best Practices

### Test Isolation
- Always use `t.Cleanup()` for resource cleanup
- Create fresh temporary directories for each test
- Reset environment variables in cleanup handlers
- Use unique names for test artifacts

### Property Testing
- Keep property tests focused on invariants
- Use meaningful property names (REQ-PBT-COMPONENT-XXX)
- Document the property being tested in comments
- Generate realistic test data with appropriate constraints

### Integration Testing
- Skip integration tests in short mode: `if testing.Short() { t.Skip() }`
- Check Docker availability before running
- Use testcontainers for service dependencies
- Always clean up containers in `t.Cleanup()`

### Performance
- Run unit tests natively for fast feedback
- Use Docker tests for CI/CD and verification
- Profile slow tests: `go test -cpuprofile=cpu.prof`
- Consider parallel test execution: `go test -parallel 4`

## References

- [Go Testing Package](https://pkg.go.dev/testing)
- [Testify Assertions](https://github.com/stretchr/testify)
- [Rapid Property Testing](https://github.com/flyingmutant/rapid)
- [Testcontainers Go](https://golang.testcontainers.org/)
- [Docker Compose](https://docs.docker.com/compose/)
