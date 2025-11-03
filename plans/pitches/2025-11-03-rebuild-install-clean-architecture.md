# Pitch: Rebuild Install Command with Clean Architecture

## Problem

The `uberman install` command mixes business logic with infrastructure concerns, making testing slow and changes risky:

- **Slow tests**: Domain logic tests require real filesystem (15-25ms per test), taking 30-60s for full suite
- **Mixed concerns**: App structure logic (`internal/appdir/manager.go`) directly calls `os.MkdirAll()`, `os.Symlink()`
- **Tight coupling**: Database creation directly executes mysql commands, web backend directly calls uberspace CLI
- **Cannot test in isolation**: Business rules require filesystem, database, external commands to test
- **No extensibility**: Adding new delivery mechanism (web UI, API) would require duplicating logic
- **Property tests too slow**: 2-3s each, limiting their use for comprehensive testing

**Impact**: Slow feedback loop, risky changes, limited test coverage, cannot extend to new interfaces.

## Appetite

**6 weeks (big batch)** - one developer, full focus on MVP (install command only).

## Solution

Delete existing `internal/` code and rebuild from scratch using Clean Architecture with workflow modeling, implementing single bounded context: "App Installation".

### Breadboard (Architecture Flow)

```
[CLI Layer]
  ├── install command
  ↓
[Application Layer]
  ├── InstallAppUseCase (orchestration)
  ↓
[Domain Layer - PURE]
  ├── Installation aggregate (workflow state machine)
  │   States: NotStarted → ManifestLoaded → Validated →
  │           ProvisioningInProgress → Configured → Installed
  │   Commands: LoadManifest, ValidatePrerequisites, ProvisionInfrastructure
  │   Events: ManifestLoaded, InfrastructureProvisioned, InstallationCompleted
  ├── Manifest aggregate (app definition)
  ├── Value Objects (AppName, Port, DatabaseName, etc.)
  ↓ (depends on interfaces)
[Ports - Repository Interfaces]
  ├── InstallationRepository
  ├── ManifestRepository
  ├── Provisioner (database, web, services, runtime)
  ↑ (implemented by)
[Infrastructure Layer]
  ├── Filesystem adapters (real I/O)
  ├── In-memory adapters (fast tests)
  └── Dry-run adapters (preview mode)
```

### Fat Marker Sketch (Directory Structure)

```
internal/
└── appinstallation/              # Bounded Context: App Installation
    ├── domain/                   # Domain Layer (100% pure)
    │   ├── workflow/
    │   │   ├── installation.go   # Aggregate root: Installation
    │   │   ├── states.go         # State machine
    │   │   ├── commands.go       # Command definitions
    │   │   └── events.go         # Domain events
    │   ├── manifest/
    │   │   ├── manifest.go       # Aggregate: Manifest
    │   │   └── parser.go         # TOML parsing (pure)
    │   ├── valueobjects/
    │   │   ├── appname.go
    │   │   ├── port.go
    │   │   ├── databasename.go
    │   │   └── directorypath.go
    │   └── ports/                # All interfaces
    │       ├── installation_repository.go
    │       ├── manifest_repository.go
    │       └── provisioner.go
    │
    ├── application/              # Application Layer
    │   └── installapp/
    │       ├── usecase.go        # Orchestrates workflow
    │       └── dto.go            # Data transfer objects
    │
    └── infrastructure/           # Infrastructure Layer
        ├── memory/               # In-memory (fast tests)
        │   ├── installation_repository.go
        │   ├── manifest_repository.go
        │   └── provisioner.go
        ├── filesystem/           # Real implementation
        │   ├── installation_repository.go
        │   ├── manifest_repository.go
        │   └── provisioner.go
        └── dryrun/               # Preview mode
            └── provisioner.go
```

### Implementation Approach

**Week 1-2: Domain Layer Foundation**
- Value objects with validation (AppName, Port, DatabaseName, DirectoryPath)
- Installation aggregate with workflow state machine
- Manifest aggregate with TOML parsing (pure)
- Repository interfaces (ports)
- Domain tests: all < 1ms per test

**Week 3: Infrastructure - In-Memory Adapters**
- In-memory implementations of all repositories
- In-memory provisioner (records calls, no side effects)
- Contract tests verifying interface compliance
- Tests: < 1ms with in-memory adapters

**Week 4: Infrastructure - Filesystem Adapters**
- Filesystem implementations (real os.MkdirAll, exec.Command, etc.)
- Integration tests with real filesystem
- Verify Uberspace command integration
- Tests: Integration tests 10-20ms each

**Week 5: Application Layer**
- InstallAppUseCase orchestrating workflow
- Dependency injection setup
- Use case tests with in-memory adapters
- Tests: < 5ms per use case test

**Week 6: CLI Integration**
- Refactor `cmd/uberman/install.go` to use InstallAppUseCase
- Wire up filesystem adapters for production
- Wire up dry-run adapters for --dry-run flag
- End-to-end CLI tests
- Verify install command works identically to old version

## Rabbit Holes

**Circular Dependencies:**
If domain entities have circular dependencies (app depends on database, database depends on app), will need to untangle with:
- Shared valueobjects package
- Dependency inversion (interfaces)
- **Mitigation**: Budget 2-3 days if discovered, scope hammer if exceeds

**Test Migration Complexity:**
Existing tests in old `internal/` will all break. Plan:
- Don't migrate old tests - delete them
- Write new tests for new code (TDD from scratch)
- Verify end-to-end CLI behavior matches
- **Risk**: If old tests had critical coverage we don't reproduce
- **Mitigation**: Review old tests for edge cases before deleting

**Workflow State Machine Complexity:**
Modeling installation as explicit state machine might be over-engineered.
- **De-risking decision**: Mark as TBD during week 1
- If too complex by week 2, switch to simpler approach
- **Fallback**: Procedural workflow without explicit states (still pure domain)

**TOML Parsing in Domain:**
Keeping manifest parsing in domain layer (pure) vs infrastructure.
- **De-risking decision**: Mark as TBD during week 1
- Domain should validate manifest structure, infrastructure should handle file I/O
- **Fallback**: Move parser to infrastructure if purity is violated

## No-Gos

This pitch explicitly excludes (defer to future cycles):

- ❌ **Other commands** - NO upgrade, backup, restore, deploy, list, status
- ❌ **Application layer** - If time is tight, CLI can call domain directly (cut orchestration)
- ❌ **All three adapters** - Focus on filesystem + in-memory only, cut dry-run if needed
- ❌ **100% test coverage** - Accept 75%+ for this cycle, improve later
- ❌ **Property-based tests for everything** - Only for critical value objects
- ❌ **Full test suite speedup** - Focus on domain tests only, integration tests can stay slow
- ❌ **Documentation updates** - Defer to cooldown
- ❌ **Perfect error messages** - Good enough error handling
- ❌ **Performance optimization** - Fast enough is sufficient

## Success Criteria

**Must-haves:**
- ✓ Domain layer: 100% pure (zero `os`, `exec`, `io`, `net` imports)
- ✓ Value objects with validation (AppName, Port, DatabaseName, DirectoryPath)
- ✓ Installation aggregate modeling workflow (TBD: state machine or simpler)
- ✓ Manifest aggregate with parsing logic
- ✓ Repository interfaces defined (ports)
- ✓ Filesystem adapter implements all interfaces
- ✓ In-memory adapter implements all interfaces
- ✓ Domain tests: all < 1ms
- ✓ Install command refactored to use new architecture
- ✓ `uberman install wordpress` works identically to old version

**Nice-to-haves (~):**
- ~ Application layer (InstallAppUseCase) - cut if time tight, CLI can call domain directly
- ~ Dry-run adapter - cut if time tight, can add next cycle
- ~ Property tests for all value objects - just AppName is enough
- ~ Integration test suite - manual testing acceptable for week 6
- ~ Detailed error messages - basic errors sufficient

**TBD (To Be Determined during shaping/building):**
- ❓ **Q1**: Is Installation the aggregate root? Or is App the root?
- ❓ **Q2**: Use explicit state machine for workflow? Or simpler procedural approach?

## Circuit Breaker

If we can't complete must-haves in 6 weeks:

**Ship partial (preferred):**
- Domain layer + in-memory adapter + tests (proves pattern works)
- Old install command still works (no regression)
- New architecture validated with fast tests
- Can extend in next 6-week cycle with filesystem adapter + CLI integration

**Kill completely (if fundamentally broken):**
- If domain layer cannot be made pure (requires I/O)
- If tests don't get faster (still requires filesystem)
- If circular dependencies cannot be resolved
- If workflow modeling is too complex to understand
- **Action**: Roll back all changes, delete new code, restore from git, stick with current architecture

**Learn (either outcome):**
- What took longer than expected?
- Was appetite too small? (Should have been 12 weeks split into 2 cycles?)
- Was scope too large? (Should have cut application layer upfront?)
- Were rabbit holes bigger than expected? (State machine complexity? Test migration?)
- Should we reshape differently for next cycle? (Simpler workflow? Less ambitious scope?)

---

**Pitch Status:** Draft (Ready for betting table review)
**Betting Table:** 2025-11-03 (Cooldown week 1)
**Proposed Start:** 2025-11-04 (Immediate - Cycle 01)
**Proposed End:** 2025-12-15 (6 weeks from start)
