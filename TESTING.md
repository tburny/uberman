# Testing Guide for Uberman

This document provides comprehensive information about testing the Uberman CLI tool.

## Test Coverage Overview

Current test coverage: **46.2% overall**

### Coverage by Package

| Package | Coverage | Test Count | Notes |
|---------|----------|------------|-------|
| internal/config | 97.2% | 8 test functions | Excellent coverage with real manifest tests |
| internal/appdir | 82.4% | 8 test functions | High coverage with integration tests |
| internal/database | 51.5% | 6 test functions | Includes testcontainers integration tests |
| internal/web | 44.0% | 5 test functions | Dry-run mode tests |
| internal/supervisor | 35.9% | 4 test functions | Template and service tests |
| internal/runtime | 25.6% | 4 test functions | Dry-run mode tests |

## Test Dependencies

The test suite uses the following testing frameworks and tools:

- **testify/assert** - Assertion library for cleaner test assertions
- **testify/require** - Assertions that stop test execution on failure
- **testcontainers-go** - Docker-based integration testing
- **testcontainers-go/modules/mysql** - MySQL container for database tests
- **pgregory.net/rapid** - Property-based testing framework

## Running Tests

### All Tests

```bash
# Run all tests
make test

# Or directly with go
go test ./...
```

### Unit Tests Only (Skip Integration)

```bash
# Run only fast unit tests, skip slow integration tests
make test-short

# Or directly
go test -short ./...
```

### With Coverage

```bash
# Generate coverage report
make test-coverage

# View detailed coverage by package
make test-coverage-detail

# Or directly
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Individual Package Tests

```bash
# Test specific packages
make test-config
make test-appdir
make test-database
make test-runtime
make test-web
make test-supervisor

# Or directly
go test -v ./internal/config
go test -v ./internal/appdir
# ... etc
```

### Property-Based Tests

```bash
# Run only property-based tests
make test-properties

# Or directly
go test -v -run Property ./...

# Run property tests for specific package
go test -v ./internal/config -run Property
go test -v ./internal/database -run Property

# Run with more iterations (default is 100)
go test -v ./internal/config -run Property -rapid.checks=1000

# Reproduce a specific failure using seed
go test -v ./internal/config -run Property -rapid.seed=12345
```

### Integration Tests

```bash
# Run integration tests (requires Docker)
make test-integration

# Or directly
go test -v -run Integration ./...
```

## Test Structure

### Unit Tests

Unit tests follow Go testing conventions with `*_test.go` files alongside the code they test:

```
internal/
├── config/
│   ├── app.go
│   ├── app_test.go                # Unit tests for config package
│   └── app_properties_test.go     # Property-based tests
├── appdir/
│   ├── manager.go
│   ├── manager_test.go            # Unit tests for appdir package
│   └── manager_properties_test.go # Property-based tests
├── database/
│   ├── mysql.go
│   ├── mysql_test.go              # Unit + integration tests
│   └── mysql_properties_test.go   # Property-based tests
└── testutil/
    └── fixtures.go                 # Shared test utilities and fixtures
```

### Test Fixtures and Utilities

The `internal/testutil` package provides shared test utilities:

- **CreateTempDir** - Creates temporary test directories with automatic cleanup
- **CreateTestFile** - Creates test files with content
- **MinimalManifest** - Returns a minimal valid app manifest in TOML format
- **NodeJSManifest** - Returns a Node.js app manifest for testing
- **InvalidManifest** - Returns an invalid manifest for error testing
- **MyCnfContent** - Generates MySQL credentials file content

### Integration Tests

Integration tests use testcontainers to run real infrastructure:

#### Database Integration Tests

The database package includes integration tests that:

1. Start a real MySQL container using testcontainers
2. Test actual database operations (create, export, import)
3. Verify database existence and queries
4. Clean up containers after tests

To run database integration tests:

```bash
# Requires Docker to be running
go test -v ./internal/database
```

Tests are skipped automatically if:
- Running in short mode (`go test -short`)
- Docker is not available

## Test Patterns and Best Practices

### Property-Based Testing

Property-based tests verify that certain properties (invariants) hold for all possible inputs. These complement example-based tests by exploring the input space automatically.

**Example**: Database name convention property

```go
// Property: Database names should always follow username_appname format
func TestProperty_DatabaseName_FollowsConvention(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
        appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

        dbName := GenerateDatabaseName(username, appName)

        // Property: Should always follow the convention
        if !strings.HasPrefix(dbName, username + "_") {
            t.Fatalf("database name doesn't follow convention: got %q", dbName)
        }
    })
}
```

**Common Property Patterns**:

1. **Idempotency**: Running operation multiple times has same effect
2. **Round-trip**: Encode → Decode returns original
3. **Invariants**: Certain conditions always hold
4. **Consistency**: Same inputs always produce same outputs

See [PROPERTY_BASED_TESTING.md](PROPERTY_BASED_TESTING.md) for comprehensive guide.

### Table-Driven Tests

Most tests use table-driven patterns for comprehensive coverage:

```go
tests := []struct {
    name        string
    input       string
    expectError bool
    validate    func(t *testing.T, result *Result)
}{
    {
        name:        "valid input",
        input:       "test",
        expectError: false,
        validate:    func(t *testing.T, result *Result) { ... },
    },
    // ... more test cases
}

for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // Test implementation
    })
}
```

### Dry-Run Testing for External Commands

Packages that execute external commands (runtime, web, supervisor) use dry-run mode for testing:

```go
// Test without executing actual uberspace commands
manager := NewManager(true, false) // dryRun=true, verbose=false
err := manager.SetVersion("php", "8.3")
assert.NoError(t, err) // Dry-run should not error
```

This allows testing the logic without requiring Uberspace-specific commands.

### Environment Isolation

Tests that depend on environment variables or filesystem state use cleanup handlers:

```go
tmpDir := testutil.CreateTempDir(t)

originalHome := os.Getenv("HOME")
os.Setenv("HOME", tmpDir)
t.Cleanup(func() {
    os.Setenv("HOME", originalHome)
})
```

### Testcontainers Integration

Database tests use testcontainers for real infrastructure:

```go
container, err := mysql.Run(ctx,
    "mysql:8.0",
    mysql.WithDatabase("testdb"),
    mysql.WithUsername("testuser"),
    mysql.WithPassword("testpass"),
)
if err != nil {
    t.Fatalf("Failed to start MySQL container: %v", err)
}

t.Cleanup(func() {
    container.Terminate(ctx)
})
```

## Test Coverage Goals

Current coverage goals by package:

- **config**: 90%+ (currently 97.2% ✓)
- **appdir**: 80%+ (currently 82.4% ✓)
- **database**: 60%+ (currently 51.5% - needs improvement)
- **web**: 50%+ (currently 44.0% - needs improvement)
- **supervisor**: 40%+ (currently 35.9% - needs improvement)
- **runtime**: 30%+ (currently 25.6% - needs improvement)

## Known Test Limitations

### External Command Dependencies

Some functionality cannot be fully tested without Uberspace-specific commands:

- **runtime**: `uberspace tools version` commands (tested in dry-run mode only)
- **web**: `uberspace web backend` commands (tested in dry-run mode only)
- **supervisor**: `supervisorctl` commands (tested in dry-run mode only)

These are tested with dry-run mode and would require running on an actual Uberspace host for full integration testing.

### Identified Bugs in Tests

Tests document current behavior, including bugs:

1. **supervisor/service.go:66-68** - AutoRestart logic bug
   - Currently: `if service.AutoRestart { service.AutoRestart = true }`
   - Should: Default AutoRestart to `true` when not set
   - Test: `TestServiceManager_AutoRestartDefault` documents this behavior

2. **supervisor/service.go:43** - Template trailing comma
   - Environment variables end with trailing comma
   - Test: `TestServiceManager_ServiceTemplateOutput` documents this

## Adding New Tests

When adding new functionality:

1. **Write tests first** (TDD approach)
2. **Use table-driven tests** for multiple scenarios
3. **Test error cases** as well as happy paths
4. **Use testutil helpers** for common operations
5. **Add integration tests** where appropriate
6. **Run coverage** to ensure adequate coverage
7. **Document known limitations** in test comments

Example new test structure:

```go
func TestNewFeature(t *testing.T) {
    tests := []struct {
        name        string
        input       Input
        expectError bool
        errorMsg    string
    }{
        {
            name:        "happy path",
            input:       Input{Valid: true},
            expectError: false,
        },
        {
            name:        "error case",
            input:       Input{Valid: false},
            expectError: true,
            errorMsg:    "expected error",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := NewFeature(tt.input)

            if tt.expectError {
                assert.Error(t, err)
                if tt.errorMsg != "" {
                    assert.Contains(t, err.Error(), tt.errorMsg)
                }
                return
            }

            require.NoError(t, err)
            require.NotNil(t, result)
        })
    }
}
```

## Continuous Integration

Tests should be run in CI/CD pipelines:

```yaml
# Example GitHub Actions workflow
- name: Run tests
  run: make test-short

- name: Run integration tests
  run: make test-integration

- name: Generate coverage
  run: make test-coverage

- name: Upload coverage
  run: bash <(curl -s https://codecov.io/bash)
```

## Makefile Targets

All available test targets:

```bash
make test                  # Run all tests (including property tests)
make test-short            # Run unit tests only (skip integration)
make test-properties       # Run only property-based tests
make test-coverage         # Run tests with coverage report
make test-coverage-detail  # Show coverage by package
make test-integration      # Run integration tests (requires Docker)
make test-config           # Test config package only
make test-appdir           # Test appdir package only
make test-database         # Test database package only
make test-runtime          # Test runtime package only
make test-web              # Test web package only
make test-supervisor       # Test supervisor package only
```

## Troubleshooting

### Tests Fail with "Docker not available"

Integration tests require Docker. Either:
- Install and start Docker
- Skip integration tests with `make test-short`

### Tests Fail with "command not found"

Some tests expect certain commands (supervisorctl, mysql, etc.). These tests:
- Use dry-run mode when possible
- Skip when commands are not available
- Are marked as integration tests

### Coverage Report Not Generated

```bash
# Ensure coverage.out exists
ls -la coverage.out

# Regenerate
go test -coverprofile=coverage.out ./...

# View HTML report
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

## Future Improvements

Planned testing enhancements:

1. **Increase coverage** to 60%+ overall
2. **Add mock HTTP server** for web backend tests
3. **Add fixture manifests** for more app types
4. **Property-based testing** for manifest validation ✓ **COMPLETED**
5. **Benchmark tests** for performance-critical paths
6. **E2E tests** for full install workflows
7. **Fuzz testing** for TOML parsing and user input
8. **Expand property tests** to cover more edge cases and invariants
