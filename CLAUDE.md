# CLAUDE.md

Guidance for Claude Code when working with this repository.

## Project Context

**Uberman** is a Go CLI tool for reproducible app installation on Uberspace hosting (uberspace.de).

**Current Status**: Clean Architecture refactoring from scratch
**MVP Scope**: `uberman install <app>` command ONLY
**Approach**: Delete existing code, rebuild with correct bounded context

## Bounded Context (DDD)

**Single Context**: "App Installation" (Go directory: `appinstallation`)

**Focus**: Model installation as workflow, not just data structures.

**Ubiquitous Language**: See UBIQUITOUS_LANGUAGE.md

**Key Concepts**:
- **Installation**: Workflow/process aggregate (root)
- **Manifest**: App definition aggregate
- **Provisioning**: Infrastructure creation (directories, database, web, services)
- **Instance**: Installed app at ~/apps/<name>

## Architecture (Clean Architecture)

### Single Bounded Context Structure

```
internal/appinstallation/
├── domain/               # 100% pure (no os, exec, io, net)
│   ├── workflow/         # Installation aggregate
│   ├── manifest/         # Manifest aggregate
│   ├── valueobjects/     # AppName, Port, DatabaseName, etc.
│   └── ports/            # Repository interfaces
├── application/          # Use cases (optional - can be cut)
│   └── installapp/
└── infrastructure/       # Adapters (filesystem, memory, dryrun)
    ├── filesystem/
    ├── memory/
    └── dryrun/
```

### Dependency Rules (ENFORCE)

1. **Dependencies point INWARD only**
2. **Domain: ZERO outward dependencies**
3. **Domain: NO imports from** os, exec, io, net, syscall
4. **Infrastructure implements domain interfaces** (ports)

**Validation**:
```bash
go list -f '{{.ImportPath}}: {{.Imports}}' ./internal/appinstallation/domain/...
```

Must see ZERO forbidden packages.

### Workflow Modeling

Installation aggregate coordinates workflow via state machine:
```
NotStarted → ManifestLoaded → PrerequisitesValidated →
ProvisioningInProgress → Configured → Installed
```

**Commands**: LoadManifest, ValidatePrerequisites, ProvisionDatabase, etc.
**Events**: ManifestLoaded, DatabaseProvisioned, InstallationCompleted, etc.

**TBD**: Explicit state machine or simpler approach? (Decide during week 1-2)

See ARCHITECTURE.md for details.

## Shape Up Methodology (PERMANENT)

### Core Principles

- **Fixed time, variable scope** (NOT fixed scope, variable time)
- **6-week cycles** (big batch) or 1-2 weeks (small batch)
- **2-week cooldown** after each cycle
- **Circuit breaker**: Kill projects that don't ship in one cycle
- **No backlogs**: Ideas tracked in IDEAS.md, resurface naturally
- **Appetite over estimates**: Set time budget first, design to fit

### Current Cycle

**Status**: Draft pitch - awaiting betting table decision
**Pitch**: plans/pitches/2025-11-03-rebuild-install-clean-architecture.md
**Appetite**: 6 weeks (big batch)
**Scope**: Install command only (MVP)
**Circuit Breaker**: Ship partial or kill if not done in 6 weeks

### Appetite Constraints

**Small batch**: 1-2 weeks (tweaks, bug fixes, small refactoring)
**Big batch**: 6 weeks (major features, large refactoring)

**NOT an estimate** - it's a creative constraint forcing hard choices.

### Scope Hammering

Forcefully question every feature to fit time box.

**When**: During shaping (primary) and building (when hitting constraints)
**How**: Must-haves vs Nice-to-haves (~), question everything
**Goal**: Good enough > perfect

### Circuit Breaker

Hard deadline at end of cycle. No extensions by default.

**Options**:
1. **Ship what's done** (most common) - use scope hammering
2. **Kill and reshape** - circuit breaker activates, learn from failure
3. **Extend** (RARE) - only if all work is downhill + nearly complete

### Hill Charts

Track progress weekly: Uphill (figuring out) vs Downhill (executing known work)

**When to worry**: Scope stuck uphill for 2+ weeks

### Betting Table

Solo decision point during cooldown: review pitches, decide what to bet on next cycle.

### Cycles and Cooldown

**6-week cycle**: Full focus on committed work
**2-week cooldown**: Bug fixes, exploration, shaping, betting table

### Plans Directory

```
plans/
├── templates/      # Pitch, hill chart, betting table templates
├── pitches/        # Shaped work ready for betting
├── cycles/         # Active cycle tracking with hill charts
└── cooldown/       # Betting decisions, exploration notes
```

See plans/README.md for Shape Up guide.

## Session Management (PERMANENT)

### Time Boxing

- **Maximum session**: 40 minutes
- **Break**: 10 minutes between sessions
- **Daily max**: 6 sessions
- **Types**: Planning | Execution | Review

### Kill Criteria (Check Every Session)

Stop immediately if:
1. Same error repeated 3x despite different approaches
2. Core requirement misunderstood (check EARS in PRD.md)
3. Architecture rule violated (check layer imports)
4. Token cost exceeds $5 for single task
5. No measurable progress in 2 consecutive sessions

### Before ANY Session

1. Read current story from PLANNING.md
2. Read EARS requirements from epic
3. Identify task to work on
4. Check session type (Planning or Execution)
5. Set 40-minute timer

### Planning Session (Read-Only Mode)

- Generate specifications only
- NO code implementation
- Output: Plan document
- Requires approval before execution

### Execution Session

- Follow approved plan
- One task only per session
- TDD: Test first, then implementation
- Update PLANNING.md with progress

### End of Session

- Update task status in PLANNING.md
- Note blockers
- Commit if complete (conventional commits)
- Update hill chart (if in cycle)

## Decision Framework (PERMANENT)

### Assumption Scoring

- **p < 60%**: Ask user to refine
- **p = 60-90%**: Ask approval if high impact, apply if low impact
- **p > 90%**: Notify user, apply automatically

### Impact Assessment

**Low Impact**:
- < 20 LOC added
- No external dependencies
- Easily removable (single method/function)
- No database schema changes
- No breaking API changes

**High Impact**:
- > 50 LOC added
- New dependencies/libraries
- Database migrations required
- API contract changes
- Cross-cutting concerns (auth, logging, caching)

### Cynefin Classification (for Initiatives)

- **CLEAR**: Best practices exist → Sense-Categorize-Respond
- **COMPLICATED**: Expert analysis needed → Sense-Analyze-Respond
- **COMPLEX**: Unknown unknowns → Probe-Sense-Respond ← **Current refactoring**
- **CHAOTIC**: Crisis → Act-Sense-Respond

**Current Initiative**: Clean Architecture Refactoring = **COMPLEX**
**Strategy**: Probe-Sense-Respond (experiment, observe, adapt)

## Quality Gates (Go-Specific)

Before committing:
- [ ] All tests pass: `go test ./...`
- [ ] No linting errors: `go test -v ./...` (using pre-commit hooks)
- [ ] Test coverage ≥ 75%: `go test -cover ./...`
- [ ] Architecture rules validated: `go list -f '{{.ImportPath}}: {{.Imports}}' ./internal/appinstallation/domain/...`
- [ ] Conventional commit format used

**Test Performance Targets**:
- Domain tests: < 1ms per test (pure, no I/O)
- Application tests: < 5ms per test (in-memory adapters)
- Integration tests: 10-20ms acceptable (testcontainers overhead)

**Testing Tools**:
- **TDD Agent**: Use `tdd-orchestrator` agent for guided test-driven development
- **Property-based**: `pgregory.net/rapid` framework for invariant testing
- **Integration**: `testcontainers-go` for isolated container tests (MySQL, filesystem)
- **Docker**: Required for tests with side effects (database, filesystem operations)

**Testing Commands**:
```bash
# Fast TDD cycle (native)
go test -v ./internal/appinstallation/domain/...

# Property-based tests
go test -v -run Property ./...

# Integration tests (requires Docker)
go test -v -run Integration ./...

# Skip slow tests
go test -short ./...
```

## Platform Constraints (Uberspace)

**Critical**:
- No Docker/containers
- No root/system package managers
- User-space only
- Supervisord for services (not systemd)
- Database naming: `${USER}_<appname>`
- Port binding: `0.0.0.0` or `::` (NOT localhost)

See UBERSPACE_INTEGRATION.md for commands and details.

## Project Structure

### Work Organization

```
.project/                  # If needed for additional tracking
plans/                     # Shape Up pitches, cycles, cooldown
├── templates/
├── pitches/
├── cycles/
└── cooldown/
```

### Status Markers (in PLANNING.md)

- [ ] Todo
- [-] In Progress (currently working on this)
- [R] Review needed
- [x] Complete

### EARS Requirements Format

All epics and stories MUST have requirements in EARS format (see PRD.md):

- **Ubiquitous**: The system shall [always do X]
- **State-Driven**: While [condition], the system shall [do X]
- **Event-Driven**: When [event], the system shall [do X]
- **Optional**: Where [feature enabled], the system shall [do X]
- **Unwanted**: If [error], then the system shall [handle X]

These are specifications to implement, not suggestions.

## References

**Core Documentation**:
- **PRD.md**: Product requirements (MVP: install command only, EARS format)
- **PLANNING.md**: Kanban board (Initiative → Epic → Story → Task)
- **ARCHITECTURE.md**: Detailed Clean Architecture design (includes testing infrastructure)
- **UBIQUITOUS_LANGUAGE.md**: Domain glossary for App Installation context

**Shape Up**:
- **plans/README.md**: Shape Up methodology guide
- **plans/pitches/**: Shaped work ready for betting
- **plans/cycles/**: Active cycle tracking with hill charts
- **plans/cooldown/**: Betting table decisions
- **plans/templates/**: Pitch, hill chart, betting table templates

**Reference**:
- **UBERSPACE_INTEGRATION.md**: Platform-specific commands
- **IDEAS.md**: Rough idea collection (NOT planned work)

**External**:
- Uberspace Manual: https://manual.uberspace.de/
- Uberspace Lab: https://lab.uberspace.de/
- Shape Up (free book): https://basecamp.com/shapeup

## Documentation Status (2025-11-03)

### Trusted (Current Architecture)

**Core** (10 files):
- CLAUDE.md, PRD.md, PLANNING.md, ARCHITECTURE.md, UBIQUITOUS_LANGUAGE.md
- UBERSPACE_INTEGRATION.md, IDEAS.md, CHANGELOG.md

**Shape Up** (6 files):
- plans/README.md (simplified to 118 lines)
- plans/pitches/2025-11-03-rebuild-install-clean-architecture.md
- plans/cooldown/2025-11-03-betting-table.md
- plans/templates/* (3 files: pitch, hill chart, betting table)

**Total Trusted**: 16 files

### Outdated (Marked with Warning Banner)

**User-Facing** (4 files):
- README.md, CONTRIBUTING.md
- docs/INSTALLATION.md, docs/SETUP_VERIFICATION.md

**Note:** These will be updated after Clean Architecture refactoring completes.

### Deleted (2025-11-03 Documentation Cleanup Spike)

**Phase 1 - Testing & Implementation Reports** (10 files, ~3,400 lines):
- Testing docs: TESTING.md, PROPERTY_BASED_TESTING.md, TEST_SUMMARY.md, docs/TESTING.md, docs/DOCKER_TESTING.md, docs/TESTING_CHEATSHEET.md, docs/CONTAINERIZED_TESTING_SUMMARY.md
- Implementation reports: docs/ARCHITECTURE_DIAGRAM.md, docs/IMPLEMENTATION_REPORT.md, PROJECT_SUMMARY.md
- Refactoring plans: CLEAN_ARCHITECTURE_REFACTORING_PLAN.md, CONTAINERIZED_TESTING_README.md

**Phase 2 - YAGNI Process Documentation** (11 files, ~2,600 lines):
- CI/CD docs: docs/CI_CD_SETUP.md, docs/RELEASE_PROCESS.md (consolidated to ARCHITECTURE.md)
- Commit guides: docs/CONVENTIONAL_COMMITS.md, docs/PRE_COMMIT_HOOKS.md
- Requirements docs: docs/EARS_GUIDE.md, docs/REQUIREMENTS_QUALITY.md, docs/REQUIREMENTS_ENGINEERING_AGENT.md
- Requirements templates: templates/requirements/* (3 files)
- User docs: docs/ROADMAP.md, docs/QUICKSTART.md

**Total Deleted**: 21 files, ~6,000 lines

**Rationale**: User has 15 years experience with TDD, testcontainers, property-based testing, CI/CD, semantic-release, conventional commits. YAGNI documentation removed. Essential constraints consolidated in ARCHITECTURE.md. Requirements engineering available via `requirements-engineer` skill and slash commands.

### Cumulative Impact

**Before cleanup**: 46 files, ~14,500 lines
**After cleanup**: 20 files, ~5,000 lines
**Reduction**: 57% fewer files (-26), 66% less content (-9,500 lines)

### Priority Rule

**When conflicts arise:** Trust "Trusted" files over "Outdated" files.

## Key Reminders

### Strictly Forbidden

- ❌ Planning future work beyond current cycle (use IDEAS.md only)
- ❌ Adding time estimates to tasks
- ❌ Creating backlogs
- ❌ Implementing features not in current pitch
- ❌ Extending deadlines without circuit breaker decision

### Always Do

- ✅ Use EARS format for requirements
- ✅ Keep domain layer 100% pure
- ✅ Work in 40-minute sessions
- ✅ Update PLANNING.md with progress
- ✅ Check architecture rules before committing
- ✅ Conventional commits format
- ✅ TDD approach (test first)

### Remember

- **Fixed time, variable scope** - cut features, never extend time
- **Probe-Sense-Respond** - experiment, observe, adapt (Complex domain)
- **Circuit breaker** - ship partial or kill, no shame in learning
- **Good enough > perfect** - scope hammer aggressively
- **Let ideas die** - most ideas in IDEAS.md will never be implemented (intentional)

---

**Last Updated**: 2025-11-03
**Current Phase**: Cooldown (shaping, awaiting betting table)
**Next Action**: Review pitch, decide to bet or not
