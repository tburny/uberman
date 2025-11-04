# How To: Use the Scope Hammer

**Audience**: Developer mid-cycle needing to cut scope

**Type**: How-to guide (task-oriented)

This guide shows you when and how to forcefully cut features to fit the time box.

---

## What is the Scope Hammer?

**Definition**: Forcefully questioning every feature to fit the time box.

**Core principle**: Good enough > perfect

**When to use**:
- Scope stuck uphill >2 weeks
- Week 5 and must-haves incomplete
- Open questions consuming too much time
- Circuit breaker approaching fast

**NOT for**:
- Week 1 (too early, still figuring out)
- Scopes moving downhill (executing known solution)
- All must-haves complete (ship what you have)

---

## When to Use Scope Hammer

### Trigger 1: Scope Stuck Uphill

**Symptom**: Hill chart shows scope uphill for >2 weeks

**Example**:
```
Week 1: Domain Layer ● (uphill - learning validation)
Week 2: Domain Layer ● (uphill - still exploring state machine)
Week 3: Domain Layer ● (uphill - blocked on circular deps)
```

**Problem**: After 2 weeks uphill, unknown unknowns persist.

**Action**: Scope hammer NOW

### Trigger 2: Week 5 and Must-Haves Incomplete

**Symptom**: Circuit breaker in 1 week, core features not working

**Example**:
```
Week 5 status:
- Domain Layer: ● Downhill (working) ✓
- In-Memory: ● Uphill (tests flaky) ✗
- Filesystem: ○ Not started ✗
- CLI: ○ Not started ✗

Circuit breaker: Friday (5 days)
```

**Problem**: Can't finish all scopes in 5 days.

**Action**: Scope hammer IMMEDIATELY

### Trigger 3: Open Questions Consuming Time

**Symptom**: 5+ sessions on same scope, no progress

**Example session log**:
```
Session 1: Explore state machine (42 min) - discarded
Session 2: Try FSM library (38 min) - discarded
Session 3: Event sourcing approach (45 min) - discarded
Session 4: Actor model (40 min) - discarded
Session 5: Refine state machine (41 min) - discarded
```

**Problem**: Rabbit hole consuming appetite.

**Action**: Scope hammer NOW (simplify or cut)

### Trigger 4: Session Metrics Show High Discard Rate

**Symptom**: >50% sessions discarded for a scope

**Example**:
```
Filesystem Adapter scope:
- 12 sessions completed
- 8 discarded (Docker testcontainers, mocking issues)
- 4 kept (basic directory creation)
- Time: 8.1 hours
```

**Problem**: Approach not working, burning time.

**Action**: Scope hammer (simplify testing strategy)

---

## How to Use Scope Hammer

### Step 1: Identify What's Eating Time

**Review**:
- Session metrics (total time per scope)
- Discarded sessions (why?)
- Hill chart (uphill vs downhill)
- Open questions (unanswered blockers)

**Example analysis**:
```markdown
Filesystem Adapter scope:
- Time: 8.1 hours (12 sessions)
- Kept: 4 sessions (directory creation works)
- Discarded: 8 sessions (Docker testcontainers flaky)
- Hill: Uphill (blocked on test infrastructure)
- Open question: How to test filesystem errors reliably?
```

**Insight**: Docker testcontainers is the problem, not the solution.

### Step 2: Question: Must-Have or Nice-to-Have?

**Framework**:
- **Must-have**: Cannot ship without this
- **Nice-to-have**: Improves quality but not essential

**Question to ask**:
> If I cut this entirely, can I still ship a working feature?

**Example**:
- Filesystem Adapter itself? **Must-have** (need production implementation)
- Docker testcontainers? **Nice-to-have** (testing improvement, not core)

### Step 3: Cut Nice-to-Have Entirely

**Action**: Remove scope from cycle

**Update cycle file**:
```markdown
## Scope Hammer Log

- 2025-12-10: Cut Docker testcontainers from Filesystem scope
  **Why**: 8 sessions discarded, circuit breaker in 3 days
  **Impact**: Filesystem tests use real temp dirs (still validated)
  **Trade-off**: Tests slower (50ms vs 10ms), but working
```

**Update scope description**:
```markdown
### 3. Filesystem Adapter (Production)

**Goal**: Real filesystem implementation for production.

**Deliverables**:
- Filesystem repository ✓
- Directory creation (~/apps/) ✓
- ~~Integration tests (testcontainers)~~ SCOPE HAMMERED
- Integration tests (real temp dirs) ✓
- Error handling (permissions, disk space)

**Hill Position**: ● Downhill (now executing)
```

### Step 4: Simplify Must-Have

**If can't cut entirely**: Simplify implementation

**Question to ask**:
> What's the simplest version that proves this works?

**Example**:
```markdown
Original scope: Application Layer
- Use case abstraction
- Command/query separation
- Transaction boundaries
- Middleware pipeline
- Logging/tracing integration

Simplified scope: Application Layer (Minimal)
- Single use case: InstallApp
- Direct call (no pipeline)
- Basic error handling only
```

**Trade-off**: Less flexible, but ships.

**Future cycle**: Can add sophistication later if pain emerges.

### Step 5: Record in Scope Hammer Log

**Format**:
```markdown
## Scope Hammer Log

- **YYYY-MM-DD**: [What was cut/simplified]
  **Why**: [Trigger reason - stuck uphill, time budget, etc.]
  **Impact**: [What changes in delivery]
  **Trade-off**: [What we give up]
  **Future**: [Can revisit in later cycle if needed]
```

**Example entries**:
```markdown
## Scope Hammer Log

- **2025-11-15**: Cut Application Layer entirely
  **Why**: Week 3, CLI can call domain directly (simpler)
  **Impact**: CLI has more logic (not ideal, but working)
  **Trade-off**: Less separation, harder to test CLI
  **Future**: Add App Layer in Cycle 02 if pain emerges

- **2025-12-10**: Simplify Filesystem tests (no Docker)
  **Why**: Week 5, Docker testcontainers flaky, circuit breaker in 3 days
  **Impact**: Tests use real temp dirs (~/.uberman-test/)
  **Trade-off**: Tests slower (50ms vs 10ms), cleanup needed
  **Future**: Revisit Docker testing if tests become bottleneck
```

### Step 6: Adjust Scope Description

**Mark cut features**:
```markdown
### 4. Application Layer (~) [SCOPE HAMMERED - Week 3]

**Original goal**: Use case orchestration layer.

**Decision**: Cut entirely (Week 3)
  **Why**: CLI can call domain directly
  **Impact**: CLI wiring simpler, less layers
  **Trade-off**: CLI has domain logic (not ideal)

~~**Deliverables**:~~
~~- InstallApp use case~~
~~- Command/query separation~~
~~- Transaction boundaries~~

**Status**: REMOVED from cycle
```

### Step 7: Continue with Reduced Scope

**Update "Next Session"**:
```markdown
### 🎯 3. Filesystem Adapter (ACTIVE)

**Goal**: Real filesystem implementation (simplified tests).

**Hill Position**: ● Downhill (now executing)
**Next Session**: Implement error handling (permissions, disk space)
**Open Questions**: None (Docker testing abandoned)
**Blocked**: None
```

**Focus**: Execute remaining must-haves.

---

## Scope Hammer Examples

### Example 1: Cut Application Layer

**Context**: Week 3, Domain + In-Memory working, App Layer not started

**Problem**: 3 weeks left, Filesystem + CLI + App Layer remaining

**Analysis**:
- App Layer: Orchestration, nice separation
- Without it: CLI can call domain directly
- Trade-off: CLI has more logic (not ideal)

**Decision**: Cut Application Layer

**Record**:
```markdown
## Scope Hammer Log

- **2025-11-18**: Cut Application Layer entirely
  **Why**: Week 3, 3 scopes remaining, App Layer is nice-to-have
  **Impact**: CLI calls domain directly (simpler wiring)
  **Trade-off**: Less separation, CLI harder to test
  **Future**: Add in Cycle 02 if CLI grows complex
```

**Outcome**: Shipped Domain + In-Memory + Filesystem + CLI (4 of 5 scopes)

### Example 2: Simplify State Machine

**Context**: Week 2, Domain Layer stuck uphill, exploring state machine

**Problem**: 6 sessions discarded, no progress on Installation workflow

**Analysis**:
- State machine: Elegant, formal transitions
- Procedural: Simple, good enough for MVP
- Trade-off: Less elegant, but working

**Decision**: Abandon state machine, use procedural

**Record**:
```markdown
## Scope Hammer Log

- **2025-11-11**: Simplify Installation workflow (no state machine)
  **Why**: Week 2, 6 sessions discarded, rabbit hole detected
  **Impact**: Procedural workflow instead of formal FSM
  **Trade-off**: Less elegant, but simple and testable
  **Future**: Refactor to state machine if complexity grows
```

**Outcome**: Scope moved downhill, completed in 4 more sessions

### Example 3: Cut Property-Based Tests

**Context**: Week 4, Domain tests working, exploring property-based testing

**Problem**: 4 sessions on property tests, core features not started

**Analysis**:
- Property tests: Excellent coverage, find edge cases
- Unit tests: Cover happy path + known errors
- Trade-off: Less coverage, but shipping

**Decision**: Cut property-based tests (nice-to-have)

**Record**:
```markdown
## Scope Hammer Log

- **2025-11-25**: Cut property-based tests from Domain
  **Why**: Week 4, 4 sessions exploring pgregory.net/rapid, other scopes waiting
  **Impact**: Unit tests only (happy path + critical errors)
  **Trade-off**: May miss edge cases in production
  **Future**: Add property tests in Cycle 02 if bugs emerge
```

**Outcome**: Freed 4 hours for Filesystem and CLI scopes

### Example 4: Simplify Error Handling

**Context**: Week 5, Filesystem Adapter working but exploring 10+ error types

**Problem**: Circuit breaker Friday, CLI not started

**Analysis**:
- Comprehensive errors: Disk space, permissions, concurrency, network, etc.
- Minimal errors: File not found, permission denied, generic error
- Trade-off: Less helpful errors, but shipping

**Decision**: Simplify to 3 error types

**Record**:
```markdown
## Scope Hammer Log

- **2025-12-09**: Simplify error handling (3 types only)
  **Why**: Week 5, circuit breaker in 4 days, exploring 10+ error types
  **Impact**: Generic errors for edge cases (not specific messages)
  **Trade-off**: Users see "operation failed" instead of detailed reason
  **Future**: Add detailed errors in Cycle 02 based on user feedback
```

**Outcome**: Completed Filesystem + CLI in 3 days, shipped on time

---

## Scope Hammer Mindset

### "Good Enough" Philosophy

**Mantra**: A working partial feature beats a perfect unfinished feature.

**Examples**:
- Filesystem tests with temp dirs > no tests (waiting for Docker perfection)
- Procedural workflow > no workflow (waiting for perfect state machine)
- CLI calling domain directly > no CLI (waiting for perfect App Layer)

**Shipping teaches**: Real users reveal what actually matters.

### Letting Go of "Perfect"

**Common resistance**:
- "But this isn't Clean Architecture if CLI calls domain directly!"
- "But property-based tests are best practice!"
- "But Docker testcontainers are the right way!"

**Responses**:
- Clean Architecture is a guide, not a religion. Ship first, refine later.
- Best practices are aspirational. Good enough practices ship features.
- Docker is great, but temp dirs work. Prove the feature, optimize later.

**Remember**: You can always improve in a future cycle.

### Circuit Breaker Safety Net

**Worst case**: Scope hammer isn't enough, circuit breaker kills project.

**Outcome**: You learned what doesn't work (valuable).

**Next cycle**: Reshape with better understanding.

**No shame**: Fast feedback is the goal.

---

## Anti-Patterns

### ❌ Anti-Pattern 1: Scope Hammer Too Early

**Symptom**: Week 1, cutting features preemptively

**Problem**: Haven't given scope a chance, may have worked fine

**Fix**: Wait until week 2 minimum, need data to decide

### ❌ Anti-Pattern 2: Cutting Must-Haves

**Symptom**: Cut core feature because "running out of time"

**Problem**: Shipped feature doesn't work, can't demonstrate value

**Fix**: If cutting must-have, question the entire cycle (may need to kill)

### ❌ Anti-Pattern 3: Not Recording Cuts

**Symptom**: Cut feature, no documentation

**Problem**: Forget why it was cut, may re-explore in future cycle

**Fix**: Always update Scope Hammer Log

### ❌ Anti-Pattern 4: Scope Hammer Instead of Kill

**Symptom**: Week 6, cutting 4 of 5 scopes to "ship something"

**Problem**: Remaining scope doesn't prove anything useful

**Fix**: Use circuit breaker to kill project, reshape for next cycle

---

## Decision Framework

**Use this flowchart**:

```
Is scope stuck uphill >2 weeks?
├─ Yes → Is it must-have or nice-to-have?
│   ├─ Must-have → Can you simplify?
│   │   ├─ Yes → Scope hammer (simplify)
│   │   └─ No → Consider circuit breaker (kill)
│   └─ Nice-to-have → Scope hammer (cut entirely)
└─ No → Keep working (not time for scope hammer yet)
```

**Example application**:
```
Is Filesystem Adapter stuck uphill >2 weeks?
└─ Yes → Is it must-have?
    └─ Yes → Can you simplify?
        └─ Yes → Cut Docker testcontainers, use temp dirs
```

---

## Summary Checklist

When using scope hammer:

- [ ] Identified trigger (stuck uphill, week 5, high discard rate)
- [ ] Reviewed session metrics (time, discarded count)
- [ ] Questioned: Must-have or nice-to-have?
- [ ] Decided: Cut entirely or simplify?
- [ ] Updated Scope Hammer Log (what, why, impact, trade-off)
- [ ] Updated scope description (mark cuts)
- [ ] Updated "Next Session" (focus on must-haves)
- [ ] Committed to git
- [ ] Continued with reduced scope

**Remember**: Good enough > perfect. Shipping teaches.

---

## Related Guides

- **Planning**: [How to plan a cycle](plan-a-cycle.md)
- **Tracking**: [How to update hill chart](update-hill-chart.md)
- **Deadline**: [How to handle circuit breaker](circuit-breaker.md)
- **After cycle**: [How to archive cycle](archive-cycle.md)

---

**Last Updated**: 2025-11-04
**Related**: [Explanation: Scope Hammer](../explanation/shape-up-and-sessions.md#scope-hammer), [Reference: Cycle File Format](../reference/cycle-file-format.md)
