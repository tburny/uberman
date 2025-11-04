# How To: Archive Cycle

**Audience**: Developer at end of cycle (after circuit breaker)

**Type**: How-to guide (task-oriented)

This guide shows you how to archive a completed cycle for future reference.

---

## When to Archive

**Trigger**: After circuit breaker decision (ship or kill)

**Timing**: Immediately after circuit breaker, before entering cooldown

**Why archive**:
- Preserve cycle history (learnings, decisions)
- Document what worked and what didn't
- Reference for future cycles
- Calculate historical velocity

---

## Archive Location

**Active cycle**: `plans/cycles/cycle-NN.md`

**Archived cycle**: `plans/cycles/archive/cycle-NN-NAME-[SHIPPED|KILLED].md`

**Naming convention**:
- `cycle-NN`: Cycle number (01, 02, etc.)
- `NAME`: Short cycle name (clean-architecture, graphql-api)
- `SHIPPED` or `KILLED`: Outcome

**Examples**:
- `plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md`
- `plans/cycles/archive/cycle-02-graphql-api-KILLED.md`

---

## Step 1: Copy Cycle File to Archive

**Command**:
```bash
cp plans/cycles/cycle-01-clean-architecture.md \
   plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md
```

**Or** (if killed):
```bash
cp plans/cycles/cycle-01-clean-architecture.md \
   plans/cycles/archive/cycle-01-clean-architecture-KILLED.md
```

**Result**: Archive file created, active cycle file still exists.

---

## Step 2: Add Outcome Section

**Open archive file**, add outcome section at the very top (before title).

### For Shipped Cycles

```markdown
# Cycle 01: Clean Architecture Refactoring

**OUTCOME**: SHIPPED PARTIAL
**Shipped**: Domain Layer + In-Memory Adapter
**Cut**: Application Layer (scope hammer week 3)
**Incomplete**: Filesystem Adapter, CLI Wiring
**Date**: 2025-12-15
**Time**: 6 weeks (18.5 hours actual working time)

## Shipped Features

- Domain layer (100% pure, zero I/O dependencies)
- Value objects: AppName, DirectoryPath, DatabaseName, Port
- Installation aggregate (procedural workflow)
- In-memory adapter (proves Clean Architecture pattern)
- Property-based tests (pgregory.net/rapid)

## What Worked

✓ **TDD with 40-minute sessions**: Sustainable pace, no burnout
✓ **Procedural workflow**: Simpler than state machine (discarded week 1)
✓ **Scope hammer week 3**: Cut Application Layer early
✓ **Session branches**: Discarded 6 experiments safely
✓ **Property-based testing**: Found validation edge cases

## What Didn't Work

✗ **Docker testcontainers**: Too complex/flaky for filesystem tests
✗ **State machine exploration**: Rabbit hole, 6 sessions discarded
✗ **Application Layer**: Nice-to-have, cut (CLI calls domain directly)

## Learnings for Next Cycle

1. **Use temp dirs for filesystem tests**: Simpler than Docker
2. **CLI can call domain directly**: Skip Application Layer
3. **Property-based testing valuable**: Continue using rapid
4. **40-minute sessions enforce focus**: Health-first protocol works
5. **Scope hammer week 3-5 ideal**: Earlier is better

## Session Metrics (Final)

- **Total sessions**: 28
- **Total working time**: 18.5 hours
- **Kept**: 22 (79%)
- **Discarded**: 6 (21%)
- **Avg session**: 40 min (target: 40)
- **Sessions >60 min**: 2 (7%) ← Improved over cycle

## Velocity

- **6 weeks**: 18.5 hours actual working time
- **Average**: 3.1 hours/week
- **Per scope**: Domain (8.2h), In-Memory (4.1h), Filesystem (6.2h discarded)

## Unfinished Scopes

### Filesystem Adapter
**Status**: Uphill (tests flaky)
**Time spent**: 6.2 hours (10 sessions)
**Why incomplete**: Docker testcontainers too complex
**Learnings**: Temp dirs simpler approach
**Next cycle**: Restart with temp dir testing

### CLI Wiring
**Status**: Not started
**Time spent**: 0 hours
**Why incomplete**: Blocked by Filesystem
**Next cycle**: Direct domain calls (no App Layer)

---

[... rest of original cycle file preserved ...]
```

### For Killed Cycles

```markdown
# Cycle 02: GraphQL API

**OUTCOME**: KILLED
**Reason**: Core assumption violated (GraphQL too complex for this domain)
**Date**: 2025-02-09
**Time**: 6 weeks (22.4 hours actual working time)

## Core Assumption (WRONG)

GraphQL provides better API than REST for Uberspace app installation.

## Why It Failed

1. **Schema complexity**: 50+ types for simple install operation
2. **Resolver overhead**: Nested resolvers for relationships overkill
3. **Client complexity**: GraphQL client heavier than fetch/axios
4. **No real benefit**: Single-operation API doesn't need graph traversal

**Conclusion**: REST API simpler, more appropriate for this use case.

## What We Learned

1. **GraphQL not always better**: Simple APIs don't need graph traversal
2. **Schema-first can over-engineer**: Started with schema, forced complexity
3. **REST is good enough**: POST /install endpoint sufficient
4. **Hexagonal architecture works**: Ports/adapters pattern valuable

## What Worked (Keep)

✓ **TDD with 40-minute sessions**: Sustainable pace maintained
✓ **Property-based testing**: Rapid library valuable
✓ **Session branches**: Discarded 12 experiments safely (40% discard rate)
✓ **Scope hammer**: Cut federation early (week 2)

## What Didn't Work (Change)

✗ **GraphQL for simple APIs**: Overkill, REST better fit
✗ **Schema-first approach**: Forces decisions too early
✗ **Apollo Server**: Heavy dependency for simple needs

## Session Metrics (Final)

- **Total sessions**: 30
- **Total working time**: 22.4 hours
- **Kept**: 18 (60%)
- **Discarded**: 12 (40%) ← High discard expected for killed project
- **Avg session**: 45 min (slightly high)
- **Sessions >60 min**: 4 (13%) ← Need to improve

## Velocity (Comparison)

- Cycle 01: 18.5 hours (SHIPPED)
- Cycle 02: 22.4 hours (KILLED)

**Insight**: Killed cycle had more time but less progress (exploration).

## Next Cycle Plan

**Reshape with**:
- REST API (POST /install endpoint)
- OpenAPI schema (not GraphQL schema)
- Fastify (lighter than Apollo)
- Same Hexagonal Architecture (ports/adapters)

**Appetite**: 6 weeks (same)
**Bet**: Will shape during cooldown, decide at betting table

---

[... rest of original cycle file preserved ...]
```

---

## Step 3: Delete Active Cycle File

**After adding outcome section** to archive:

```bash
rm plans/cycles/cycle-01-clean-architecture.md
```

**Result**: Only archive file remains.

**Why delete active**: Prevents confusion (only one source of truth).

---

## Step 4: Update plans/cycles/README.md

**Add link to archived cycle**:

```markdown
# Cycles

Active cycles and archived completed cycles.

## Active Cycles

- (None - currently in cooldown)

## Archived Cycles

### Shipped

- [Cycle 01: Clean Architecture Refactoring](archive/cycle-01-clean-architecture-SHIPPED.md) - Domain + In-Memory (Dec 2025)

### Killed

- [Cycle 02: GraphQL API](archive/cycle-02-graphql-api-KILLED.md) - REST simpler (Feb 2025)

## Cooldown

Currently shaping pitches for next cycle.

**Betting table**: 2025-02-23
```

---

## Step 5: Commit Archive

**Commit message**:

### For Shipped Cycle

```bash
git add plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md
git rm plans/cycles/cycle-01-clean-architecture.md
git add plans/cycles/README.md
git commit -m "$(cat <<'EOF'
docs(planning): archive Cycle 01 (SHIPPED PARTIAL)

Shipped:
- Domain layer (pure, zero I/O)
- In-Memory adapter (proves pattern)

Cut:
- Application Layer (scope hammer week 3)

Incomplete:
- Filesystem Adapter (Docker too complex)
- CLI Wiring (blocked by Filesystem)

Learnings:
- TDD with 40-minute sessions sustainable
- Procedural workflow simpler than state machine
- Temp dirs better than Docker for filesystem tests
- Property-based testing finds edge cases

Metrics:
- 28 sessions (18.5 hours)
- 79% kept, 21% discarded
- Average 40 min/session

Archive: plans/cycles/archive/cycle-01-clean-architecture-SHIPPED.md

Entering cooldown (2 weeks).

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
EOF
)"
```

### For Killed Cycle

```bash
git add plans/cycles/archive/cycle-02-graphql-api-KILLED.md
git rm plans/cycles/cycle-02-graphql-api.md
git add plans/cycles/README.md
git commit -m "$(cat <<'EOF'
docs(planning): archive Cycle 02 (KILLED)

Reason: GraphQL too complex for simple install API

Core assumption wrong:
- GraphQL provides better API than REST
- Reality: 50+ types for single operation overkill
- REST simpler, more appropriate

Learnings:
- GraphQL not always better (simple APIs don't need graphs)
- REST good enough for single-operation APIs
- Schema-first can over-engineer
- Hexagonal architecture still valuable

Metrics:
- 30 sessions (22.4 hours)
- 60% kept, 40% discarded (high discard = exploration)
- Average 45 min/session (slightly high)

Next cycle: Reshape with REST API

Archive: plans/cycles/archive/cycle-02-graphql-api-KILLED.md

Entering cooldown (2 weeks).

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
EOF
)"
```

---

## Step 6: Enter Cooldown

**Duration**: 2 weeks (mandatory)

**Activities**:
- Week 1: Recovery, bug fixes (if shipped)
- Week 2: Shaping next cycle, betting table

**No skipping**: Cooldown is required regardless of outcome.

---

## Archive Format

**Complete structure**:

```markdown
# Cycle NN: [Name]

**OUTCOME**: SHIPPED | SHIPPED PARTIAL | KILLED
**[Outcome-specific metadata]**
**Date**: YYYY-MM-DD
**Time**: N weeks (N.N hours actual)

## [Outcome-Specific Sections]
[Shipped Features | Core Assumption | Why It Failed]

## What Worked
[Keep these practices]

## What Didn't Work
[Change these practices]

## Learnings for Next Cycle
[Specific insights]

## Session Metrics (Final)
[Complete metrics]

## Velocity
[Time breakdown]

## [Outcome-Specific Sections]
[Unfinished Scopes | Next Cycle Plan]

---

[... ORIGINAL CYCLE FILE CONTENT BELOW ...]

# Cycle NN: [Name] (ORIGINAL)

**Status**: ACTIVE → ARCHIVED
**Dates**: [Original dates]
**Pitch**: [Link]
**Circuit Breaker**: [Date]

[... all original sections preserved ...]
```

**Key principle**: Preserve entire original cycle file below outcome section.

---

## Analyzing Archives

**Use archives to**:

### 1. Calculate Historical Velocity

```bash
grep "Total working time:" plans/cycles/archive/*.md

# Example output:
cycle-01-clean-architecture-SHIPPED.md:18.5 hours
cycle-02-graphql-api-KILLED.md:22.4 hours
cycle-03-rest-api-SHIPPED.md:19.2 hours

# Average: 20 hours per 6-week cycle
# = 3.3 hours/week
```

### 2. Review Learnings

```bash
grep -A 10 "## Learnings" plans/cycles/archive/*.md

# Patterns emerge:
- TDD with 40-minute sessions (repeated)
- Temp dirs better than Docker (Cycle 01)
- REST simpler than GraphQL (Cycle 02)
```

### 3. Compare Shipped vs Killed

```bash
ls plans/cycles/archive/*SHIPPED*.md | wc -l  # Shipped count
ls plans/cycles/archive/*KILLED*.md | wc -l   # Killed count

# Example: 6 shipped, 2 killed = 75% ship rate
```

### 4. Identify Patterns

**Questions**:
- Which scopes consistently get cut? (Application Layer cut 3 times)
- Which practices appear in "What Worked"? (TDD, session branches)
- Average session time? (40-45 min typical)
- Discard rate? (20-30% healthy, >40% for killed projects)

---

## Common Pitfalls

### ❌ Pitfall 1: Not Preserving Original

**Symptom**: Archive only has outcome section, loses cycle details

**Problem**: Can't review original hill charts, metrics, decisions

**Fix**: Preserve entire original cycle file below outcome

### ❌ Pitfall 2: Vague "What Worked"

**Bad**: "TDD was good"

**Good**: "TDD with 40-minute sessions sustainable, no burnout"

**Fix**: Be specific (quantify when possible)

### ❌ Pitfall 3: Not Recording Learnings

**Symptom**: Archive just copies cycle file, no reflection

**Problem**: Repeat same mistakes in future cycles

**Fix**: Always add "Learnings for Next Cycle" section

### ❌ Pitfall 4: Deleting Active Before Archiving

**Symptom**: Delete cycle file, lose history

**Problem**: Can't create archive (file gone)

**Fix**: Copy to archive first, THEN delete active

---

## Archive Checklist

After circuit breaker:

- [ ] Copy cycle file to archive directory
- [ ] Rename with -SHIPPED or -KILLED suffix
- [ ] Add outcome section at top
- [ ] Document what worked
- [ ] Document what didn't work
- [ ] Record learnings for next cycle
- [ ] Preserve original cycle file content
- [ ] Delete active cycle file
- [ ] Update plans/cycles/README.md
- [ ] Commit archive
- [ ] Enter cooldown (2 weeks)

---

## Summary

Archiving cycles:
- **Preserves history** (learnings, decisions, velocity)
- **Informs future cycles** (what works, what doesn't)
- **Calculates velocity** (historical averages)
- **Enables reflection** (patterns over time)

**Key insight**: Archives are learning database, not just storage.

---

## Related Guides

- **Before archiving**: [How to handle circuit breaker](circuit-breaker.md)
- **During cycle**: [How to update hill chart](update-hill-chart.md)
- **After cooldown**: [How to plan a cycle](plan-a-cycle.md)

---

**Last Updated**: 2025-11-04
**Related**: [Explanation: Circuit Breaker](../explanation/shape-up-and-sessions.md#circuit-breaker), [Reference: Cycle File Format](../reference/cycle-file-format.md)
