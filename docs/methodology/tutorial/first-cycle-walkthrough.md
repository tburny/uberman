# First Cycle Walkthrough Tutorial

Learn the complete Shape Up + session management methodology by following Cycle 01 from start to finish.

**Audience**: New contributors learning the system hands-on.

**What you'll learn**:
- Shaping a pitch during cooldown
- Making a betting table decision
- Starting a cycle with scopes and hill chart
- Running your first session
- Weekly hill chart updates
- Using the scope hammer when needed
- Circuit breaker decisions
- Archiving and reflecting

**Time investment**: ~1 hour reading, lifetime of effective development

---

## Overview: Cycle 01 Story

**Context**: November 2025, existing uberman codebase has mixed concerns. Decision made to rebuild `install` command using Clean Architecture.

**Appetite**: 6 weeks (big batch)

**Outcome**: Shipped partial (4/5 scopes complete, integration tests deferred)

**Follow along**: Real examples from actual Cycle 01 implementation.

---

## Part 1: Shaping (Cooldown Week 1)

**Date**: Monday, November 3, 2025
**Phase**: Cooldown after previous work
**Activity**: Shaping new pitch

### The Problem

Existing code mixes all concerns:
- CLI commands directly call shell commands
- No domain model
- Hard to test (no dependency injection)
- Can't swap implementations (no adapters)

**Pain point**: Adding new apps requires touching multiple layers simultaneously.

### Setting Appetite

**Question**: How much time is this worth?

**Options considered**:
- 1 week (small batch) - Too tight for full refactoring
- 6 weeks (big batch) - Right size for architectural change
- 12 weeks - Too much, risk creep

**Decision**: 6 weeks (big batch)

**Rationale**: Architectural refactoring is complex domain. Need time for:
- Learning Clean Architecture patterns in Go
- Property-based testing experimentation
- Discovery of correct bounded context

### Breadboarding the Solution

**Layers** (Clean Architecture):
1. **Domain**: Pure business logic (Install app installation workflow, Manifest app definition)
2. **Application**: Use cases (optional, may be YAGNI)
3. **Infrastructure**: Adapters (filesystem, in-memory for testing, future: dry-run)
4. **CLI**: User interface (cobra commands)

**Flow**:
```
User → CLI → Application → Domain ← Infrastructure
                              ↓
                         Pure Logic
```

**Key insight**: Domain must be 100% pure (no os, exec, io, net imports).

### Identifying Rabbit Holes

**Potential time sinks**:
1. **Circular dependencies** - Domain depending on infrastructure
   - **No-go**: Enforce via `go list` check in tests
2. **Over-engineering state machine** - Explicit vs implicit
   - **Acceptable risk**: Start simple, add if needed
3. **Test migration complexity** - Moving tests to new structure
   - **Time box**: Rewrite tests, don't try to preserve old ones

### Defining No-Gos

**Out of scope for this cycle**:
- Other commands (`upgrade`, `remove`, `list`) - MVP is `install` only
- 100% test coverage - 75% is sufficient
- Performance optimization - Focus on correctness first
- CLI output beautification - Functional output is enough

### Writing the Pitch

**Created**: `plans/pitches/2025-11-03-rebuild-install-clean-architecture.md`

**Structure**:
```markdown
# Rebuild Install Command with Clean Architecture

## Problem
[Detailed problem statement]

## Appetite
6 weeks (big batch)

## Solution
[Breadboard layers, flow diagram]

## Rabbit Holes
[3 identified risks with mitigations]

## No-Gos
[4 explicit exclusions]
```

**Key sections**:
- **Problem**: Clear pain point (mixed concerns)
- **Appetite**: Explicit time budget (6 weeks)
- **Solution**: Enough detail to bet, not implementation
- **Rabbit holes**: Acknowledge risks
- **No-gos**: Prevent scope creep

**Outcome**: Pitch ready for betting table.

---

## Part 2: Betting (Cooldown Week 2)

**Date**: Friday, November 8, 2025
**Phase**: Still cooldown
**Activity**: Betting table decision

### Review the Pitch

**Self-reflection questions**:
1. Is the problem important enough to spend 6 weeks?
   - **Answer**: Yes - technical debt blocking new features
2. Is the solution approach sound?
   - **Answer**: Yes - Clean Architecture is proven pattern
3. Do I have the knowledge/skills?
   - **Answer**: Mostly - will learn property-based testing
4. Is the appetite realistic?
   - **Answer**: Yes - 6 weeks for 4-5 scopes seems feasible
5. What's the worst case?
   - **Answer**: Kill at circuit breaker, keep learning

### Making the Bet

**Decision**: BET (commit to 6 weeks)

**Rationale documented** in `plans/cooldown/2025-11-03-betting-table.md`:
```markdown
# Betting Table Decision: 2025-11-03

## Pitch
Rebuild Install Command with Clean Architecture

## Decision: BET ✓

## Reasoning
- Problem is important (blocking new features)
- Solution is sound (proven pattern)
- Appetite is realistic (6 weeks for architectural change)
- Worst case: Learn Clean Architecture even if killed

## Concerns
- Property-based testing learning curve
- State machine complexity unknown

## Commitment
- Start: Monday, November 11, 2025
- Circuit breaker: Friday, December 20, 2025
- Will ship partial or kill - no extensions
```

**Outcome**: Committed to Cycle 01 start on Monday.

---

## Part 3: Cycle Start (Week 1, Day 1)

**Date**: Monday, November 11, 2025
**Activity**: Create cycle file, define scopes, initialize hill chart

### Creating Cycle File

**File**: `plans/cycles/cycle-01.md`

**Header**:
```markdown
# Cycle 01: Clean Architecture Refactoring

**Status**: ACTIVE
**Dates**: 2025-11-11 to 2025-12-20 (6 weeks)
**Pitch**: ../pitches/2025-11-03-rebuild-install-clean-architecture.md
**Circuit Breaker**: 2025-12-20
```

### Defining Scopes

**Break pitch into 5 integrated slices**:

1. **Domain Layer Foundation**
   - Value objects: AppName, DatabaseName, Port, DirectoryPath
   - Installation aggregate
   - Manifest aggregate
   - Repository interfaces (ports)

2. **In-Memory Adapter**
   - Prove domain works without I/O
   - Implement repository interfaces
   - Fast test feedback

3. **Filesystem Adapter**
   - Real persistence
   - Directory creation, permissions
   - Database provisioning (via uberspace commands)

4. **CLI Layer**
   - Cobra command structure
   - Flag parsing
   - Wire to use cases/domain

5. **Integration Tests**
   - End-to-end workflow tests
   - Testcontainers for isolation
   - Docker-based filesystem tests

**Rationale**: Each scope is a vertical slice that adds value.

### Initializing Hill Chart

**All scopes start at Position 0** (not started):

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Domain Layer**: ○ Position 0/10 - Not started
**In-Memory Adapter**: ○ Position 0/10 - Not started
**Filesystem Adapter**: ○ Position 0/10 - Not started
**CLI Layer**: ○ Position 0/10 - Not started
**Integration Tests**: ○ Position 0/10 - Not started

### Marking First Scope Active

**Active scope**: Domain Layer Foundation 🎯

**Next Session**: Create AppName value object

**Goal**: Get first meaningful work done today.

### Committing Cycle File

```bash
git add plans/cycles/cycle-01.md
git commit -m "chore: start Cycle 01 - Clean Architecture Refactoring

6-week cycle beginning 2025-11-11.
5 scopes defined, Domain Layer active.
Circuit breaker: 2025-12-20.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
"
```

**Outcome**: Cycle is live, ready for first session.

---

## Part 4: First Session (Week 1, Day 1, 10:00)

**Date**: Monday, November 11, 2025, 10:00 AM
**Activity**: Implement AppName value object using TDD

### Starting the Session

**Command**:
```bash
/session start domain-appname
```

**Output**:
```
✓ Session started: domain-appname
✓ Branch created: session/20251111-domain-appname
✓ Timer started: 00:00

Active session:
- Scope: Domain Layer Foundation
- Task: Create AppName value object
- Started: 2025-11-11 10:00

Working on session/20251111-domain-appname
```

**Current branch**: `session/20251111-domain-appname`
**Status**: Timer running, monitoring active

### TDD Cycle 1: Failing Test

**Create**: `internal/appinstallation/domain/valueobjects/appname_test.go`

```go
package valueobjects_test

import (
	"testing"
	"github.com/yourusername/uberman/internal/appinstallation/domain/valueobjects"
)

func TestAppName_ValidName(t *testing.T) {
	name, err := valueobjects.NewAppName("wordpress")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if name.String() != "wordpress" {
		t.Errorf("expected 'wordpress', got '%s'", name.String())
	}
}
```

**Run test**:
```bash
go test ./internal/appinstallation/domain/valueobjects/
```

**Result**: Compilation error (AppName doesn't exist yet) ✓ RED

**Time elapsed**: 5 minutes

### TDD Cycle 2: Make It Pass

**Create**: `internal/appinstallation/domain/valueobjects/appname.go`

```go
package valueobjects

import "fmt"

type AppName struct {
	value string
}

func NewAppName(name string) (AppName, error) {
	if name == "" {
		return AppName{}, fmt.Errorf("app name cannot be empty")
	}
	// TODO: Add more validation
	return AppName{value: name}, nil
}

func (a AppName) String() string {
	return a.value
}
```

**Run test**:
```bash
go test ./internal/appinstallation/domain/valueobjects/
```

**Result**: PASS ✓ GREEN

**Time elapsed**: 12 minutes

### TDD Cycle 3: Refactor + More Tests

**Add validation tests**:
```go
func TestAppName_EmptyName(t *testing.T) {
	_, err := valueobjects.NewAppName("")
	if err == nil {
		t.Error("expected error for empty name")
	}
}

func TestAppName_InvalidCharacters(t *testing.T) {
	_, err := valueobjects.NewAppName("word press")  // Space invalid
	if err == nil {
		t.Error("expected error for space in name")
	}
}
```

**Improve implementation**:
```go
import "regexp"

var appNamePattern = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

func NewAppName(name string) (AppName, error) {
	if name == "" {
		return AppName{}, fmt.Errorf("app name cannot be empty")
	}
	if !appNamePattern.MatchString(name) {
		return AppName{}, fmt.Errorf("app name must be lowercase alphanumeric with hyphens")
	}
	if len(name) > 64 {
		return AppName{}, fmt.Errorf("app name too long (max 64 chars)")
	}
	return AppName{value: name}, nil
}
```

**Run tests**: PASS ✓ GREEN

**Time elapsed**: 25 minutes

### Adding Property-Based Tests

**Learning property-based testing** (new technique):

```go
import "pgregory.net/rapid"

func TestAppName_Properties(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		name := rapid.StringMatching(`^[a-z][a-z0-9-]{0,63}$`).Draw(t, "name")

		appName, err := valueobjects.NewAppName(name)

		if err != nil {
			t.Fatalf("valid name rejected: %s", name)
		}

		// Invariant: String() returns original value
		if appName.String() != name {
			t.Errorf("expected %s, got %s", name, appName.String())
		}
	})
}
```

**Run property tests**:
```bash
go test -run Property ./internal/appinstallation/domain/valueobjects/
```

**Result**: PASS (100 random cases) ✓

**Time elapsed**: 35 minutes

### Stopping the Session

**Command**:
```bash
/session stop
```

**Output**:
```
Session complete: domain-appname
- Working time: 35 min
- Paused time: 0 min
- Total time: 35 min ✓ (target reached)

Keep or discard this session?
[K]eep / [D]iscard: K
```

**Decision**: Keep (tests pass, value object complete)

**Squash merge**:
```
✓ Squash merging session/20251111-domain-appname to main
✓ Committed with session metadata
✓ Updated cycle metrics
✓ Session state cleaned up

Ready for 10-minute break.
```

### Commit Created

```
feat(domain): add AppName value object

Session: 20251111 10:00-10:35 (35 min actual, 0 min paused)

Created AppName value object with validation:
- Pattern: lowercase alphanumeric with hyphens
- Length: 1-64 characters
- Must start with letter

Added property-based tests using rapid framework.
Tests verify invariants across 100 random inputs.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Taking Break

**Duration**: 10 minutes (required)
**Activity**: Walk, coffee, stretch
**Purpose**: Mental reset before next session

**Time**: 10:35 - 10:45

**Outcome**: First session complete, healthy rhythm established.

---

## Part 5: Weekly Update (Week 1, Friday)

**Date**: Friday, November 15, 2025, 4:00 PM
**Activity**: Weekly hill chart review ritual

### Review Week's Sessions

**Sessions completed**: 8

**Breakdown**:
1. 20251111-domain-appname: 35 min - kept
2. 20251111-domain-databasename: 38 min - kept
3. 20251112-domain-port: 40 min - kept
4. 20251112-domain-directorypath: 42 min - kept
5. 20251113-domain-installation-thinking: 45 min - **discarded**
6. 20251113-domain-installation: 52 min - kept
7. 20251114-domain-manifest: 38 min - kept
8. 20251115-domain-ports: 35 min - kept

**Total working time**: 5.4 hours
**Avg session**: 40.6 min (slightly over, acceptable)
**Discarded**: 1 session (initial Installation aggregate - too complex)

### Assessing Scope Progress

**Domain Layer Foundation**:

**Completed work**:
- ✓ All 4 value objects (AppName, DatabaseName, Port, DirectoryPath)
- ✓ Installation aggregate (simplified, no explicit state machine)
- ✓ Manifest aggregate
- ✓ Repository interfaces (ports)
- ✓ Property-based tests for all value objects

**Current status**: Still figuring out validation patterns

**Assessment**: **Uphill**, but making progress

**Position**: Move from 0/10 to 3/10

**Rationale**:
- Completed core value objects (concrete progress)
- Still learning property-based testing (uphill = learning)
- Installation aggregate needed iteration (discarded one attempt)
- Not at crest yet (haven't proven domain works end-to-end)

### Updating Hill Chart

**New hill chart**:

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Domain Layer**: ● Position 3/10 - Learning validation patterns, value objects complete
**In-Memory Adapter**: ○ Position 0/10 - Not started
**Filesystem Adapter**: ○ Position 0/10 - Not started
**CLI Layer**: ○ Position 0/10 - Not started
**Integration Tests**: ○ Position 0/10 - Not started

**Change**: Domain Layer moved from 0 to 3 (progress visible).

### Updating Next Session

**Domain Layer - Next Session**:
```markdown
**Next Session**: Add domain-level integration test (Installation workflow without I/O)
```

**Rationale**: Prove Installation aggregate works before moving to adapters.

### Updating Session Metrics

**Cycle file addition**:
```markdown
## Session Metrics

### Week 1 (2025-11-11 to 2025-11-15)

- **Sessions completed**: 8
- **Total working time**: 5.4 hours
- **Avg session**: 41 min (target: 40) ✓
- **Sessions >60 min**: 0
- **Sessions discarded**: 1 (domain-installation-thinking: too complex)

**Recent sessions**:
1. 20251111-domain-appname: 35 min (0 min paused) - kept
2. 20251111-domain-databasename: 38 min (2 min paused) - kept
[... 6 more ...]

**Notes**:
- Property-based testing learning curve slowed Week 1 start
- One discarded session expected (exploring state machine approach)
- Healthy pace overall, no long sessions

### All-Time

- **Total sessions**: 8
- **Total working time**: 5.4 hours
- **Avg session**: 41 min
- **Sessions kept**: 7 (88%)
- **Sessions discarded**: 1
```

### Committing Update

```bash
git add plans/cycles/cycle-01.md
git commit -m "chore: update Cycle 01 hill chart - Week 1 complete

Domain Layer: 0/10 → 3/10 (uphill, learning validation patterns)
8 sessions completed, 5.4 hours, avg 41 min.
1 session discarded (state machine complexity).

Next: Domain integration test to prove workflow.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
"
```

**Outcome**: Hill chart updated, visible progress, plan for Week 2 clear.

---

## Part 6: Scope Hammer (Week 5, Tuesday)

**Date**: Tuesday, December 10, 2025 (Week 5, Day 2)
**Context**: 1.5 weeks until circuit breaker
**Activity**: Forcefully question scope to fit time box

### Observing the Situation

**Current hill chart** (Week 5):

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Domain Layer**: ✓ Position 10/10 - Complete
**In-Memory Adapter**: ● Position 9/10 - Almost complete (downhill, executing final tests)
**Filesystem Adapter**: ● Position 4/10 - Uphill (stuck for 2 weeks figuring out permissions)
**CLI Layer**: ○ Position 0/10 - Not started
**Integration Tests**: ○ Position 0/10 - Not started

**Sessions this week**: 6
**Total sessions**: 52
**Time remaining**: 10 days (max 20 more sessions at current pace)

### The Problem

**Math**:
- 2 scopes not started (CLI Layer, Integration Tests)
- 1 scope stuck uphill (Filesystem Adapter)
- 10 days remaining
- ~20 sessions remaining (if healthy pace maintained)

**Estimate per scope**:
- Filesystem Adapter: 8-10 sessions to complete (downhill would be 5-6, but still uphill)
- CLI Layer: 8-10 sessions (new scope)
- Integration Tests: 6-8 sessions (new scope)

**Total needed**: 22-28 sessions
**Total available**: ~20 sessions

**Conclusion**: Can't complete all 5 scopes. Scope hammer needed.

### Questioning Assumptions

**Question 1**: Is Integration Tests scope a must-have?

**Analysis**:
- Already have unit tests (domain layer)
- Already have adapter tests (in-memory, filesystem)
- Integration tests prove end-to-end workflow
- **But**: Can ship without them (defer to cooldown)

**Conclusion**: **Nice-to-have, not must-have**. Cut entirely.

**Question 2**: Is CLI Layer a must-have?

**Analysis**:
- Without CLI, can't use the tool (fatal)
- CLI is thin layer (wiring only, no logic)
- **Must-have**: CLI is the user interface

**Conclusion**: **Must-have**. Keep.

**Question 3**: Can Filesystem Adapter be simplified?

**Analysis**:
- Stuck on permission handling edge cases
- Core functionality works (directory creation, basic provisioning)
- Edge cases are for unusual setups
- **Simplification**: Ship "good enough" permissions, defer edge cases

**Conclusion**: **Simplify to good enough**. Mark as scope hammer.

### Making the Cuts

**Scope Hammer Log entry**:
```markdown
## Scope Hammer Log

- **2025-12-10**: Cut Integration Tests entirely - Nice-to-have, defer to cooldown - Saves 6-8 sessions
- **2025-12-10**: Simplify Filesystem Adapter permissions - Ship good enough (0755), defer edge cases - Saves 3-4 sessions
```

**Updated scopes**:
- ✓ Domain Layer (complete)
- ✓ In-Memory Adapter (complete next session)
- ~ Filesystem Adapter (simplified, good enough)
- CLI Layer (focus here)
- ✗ Integration Tests (cut)

**Remaining work**: ~15 sessions needed, ~20 available. **Feasible.**

### Updating Hill Chart

**Revised plan** reflected in hill chart:

**Filesystem Adapter**: ● Position 6/10 - Simplified to good enough, moving downhill
**CLI Layer**: ○ Position 0/10 - Starting this week
**Integration Tests**: ~~Removed from hill chart~~ (cut)

### Committing Scope Hammer Decision

```bash
git add plans/cycles/cycle-01.md
git commit -m "chore: apply scope hammer - Cycle 01 Week 5

Cut Integration Tests entirely (nice-to-have, defer to cooldown).
Simplify Filesystem Adapter (good enough permissions, defer edge cases).

Remaining work: Domain ✓, In-Memory ✓, Filesystem ~, CLI.
Feasible for circuit breaker in 10 days.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
"
```

**Outcome**: Scope reduced to fit time box, circuit breaker achievable.

---

## Part 7: Circuit Breaker (Week 6, Friday)

**Date**: Friday, December 20, 2025 (Week 6, Day 5)
**Activity**: End-of-cycle decision: ship, kill, or extend?

### Reviewing Final Status

**Hill chart** (end of Week 6):

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Domain Layer**: ✓ Position 10/10 - Complete
**In-Memory Adapter**: ✓ Position 10/10 - Complete
**Filesystem Adapter**: ✓ Position 10/10 - Complete (simplified)
**CLI Layer**: ✓ Position 10/10 - Complete (core functionality)
**Integration Tests**: ✗ Cut via scope hammer

**Final metrics**:
- **Duration**: 6 weeks (full cycle)
- **Sessions**: 67 total
- **Total working time**: 44.2 hours
- **Avg session**: 39.6 min (healthy ✓)
- **Sessions >60 min**: 3 (4.5%, acceptable)
- **Scopes completed**: 4/5 (80%)
- **Scopes cut**: 1 (Integration Tests)

### Evaluating Options

**Option 1: Ship Partial**

**Pros**:
- 4/5 scopes complete and working
- MVP functional (install command works)
- Clean Architecture proven
- Tests comprehensive (unit + adapter tests)

**Cons**:
- No end-to-end integration tests (risk unknown)
- Filesystem edge cases deferred

**Acceptable?**: **Yes**. MVP is functional, tests sufficient for v0.1.

**Option 2: Kill Project**

**Pros**:
- Would free time for other work

**Cons**:
- 4/5 scopes complete - wasteful to discard
- Architecture works (proven by in-memory adapter)
- Problem still exists (mixed concerns in old code)

**Acceptable?**: **No**. Too much working code to kill.

**Option 3: Extend**

**Pros**:
- Could complete Integration Tests

**Cons**:
- Violates Shape Up principle (fixed time, variable scope)
- Scope already cut - extension not justified
- Integration Tests are nice-to-have (defer to cooldown)

**Acceptable?**: **No**. Unnecessary extension.

### Decision: Ship Partial

**Rationale**:
1. MVP complete (install command functional)
2. 4/5 scopes working (80% complete)
3. Tests sufficient (unit + adapter coverage)
4. Integration Tests can be added in cooldown (not blocking)
5. Scope hammer worked (reduced to fit time box)

**Commit to**:
- Package what's complete
- Deploy as v0.1.0
- Mark Integration Tests as "deferred to cooldown"
- Archive cycle with SHIPPED status

### Packaging for Ship

**Actions**:

1. **Run full test suite**:
```bash
go test ./...
```
Result: PASS (all tests green)

2. **Update CHANGELOG.md**:
```markdown
## [0.1.0] - 2025-12-20

### Added
- Clean Architecture refactoring of `install` command
- Pure domain layer with value objects and aggregates
- In-memory adapter for testing
- Filesystem adapter for production
- CLI layer with improved error handling

### Changed
- Rebuilt `install` command from scratch
- Separated concerns across architecture layers

### Technical
- Added property-based testing using rapid framework
- Enforced domain purity (no I/O in domain layer)
- 67 sessions, 44.2 hours of focused work
```

3. **Tag release**:
```bash
git tag -a v0.1.0 -m "Release v0.1.0: Clean Architecture refactoring

MVP: install command rebuilt with Clean Architecture.
4/5 scopes shipped, integration tests deferred to cooldown.

Cycle 01: 6 weeks, 67 sessions, 44.2 hours.
"
git push origin v0.1.0
```

4. **Archive cycle** (see Part 8)

### Outcome

**Status**: SHIPPED ✓

**What shipped**:
- Domain Layer: Complete
- In-Memory Adapter: Complete
- Filesystem Adapter: Complete (good enough)
- CLI Layer: Complete (core functionality)

**What's deferred**:
- Integration Tests: Cooldown work (not blocking)
- Filesystem edge cases: Future enhancement

**Learnings**: See Part 8 (Archive).

---

## Part 8: Cooldown (Week 7-8)

**Dates**: December 23, 2025 - January 3, 2026 (2 weeks)
**Activity**: Reflection, bug fixes, shaping next pitch

### Archiving the Cycle

**Action**: Move cycle file to archive with outcome.

**Command**:
```bash
mv plans/cycles/cycle-01.md plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md
```

**Add Outcome section** at top of archived file:

```markdown
# Cycle 01: Clean Architecture Refactoring

## Outcome

**Status**: SHIPPED (Partial)
**Date**: 2025-12-20
**Decision**: Ship what's complete, defer integration tests to cooldown

### What Shipped
- ✓ Domain Layer: Complete (100%)
- ✓ In-Memory Adapter: Complete (100%)
- ✓ Filesystem Adapter: Complete (simplified, good enough)
- ✓ CLI Layer: Core functionality (95%)
- ✗ Integration Tests: Deferred to cooldown

### Final Metrics
- **Duration**: 6 weeks (full cycle)
- **Sessions**: 67 total
- **Working time**: 44.2 hours
- **Avg session**: 39.6 min (within target ✓)
- **Sessions >60 min**: 3 (4.5%)
- **Scopes completed**: 4/5 (80%)
- **Scopes cut**: 1 (Integration Tests)
- **Sessions kept**: 64 (96%)
- **Sessions discarded**: 3 (4%)

### Learnings

**What Worked**:
- ✓ Domain purity enforcement - HIGH value (fast tests, confidence)
- ✓ Property-based testing - Found edge cases we missed
- ✓ In-memory adapter first - Proved architecture before filesystem complexity
- ✓ Scope hammer in Week 5 - Honest assessment prevented deadline panic
- ✓ TDD rhythm - 40-min sessions kept focus sharp
- ✓ Session discarding - Freedom to throw away experiments (3 discarded, no guilt)

**What Didn't Work**:
- ⚠️ Initial underestimation - Property-based testing curve steeper than expected (Week 1-2 slower)
- ⚠️ Filesystem permissions - Rabbit hole as predicted, scope hammer needed
- ⚠️ Integration Tests appetite - Should have marked nice-to-have from start

**Surprises**:
- Domain purity easier to enforce than expected (`go list` check worked perfectly)
- State machine was YAGNI (procedural approach sufficient)
- rapid framework better than gopter (simpler API)

**Future Improvements**:
- Mark nice-to-haves explicitly in pitch (prevent Week 5 scope hammer drama)
- Budget learning curve time for new techniques (property-based testing took 2 weeks to click)
- Start filesystem work earlier (integration work revealed hidden complexity)

---

[... rest of cycle file preserved as-is ...]
```

### Cooldown Activities

#### Week 7: Bug Fixes and Integration Tests

**Activity**: Add deferred integration tests

**Sessions**:
1. Setup testcontainers infrastructure
2. Write end-to-end workflow test
3. Add filesystem integration test with Docker
4. Document integration test patterns

**Outcome**: Integration Tests scope completed (no cycle pressure).

#### Week 7-8: Reflection

**Questions**:
1. **What did I learn?**
   - Clean Architecture in Go is achievable
   - Property-based testing is powerful but has learning curve
   - Domain purity is enforceable and valuable
   - Scope hammer works (necessary intervention)

2. **What would I do differently?**
   - Mark nice-to-haves explicitly in pitch
   - Budget time for learning curves
   - Start integration work earlier

3. **Am I proud of the work?**
   - **Yes**. Shipped working code, healthy pace, no burnout.

#### Week 8: Shaping Next Pitch

**Idea exploration**:
- Add `upgrade` command?
- Add `remove` command?
- Improve CLI output (colors, progress bars)?

**Appetite setting**:
- `upgrade`: 6 weeks (big batch) - Complex workflow
- `remove`: 2 weeks (small batch) - Simpler, just cleanup
- CLI beautification: 1 week (small batch) - Polish

**Decision**: Defer to later. Take full cooldown, return refreshed.

### Betting Table (End of Week 8)

**Next cycle decision**: Deferred to January.

**Rationale**: Cooldown is for rest and shaping, not committing.

**Outcome**: No bet made yet. Ideas captured in IDEAS.md for future consideration.

---

## Learnings from This Tutorial

### Shape Up Methodology

**Core principles demonstrated**:
1. **Fixed time, variable scope** - 6 weeks fixed, scopes cut to fit
2. **Appetite over estimates** - Set time budget first, designed solution to fit
3. **Circuit breaker** - Hard deadline at Week 6, no exceptions
4. **Scope hammer** - Week 5 intervention prevented deadline panic
5. **Cooldown** - 2 weeks rest/reflection after intense focus

### Session Management

**Automatic tracking benefits**:
- No manual timers (auto-pause for idle/Task agent)
- Progressive reminders (40/60/90 min)
- Session branches (freedom to discard)
- Metrics visibility (avg 39.6 min, healthy)

**Session outcomes**:
- 64 kept (96%) - High success rate
- 3 discarded (4%) - Freedom to experiment without guilt

### TDD + Property-Based Testing

**TDD rhythm**:
1. Write failing test (RED)
2. Make it pass (GREEN)
3. Refactor (CLEAN)
4. Add property tests (PROVE INVARIANTS)

**Property-based testing value**:
- Found edge cases (property tests failed where unit tests passed)
- Confidence in invariants (100 random inputs per test)
- Learning curve (~2 weeks to become proficient)

### Clean Architecture in Go

**Domain purity**:
- Enforced via `go list` check (automated validation)
- Fast tests (domain tests <1ms, no I/O)
- Confidence (can test without infrastructure)

**Adapter pattern**:
- In-memory first (proved domain worked)
- Filesystem second (real implementation)
- Swap implementations easily (future: dry-run adapter)

### Scoping and Planning

**Vertical slices work**:
- Each scope delivered value
- Integration prevented "big bang" integration
- Scope hammer was surgical (cut 1 scope, simplified 1)

**Nice-to-haves are dangerous**:
- Integration Tests felt like must-have
- Week 5 revealed: nice-to-have
- **Lesson**: Mark explicitly in pitch

### Health and Sustainability

**40-minute sessions**:
- Avg 39.6 min (target achieved)
- Only 3 sessions >60 min (4.5%)
- 10-minute breaks enforced
- **Result**: No burnout, sustainable pace

**Discarding sessions**:
- 3 discarded (4%)
- No guilt (freedom to experiment)
- Learnings captured in cycle file

---

## Next Steps

**You've learned**:
- ✓ How to shape a pitch during cooldown
- ✓ How to make betting table decisions
- ✓ How to start a cycle with scopes
- ✓ How to run sessions with TDD
- ✓ How to update hill charts weekly
- ✓ How to use scope hammer when needed
- ✓ How to make circuit breaker decisions
- ✓ How to archive and reflect on cycles

**Now you can**:
1. Shape your own pitches
2. Start your own cycles
3. Run healthy sessions
4. Ship with confidence

**Recommended reading**:
- [Shape Up (Basecamp)](https://basecamp.com/shapeup) - Full methodology book
- [How-To Guides](../how-to/) - Practical guides for each step
- [Reference](../reference/) - Command syntax and formats

**Start your first cycle**: You're ready. Shape a pitch, make a bet, start your Cycle 01.

**Welcome to the methodology**. Ship great software, sustainably.
