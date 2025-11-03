# Planning Board - Uberman Clean Architecture

**Last Updated**: 2025-11-03
**Cycle**: Not started (reviewing pitch for betting table)
**WIP Limits**: Initiative: 1/1 | Epic: 1/2 | Story: 2/4 | Task: 1/1 per person
**NO TIME ESTIMATES** - Tasks scoped to 40-minute sessions

---

## 🎯 Initiative: Clean Architecture from Scratch (WIP: 0/1)

**Cynefin**: COMPLEX (unknown unknowns in domain boundaries)
**Strategy**: Probe-Sense-Respond (experiment, observe, adapt)
**Priority**: P0
**Appetite**: 6 weeks (big batch)
**Status**: Draft pitch - awaiting betting table decision

**Pitch**: See plans/pitches/2025-11-03-rebuild-install-clean-architecture.md

**Circuit Breaker**: Ship partial or kill if not done in 6 weeks

---

## 📊 Epic: Installation Workflow Implementation (WIP: 0/2)

**EARS Requirements**: See PRD.md FR-001 through FR-010
**Bounded Context**: App Installation (directory: `appinstallation`)
**Status**: Not started

### Story: Delete Existing Code (WIP: 0/1)

**EARS**: When refactoring begins, then the system shall delete all existing internal/ and apps/ code.

#### Tasks:
- [x] Delete internal/ directory
- [x] Delete apps/ directory
- [ ] Create new directory: internal/appinstallation/domain/
- [ ] Create new directory: internal/appinstallation/application/
- [ ] Create new directory: internal/appinstallation/infrastructure/
- [ ] Commit: "refactor: delete existing code for clean slate approach"

### Story: Domain Value Objects (WIP: 0/2)

**EARS**: Domain value objects shall validate input and be immutable.

**Dependencies**: Delete Existing Code

#### Tasks (40-minute sessions):
- [ ] Create appinstallation/domain/valueobjects/appname.go with validation (TDD)
- [ ] Create appinstallation/domain/valueobjects/appname_test.go
- [ ] Create appinstallation/domain/valueobjects/port.go with validation (TDD)
- [ ] Create appinstallation/domain/valueobjects/port_test.go
- [ ] Create appinstallation/domain/valueobjects/databasename.go with validation (TDD)
- [ ] Create appinstallation/domain/valueobjects/databasename_test.go
- [ ] Create appinstallation/domain/valueobjects/directorypath.go with validation (TDD)
- [ ] Create appinstallation/domain/valueobjects/directorypath_test.go

### Story: Installation Workflow Aggregate (WIP: 0/2)

**EARS**: Installation aggregate shall model workflow states and transitions.

**Dependencies**: Domain Value Objects

**TBD**: Explicit state machine or simpler approach? (Decide during week 1-2)

#### Tasks (40-minute sessions):
- [ ] Define workflow states enum (NotStarted → Installed)
- [ ] Create Installation aggregate with state field
- [ ] Implement LoadManifest command (NotStarted → ManifestLoaded)
- [ ] Write tests for LoadManifest command (TDD)
- [ ] Implement ValidatePrerequisites command (ManifestLoaded → PrerequisitesValidated)
- [ ] Write tests for ValidatePrerequisites (TDD)
- [ ] Implement ProvisionDirectories command
- [ ] Write tests for ProvisionDirectories (TDD)
- [ ] Implement ProvisionDatabase command
- [ ] Write tests for ProvisionDatabase (TDD)
- [ ] Implement ProvisionWebBackend command
- [ ] Write tests for ProvisionWebBackend (TDD)
- [ ] Implement ProvisionService command
- [ ] Write tests for ProvisionService (TDD)
- [ ] Implement SetRuntimeVersion command
- [ ] Write tests for SetRuntimeVersion (TDD)
- [ ] Implement FinalizeInstallation command
- [ ] Write tests for FinalizeInstallation (TDD)

### Story: Manifest Aggregate (WIP: 0/2)

**EARS**: Manifest shall be immutable once loaded and validated.

**Dependencies**: Domain Value Objects

**TBD**: TOML parsing in domain (pure) or infrastructure? (Decide during week 1-2)

#### Tasks (40-minute sessions):
- [ ] Create Manifest aggregate struct (immutable fields)
- [ ] Write tests for Manifest creation (TDD)
- [ ] Implement TOML parsing logic
- [ ] Write tests for parsing (TDD)
- [ ] Implement validation rules for manifest fields
- [ ] Write tests for validation (TDD)

### Story: Repository Interfaces (Ports) (WIP: 0/1)

**EARS**: Domain shall define interfaces, infrastructure shall implement.

**Dependencies**: Installation Workflow Aggregate, Manifest Aggregate

#### Tasks (40-minute sessions):
- [ ] Define Installation repository interface
- [ ] Define Manifest repository interface
- [ ] Define Provisioner interface (database, web, services, runtime)
- [ ] Document interface contracts with examples

### Story: In-Memory Adapters (WIP: 0/2)

**EARS**: Where adapter is in-memory, then tests shall run in < 1ms.

**Dependencies**: Repository Interfaces

#### Tasks (40-minute sessions):
- [ ] Implement in-memory Installation repository
- [ ] Write contract tests for Installation repository
- [ ] Implement in-memory Manifest repository
- [ ] Write contract tests for Manifest repository
- [ ] Implement in-memory Provisioner (records calls, no-op)
- [ ] Write contract tests for Provisioner

### Story: Application Layer (Optional ~) (WIP: 0/2)

**EARS**: When use case executes, then it shall orchestrate domain workflow.

**Dependencies**: In-Memory Adapters

**Note**: Can be cut if time runs short - CLI can call domain directly

#### Tasks (40-minute sessions):
- [ ] Create InstallAppUseCase struct
- [ ] Implement orchestration logic (call domain commands in sequence)
- [ ] Write tests using in-memory adapters (TDD)
- [ ] Verify all workflow states covered in tests
- [ ] Verify all domain events emitted

### Story: Filesystem Adapters (WIP: 0/2)

**EARS**: Where adapter is filesystem, then it shall execute real infrastructure operations.

**Dependencies**: In-Memory Adapters (for contract test comparison)

#### Tasks (40-minute sessions):
- [ ] Implement filesystem Installation repository
- [ ] Write integration tests
- [ ] Implement filesystem Manifest repository (loads from disk)
- [ ] Write integration tests
- [ ] Implement real Provisioner (os.MkdirAll, exec.Command for uberspace)
- [ ] Write integration tests
- [ ] Verify contract tests pass for filesystem adapters

### Story: CLI Integration (WIP: 0/2)

**EARS**: When user runs install command, then CLI shall invoke use case with filesystem adapters.

**Dependencies**: Filesystem Adapters

#### Tasks (40-minute sessions):
- [ ] Refactor cmd/uberman/install.go to wire up new architecture
- [ ] Implement dependency injection
- [ ] Add --dry-run mode support (use dry-run adapter)
- [ ] Add --verbose mode support (structured logging)
- [ ] Write end-to-end CLI test
- [ ] Verify `uberman install wordpress` works identically

---

## 🔧 Technical Debt (15% capacity reserved)

**WIP**: 0/1

- [ ] Add golangci-lint configuration
- [ ] Set up CI with quality gates (go test, golangci-lint, coverage)
- [ ] Fix any existing linting issues in cmd/

---

## 🚫 Blocked / Dependencies

**None currently**

---

## ✅ Completed

**Story: Delete Existing Code** (2 tasks completed):
- [x] Delete internal/ directory - 2025-11-03
- [x] Delete apps/ directory - 2025-11-03

---

## ❌ Out of Scope (Strictly Forbidden to Plan)

**Do NOT create tasks for these** - see IDEAS.md only:
- upgrade, backup, restore, deploy, list, status commands
- Multi-app orchestration
- App dependencies
- Web UI
- API
- Remote manifests
- Any "future enhancement" ideas

---

## Five Thieves Mitigation

**1. Too Much WIP**: Strict WIP limits enforced (see top)
**2. Unknown Dependencies**: Explicit dependency tracking in story headers
**3. Unplanned Work**: No emergency lane (no production system yet)
**4. Conflicting Priorities**: Single initiative focus (P0)
**5. Neglected Work**: 15% capacity reserved for technical debt

---

## Probe-Sense-Respond Strategy

After completing first 2-3 stories (value objects + workflow aggregate):

**Probe**: Implemented workflow modeling with state machine
**Sense**:
- Does it feel natural to work with?
- Are tests fast (< 1ms)?
- Is domain pure (no I/O imports)?
- Is state machine too complex?

**Respond**:
- If state machine too complex → simplify to procedural workflow (still pure domain)
- If tests slow → investigate why domain depends on I/O
- If feeling good → continue with current approach

**Review Point**: End of Week 2 (before committing to filesystem adapters)

---

## Shape Up Integration

**Cycle Status**: Draft pitch, awaiting betting table decision
**Appetite**: 6 weeks (big batch)
**Circuit Breaker**: Active at end of week 6
**Hill Chart**: Will be created when cycle starts (see plans/cycles/)

**Betting Table**: Should happen by 2025-11-04 (review pitch, decide to bet or not)

---

**Next Update**: After betting table decision
