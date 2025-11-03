# Product Requirements Document
# Uberman CLI - Clean Architecture Refactoring

## Objective

Rebuild Uberman install command from scratch using Clean Architecture with workflow modeling to achieve testability, maintainability, and extensibility.

**Success Metrics**:
- Domain layer: 100% pure (zero infrastructure imports)
- Domain tests: < 1ms per test
- Test coverage: ≥ 75%
- Install command works identically to current version

## Scope: MVP - Install Command ONLY

**IN SCOPE**:
- `uberman install <app>` command
- Manifest loading and validation
- Installation workflow modeling
- Directory structure provisioning
- Database provisioning (MySQL)
- Web backend configuration
- Supervisord service setup
- Runtime version management
- Dry-run mode support
- Clean Architecture from scratch

**OUT OF SCOPE** (see IDEAS.md):
- upgrade, backup, restore, deploy, list, status commands
- Multi-app orchestration
- App dependencies
- All other features

## Bounded Context

**Name**: "App Installation" (Go directory: `appinstallation`)

**Core Focus**: Model the installation process as domain workflow, not just data structures.

**Ubiquitous Language** (see UBIQUITOUS_LANGUAGE.md):
- **Installation**: The workflow/process of setting up an app
- **App**: Installable software package (WordPress, Ghost, Nextcloud)
- **Manifest**: TOML definition of app requirements
- **Provisioning**: Creating infrastructure (directories, database, web, services)
- **Instance**: Installed app at ~/apps/<name>

## Installation Workflow (Domain Model)

### States
```
NotStarted
  → ManifestLoaded
  → PrerequisitesValidated
  → ProvisioningInProgress
  → Provisioned
  → Configured
  → Installed
  → (Failed - error state)
```

### Commands (Drive state transitions)
- LoadManifest(appName)
- ValidatePrerequisites()
- ProvisionDirectories()
- ProvisionDatabase()
- ProvisionWebBackend()
- ProvisionService()
- SetRuntimeVersion()
- FinalizeInstallation()

### Domain Events (Outcomes)
- ManifestLoaded
- PrerequisitesValidated
- DirectoriesProvisioned
- DatabaseProvisioned
- WebBackendProvisioned
- ServiceProvisioned
- RuntimeVersionSet
- InstallationCompleted
- InstallationFailed(error)

### Aggregates
- **Installation** (root): Coordinates workflow, maintains state
- **Manifest**: App definition (immutable once loaded)
- **AppInstance**: Installed app representation

## Requirements (EARS Format)

### Epic: Installation Workflow

#### FR-001: Start Installation
When user runs `uberman install <appname>`, then the system shall create Installation aggregate in NotStarted state.

#### FR-002: Load Manifest
When Installation receives LoadManifest command, then the system shall load manifest from apps/<appname>/app.toml and transition to ManifestLoaded state.

If manifest not found, then the system shall transition to Failed state with error "manifest not found for <appname>".

#### FR-003: Validate Prerequisites
When Installation is in ManifestLoaded state, then the system shall validate app is not already installed at ~/apps/<appname>.

If app already exists, then the system shall transition to Failed state with error "app <appname> already installed at ~/apps/<appname>".

When validation passes, then the system shall transition to PrerequisitesValidated state and emit PrerequisitesValidated event.

#### FR-004: Provision Directories
When Installation is in PrerequisitesValidated state, then the system shall create directory structure with subdirectories: app/, data/, logs/, backups/, tmp/.

When directories created successfully, then the system shall emit DirectoriesProvisioned event and transition to ProvisioningInProgress state.

#### FR-005: Provision Database
Where manifest specifies database.required = true, then the system shall provision MySQL database with name ${USER}_<appname>.

If database creation fails, then the system shall transition to Failed state with error details.

When database provisioned successfully, then the system shall emit DatabaseProvisioned event.

#### FR-006: Provision Web Backend
Where manifest specifies web.backend, then the system shall configure Uberspace web backend (apache or http).

Where backend = "http", then the system shall allocate available port in range 1024-65535.

If port allocation fails, then the system shall transition to Failed state with error "no available ports".

When backend configured successfully, then the system shall emit WebBackendProvisioned event.

#### FR-007: Provision Service
Where manifest defines services array, then the system shall create supervisord configuration in ~/etc/services.d/<appname>.ini.

When service configured successfully, then the system shall emit ServiceProvisioned event.

#### FR-008: Set Runtime Version
Where manifest specifies runtime.version, then the system shall set runtime version via uberspace tools.

When version set successfully, then the system shall emit RuntimeVersionSet event.

#### FR-009: Finalize Installation
When all required provisioning steps complete, then the system shall transition to Installed state and emit InstallationCompleted event.

Then the system shall output "App <appname> installed successfully at ~/apps/<appname>".

#### FR-010: Dry-Run Mode
Where --dry-run flag is set, then the system shall execute workflow without side effects and log all operations.

Then the system shall transition through states but skip infrastructure commands.

### Non-Functional Requirements

#### NFR-001: Performance
The installation workflow shall complete in < 30 seconds (excluding app download time).

#### NFR-002: Testability
The Installation aggregate shall be testable without filesystem, database, or external commands.

#### NFR-003: Pure Domain
The domain layer shall have zero imports from os, exec, io, net, syscall packages.

The domain layer shall validate with: `go list -f '{{.Imports}}' ./internal/appinstallation/domain/...`

#### NFR-004: Test Speed
Domain tests shall execute in < 1ms per test.

Application layer tests (with in-memory adapters) shall execute in < 5ms per test.

## Architecture Constraints

- Clean Architecture layers (Presentation → Application → Domain → Infrastructure)
- Single bounded context: "App Installation"
- Dependency injection throughout
- Repository pattern (ports + adapters)
- TDD approach (test-first development)

## Approach: Clean Slate

**Delete**:
- All existing `internal/` code
- All existing `apps/` manifests
- Start fresh with correct architecture

**Keep**:
- `cmd/` (CLI entry point - will refactor)
- Tests (as reference for behavior - will rewrite)

**Build Order** (Probe-Sense-Respond):
1. Domain layer (value objects + workflow aggregate)
2. In-memory adapter (fast tests)
3. Application layer (use case - if time permits)
4. Filesystem adapter (real implementation)
5. CLI integration

## Assumptions

**ASSUMPTION 1** (p=90%, Impact: Medium, Notify ✅)
Installation is the aggregate root that models the workflow, not App or Manifest.

**ASSUMPTION 2** (p=88%, Impact: Low, Notify ✅)
Domain events (ManifestLoaded, DatabaseProvisioned, etc.) provide sufficient observability.

**ASSUMPTION 3** (p=85%, Impact: Medium, TBD ❓)
Workflow should model explicit states + commands + events (vs simpler procedural approach).

**Decision**: Determine during week 1-2 of implementation.

## Out of Scope

**Strictly forbidden** to plan or implement:
- Other commands (upgrade, backup, restore, deploy, list, status)
- Multi-app features
- App dependencies
- Web UI
- API
- Remote manifests

See IDEAS.md for rough idea collection (not planned work).
