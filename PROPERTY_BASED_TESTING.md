# Property-Based Testing with Rapid

This document explains property-based testing in the Uberman project using the [pgregory.net/rapid](https://github.com/flyingmutant/rapid) library.

## Table of Contents

- [What is Property-Based Testing?](#what-is-property-based-testing)
- [Why Rapid?](#why-rapid)
- [Quick Start](#quick-start)
- [Writing Property Tests](#writing-property-tests)
- [Property Patterns](#property-patterns)
- [Best Practices](#best-practices)
- [Examples from Uberman](#examples-from-uberman)
- [Running Property Tests](#running-property-tests)
- [Debugging Failed Properties](#debugging-failed-properties)
- [References](#references)

## What is Property-Based Testing?

Property-based testing is a testing methodology where you define **properties** (invariants) that should hold true for all possible inputs, rather than testing specific example cases.

### Traditional Example-Based Testing

```go
func TestDatabaseName(t *testing.T) {
    result := GenerateDatabaseName("john", "wordpress")
    assert.Equal(t, "john_wordpress", result)

    result = GenerateDatabaseName("alice", "ghost")
    assert.Equal(t, "alice_ghost", result)
}
```

### Property-Based Testing

```go
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

**Key Difference**: Instead of testing 2 specific cases, property-based testing automatically generates hundreds of random test cases and verifies the property holds for all of them.

## Why Rapid?

We chose `pgregory.net/rapid` for Uberman because:

1. **Native Go Integration**: Works seamlessly with Go's testing package
2. **Powerful Generators**: Rich set of built-in generators for common types
3. **Regex Support**: `rapid.StringMatching()` for pattern-based string generation
4. **Shrinking**: Automatically minimizes failing test cases
5. **Deterministic**: Uses seeds for reproducible test failures
6. **Simple API**: Clean, idiomatic Go API

### Alternatives Considered

- **gopter**: More complex API, less active development
- **go-fuzz**: Focused on fuzzing rather than property testing
- **testing/quick**: Limited generators, no regex support

## Quick Start

### 1. Import Rapid

```go
import (
    "testing"
    "pgregory.net/rapid"
)
```

### 2. Write a Property Test

```go
func TestProperty_MyFunction_Property(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        // Generate random inputs
        input := rapid.String().Draw(t, "input")

        // Call function under test
        result := MyFunction(input)

        // Assert property
        if !someProperty(result) {
            t.Fatalf("property violated: %v", result)
        }
    })
}
```

### 3. Run Tests

```bash
# Run all tests including properties
make test

# Run only property-based tests
make test-properties

# Run with verbose output
go test -v ./internal/config -run Property
```

## Writing Property Tests

### Basic Structure

Every property test follows this structure:

```go
func TestProperty_ComponentName_PropertyDescription(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        // 1. Generate random inputs
        input := rapid.Int().Draw(t, "input")

        // 2. Execute the code under test
        result := FunctionUnderTest(input)

        // 3. Assert the property holds
        if !PropertyHolds(result) {
            t.Fatalf("property violated")
        }
    })
}
```

### Naming Convention

Property tests in Uberman follow this naming pattern:

```
TestProperty_<Package>_<PropertyDescription>
```

Examples:
- `TestProperty_DatabaseName_FollowsConvention`
- `TestProperty_ServiceConfig_ValidINI`
- `TestProperty_AppRoot_FollowsConvention`

### Generators

Rapid provides many built-in generators:

#### Basic Types

```go
// Integers
i := rapid.Int().Draw(t, "i")                    // Any int
i := rapid.IntRange(1, 100).Draw(t, "i")        // Range [1, 100]
i := rapid.Int64().Draw(t, "i")                 // int64

// Strings
s := rapid.String().Draw(t, "s")                // Any string
s := rapid.StringN(5, 20, -1).Draw(t, "s")     // Length 5-20
s := rapid.StringMatching(`^[a-z]+$`).Draw(t, "s") // Regex pattern

// Booleans
b := rapid.Bool().Draw(t, "b")

// Floats
f := rapid.Float64().Draw(t, "f")
```

#### Collections

```go
// Slices
slice := rapid.SliceOf(rapid.Int()).Draw(t, "slice")
slice := rapid.SliceOfN(rapid.String(), 1, 10).Draw(t, "slice")

// Maps
m := rapid.MapOf(
    rapid.String(),
    rapid.Int(),
).Draw(t, "m")
```

#### Sampling

```go
// Pick from specific values
choice := rapid.SampledFrom([]string{"php", "python", "nodejs"}).Draw(t, "choice")
```

#### Custom Generators

```go
func generateValidManifest(t *rapid.T) *AppManifest {
    return &AppManifest{
        App: AppConfig{
            Name: rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "name"),
            Type: rapid.SampledFrom([]string{"php", "python", "nodejs"}).Draw(t, "type"),
        },
    }
}
```

### Filters

Filter out invalid inputs:

```go
// Generate strings that are NOT "php"
s := rapid.String().
    Filter(func(s string) bool { return s != "php" }).
    Draw(t, "s")
```

## Property Patterns

### 1. Idempotency

**Property**: Applying an operation multiple times has the same effect as applying it once.

```go
func TestProperty_Validate_Idempotent(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        manifest := generateValidManifest(t)

        err1 := manifest.Validate()
        err2 := manifest.Validate()

        // Both should have same result
        if (err1 == nil) != (err2 == nil) {
            t.Fatalf("validation is not idempotent")
        }
    })
}
```

### 2. Round-Trip / Encoding-Decoding

**Property**: Encoding then decoding should return the original value.

```go
func TestProperty_ManifestSerialization_RoundTrip(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        original := generateValidManifest(t)

        // Encode
        encoded, _ := toml.Marshal(original)

        // Decode
        var decoded AppManifest
        toml.Unmarshal(encoded, &decoded)

        // Should match
        if decoded.App.Name != original.App.Name {
            t.Fatalf("round-trip failed")
        }
    })
}
```

### 3. Invariants

**Property**: Certain conditions should always hold.

```go
func TestProperty_DatabaseName_WithinMySQLLimit(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        username := rapid.String().Draw(t, "username")
        appName := rapid.String().Draw(t, "appName")

        dbName := GenerateDatabaseName(username, appName)

        // Invariant: Must be <= 64 characters
        if len(dbName) > 64 {
            t.Fatalf("exceeds MySQL limit: %d chars", len(dbName))
        }
    })
}
```

### 4. Consistency / Determinism

**Property**: Same inputs should always produce same outputs.

```go
func TestProperty_GenerateDatabaseName_Deterministic(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        username := rapid.String().Draw(t, "username")
        appName := rapid.String().Draw(t, "appName")

        name1 := GenerateDatabaseName(username, appName)
        name2 := GenerateDatabaseName(username, appName)

        if name1 != name2 {
            t.Fatalf("not deterministic: %q vs %q", name1, name2)
        }
    })
}
```

### 5. Inverse Functions

**Property**: Applying a function and its inverse should return the original.

```go
func TestProperty_Encryption_Inverse(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        plaintext := rapid.String().Draw(t, "plaintext")

        encrypted := Encrypt(plaintext)
        decrypted := Decrypt(encrypted)

        if decrypted != plaintext {
            t.Fatalf("encrypt/decrypt not inverse")
        }
    })
}
```

### 6. Structural Properties

**Property**: Output structure should match expected format.

```go
func TestProperty_ServiceConfig_ValidINI(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        service := generateValidService(t)

        config := GenerateServiceConfig(service)

        // Should start with [program:name]
        if !strings.HasPrefix(config, "[program:") {
            t.Fatalf("invalid INI format")
        }
    })
}
```

## Best Practices

### 1. Name Tests Clearly

✅ **Good**: `TestProperty_DatabaseName_FollowsConvention`

❌ **Bad**: `TestDatabaseNameProperty`

### 2. Document Properties

Always add a comment explaining what property is being tested:

```go
// Property: Database names should never exceed MySQL's 64-character limit
// REQ-PBT-DB-002
func TestProperty_DatabaseName_WithinMySQLLimit(t *testing.T) {
    // ...
}
```

### 3. Use Meaningful Generator Names

The second argument to `Draw()` helps with debugging:

```go
username := rapid.String().Draw(t, "username")  // ✅ Good
s := rapid.String().Draw(t, "s")                // ❌ Bad
```

### 4. Keep Properties Simple

Each test should verify **one** property. Split complex properties into multiple tests.

### 5. Use Filters Sparingly

Filters can make test generation slow. Prefer specific generators:

```go
// ❌ Slow
s := rapid.String().Filter(func(s string) bool {
    return len(s) > 5 && len(s) < 20
}).Draw(t, "s")

// ✅ Fast
s := rapid.StringN(6, 19, -1).Draw(t, "s")
```

### 6. Combine with Example-Based Tests

Property tests complement, not replace, example-based tests:

- **Example tests**: Edge cases, specific regressions, documentation
- **Property tests**: General behavior, invariants, random edge cases

### 7. Use t.TempDir() for File Operations

Always use temporary directories for tests that create files:

```go
rapid.Check(t, func(t *rapid.T) {
    tmpDir := t.TempDir() // Automatically cleaned up
    // ...
})
```

### 8. Set Appropriate Test Counts

Rapid runs 100 iterations by default. Adjust if needed:

```go
// Run more iterations for critical code
rapid.Check(t, func(t *rapid.T) {
    // ...
}, rapid.CheckIterations(1000))
```

## Examples from Uberman

### Config Package

**Property**: Validation is idempotent

```go
func TestProperty_ValidManifestValidation_Idempotent(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        manifest := generateValidManifest(t)

        err1 := manifest.Validate()
        err2 := manifest.Validate()

        if (err1 == nil) != (err2 == nil) {
            t.Fatalf("validation not idempotent")
        }
    })
}
```

### Appdir Package

**Property**: Created structure always validates

```go
func TestProperty_CreatedStructure_AlwaysValidates(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

        manager, _ := NewManager(appName, false, false)
        manager.Create()

        // Create required config file
        os.WriteFile(manager.ConfigFile(), []byte("[app]\n"), 0644)

        if err := manager.Validate(); err != nil {
            t.Fatalf("created structure doesn't validate: %v", err)
        }
    })
}
```

### Database Package

**Property**: Database names follow convention

```go
func TestProperty_DatabaseName_FollowsConvention(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
        appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

        dbName := GenerateDatabaseName(username, appName)

        if !strings.HasPrefix(dbName, username + "_") {
            t.Fatalf("doesn't follow convention: %q", dbName)
        }
    })
}
```

### Supervisor Package

**Property**: Generated config is valid INI

```go
func TestProperty_ServiceConfig_ValidINI(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        service := generateValidService(t)

        manager := NewServiceManager(false, false)
        manager.CreateService(service)

        // Read config file
        content, _ := os.ReadFile(configPath)

        // Should start with [program:name]
        if !strings.HasPrefix(string(content), "[program:") {
            t.Fatalf("invalid INI format")
        }
    })
}
```

### Web Package

**Property**: Found ports are always in valid range

```go
func TestProperty_FindAvailablePort_ValidRange(t *testing.T) {
    rapid.Check(t, func(t *rapid.T) {
        start := rapid.IntRange(1024, 32767).Draw(t, "start")
        end := rapid.IntRange(start+1, 65535).Draw(t, "end")

        manager := NewBackendManager(false, false)
        port, err := manager.FindAvailablePort(start, end)

        if err == nil {
            if port < start || port > end {
                t.Fatalf("port %d not in range [%d, %d]", port, start, end)
            }
        }
    })
}
```

## Running Property Tests

### Run All Tests

```bash
make test
```

### Run Only Property Tests

```bash
make test-properties
```

### Run Property Tests for Specific Package

```bash
go test -v ./internal/config -run Property
go test -v ./internal/database -run Property
```

### Run Single Property Test

```bash
go test -v ./internal/config -run TestProperty_DatabaseName_FollowsConvention
```

### Run with More Iterations

```bash
go test -v ./internal/config -run Property -rapid.checks=1000
```

### Run with Specific Seed (for reproducing failures)

```bash
go test -v ./internal/config -run Property -rapid.seed=12345
```

## Debugging Failed Properties

When a property test fails, Rapid automatically **shrinks** the input to find the minimal failing case.

### Example Failure Output

```
--- FAIL: TestProperty_DatabaseName_WithinMySQLLimit (0.03s)
    rapid.go:123: failed after 47 checks: database name exceeds limit:
        "verylongusername_verylongappnamewithmanychars" (67 chars)
    rapid.go:124: minimal failing case:
        username: "user123456789012"
        appName: "app12345678901234567890123456789012345678901234567"
    rapid.go:125: seed: 1638234567890
```

### Steps to Debug

1. **Copy the seed**: The failure output shows the random seed used
2. **Reproduce**: Run with that seed: `go test -rapid.seed=1638234567890`
3. **Examine minimal case**: Rapid shows the simplest input that causes failure
4. **Fix the bug**: Update your code to handle the edge case
5. **Verify**: Run tests again to ensure fix works

### Adding Diagnostic Output

```go
rapid.Check(t, func(t *rapid.T) {
    input := rapid.String().Draw(t, "input")

    result := MyFunction(input)

    if !PropertyHolds(result) {
        t.Logf("Input: %q", input)
        t.Logf("Result: %v", result)
        t.Fatalf("property violated")
    }
})
```

## References

- [Rapid GitHub Repository](https://github.com/flyingmutant/rapid)
- [Rapid Documentation](https://pkg.go.dev/pgregory.net/rapid)
- [Property-Based Testing Intro](https://hypothesis.works/articles/what-is-property-based-testing/)
- [QuickCheck Paper](https://www.cs.tufts.edu/~nr/cs257/archive/john-hughes/quick.pdf) (original PBT paper)
- [Test-Driven Development by Example](https://www.amazon.com/Test-Driven-Development-Kent-Beck/dp/0321146530) by Kent Beck

## Contributing

When adding new property tests to Uberman:

1. Follow the naming convention: `TestProperty_<Component>_<Property>`
2. Add a comment explaining the property being tested
3. Reference the requirement ID from the specification (e.g., `REQ-PBT-DB-002`)
4. Use meaningful generator names in `Draw()` calls
5. Keep properties simple and focused
6. Add examples to this documentation for novel patterns

## Questions?

For questions about property-based testing in Uberman:

1. Check this documentation first
2. Look at existing property tests in the codebase
3. Consult the [Rapid documentation](https://pkg.go.dev/pgregory.net/rapid)
4. Open an issue with the `testing` label
