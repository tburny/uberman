# Cycle File Format Reference

Complete cycle file structure and format specification.

**Audience**: Developers creating or updating cycle files during development.

**See also**:
- [Plan a Cycle](../how-to/plan-a-cycle.md) - How to create cycle files
- [Update Hill Chart](../how-to/update-hill-chart.md) - Weekly updates
- [Run a Session](../how-to/run-a-session.md) - Session metrics tracking

---

## File Location

```
plans/cycles/cycle-NN.md
```

Where `NN` is the cycle number (zero-padded):
- `cycle-01.md` - First cycle
- `cycle-02.md` - Second cycle
- `cycle-10.md` - Tenth cycle

**Active cycle**: Only ONE cycle file should be active (not in archive).

**Archive location**: `plans/cycles/archive/cycle-NN-NAME-[SHIPPED|KILLED].md`

---

## Complete Template

```markdown
# Cycle NN: [Cycle Name]

**Status**: ACTIVE | SHIPPED | KILLED
**Dates**: YYYY-MM-DD to YYYY-MM-DD (N weeks)
**Pitch**: ../pitches/YYYY-MM-DD-title.md
**Circuit Breaker**: YYYY-MM-DD

---

## Current Status

**Active Scope**: [Scope name]
**Active Session**: session/YYYYMMDD-scope-task (NN min working) | None
**Last Update**: YYYY-MM-DD HH:MM

---

## Hill Chart (Week N: Mon DD - Fri DD)

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Scope 1**: ● Position 5/10 - [description]
**Scope 2**: ○ Position 0/10 - Not started
**Scope 3**: ● Position 8/10 - [description]
**Scope 4**: ◐ Position 6/10 - [description]
**Scope 5**: ● Position 10/10 - Complete

**Legend**:
- ○ Not started
- ● Active/in progress
- ◐ Blocked/stuck
- ✓ Complete

---

## Scopes

### 🎯 [Active Scope Name] (ACTIVE)

**Goal**: [What this scope accomplishes - 1 sentence]

**Hill Position**: [Uphill/Downhill] (Position N/10) - [what you're figuring out or executing]

**Next Session**: [Concrete next task - specific enough to `/session start`]

**Completed**:
- [x] Task 1
- [x] Task 2

**In Progress**:
- [ ] Current task

**Remaining**:
- [ ] Future task 1
- [ ] Future task 2

**Open Questions**:
- [Question 1]?
- [Question 2]?

**Blocked**: None | [Blocker description]

**Discoveries**:
- [Key learning or insight]

---

### [Scope Name 2]

**Goal**: [What this scope accomplishes]

**Hill Position**: [Uphill/Downhill] (Position N/10) - [status]

**Next Session**: [Next concrete task]

**Completed**:
- [x] Completed work

**Remaining**:
- [ ] Future work

**Open Questions**: None

**Blocked**: None

---

### [Scope Name 3]

[Same structure as above]

---

## Session Metrics

### Week N (YYYY-MM-DD to YYYY-MM-DD)

- **Sessions completed**: N
- **Total working time**: N.N hours
- **Avg session**: NN min (target: 40)
- **Sessions >60 min**: N ⚠️
- **Sessions discarded**: N

**Recent sessions**:
1. YYYYMMDD-scope-task: NN min working (NN min paused) - kept | discarded
2. YYYYMMDD-scope-task: NN min working (NN min paused) - kept | discarded
3. [...]

**Notes**:
- [Any observations about session patterns]

---

### All-Time (Cycle Total)

- **Total sessions**: NN
- **Total working time**: NN.N hours
- **Avg session**: NN min
- **Sessions kept**: NN (NN%)
- **Sessions discarded**: NN (NN%)
- **Longest session**: NN min (YYYYMMDD-scope-task)

---

## Scope Hammer Log

Record of features cut or simplified during cycle.

- **YYYY-MM-DD**: [What was cut] - [Why] - [Impact on scope]
- **YYYY-MM-DD**: [What was simplified] - [Why] - [Impact on scope]

**Example**:
- **2025-11-15**: Cut Application Layer entirely - Nice-to-have, CLI can call domain directly - Reduces 1 full scope, simplifies architecture

---

## Decisions

Architecture decision records made during this cycle.

- **ADR-NNN**: [Title] - [Brief description and rationale]

**Example**:
- **ADR-001**: Use rapid for property-based testing - Simpler API than gopter, sufficient for domain validation

---

## Notes

Free-form notes about cycle progress, challenges, learnings.

**Example**:
- Week 2: Property-based tests taking longer than expected (learning curve)
- Week 4: In-memory adapter proved domain purity - big confidence boost
- Week 5: Scope hammer necessary - acknowledged Application Layer was YAGNI

---

## Archive

After circuit breaker, this file moves to:

`plans/cycles/archive/cycle-NN-[NAME]-[SHIPPED|KILLED].md`

**Outcome Section Added** (at top of archived file):

```markdown
## Outcome

**Status**: SHIPPED | KILLED
**Date**: YYYY-MM-DD
**Decision**: [Ship partial / Kill project / Extend (rare)]

### What Shipped (if SHIPPED)
- Scope 1: Complete ✓
- Scope 2: Complete ✓
- Scope 3: Cut (scope hammer)
- Scope 4: Partial (core functionality only)
- Scope 5: Not started

### Why Killed (if KILLED)
[Explanation of why project was killed]
[Learnings from the attempt]

### Final Metrics
- Duration: N weeks
- Sessions: NN total
- Working time: NN.N hours
- Scopes completed: N/N
- Scopes cut: N
```

---
```

---

## Section Details

### Header

**Cycle Name**: Descriptive, concise (2-4 words)
- ✅ "Clean Architecture Refactoring"
- ✅ "Upgrade Command"
- ❌ "Implementing the upgrade command with version detection and rollback"

**Status**:
- `ACTIVE` - Current cycle in progress
- `SHIPPED` - Completed and deployed
- `KILLED` - Abandoned via circuit breaker

**Dates**: ISO 8601 format (YYYY-MM-DD)
- Start date: First day of Week 1
- End date: Last day of final week (Week 6 for big batch)

**Pitch**: Relative path to pitch file that started this cycle

**Circuit Breaker**: Date when cycle MUST end (no extensions by default)

### Current Status

**Active Scope**: Which scope from hill chart is currently being worked on (marked with 🎯)

**Active Session**: Current session branch name and working time, or "None" if no session active

**Last Update**: Timestamp of last cycle file modification (helpful for staleness detection)

### Hill Chart

ASCII visualization of scope progress.

**Positions** (0-10 scale):
- `0`: Not started
- `1-4`: Uphill (figuring things out, many unknowns)
- `5`: Crest (transition point, clarity emerging)
- `6-9`: Downhill (executing known work, few unknowns)
- `10`: Complete

**Symbols**:
- `○` Not started
- `●` Active/in progress
- `◐` Blocked/stuck (>2 weeks in same position)
- `✓` Complete

**Update frequency**: Weekly (Friday afternoon ritual)

**Example**:
```
**Scope 1: Domain Layer**: ● Position 8/10 - Executing final value objects
**Scope 2: In-Memory Adapter**: ● Position 5/10 - Transitioning to implementation
**Scope 3: Filesystem Adapter**: ○ Position 0/10 - Not started
```

### Scopes

**Active scope marker**: 🎯 emoji marks the scope currently being worked on.

**Goal**: One-sentence description of what scope achieves. Should be testable/verifiable.

**Hill Position**:
- Phase: Uphill (figuring out) or Downhill (executing)
- Position: N/10 numeric position
- Description: What you're currently doing in this phase

**Next Session**: Concrete task specific enough to use in `/session start <task>`. Should be actionable immediately.

**Task Lists**:
- Completed: Checked boxes `[x]`
- In Progress: Unchecked with marker `[-]` or `[ ]`
- Remaining: Unchecked `[ ]`

**Open Questions**: Blockers or unknowns preventing downhill movement. Empty if all clear.

**Blocked**: "None" or description of blocker (external dependency, technical unknown, etc.)

**Discoveries**: Key learnings during scope work. Especially important for understanding "why uphill."

### Session Metrics

Track session statistics for cycle transparency and health monitoring.

**Weekly Metrics**:
- Sessions completed: Count for the week
- Total working time: Sum in hours (excludes pauses)
- Avg session: Average working time per session
- Sessions >60 min: Count of sessions exceeding 60 min (⚠️ warning indicator)
- Sessions discarded: Count of sessions where work was thrown away

**Recent Sessions**: Last 5-10 sessions with:
- Date (YYYYMMDD)
- Scope-task identifier
- Working time and paused time
- Outcome (kept or discarded)

**All-Time Metrics**:
- Cumulative statistics across entire cycle
- Kept/discarded ratio (quality indicator)
- Longest session (outlier detection)

**Purpose**:
- Detect unhealthy patterns (too many long sessions)
- Celebrate progress (session count growth)
- Identify scope difficulty (many discarded sessions in one scope = trouble)

### Scope Hammer Log

Chronicle of cuts made during cycle.

**When to record**:
- Feature cut entirely
- Feature simplified significantly
- Scope redefined to reduce appetite

**Format**: Date, What, Why, Impact

**Example**:
```
- 2025-11-15: Cut Application Layer - YAGNI, CLI can call domain directly - Saves 1 full scope
- 2025-11-20: Simplify validation - Good enough > perfect - Saves 2-3 sessions
```

**Purpose**:
- Transparency about compromises
- Learning for future appetite setting
- Justification for partial shipping

### Decisions

Architecture decisions made during cycle execution.

**Link to ADRs**: If using formal ADRs, reference them here.

**Inline decisions**: Brief rationale for non-formal decisions.

**Example**:
```
- Use rapid framework: Simpler API than gopter, meets needs
- Skip dry-run initially: Focus on real adapters, add later if needed
- Domain must be 100% pure: Enforce via go list check in tests
```

**Purpose**:
- Remember why choices were made
- Avoid revisiting settled decisions
- Provide context for future maintainers

### Notes

Free-form observations about cycle health.

**Good note topics**:
- Surprising discoveries
- Pace observations (ahead/behind expectations)
- Morale notes (stuck, frustrated, breakthrough)
- External factors (holidays, interruptions)

**Example**:
```
- Week 2: Property-based testing curve steeper than expected, slowing domain scope
- Week 3: Breakthrough - domain purity works! Tests run in <1ms
- Week 5: Realistic assessment - Application Layer is YAGNI, cutting it
```

---

## Update Frequency

| Section | Frequency | Trigger |
|---------|-----------|---------|
| Current Status | Every session | `/session start` and `/session stop` |
| Hill Chart | Weekly | Friday afternoon ritual |
| Scopes | As needed | Discoveries, questions, blockers |
| Session Metrics | Every session | `/session stop` |
| Scope Hammer Log | On cuts | Scope hammer decision |
| Decisions | On decisions | ADR creation or significant choice |
| Notes | Ad-hoc | Any noteworthy observation |

---

## Examples

### Cycle Start (Week 1, Day 1)

```markdown
# Cycle 01: Clean Architecture Refactoring

**Status**: ACTIVE
**Dates**: 2025-11-04 to 2025-12-15 (6 weeks)
**Pitch**: ../pitches/2025-11-03-rebuild-install-clean-architecture.md
**Circuit Breaker**: 2025-12-15

---

## Current Status

**Active Scope**: Domain Layer Foundation
**Active Session**: None
**Last Update**: 2025-11-04 09:00

---

## Hill Chart (Week 1: Mon 04 - Fri 08)

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

---

## Scopes

### 🎯 Domain Layer Foundation (ACTIVE)

**Goal**: Create pure domain layer with value objects and Installation aggregate.

**Hill Position**: Uphill (Position 0/10) - Starting from scratch

**Next Session**: Create AppName value object with validation

**Completed**: None

**Remaining**:
- [ ] Value objects: AppName, DatabaseName, Port, DirectoryPath
- [ ] Installation aggregate
- [ ] Manifest aggregate
- [ ] Domain services
- [ ] Repository interfaces (ports)

**Open Questions**:
- Should Installation be aggregate or just entity?
- State machine explicit or implicit?

**Blocked**: None

[... other scopes ...]

---

## Session Metrics

### Week 1 (2025-11-04 to 2025-11-08)

- **Sessions completed**: 0
- **Total working time**: 0.0 hours
- **Avg session**: - min
- **Sessions >60 min**: 0
- **Sessions discarded**: 0

**Recent sessions**: None yet

---

### All-Time

- **Total sessions**: 0
- **Total working time**: 0.0 hours

---

## Scope Hammer Log

None yet.

---

## Decisions

- **Domain purity**: Enforce 100% pure domain (no os, exec, io, net imports)

---

## Notes

- Starting fresh, excitement high
- Open questions about aggregate design - will probe during Week 1

```

### Mid-Cycle (Week 3, Active Work)

```markdown
# Cycle 01: Clean Architecture Refactoring

**Status**: ACTIVE
**Dates**: 2025-11-04 to 2025-12-15 (6 weeks)
**Pitch**: ../pitches/2025-11-03-rebuild-install-clean-architecture.md
**Circuit Breaker**: 2025-12-15

---

## Current Status

**Active Scope**: In-Memory Adapter
**Active Session**: session/20251118-infrastructure-inmemory (18 min working)
**Last Update**: 2025-11-18 10:35

---

## Hill Chart (Week 3: Mon 18 - Fri 22)

```
      /\
     /  \____
    /        \____
Uphill      Downhill
```

**Domain Layer**: ✓ Position 10/10 - Complete
**In-Memory Adapter**: ● Position 6/10 - Implementing repository interfaces
**Filesystem Adapter**: ○ Position 0/10 - Not started
**CLI Layer**: ○ Position 0/10 - Not started
**Integration Tests**: ○ Position 0/10 - Not started

---

## Scopes

### Domain Layer Foundation (COMPLETE)

**Goal**: Create pure domain layer with value objects and Installation aggregate.

**Hill Position**: Complete (Position 10/10)

**Completed**:
- [x] AppName value object
- [x] DatabaseName value object
- [x] Port value object
- [x] DirectoryPath value object
- [x] Installation aggregate (simplified - no explicit state machine)
- [x] Manifest aggregate
- [x] Repository interfaces (ports)
- [x] Property-based tests for all value objects

**Discoveries**:
- State machine was YAGNI - procedural approach works fine
- Property-based tests found edge cases we missed
- Domain purity verified via `go list` check

---

### 🎯 In-Memory Adapter (ACTIVE)

**Goal**: Implement in-memory adapter to prove domain works without I/O.

**Hill Position**: Downhill (Position 6/10) - Executing repository implementations

**Next Session**: Implement InstallationRepository in-memory

**Completed**:
- [x] ManifestRepository in-memory implementation
- [x] Basic test suite

**In Progress**:
- [-] InstallationRepository implementation

**Remaining**:
- [ ] ProvisioningService mock
- [ ] Integration test with domain

**Open Questions**: None

**Blocked**: None

[... other scopes ...]

---

## Session Metrics

### Week 3 (2025-11-18 to 2025-11-22)

- **Sessions completed**: 8
- **Total working time**: 5.2 hours
- **Avg session**: 39 min (target: 40) ✓
- **Sessions >60 min**: 1 ⚠️
- **Sessions discarded**: 0

**Recent sessions**:
1. 20251118-infrastructure-inmemory: 38 min (2 min paused) - kept
2. 20251117-infrastructure-setup: 42 min (0 min paused) - kept
3. 20251117-domain-manifest: 35 min (3 min paused) - kept
4. 20251116-domain-installation: 67 min (5 min paused) - kept ⚠️
5. 20251116-domain-directorypath: 40 min (1 min paused) - kept

**Notes**:
- One long session (Installation aggregate) - complex design, expected
- Otherwise healthy pace

---

### All-Time

- **Total sessions**: 23
- **Total working time**: 15.1 hours
- **Avg session**: 39 min
- **Sessions kept**: 22 (96%)
- **Sessions discarded**: 1 (State machine attempt - too complex)
- **Longest session**: 67 min (20251116-domain-installation)

---

## Scope Hammer Log

- **2025-11-16**: Simplified Installation aggregate - Dropped explicit state machine, use procedural approach - Saves 2-3 sessions, sufficient for MVP

---

## Decisions

- **Domain purity**: Enforced via `go list` check (success!)
- **Property-based testing**: Use rapid framework (working well)
- **State machine**: Skip explicit state machine (YAGNI for MVP)
- **In-memory first**: Proves domain before filesystem complexity

---

## Notes

- Week 1: Slower start (property-based testing learning curve)
- Week 2: Breakthrough - domain purity works, tests fast
- Week 3: Confidence high, in-memory adapter proving design
- On track for circuit breaker

```

### End of Cycle (Week 6, Circuit Breaker)

```markdown
# Cycle 01: Clean Architecture Refactoring

**Status**: SHIPPED
**Dates**: 2025-11-04 to 2025-12-15 (6 weeks)
**Pitch**: ../pitches/2025-11-03-rebuild-install-clean-architecture.md
**Circuit Breaker**: 2025-12-15 (REACHED)

---

## Outcome

**Status**: SHIPPED (Partial)
**Date**: 2025-12-15
**Decision**: Ship what's complete, defer integration tests to cooldown

### What Shipped
- ✓ Domain Layer: Complete (100%)
- ✓ In-Memory Adapter: Complete (100%)
- ✓ Filesystem Adapter: Complete (100%)
- ✓ CLI Layer: Core functionality (80%)
- ✗ Integration Tests: Not started (cut via scope hammer)

### Final Metrics
- Duration: 6 weeks (full cycle)
- Sessions: 67 total
- Working time: 44.2 hours
- Avg session: 39.6 min (within target ✓)
- Scopes completed: 4/5 (80%)
- Scopes cut: 1 (Integration Tests deferred)

### Learnings
- Property-based testing curve steeper than expected (Week 1-2 slower)
- Domain purity enforcement was HIGH value (confidence + speed)
- In-memory adapter breakthrough (Week 3) - proved architecture early
- Scope hammer necessary (Week 5) - acknowledged Integration Tests deferrable
- Circuit breaker worked - shipped partial with confidence

---

[... rest of cycle file preserved as-is ...]
```

---

## Related Documentation

- [Plan a Cycle](../how-to/plan-a-cycle.md) - Creating cycle files
- [Update Hill Chart](../how-to/update-hill-chart.md) - Weekly rituals
- [Run a Session](../how-to/run-a-session.md) - Session tracking
- [Scope Hammer](../how-to/scope-hammer.md) - Cutting features
- [Circuit Breaker](../how-to/circuit-breaker.md) - Ending cycles
- [Archive Cycle](../how-to/archive-cycle.md) - Post-cycle process
