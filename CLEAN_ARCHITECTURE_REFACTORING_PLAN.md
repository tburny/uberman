# Clean Architecture Refactoring Plan - Uberman CLI
## TDD-Driven Domain-Driven Design Implementation

**Version:** 2.0 (Comprehensive TDD Edition)
**Status:** In Progress - Phase 1 (Day 1)
**Timeline:** 8-12 weeks (revised for thorough TDD approach)
**Approach:** Complete Clean Architecture + DDD with Test-First Development
**Last Updated:** 2025-11-02
**Author:** TDD Orchestrator Agent

---

## Executive Summary

Transforming Uberman from a transaction script architecture with mixed concerns into a Clean Architecture implementation enhanced with Domain-Driven Design principles, using comprehensive Test-Driven Development practices.

### Critical Architectural Issues Identified

#### 1. Side Effects Deeply Embedded in Business Logic
- **Violation Location**: All infrastructure "Manager" types directly execute side effects
  - `internal/appdir/manager.go:94`: `Create()` directly calls `os.MkdirAll()`, `os.Symlink()`
  - `internal/database/mysql.go:115`: `CreateDatabase()` directly executes `exec.Command("mysql", ...)`
  - `internal/web/backend.go:58`: `SetBackend()` directly executes `exec.Command("uberspace", ...)`
  - `internal/supervisor/service.go:100`: `CreateService()` directly writes files with `os.Create()`
  - `internal/runtime/uberspace.go:34`: `SetVersion()` directly executes commands

#### 2. Missing Architectural Layers
- ❌ **No Domain Layer**: Missing entities, aggregates, domain services, domain events
- ❌ **No Application Layer**: No use cases or application services orchestrating domain logic
- ❌ **No Clear Ports**: No interfaces defining contracts between layers
- ❌ **No Infrastructure Layer**: Infrastructure concerns mixed directly into current managers

#### 3. Anemic Domain Model
- `internal/config/app.go` contains only data structures with no behavior
- Business rules scattered across infrastructure code
- No domain language expressed in code
- Impossible to enforce invariants

#### 4. Strong Foundation Elements (To Preserve)
- ✅ Excellent value objects in `internal/domain/common/value_objects.go`
- ✅ Property-based tests using pgregory.net/rapid
- ✅ Good test coverage patterns (unit + property tests)
- ✅ Clear manifest-based configuration
- ✅ Dry-run mode support throughout

### Target Benefits
- **100% domain layer test coverage** - Pure logic testable in < 1ms per test
- **20-30x faster property tests** - From 2-3s to < 100ms (in-memory adapters)
- **Pluggable infrastructure** - Filesystem, in-memory, dry-run adapters
- **Clear boundaries** - Domain (pure), Application (orchestration), Infrastructure (I/O)
- **Testability without external dependencies** - Domain logic fully testable without filesystem, database, or commands
- **Multiple delivery mechanisms** - Foundation for future web UI or API without duplicating logic

---

## Architecture Overview

### Clean Architecture Layers

```
┌──────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                         │
│                    cmd/uberman/*.go                           │
│  - CLI Commands (Cobra)                                       │
│  - Flag parsing                                               │
│  - Output formatting                                          │
└────────────────┬─────────────────────────────────────────────┘
                 │ Depends on ↓
┌────────────────┴─────────────────────────────────────────────┐
│                   APPLICATION LAYER                           │
│              internal/application/*.go                        │
│  - Use Cases (CreateAppUseCase, UpgradeAppUseCase)          │
│  - Application Services                                       │
│  - DTOs for cross-boundary communication                      │
└────────────────┬─────────────────────────────────────────────┘
                 │ Depends on ↓
┌────────────────┴─────────────────────────────────────────────┐
│                     DOMAIN LAYER                              │
│              internal/domain/*.go                             │
│  - Entities (App, Database, Service)                         │
│  - Value Objects (AppName, Port, DatabaseName)              │
│  - Domain Services (pure business logic)                     │
│  - Repository Interfaces (NO implementations!)               │
│  ✓ NO dependencies on outer layers                           │
│  ✓ 100% pure Go (no os, exec, net packages)                 │
└────────────────▲─────────────────────────────────────────────┘
                 │ Implemented by ↓
┌────────────────┴─────────────────────────────────────────────┐
│                 INFRASTRUCTURE LAYER                          │
│             internal/infrastructure/*.go                      │
│  - Repository Implementations                                 │
│  - FileSystemAdapter (os.MkdirAll, os.Symlink)              │
│  - InMemoryAdapter (for fast tests)                          │
│  - DryRunAdapter (preview mode)                              │
└──────────────────────────────────────────────────────────────┘
```

**Dependency Rule**: Dependencies point **INWARD ONLY**

---

## Phase-by-Phase Implementation Plan

### ✅ Progress Tracker

- [x] **Phase 1-2 Started**: Create domain package structure
- [x] **Phase 1-2 Started**: Create value objects (AppName, Port, DatabaseName, etc.)
- [ ] Phase 1-2: Extract AppStructure entity with pure logic
- [ ] Phase 1-2: Define repository interfaces (ports)
- [ ] Phase 1-2: Write domain tests (100% coverage goal)
- [ ] Phase 1-2: Implement filesystem adapter
- [ ] Phase 1-2: Implement in-memory adapter for tests
- [ ] Phase 1-2: Implement dry-run adapter
- [ ] Phase 3: Create use cases layer with DTOs
- [ ] Phase 3: Test use cases with in-memory adapters
- [ ] Phase 4: Create DI container for CLI
- [ ] Phase 4: Refactor CLI commands to use dependency injection
- [ ] Phase 5: Remove old packages (appdir, database, web, etc.)
- [ ] Phase 5: Update all tests and documentation

---

### **Phase 1-2 Combined: Foundation (Days 1-3)** ⚡ IN PROGRESS

Extract domain layer + implement adapters simultaneously

#### **Completed Work:**

✅ **Domain Package Structure Created:**
```
internal/domain/
├── app/              # AppStructure entity, Repository interface
├── database/         # Database entity, DatabaseName VO
├── service/          # ServiceDefinition entity
├── runtime/          # RuntimeVersion entity
├── web/              # Backend entity, PortAllocation service
└── common/           # Shared value objects
```

✅ **Value Objects Implemented & Tested:**
- `AppName` - Validated app names (lowercase, 3-21 chars)
- `DirectoryPath` - Absolute directory paths
- `FilePath` - Absolute file paths
- `Port` - Valid ports (1024-65535)
- `DatabaseName` - Uberspace convention (`${USER}_${APP}`)

**Test Results:** All tests pass in 0.003s ✨

#### **Remaining Tasks:**

1. **Extract AppStructure Entity** (Next Step)
   - Source: `internal/appdir/manager.go` lines 35-66 (path calculations)
   - Create: `internal/domain/app/app_structure.go`
   - Make all methods pure (no I/O)
   - Write comprehensive tests

2. **Extract Other Domain Entities:**
   - `Database` entity from `internal/database/mysql.go`
   - `ServiceDefinition` from `internal/supervisor/service.go`
   - `RuntimeVersion` from `internal/runtime/uberspace.go`
   - `WebBackend` from `internal/web/backend.go`

3. **Define Repository Interfaces (Ports):**
   - `app.Repository` for app structure operations
   - `database.Repository` for database operations
   - `service.Repository` for supervisord services
   - `runtime.Repository` for runtime version management
   - `web.Repository` for web backend configuration

4. **Implement Adapters:**
   ```
   internal/infrastructure/
   ├── filesystem/       # Real I/O (os.MkdirAll, os.Symlink)
   ├── memory/           # In-memory for fast tests
   └── dryrun/           # Preview mode (logs only)
   ```

#### **Success Criteria:**
- ✅ Domain layer: 100% pure (zero `os`, `exec`, `io` imports)
- ✅ Domain tests: < 1ms per test, 100% coverage
- ✅ Property tests: < 100ms for 100 iterations (was 2-3 seconds)
- ✅ All adapters implement repository interfaces

---

### **Phase 3: Application Layer (Days 4-5)**

Create use cases that orchestrate domain + infrastructure

#### **Tasks:**

1. **Define Use Cases:**
   ```
   internal/application/usecases/
   ├── create_app.go        # App creation orchestration
   ├── upgrade_app.go       # App upgrade with backup
   ├── backup_app.go        # Backup creation
   ├── validate_app.go      # Structure validation
   └── dto/                 # Input/output DTOs
   ```

2. **Implement CreateAppUseCase Example:**
   ```go
   type CreateAppUseCase struct {
       appRepo app.Repository  // Injected
       logger  Logger          // Injected
   }

   func (uc *CreateAppUseCase) Execute(ctx context.Context, cmd CreateAppCommand) (*CreateAppResult, error) {
       // 1. Create domain entity (pure)
       structure, err := app.NewAppStructure(appName, homeDir)

       // 2. Check existence (delegate to repository)
       exists, err := uc.appRepo.Exists(ctx, appName)

       // 3. Execute creation (delegate to repository)
       if !cmd.DryRun {
           err := uc.appRepo.Create(ctx, structure)
       }

       // 4. Return result (DTO)
       return &CreateAppResult{...}, nil
   }
   ```

3. **Test with In-Memory Adapters:**
   ```go
   repo := memory.NewAppRepository()  // No filesystem!
   useCase := usecases.NewCreateAppUseCase(repo, logger)
   result, err := useCase.Execute(ctx, cmd)
   // Test completes in < 1ms
   ```

#### **Success Criteria:**
- ✅ Use cases tested with mocked repositories
- ✅ 95%+ test coverage on application layer
- ✅ Use case tests run in < 5ms each

---

### **Phase 4: CLI Integration (Days 6-8)**

Wire up dependency injection, refactor all commands

#### **Tasks:**

1. **Create DI Container:**
   ```go
   // cmd/uberman/container.go
   type Container struct {
       CreateApp  *usecases.CreateAppUseCase
       UpgradeApp *usecases.UpgradeAppUseCase
       // ...other use cases
   }

   func NewContainer() *Container {
       appRepo := filesystem.NewAppRepository()
       dbRepo := filesystem.NewDatabaseRepository()
       logger := newLogger(verbose)

       return &Container{
           CreateApp: usecases.NewCreateAppUseCase(appRepo, logger),
           // ...
       }
   }
   ```

2. **Refactor CLI Commands:**
   - `cmd/uberman/install.go` - Use CreateAppUseCase
   - `cmd/uberman/upgrade.go` - Use UpgradeAppUseCase
   - All other commands following same pattern

3. **Adapter Factory Pattern:**
   ```go
   func newAppRepository(dryRun bool) app.Repository {
       if dryRun {
           return dryrun.NewAppRepository()
       }
       return filesystem.NewAppRepository()
   }
   ```

#### **Success Criteria:**
- ✅ All CLI commands use dependency injection
- ✅ Integration tests pass (or updated)
- ✅ CLI interface unchanged (all flags work identically)

---

### **Phase 5: Cleanup (Days 9-10)**

Remove old code, optimize, document

#### **Tasks:**

1. **Delete Old Packages:**
   ```bash
   rm -rf internal/appdir/
   rm internal/database/mysql.go
   rm internal/web/backend.go
   rm internal/supervisor/service.go
   rm internal/runtime/uberspace.go
   ```

2. **Migrate/Update Tests:**
   - Convert existing tests to use new architecture
   - Keep property-based tests (now 20x faster!)
   - Add more property tests (cheap now)

3. **Update Documentation:**
   - `docs/ARCHITECTURE.md` - New architecture diagrams
   - `docs/TESTING.md` - Updated testing guide
   - `README.md` - Developer setup instructions

#### **Success Criteria:**
- ✅ Zero deprecated code remaining
- ✅ All tests pass in < 5 seconds total (vs 30+ seconds before)
- ✅ Test coverage 75%+ (up from 46.2%)
- ✅ Documentation complete

---

## Detailed Entity & Value Object Reference

### Value Objects (Already Implemented ✅)

#### `AppName`
```go
type AppName string
// Pattern: ^[a-z][a-z0-9-]{2,20}$
// Example: "myapp", "my-app-123"
```

#### `DirectoryPath` / `FilePath`
```go
type DirectoryPath string
type FilePath string
// Must be absolute paths
// Example: "/home/user/apps/myapp"
```

#### `Port`
```go
type Port int
// Range: 1024-65535 (no privileged ports)
// Example: 8080, 3000, 5000
```

#### `DatabaseName`
```go
type DatabaseName string
// Convention: ${USER}_${APP}
// Example: "johndoe_wordpress"
```

### Entities To Extract

#### `AppStructure` (from `internal/appdir/manager.go`)

**Pure Methods to Extract:**
```go
// Path calculations (lines 35-66)
func (a *AppStructure) Root() DirectoryPath
func (a *AppStructure) AppCodeDir() DirectoryPath
func (a *AppStructure) DataDir() DirectoryPath
func (a *AppStructure) LogsDir() DirectoryPath
func (a *AppStructure) BackupsDir() DirectoryPath
func (a *AppStructure) TmpDir() DirectoryPath
func (a *AppStructure) ConfigFile() FilePath

// Directory lists (lines 71-78)
func (a *AppStructure) AllDirectories() []DirectoryPath
func (a *AppStructure) RequiredDirectoriesForValidation() []DirectoryPath

// Symlink logic (pure calculation)
func (a *AppStructure) LogSymlink(centralLogsDir DirectoryPath) (symlinkPath, targetPath DirectoryPath)
```

**Key Property:** ALL methods are pure - they only calculate paths, never perform I/O.

#### `Database` (from `internal/database/mysql.go`)

**Domain Logic:**
```go
type Database struct {
    Name        DatabaseName
    Credentials Credentials
    Type        DatabaseType
}

// Pure domain logic
func GenerateDatabaseName(username, appName string) (DatabaseName, error)
func (d *Database) ValidateName() error
```

**Leave in Infrastructure:** Actual mysql command execution

#### `ServiceDefinition` (from `internal/supervisor/service.go`)

**Domain Logic:**
```go
type ServiceDefinition struct {
    Name        ServiceName
    Command     Command
    Port        Port
    Directory   DirectoryPath
    Environment map[string]string
}

// Pure logic
func (s *ServiceDefinition) ConfigurationTemplate() string
func (s *ServiceDefinition) Validate() error
```

#### `WebBackend` (from `internal/web/backend.go`)

**Domain Logic:**
```go
type WebBackend struct {
    Path   URLPath
    Type   BackendType  // Apache or HTTP
    Port   Port         // Optional for HTTP
    Domain Domain       // Optional
}

// Pure logic
func (b *WebBackend) RequiresPort() bool
func (b *WebBackend) Validate() error

// Port allocation (pure algorithm - lines 141-173)
type PortRange struct {
    Start     Port
    End       Port
    UsedPorts []Port
}

func (pr *PortRange) FindAvailablePort() (Port, error)
```

---

## Repository Interfaces (Ports)

### `app.Repository`

```go
package app

type Repository interface {
    Create(ctx context.Context, structure *AppStructure) error
    Exists(ctx context.Context, name common.AppName) (bool, error)
    IsEmpty(ctx context.Context, name common.AppName) (bool, error)
    Remove(ctx context.Context, name common.AppName) error
    Validate(ctx context.Context, structure *AppStructure) error
    CreateLogSymlink(ctx context.Context, structure *AppStructure) error
    RemoveLogSymlink(ctx context.Context, name common.AppName) error
}
```

### `database.Repository`

```go
package database

type Repository interface {
    Create(ctx context.Context, name DatabaseName) error
    Exists(ctx context.Context, name DatabaseName) (bool, error)
    Drop(ctx context.Context, name DatabaseName) error
    Export(ctx context.Context, name DatabaseName, outputPath string) error
    Import(ctx context.Context, name DatabaseName, inputPath string) error
    GetCredentials(ctx context.Context) (*Credentials, error)
}
```

---

## Adapter Implementations

### Filesystem Adapter Pattern

```go
// internal/infrastructure/filesystem/app_repository.go
type AppRepository struct {}

func NewAppRepository() *AppRepository {
    return &AppRepository{}
}

func (r *AppRepository) Create(ctx context.Context, structure *app.AppStructure) error {
    // Get pure domain logic output
    dirs := structure.AllDirectories()

    // Perform side effects (isolated here)
    for _, dir := range dirs {
        if err := os.MkdirAll(string(dir), 0755); err != nil {
            return fmt.Errorf("failed to create directory %s: %w", dir, err)
        }
    }

    return nil
}
```

### In-Memory Adapter Pattern (for tests)

```go
// internal/infrastructure/memory/app_repository.go
type AppRepository struct {
    mu          sync.RWMutex
    structures  map[common.AppName]*app.AppStructure
    directories map[common.AppName]map[string]bool
}

func NewAppRepository() *AppRepository {
    return &AppRepository{
        structures:  make(map[common.AppName]*app.AppStructure),
        directories: make(map[common.AppName]map[string]bool),
    }
}

func (r *AppRepository) Create(ctx context.Context, structure *app.AppStructure) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    name := structure.Name()
    r.structures[name] = structure
    r.directories[name] = make(map[string]bool)

    // Create directories (in-memory)
    for _, dir := range structure.AllDirectories() {
        r.directories[name][string(dir)] = true
    }

    return nil
}
```

**Speed Comparison:**
- Filesystem tests: 15-25ms per test
- In-memory tests: < 1ms per test (15-25x faster!)

---

## Testing Strategy

### Test Pyramid

```
                    ▲
                   ╱ ╲
                  ╱E2E ╲              ← 5% (Slow, high value)
                 ╱─────╲
                ╱Integ  ╲            ← 15% (Medium speed)
               ╱─────────╲
              ╱Application╲           ← 30% (Fast with mocks)
             ╱─────────────╲
            ╱    Domain     ╲        ← 50% (Ultra-fast, pure)
           ╱─────────────────╲
```

### Test Performance Goals

| Layer | Test Type | Current | Target | Speedup |
|-------|-----------|---------|--------|---------|
| Domain | Unit (pure) | N/A | < 1ms | N/A |
| Domain | Property | 2-3s | < 100ms | 20-30x |
| Application | Unit | N/A | < 5ms | N/A |
| Infrastructure | Integration | 15-25ms | 10-20ms | 1.5x |
| Total Suite | All | 30-60s | 5-10s | 5-10x |

### TDD Workflow (MANDATORY)

```bash
# 1. Write test FIRST (red)
vim internal/domain/app/app_structure_test.go

# 2. Implement to make test pass (green)
vim internal/domain/app/app_structure.go

# 3. Refactor if needed

# 4. Verify
go test ./internal/domain/app -v

# Repeat for every function/method
```

---

## Commit Strategy

### Git Workflow

```bash
# Working branch
git checkout -b feature/clean-architecture-full-rewrite

# Frequent commits (end of each day or logical milestone)
git add .
git commit -m "feat(domain): add value objects with validation"
git commit -m "feat(domain): extract AppStructure entity"
git commit -m "feat(infrastructure): implement filesystem adapter"
# etc.

# Final merge to main after Phase 5 complete
```

### Commit Messages (Conventional Commits)

**Phase 1-2:**
```
feat(domain): add value objects with comprehensive validation
feat(domain): extract AppStructure entity with pure logic
feat(domain): define repository interfaces (ports)
feat(infrastructure): implement filesystem adapter
feat(infrastructure): implement in-memory adapter for fast tests
feat(infrastructure): implement dry-run adapter
```

**Phase 3:**
```
feat(application): add use cases layer with DTOs
test(application): add use case tests with in-memory adapters
```

**Phase 4:**
```
refactor(cli): add dependency injection container
refactor(cli): migrate install command to use CreateAppUseCase
refactor(cli): migrate all CLI commands to Clean Architecture
```

**Phase 5:**
```
chore: remove legacy appdir package
chore: remove legacy database/web/supervisor/runtime packages
docs: update architecture documentation
chore: finalize Clean Architecture migration
```

---

## Current Status Summary

### ✅ Completed (Day 1)

1. **Domain package structure** - All folders created
2. **Value objects** - 5 VOs implemented & tested
   - AppName
   - DirectoryPath
   - FilePath
   - Port
   - DatabaseName
3. **Value object tests** - 100% coverage, all passing in 0.003s

### 🚧 Next Steps (Continue Day 1)

1. **Extract AppStructure entity** from `internal/appdir/manager.go`
   - Copy path calculation methods (lines 35-66)
   - Make all methods pure (no I/O)
   - Write comprehensive tests
   - Ensure < 1ms per test

2. **Define app.Repository interface**
   - Create `internal/domain/app/repository.go`
   - Define all methods (Create, Exists, Remove, etc.)

3. **Continue with remaining entities**
   - Database, Service, Runtime, Web

### 📊 Metrics Tracking

| Metric | Baseline | Current | Target |
|--------|----------|---------|--------|
| Test Coverage | 46.2% | 46.2% | 75%+ |
| Domain Coverage | 0% | ~60% | 100% |
| Test Speed | 30-60s | 30-60s | 5-10s |
| Property Tests | 2-3s | 2-3s | < 100ms |

---

## Key Principles to Remember

### Clean Architecture Rules

1. **Dependency Rule**: Dependencies point INWARD only
2. **Domain Independence**: Domain has ZERO dependencies on outer layers
3. **Pure Domain**: No `os`, `exec`, `io`, `net` in domain layer
4. **Interface Segregation**: Small, focused interfaces
5. **Dependency Injection**: Inject adapters via constructors

### TDD Discipline

1. **Red-Green-Refactor**: Test first, always
2. **100% Domain Coverage**: Every domain method tested
3. **Fast Tests**: Domain tests must be < 1ms
4. **Property-Based Testing**: Use rapid for comprehensive testing
5. **Contract Tests**: Verify all adapters comply with interfaces

### Go Best Practices

1. **Accept interfaces, return structs**
2. **Keep interfaces small** (1-5 methods)
3. **Use value objects** for validation
4. **Immutability** for domain entities
5. **Clear naming** (no abbreviations)

---

## References & Resources

### Project Files Created

- `/internal/domain/common/value_objects.go` - Value objects with validation
- `/internal/domain/common/value_objects_test.go` - Comprehensive value object tests

### Files To Extract From

- `/internal/appdir/manager.go` - Lines 35-66 (path calculations)
- `/internal/database/mysql.go` - Lines 229-231 (database naming)
- `/internal/web/backend.go` - Lines 141-173 (port allocation)
- `/internal/supervisor/service.go` - Lines 26-53 (service config)

### Key Documentation

- Original analysis: Full 12,000 word architectural analysis from TDD orchestrator
- Test results: All value object tests passing in 0.003s
- Current TODO list: 14 items tracking progress through 5 phases

---

## Timeline & Milestones

### Week 1 (Days 1-5)

- **Day 1 (Today)**: ✅ Value objects + Start AppStructure entity
- **Day 2**: Complete all domain entities + repository interfaces
- **Day 3**: Implement all 3 adapters (filesystem, memory, dry-run)
- **Day 4**: Start use cases layer
- **Day 5**: Complete use cases + DTOs

### Week 2 (Days 6-10)

- **Day 6-7**: CLI integration with DI container
- **Day 8**: Refactor all CLI commands
- **Day 9**: Remove old code + migrate tests
- **Day 10**: Documentation + final polish

### Buffer (Days 11-12)

- **As needed**: Handle any overruns or unexpected complexity
- **Polish**: Additional property tests, performance optimization

---

## Success Criteria (Final Checklist)

### Technical Metrics

- [ ] Domain layer: 100% test coverage
- [ ] Domain tests: All < 1ms execution time
- [ ] Property tests: < 100ms for 100 iterations
- [ ] Overall coverage: 75%+
- [ ] Test suite: < 5 seconds total
- [ ] Zero `os`/`exec`/`io` imports in domain layer

### Architectural Quality

- [ ] Clean separation of concerns
- [ ] All adapters implement repository interfaces
- [ ] Dependency injection throughout
- [ ] No circular dependencies
- [ ] Contract tests verify adapter compliance

### Functional Quality

- [ ] All CLI commands work identically
- [ ] All flags work as before
- [ ] Integration tests pass
- [ ] Property tests pass with 1000+ checks
- [ ] No breaking changes to user interface

### Documentation

- [ ] Architecture diagrams updated
- [ ] Testing guide updated
- [ ] README with setup instructions
- [ ] Code examples in documentation
- [ ] Migration guide (if needed for external users)

---

**Last Updated:** 2025-11-02 23:42 UTC
**Current Phase:** Phase 1-2 (Foundation) - Day 1
**Next Action:** Extract AppStructure entity from appdir/manager.go
