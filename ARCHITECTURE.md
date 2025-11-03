# Architecture - Clean Architecture from Scratch

## Clean Slate Approach

**Status**: Existing `internal/` and `apps/` deleted
**Approach**: Rebuild from scratch with correct bounded context and Clean Architecture

## Single Bounded Context

**Name**: "App Installation" (directory: `appinstallation`)
**Focus**: Model installation as workflow, not just data structures

## Clean Architecture Layers

```
┌─────────────────────────────────────┐
│     Presentation Layer (CLI)        │
│     cmd/uberman/                    │
│     - Cobra commands                │
│     - Flag parsing                  │
│     - Output formatting             │
└──────────────┬──────────────────────┘
               │ depends on
┌──────────────▼──────────────────────┐
│     Application Layer                │
│     internal/appinstallation/       │
│                 application/         │
│     - InstallAppUseCase             │
│     - DTOs (data transfer objects)  │
└──────────────┬──────────────────────┘
               │ depends on
┌──────────────▼──────────────────────┐
│     Domain Layer (100% PURE)         │
│     internal/appinstallation/domain/ │
│                                      │
│     workflow/                        │
│     - Installation (aggregate root)  │
│     - States, Commands, Events       │
│                                      │
│     manifest/                        │
│     - Manifest (aggregate)           │
│     - Parser (pure TOML parsing)     │
│                                      │
│     valueobjects/                    │
│     - AppName, Port, DatabaseName    │
│     - DirectoryPath, FilePath        │
│                                      │
│     ports/                           │
│     - Repository interfaces          │
│     - Provisioner interface          │
└──────────────▲──────────────────────┘
               │ implemented by
┌──────────────┴──────────────────────┐
│     Infrastructure Layer             │
│     internal/appinstallation/       │
│                infrastructure/       │
│                                      │
│     filesystem/                      │
│     - Real I/O (os.MkdirAll, etc.)  │
│     - Uberspace command execution    │
│                                      │
│     memory/                          │
│     - In-memory for fast tests       │
│     - No side effects                │
│                                      │
│     dryrun/                          │
│     - Preview mode (logs only)       │
└─────────────────────────────────────┘
```

## Dependency Rules (ENFORCE)

### Rule 1: Dependencies Point INWARD Only
- Presentation depends on Application
- Application depends on Domain
- Infrastructure depends on Domain (implements ports)
- Domain depends on NOTHING

### Rule 2: Domain is 100% Pure
**NO imports allowed in domain layer:**
- ❌ `os` - File system operations
- ❌ `exec` - Command execution
- ❌ `io` - I/O operations
- ❌ `net` - Network operations
- ❌ `syscall` - System calls
- ❌ Any infrastructure concerns

**Validation**:
```bash
go list -f '{{.ImportPath}}: {{.Imports}}' ./internal/appinstallation/domain/...
```

Should see ZERO forbidden packages.

### Rule 3: Use Dependency Inversion
Domain defines interfaces (ports).
Infrastructure implements interfaces (adapters).
Application injects adapters into domain.

## Directory Structure

```
internal/
└── appinstallation/                     # Bounded Context
    ├── domain/                          # Domain Layer
    │   ├── workflow/
    │   │   ├── installation.go          # Aggregate root
    │   │   ├── installation_test.go     # < 1ms tests
    │   │   ├── states.go                # State enum
    │   │   ├── commands.go              # Command definitions
    │   │   └── events.go                # Domain events
    │   ├── manifest/
    │   │   ├── manifest.go              # Aggregate
    │   │   ├── manifest_test.go
    │   │   ├── parser.go                # Pure TOML parsing
    │   │   ├── parser_test.go
    │   │   └── repository.go            # Interface (port)
    │   ├── valueobjects/
    │   │   ├── appname.go
    │   │   ├── appname_test.go
    │   │   ├── port.go
    │   │   ├── port_test.go
    │   │   ├── databasename.go
    │   │   ├── databasename_test.go
    │   │   ├── directorypath.go
    │   │   └── directorypath_test.go
    │   └── ports/                       # All interfaces
    │       ├── installation_repository.go
    │       ├── manifest_repository.go
    │       └── provisioner.go
    │
    ├── application/                     # Application Layer
    │   └── installapp/
    │       ├── usecase.go               # Orchestrates workflow
    │       ├── usecase_test.go          # Uses in-memory adapters
    │       └── dto.go                   # Data transfer objects
    │
    └── infrastructure/                  # Infrastructure Layer
        ├── memory/                      # Fast tests (<1ms)
        │   ├── installation_repository.go
        │   ├── installation_repository_test.go
        │   ├── manifest_repository.go
        │   ├── manifest_repository_test.go
        │   ├── provisioner.go
        │   └── provisioner_test.go
        ├── filesystem/                  # Real implementation
        │   ├── installation_repository.go
        │   ├── installation_repository_test.go
        │   ├── manifest_repository.go
        │   ├── manifest_repository_test.go
        │   ├── provisioner.go
        │   └── provisioner_test.go
        └── dryrun/                      # Preview mode
            ├── provisioner.go
            └── provisioner_test.go
```

## Workflow Modeling

### Installation Aggregate (Root)

**Purpose**: Coordinate installation workflow, maintain state

**State Machine**:
```
NotStarted
  → LoadManifest() → ManifestLoaded
  → ValidatePrerequisites() → PrerequisitesValidated
  → ProvisionDirectories() → ProvisioningInProgress
  → ProvisionDatabase() → ProvisioningInProgress
  → ProvisionWebBackend() → ProvisioningInProgress
  → ProvisionService() → ProvisioningInProgress
  → SetRuntimeVersion() → Configured
  → FinalizeInstallation() → Installed

(Any command can transition to → Failed on error)
```

**Interface Example**:
```go
// internal/appinstallation/domain/workflow/installation.go

type Installation struct {
    id       InstallationID
    appName  valueobjects.AppName
    state    State
    manifest *manifest.Manifest
    events   []Event
}

// Commands (drive state transitions)
func (i *Installation) LoadManifest(m *manifest.Manifest) error
func (i *Installation) ValidatePrerequisites() error
func (i *Installation) ProvisionDirectories(paths []valueobjects.DirectoryPath) error
func (i *Installation) ProvisionDatabase(name valueobjects.DatabaseName) error
func (i *Installation) ProvisionWebBackend(backend Backend, port valueobjects.Port) error
func (i *Installation) ProvisionService(service Service) error
func (i *Installation) SetRuntimeVersion(runtime Runtime) error
func (i *Installation) FinalizeInstallation() error

// Queries
func (i *Installation) State() State
func (i *Installation) Events() []Event
func (i *Installation) IsComplete() bool
func (i *Installation) HasFailed() bool
```

**Key Property**: ALL methods are pure logic - they update state, validate, emit events. NO I/O.

### Repository Pattern (Ports & Adapters)

**Port (Domain Interface)**:
```go
// internal/appinstallation/domain/ports/provisioner.go

type Provisioner interface {
    CreateDirectories(ctx context.Context, paths []valueobjects.DirectoryPath) error
    CreateDatabase(ctx context.Context, name valueobjects.DatabaseName) error
    ConfigureWebBackend(ctx context.Context, backend Backend, port valueobjects.Port) error
    CreateService(ctx context.Context, service Service) error
    SetRuntimeVersion(ctx context.Context, runtime Runtime) error
}
```

**Adapter (Infrastructure Implementation)**:
```go
// internal/appinstallation/infrastructure/filesystem/provisioner.go

type Provisioner struct {
    // Real dependencies (exec.Cmd, os filesystem, etc.)
}

func (p *Provisioner) CreateDirectories(ctx context.Context, paths []valueobjects.DirectoryPath) error {
    // Here we can use os.MkdirAll - infrastructure layer allows it
    for _, path := range paths {
        if err := os.MkdirAll(string(path), 0755); err != nil {
            return fmt.Errorf("failed to create directory %s: %w", path, err)
        }
    }
    return nil
}
```

## Testing Infrastructure

### Platform Constraints

**CRITICAL**: Cannot test on Uberspace itself (no CI/CD access to hosting environment).

**Required Testing Approach**:
- **Domain layer**: Pure tests (no I/O), < 1ms per test
- **Infrastructure layer**: Isolated tests with testcontainers
- **Integration**: ALL side effects MUST run in Docker containers

### Testing Tools & Frameworks

**Property-Based Testing**:
- **Framework**: `pgregory.net/rapid`
- **Purpose**: Verify invariants across generated inputs
- **Common patterns**: Idempotency, round-trip, invariants, determinism
- **Usage**: `rapid.Check(t, func(t *rapid.T) { ... })`

**Container-Based Testing**:
- **Framework**: `testcontainers-go` (MySQL module, generic containers)
- **Purpose**: Test infrastructure adapters with real dependencies
- **Required for**: Database operations, filesystem side effects, external commands
- **Skip in short mode**: `if testing.Short() { t.Skip() }`

**TDD Agent Integration**:
For test-driven development assistance, use the specialized agent:
```
@tdd-orchestrator

I need to implement [feature] in the [domain/infrastructure] layer
```

Agent provides:
- Red-Green-Refactor cycle guidance
- Test structure recommendations
- Property-based test generation
- Container setup for integration tests

### Testing Commands

```bash
# Fast TDD cycle (native execution)
go test -v ./internal/appinstallation/domain/...

# Property-based tests only
go test -v -run Property ./...

# Integration tests (requires Docker)
go test -v -run Integration ./...

# Skip integration tests
go test -short ./...

# All tests with coverage
go test -cover ./...

# Makefile targets
make test              # All tests
make test-short        # Skip integration tests
make test-properties   # Property-based only
```

### Testing Rules (Architecture Enforcement)

**Domain Layer Rules**:
- ✅ **ZERO imports from**: `os`, `exec`, `io`, `net`, `syscall`
- ✅ **Tests MUST be pure** (no side effects, no I/O)
- ✅ **Performance target**: < 1ms per test
- ✅ **No mocks needed** (pure functions)

**Infrastructure Layer Rules**:
- ✅ **Use testcontainers** for all external dependencies
- ✅ **Tests run in isolated Docker containers**
- ✅ **Skip in short mode** for faster TDD cycles
- ✅ **Performance target**: 10-20ms per test acceptable

**Validation Command**:
```bash
# Verify domain layer has no forbidden imports
go list -f '{{.ImportPath}}: {{.Imports}}' ./internal/appinstallation/domain/...

# Should see ZERO: os, exec, io, net, syscall
```

### Test Performance Targets

| Layer | Target | Constraint |
|-------|--------|------------|
| Domain | < 1ms per test | Pure functions, no I/O |
| Application | < 5ms per test | In-memory adapters |
| Infrastructure | 10-20ms per test | Testcontainers overhead acceptable |
| E2E | 100ms+ per test | CLI integration acceptable |

### Testing Strategy

#### Test Pyramid

```
         ▲
        / \
       /E2E\              ← 5% (End-to-end CLI tests)
      /─────\
     /Integ  \            ← 15% (Filesystem adapter tests)
    /─────────\
   /Application\          ← 30% (Use case tests with in-memory)
  /─────────────\
 /    Domain     \        ← 50% (Pure unit tests)
/─────────────────\
```

#### Domain Layer Tests (50%)
- **Pure unit tests**
- **< 1ms per test**
- **No mocks needed** (pure functions)
- **100% coverage goal**
- Test workflow state transitions
- Test value object validation
- Property-based tests for invariants

**Example**:
```bash
$ go test ./internal/appinstallation/domain/...
ok    internal/appinstallation/domain/workflow      0.003s
ok    internal/appinstallation/domain/manifest      0.002s
ok    internal/appinstallation/domain/valueobjects  0.001s
```

#### Application Layer Tests (30%)
- **Use case tests with in-memory adapters**
- **< 5ms per test**
- Verify orchestration logic
- Test error handling
- Test all workflow paths

#### Infrastructure Layer Tests (15%)
- **Contract tests** (all adapters implement interfaces)
- **Integration tests** (with real filesystem/testcontainers)
- **Slower but thorough** (10-20ms per test acceptable)
- **Always use testcontainers** for database/filesystem isolation

**Example with testcontainers**:
```go
func TestDatabaseProvisioner_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test in short mode")
    }

    ctx := context.Background()
    container, err := mysql.Run(ctx, "mysql:8.0",
        mysql.WithDatabase("testdb"),
        mysql.WithUsername("testuser"),
        mysql.WithPassword("testpass"),
    )
    require.NoError(t, err)
    defer container.Terminate(ctx)

    // Test database provisioner with real MySQL...
}
```

#### End-to-End Tests (5%)
- **CLI integration tests**
- **Verify user-facing behavior**
- **Slowest** (100ms+ acceptable)
- Test happy path + critical errors

## Build Order (Probe-Sense-Respond)

Following Shape Up methodology for COMPLEX domain (Cynefin):

### Phase 1: PROBE (Weeks 1-2)
**Build**: Domain layer (value objects + Installation aggregate + Manifest aggregate)
**Goal**: Validate workflow modeling approach

### Phase 2: SENSE (End of Week 2)
**Questions**:
- Are tests fast (< 1ms)?
- Is domain pure (no I/O imports)?
- Does workflow feel natural?
- Is state machine too complex?

**Decisions**:
- Continue with state machine? Or simplify?
- TOML parsing in domain? Or move to infrastructure?
- Installation as aggregate root? Or something else?

### Phase 3: RESPOND (Weeks 3-4)
**Build**: Infrastructure adapters based on learnings

**Options**:
- If domain is clean → build filesystem + in-memory adapters
- If domain too complex → simplify before continuing
- If tests slow → investigate why (domain not pure?)

### Phase 4: PROBE (Week 5)
**Build**: Application layer (use case)

**Goal**: Validate orchestration pattern

### Phase 5: RESPOND (Week 6)
**Build**: CLI integration

**Goal**: Ship working install command

**Scope Hammer**: Cut application layer if needed - CLI can call domain directly

## Validation Commands

### Check Domain Purity
```bash
go list -f '{{.ImportPath}}: {{.Imports}}' ./internal/appinstallation/domain/...

# Should see ZERO: os, exec, io, net, syscall
```

### Run Tests by Layer
```bash
# Domain (should be < 1ms per test)
go test -v ./internal/appinstallation/domain/...

# Application (should be < 5ms per test)
go test -v ./internal/appinstallation/application/...

# Infrastructure (10-20ms acceptable)
go test -v ./internal/appinstallation/infrastructure/...

# All tests
go test -v ./internal/appinstallation/...
```

### Test Coverage
```bash
# Overall coverage
go test -cover ./internal/appinstallation/...

# Detailed coverage report
go test -coverprofile=coverage.out ./internal/appinstallation/...
go tool cover -html=coverage.out
```

### Lint
```bash
golangci-lint run ./internal/appinstallation/...
```

## Architecture Decision Records (ADRs)

### ADR-001: Single Bounded Context
**Status**: Accepted
**Decision**: Use single bounded context "App Installation" vs multiple contexts
**Rationale**: No linguistic boundaries exist - all terms have consistent meanings
**Date**: 2025-11-03

### ADR-002: Workflow as Aggregate Root
**Status**: TBD (To be determined during week 1-2)
**Decision**: Installation aggregate as root vs App aggregate as root
**Rationale**: Need to validate during implementation
**Date**: TBD

### ADR-003: Explicit State Machine
**Status**: TBD (To be determined during week 1-2)
**Decision**: Use explicit state machine vs simpler procedural workflow
**Rationale**: Need to validate complexity vs value during implementation
**Date**: TBD

## References

- **PRD.md**: Requirements using EARS format
- **UBIQUITOUS_LANGUAGE.md**: Domain glossary
- **PLANNING.md**: Implementation tasks
- **plans/pitches/2025-11-03-rebuild-install-clean-architecture.md**: Shape Up pitch
