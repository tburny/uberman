# How To: Handle the Circuit Breaker

**Audience**: Developer at end of cycle facing deadline

**Type**: How-to guide (task-oriented)

This guide walks you through the circuit breaker decision: ship, kill, or extend.

---

## What is the Circuit Breaker?

**Definition**: Hard deadline at end of cycle. No automatic extensions.

**Purpose**: Prevents sunk cost fallacy, enforces shipping discipline.

**When**: Last day of cycle (typically Friday of week 6)

**Options**:
1. **Ship Partial** (most common) - Package what works, deploy
2. **Kill Project** - Discard work, record learnings, move on
3. **Extend** (RARE) - Only if work is downhill and 90% done

**Default behavior**: No extensions. Ship or kill.

---

## When Circuit Breaker Triggers

**Cycle timeline**:
```
Week 1-5: Building
Week 6 Friday: Circuit breaker decision day
Week 7-8: Cooldown (mandatory)
```

**Example dates**:
```
Cycle 01 started: Nov 4 (Monday)
Circuit breaker: Dec 15 (Friday, week 6)
Decision time: Dec 15, 5pm
Cooldown starts: Dec 16 (regardless of decision)
```

**No negotiation**: Circuit breaker date is set at cycle start, non-negotiable.

---

## Pre-Circuit Breaker Assessment

**Thursday before circuit breaker**: Review cycle status.

### Assessment Checklist

**1. Review Hill Chart**:
```markdown
Hill Chart (Week 6: Dec 9-15)

Domain Layer:      ●● Downhill (complete, tests passing) ✓
In-Memory:         ●● Downhill (complete, tests passing) ✓
Filesystem:        ● Uphill (tests flaky) ✗
Application Layer: ○ Not started (cut week 3) —
CLI:               ○ Not started ✗
```

**2. Check Session Metrics**:
```markdown
### All-Time
- **Total sessions**: 28
- **Total working time**: 18.5 hours
- **Kept**: 22 (79%)
- **Discarded**: 6 (21%)
- **Avg session**: 40 min
```

**3. Review Must-Haves**:
- Domain Layer: ✓ Working
- In-Memory Adapter: ✓ Working
- Filesystem Adapter: ✗ Tests flaky
- CLI: ✗ Not started

**4. Assess Completeness**:
- Can I demonstrate working feature? Partially (domain + in-memory)
- Are must-haves complete? No (filesystem broken, CLI missing)
- What's blocking? Filesystem tests, time running out

---

## Option 1: Ship Partial (Most Common)

**When to choose**:
- At least one scope complete and working
- Partial feature demonstrates value
- Can deploy what's working
- Learned from cycle

**NOT when**:
- Nothing works (consider kill instead)
- Core assumption wrong (consider kill instead)

### Steps for Shipping Partial

#### 1.1 Identify What's Complete and Working

**Review each scope**:
```markdown
✓ Domain Layer: Value objects, validation, pure logic (tests passing)
✓ In-Memory: Proves pattern works, fast tests
✗ Filesystem: Tests flaky, permission errors not handled
✗ CLI: Not started
```

**Question**: What can I deploy?

**Answer**: Domain + In-Memory (proves Clean Architecture works)

#### 1.2 Package Into Deployable Unit

**If code complete**:
```bash
# Run full test suite
go test ./...

# Build binary
go build -o bin/uberman cmd/uberman/main.go

# Tag version (if shipping)
git tag v0.1.0-alpha
```

**If code incomplete** (domain only):
```bash
# Run domain tests
go test ./internal/appinstallation/domain/...

# Commit documentation showing approach
git add docs/ARCHITECTURE.md
git commit -m "docs: document Clean Architecture approach"
```

**Deliverable**: Working code OR documented approach.

#### 1.3 Update Documentation

**What to update**:
- README.md: What works, what's planned
- CHANGELOG.md: What shipped this cycle
- ARCHITECTURE.md: Current state, future plans

**Example CHANGELOG.md**:
```markdown
## [0.1.0-alpha] - 2025-12-15

### Added (Cycle 01 - Shipped Partial)
- Domain layer with value objects (AppName, DirectoryPath, DatabaseName, Port)
- Installation aggregate (pure, no I/O)
- In-memory adapter for testing
- Property-based tests for validation

### Incomplete (Cut at Circuit Breaker)
- Filesystem adapter (tests flaky, scope hammered)
- Application layer (cut week 3, CLI calls domain directly)
- CLI wiring (not started)

### Learnings
- Clean Architecture pattern works for Uberspace domain
- Procedural workflow simpler than state machine
- Docker testcontainers too flaky, temp dirs better

### Next Cycle
- Filesystem adapter (simplified testing)
- CLI wiring (direct domain calls)
```

#### 1.4 Record Unfinished Scopes

**In cycle archive** (created in step 1.5):
```markdown
## Unfinished Scopes

### Filesystem Adapter
**Status**: Uphill (tests flaky)
**Time spent**: 8.1 hours (12 sessions)
**Why incomplete**: Docker testcontainers too complex, ran out of time
**Learnings**: Temp dirs simpler than Docker for filesystem tests
**Next cycle**: Restart with temp dir approach

### CLI Wiring
**Status**: Not started
**Time spent**: 0 hours
**Why incomplete**: Filesystem blocked CLI start
**Next cycle**: Direct domain calls (no App Layer)
```

#### 1.5 Create Cycle Archive

**Filename**: `plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md`

**Content**: Copy entire cycle file, add outcome section at top.

**Outcome section**:
```markdown
# Cycle 01: Clean Architecture Refactoring

**OUTCOME**: SHIPPED PARTIAL
**Shipped**: Domain Layer + In-Memory Adapter
**Cut**: Application Layer (scope hammer week 3)
**Incomplete**: Filesystem Adapter, CLI Wiring
**Date**: 2025-12-15
**Time**: 6 weeks (18.5 hours actual)

## Shipped Features

- Domain layer (100% pure, zero I/O dependencies)
- Value objects with validation
- In-memory adapter (proves pattern)
- Property-based tests

## What Worked

- TDD with 40-minute sessions (sustainable pace)
- Procedural workflow (simpler than state machine)
- Scope hammer week 3 (cut App Layer early)
- Session branches (discarded 6 experiments safely)

## What Didn't Work

- Docker testcontainers (too complex for filesystem tests)
- State machine exploration (rabbit hole, 6 sessions discarded)
- Application Layer (nice-to-have, cut)

## Learnings for Next Cycle

- Use temp dirs for filesystem tests (simpler)
- CLI can call domain directly (skip App Layer)
- Property-based testing valuable (found edge cases)
- 40-minute sessions enforce focus

---

[... rest of cycle file preserved ...]
```

**Move file**:
```bash
cp plans/cycles/cycle-01-clean-architecture.md \
   plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md

# Add outcome section at top (edit file)

# Delete active cycle file
rm plans/cycles/cycle-01-clean-architecture.md
```

#### 1.6 Update CHANGELOG.md

**Add entry** (as shown in step 1.3).

#### 1.7 Commit and Tag

```bash
git add .
git commit -m "feat(cycle01): ship partial - domain + in-memory adapter complete

Shipped Features:
- Domain layer (pure, zero I/O)
- In-memory adapter (proves pattern)
- Property-based tests

Incomplete:
- Filesystem adapter (tests flaky, scope hammered)
- CLI wiring (not started)

Cycle 01 circuit breaker: Shipped partial (2 of 5 scopes)
Archive: plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>"

# Tag if deploying binary
git tag v0.1.0-alpha
```

#### 1.8 Enter Cooldown

**Cooldown starts**: Next day (Saturday, Dec 16)

**Duration**: 2 weeks

**Activities**:
- Week 1: Recovery, bug fixes (if any)
- Week 2: Shaping next cycle (Filesystem + CLI)

---

## Option 2: Kill Project

**When to choose**:
- Core assumption wrong (e.g., domain can't be made pure)
- Nothing works (all scopes failed)
- Learned that problem doesn't matter
- Better approach discovered

**NOT when**:
- Partial feature works (ship partial instead)
- Just ran out of time (ship what works or extend if 90% done)

### Steps for Killing Project

#### 2.1 Reflect: Why Didn't It Work?

**Questions**:
- What was the core assumption?
- Why did it fail?
- What did we learn?
- Was the problem real?
- Was the appetite wrong?

**Example reflection**:
```markdown
## Why Kill?

**Core assumption**: Clean Architecture can make domain 100% pure

**Why it failed**:
- Uberspace CLI commands require os/exec
- Directory operations require os (can't abstract away)
- Even value objects need filesystem validation (directory exists?)

**What we learned**:
- Pure domain is ideal, but not always achievable
- Pragmatic Clean Architecture allows os in domain (with interfaces)
- Hexagonal Architecture more flexible (ports in domain OK)

**Was problem real?**: Yes, current code is messy

**Appetite wrong?**: No, 6 weeks was fine. Approach was wrong.

**Next steps**: Reshape with Hexagonal Architecture instead
```

#### 2.2 Record Learnings in Archive

**Filename**: `plans/cycles/archive/cycle-01-clean-architecture-KILLED.md`

**Outcome section**:
```markdown
# Cycle 01: Clean Architecture Refactoring

**OUTCOME**: KILLED
**Reason**: Core assumption violated (domain can't be pure)
**Date**: 2025-12-15
**Time**: 6 weeks (22.4 hours actual)

## Core Assumption (WRONG)

Clean Architecture with 100% pure domain (zero I/O dependencies).

## Why It Failed

- Uberspace CLI requires os/exec
- Directory operations require os
- Filesystem validation needed (path exists?)

**Conclusion**: Pure domain not achievable for this domain.

## What We Learned

1. **Clean Architecture too strict**: Pure domain forces awkward abstractions
2. **Hexagonal Architecture better fit**: Ports in domain OK
3. **Pragmatic over purist**: Allow os interfaces in domain
4. **TDD works**: 40-minute sessions sustainable
5. **Property-based testing valuable**: Found validation edge cases

## What Worked (Keep)

- TDD with 40-minute sessions (sustainable)
- Property-based tests (rapid library)
- Session branches (safe experimentation)
- Scope hammer (cut App Layer week 3)

## What Didn't Work (Change)

- 100% pure domain (unrealistic for this problem)
- State machine exploration (rabbit hole)
- Docker testcontainers (too complex)

## Session Metrics

- **Total sessions**: 30
- **Total time**: 22.4 hours
- **Kept**: 18 (60%)
- **Discarded**: 12 (40%) ← High discard = exploration
- **Avg session**: 45 min (slightly high)

**High discard rate expected** for killed project (lots of exploration).

## Next Cycle Plan

**Reshape with**:
- Hexagonal Architecture (ports in domain OK)
- Pragmatic approach (allow os interfaces)
- Skip state machine (procedural workflow)
- Temp dirs for tests (not Docker)

**Appetite**: 6 weeks (same)
**Bet**: Will shape during cooldown, decide at betting table

---

[... rest of cycle file preserved ...]
```

#### 2.3 Delete Working Branches

**If session branches still exist**:
```bash
# List session branches
git branch | grep 'session/'

# Delete all
git branch | grep 'session/' | xargs git branch -D

# Or delete individually
git branch -D session/20251104-domain-databasename
git branch -D session/20251105-domain-state-machine
```

#### 2.4 Reset to Pre-Cycle State (Optional)

**If you want clean slate**:
```bash
# View commits from this cycle
git log --oneline --since="2025-11-04"

# Optionally revert cycle commits
git revert <commit1> <commit2> ...

# Or keep exploratory code (may be useful later)
# No action needed if keeping
```

**Recommendation**: Keep exploratory code (commit count shows work done).

#### 2.5 Create Archive and Delete Cycle File

```bash
cp plans/cycles/cycle-01-clean-architecture.md \
   plans/cycles/archive/cycle-01-clean-architecture-KILLED.md

# Add outcome section (edit file)

# Delete active cycle file
rm plans/cycles/cycle-01-clean-architecture.md
```

#### 2.6 Commit

```bash
git add .
git commit -m "feat(cycle01): kill project - core assumption violated

Cycle 01 circuit breaker: KILLED

Reason: Clean Architecture with pure domain not achievable
- Uberspace CLI requires os/exec
- Directory operations require os
- Filesystem validation needs path checking

Learnings:
- Hexagonal Architecture better fit (ports in domain OK)
- Pragmatic over purist (allow os interfaces)
- TDD with 40-minute sessions works well
- Property-based testing valuable

Next cycle: Reshape with Hexagonal Architecture

Archive: plans/cycles/archive/cycle-01-clean-architecture-KILLED.md

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>"
```

#### 2.7 Enter Cooldown

**Cooldown starts**: Next day

**Focus**:
- Recovery (killing projects is emotionally hard)
- Reflection (what did we learn?)
- Reshaping (Hexagonal Architecture pitch)

**No shame**: Fast feedback is valuable. Killing projects prevents wasting months on wrong approach.

---

## Option 3: Extend (RARE)

**When to choose** (ALL must be true):
- All must-haves working (tests passing)
- Work is downhill (executing, not figuring out)
- < 10% remains (e.g., polish, error messages, docs)
- Extension ≤ 1 week

**NOT when**:
- Any scope still uphill (unknown unknowns persist)
- Must-haves incomplete (ship partial instead)
- > 10% remains (ship partial, not extend)

**Why rare**: Violates circuit breaker discipline. Slippery slope.

### Steps for Extension

#### 3.1 Assess Criteria

**Checklist**:
```markdown
Hill Chart (Week 6):
- Domain Layer: ●● Downhill (complete) ✓
- In-Memory: ●● Downhill (complete) ✓
- Filesystem: ●● Downhill (complete) ✓
- CLI: ● Downhill (90% done, polish remaining) ✓

Must-haves complete? ✓ YES (all working)
Work downhill? ✓ YES (all executing)
< 10% remains? ✓ YES (error messages + docs only)
Extension ≤ 1 week? ✓ YES (3 days estimated)

**Criteria met**: ALL ✓
**Decision**: EXTEND (rare exception)
```

**If ANY criteria fails**: Ship partial or kill instead.

#### 3.2 Document Why Extension Needed

**In cycle file**:
```markdown
## Circuit Breaker Extension

**Date**: 2025-12-15 (original circuit breaker)
**Reason**: All must-haves complete, 3 days polish remaining
**Extension**: 3 days (new circuit breaker: Dec 18)

### Criteria Met

- ✓ Must-haves working (Domain, In-Memory, Filesystem, CLI)
- ✓ Work downhill (all executing)
- ✓ < 10% remains (error messages, docs)
- ✓ Extension ≤ 1 week (3 days)

### Remaining Work

- Polish CLI error messages (1 day)
- Update README.md (1 day)
- Integration test edge cases (1 day)

### Commitment

If not done by Dec 18:
- Ship as-is (no further extensions)
- Circuit breaker HARD STOP
```

#### 3.3 Set New Circuit Breaker Date

**Original**: Dec 15 (Friday)
**Extension**: Dec 18 (Monday, +3 days)

**Commit**:
```bash
git add plans/cycles/cycle-01-clean-architecture.md
git commit -m "feat(cycle01): extend circuit breaker by 3 days (rare exception)

All must-haves complete and downhill:
- Domain Layer ✓
- In-Memory ✓
- Filesystem ✓
- CLI ✓ (90% done)

Extension for polish only:
- CLI error messages
- README update
- Integration test edge cases

New circuit breaker: Dec 18 (HARD STOP, no further extensions)

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>"
```

#### 3.4 Work Extension Period

**Focus**: Remaining 10% only.

**If new unknowns emerge**: Stop immediately, ship as-is.

**No scope changes**: Polish only, no new features.

#### 3.5 Extended Circuit Breaker (Hard Stop)

**On extended date** (Dec 18):
- Ship as-is (no further extensions)
- If incomplete: Ship partial
- Enter cooldown immediately

**No negotiation**: Extended circuit breaker is HARD STOP.

---

## Decision Framework

**Use this flowchart**:

```
Is work done?
├─ Yes → Ship (celebrate!) → Cooldown
└─ No → Are must-haves working?
    ├─ Yes → Is work downhill and <10% remains?
    │   ├─ Yes → Extend 1 week (rare) → Extended circuit breaker
    │   └─ No → Ship partial → Cooldown
    └─ No → Is core assumption wrong?
        ├─ Yes → Kill → Cooldown
        └─ No → Ship partial → Cooldown
```

**Example application**:
```
Is work done?
└─ No → Are must-haves working?
    └─ Yes (Domain + In-Memory) → Is work downhill?
        └─ No (Filesystem uphill) → Ship partial
```

---

## Circuit Breaker Scenarios

### Scenario A: Shipped Complete

**Status**:
- All 5 scopes complete
- Tests passing
- Documentation updated
- Ready to deploy

**Decision**: Ship complete (celebrate!)

**Action**: Tag version, deploy, archive cycle, enter cooldown.

### Scenario B: Shipped Partial (Typical)

**Status**:
- 2 of 5 scopes complete (Domain + In-Memory)
- Proves Clean Architecture works
- Remaining scopes cut or incomplete

**Decision**: Ship partial

**Action**: Archive cycle with SHIPPED status, document incomplete scopes, enter cooldown.

### Scenario C: Killed Project

**Status**:
- Core assumption violated (pure domain impossible)
- All scopes failed or blocked
- Learned approach doesn't work

**Decision**: Kill

**Action**: Archive cycle with KILLED status, record learnings, reshape in cooldown.

### Scenario D: Extended (Rare)

**Status**:
- 4 of 5 scopes complete
- 5th scope 90% done (polish remaining)
- All work downhill

**Decision**: Extend 3 days

**Action**: Document extension, commit to hard stop, finish in 3 days.

---

## After Circuit Breaker

**Regardless of decision**: Enter cooldown immediately.

**Cooldown activities**:
- Week 1: Recovery, bug fixes, reflection
- Week 2: Shaping next cycle, betting table

**No skipping cooldown**: 2 weeks mandatory (prevents burnout).

---

## Common Pitfalls

### ❌ Pitfall 1: "Just One More Day"

**Symptom**: Friday circuit breaker, "I'll finish over weekend"

**Problem**: Violates circuit breaker discipline, slippery slope

**Fix**: Ship partial Friday, rest weekend, review Monday in cooldown

### ❌ Pitfall 2: Shipping Broken Code

**Symptom**: Tests failing, "ship anyway"

**Problem**: Broken code in main branch, damages trust

**Fix**: Ship only what works, or kill if nothing works

### ❌ Pitfall 3: Extending When Uphill

**Symptom**: Scope still uphill, "just need more time to figure out"

**Problem**: Unknown unknowns persist, extension won't help

**Fix**: Kill or ship partial, reshape in cooldown

### ❌ Pitfall 4: Not Recording Learnings

**Symptom**: Kill project, no documentation

**Problem**: Repeat same mistakes in future cycles

**Fix**: Always update archive with learnings section

---

## Summary Checklist

Before circuit breaker:

- [ ] Review hill chart (uphill vs downhill)
- [ ] Check session metrics (time, kept vs discarded)
- [ ] Assess must-haves (working or broken)
- [ ] Decide: Ship, kill, or extend?

If shipping partial:

- [ ] Identify what works
- [ ] Package deployable unit (or document approach)
- [ ] Update documentation (README, CHANGELOG, ARCHITECTURE)
- [ ] Record unfinished scopes
- [ ] Create cycle archive (SHIPPED)
- [ ] Commit and tag (if deploying)
- [ ] Enter cooldown

If killing:

- [ ] Reflect why it failed
- [ ] Record learnings in archive
- [ ] Delete working branches (optional)
- [ ] Create cycle archive (KILLED)
- [ ] Commit
- [ ] Enter cooldown

If extending (rare):

- [ ] Verify ALL criteria met
- [ ] Document why extension needed
- [ ] Set new circuit breaker (hard stop)
- [ ] Commit extension
- [ ] Work extension period
- [ ] Ship at extended deadline (no further extensions)

**Always enter cooldown**: 2 weeks, mandatory.

---

## Related Guides

- **Planning**: [How to plan a cycle](plan-a-cycle.md)
- **Mid-cycle**: [How to use scope hammer](scope-hammer.md)
- **Archival**: [How to archive cycle](archive-cycle.md)
- **After cooldown**: [How to plan a cycle](plan-a-cycle.md)

---

**Last Updated**: 2025-11-04
**Related**: [Explanation: Circuit Breaker](../explanation/shape-up-and-sessions.md#circuit-breaker), [Reference: Cycle File Format](../reference/cycle-file-format.md)
