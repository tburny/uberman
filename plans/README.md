# Plans Directory - Shape Up Methodology

This directory contains shaped work, active cycles, and betting decisions following the Shape Up methodology by Basecamp.

## Directory Structure

```
plans/
├── README.md               # This file - Shape Up guide
├── templates/              # Reusable templates
│   ├── pitch-template.md
│   ├── hill-chart-template.md
│   └── betting-table-template.md
├── pitches/                # Shaped work ready for betting
│   └── YYYY-MM-DD-title.md
├── cycles/                 # Active cycle tracking
│   └── cycle-NN-name.md
└── cooldown/               # Cooldown notes and decisions
    └── YYYY-MM-DD-betting-table.md
```

## Shape Up Core Principles

### Fixed Time, Variable Scope

**NOT** fixed scope with variable time (traditional waterfall/agile).
**YES** fixed time budget with flexible scope.

- Set appetite FIRST (1-2 weeks or 6 weeks)
- Design solution to fit appetite
- Cut scope aggressively if needed
- Never extend deadlines by default

### Cycles and Cooldown

**6-week cycle (big batch):**
- Full focus on one major project
- No context switching
- No other commitments

**1-2 week cycle (small batch):**
- Quick wins, bug fixes, small features
- Can run multiple in parallel

**2-week cooldown after each cycle:**
- Fix bugs from previous cycle
- Explore new ideas
- Shape work for next cycle
- Betting table decision
- NO commitments - self-directed work

### Appetite Constraints

**Appetite** = time budget you're willing to spend (set BEFORE shaping).

- **Small batch**: 1-2 weeks (tweaks, minor features, small refactoring)
- **Big batch**: 6 weeks (major features, large refactoring)

**NOT an estimate** - it's a creative constraint that forces hard choices.

## Shaping Phase (During Cooldown)

Transform raw ideas into pitches ready for betting. Four steps:

### 1. Set Boundaries

Define appetite (time budget) and frame the problem (not the solution).

### 2. Rough Out Elements

Use two techniques:

**Breadboarding** (for workflows):
```
[Place 1] → (Action) → [Place 2] → (Action) → [Place 3]
```

**Fat Marker Sketches** (for structure):
- Use thick marker to prevent detail
- Show arrangement, not specifics
- Leave room for implementation creativity

### 3. Address Risks and Rabbit Holes

- **Rabbit holes**: Technical unknowns that could derail the project
- **De-risking**: Spike, prototype, or explicitly exclude risky areas
- **Mark as "out of scope"** if too risky for appetite

### 4. Write the Pitch

Document shaped work in `pitches/YYYY-MM-DD-title.md` using template.

## Scope Hammering

**Forcefully question** every feature to fit within time box.

**When**: During shaping (primary) and during building (when hitting constraints).

**Techniques**:
- Must-haves vs Nice-to-haves (~)
- Core vs Edge cases (focus on 80%)
- Question everything: "Do we really need this?"
- Good enough > perfect

**Cut nice-to-haves first** when time runs short.

## Circuit Breaker

**Hard deadline** that kills projects that don't ship within one cycle. No extensions by default.

### When Work Takes Longer

**Option 1: Ship What's Done** (most common)
- Use scope hammering to cut non-essentials
- Ship working (even if reduced) feature

**Option 2: Kill and Reshape**
- Circuit breaker activates
- Project ends, incomplete work discarded
- Learn from failure
- Reshape for next cycle if still valuable

**Option 3: Extend** (RARE)
Only when:
- All outstanding work is "downhill" (no unknowns)
- All tasks are true must-haves
- Work is nearly complete (1-2 days remaining)

## Hill Charts

Visual progress tracking showing uncertainty vs certainty.

```
      /\
     /  \
    /    \
   /      \
  /        \
Uphill   Downhill
(Unknown) (Known)
```

**Uphill (Figuring things out)**:
- Research, prototyping, unknowns
- Risky phase - might get stuck

**Downhill (Executing known work)**:
- Implementation, testing, polish
- Low risk - just execution

**Weekly check-in**:
- List scopes (major pieces of work)
- Place each on hill
- When to worry: stuck uphill for 2+ weeks

## Betting Table (Solo Adaptation)

Self-reflection session during cooldown to decide what to work on next cycle.

### Process

1. **Review shaped pitches** from cooldown
2. **Evaluate each pitch**:
   - Does the problem matter enough?
   - Is the appetite right?
   - Is the solution compelling?
   - Is timing good?
   - Do I have capacity?
3. **Make bets (commitments)**:
   - One big batch (6 weeks), OR
   - 3-4 small batches (1-2 weeks each)
4. **Let other ideas die**:
   - If not bet on, idea is discarded
   - Important ideas will resurface naturally

Document decision in `cooldown/YYYY-MM-DD-betting-table.md`.

## No Backlogs

Shape Up avoids backlogs. Instead:
- Ideas tracked individually (personal notes, IDEAS.md)
- Important ideas resurface naturally
- Unimportant ideas die on the vine (intentionally)
- No constant grooming (time waster)

## Integration with This Project

### Clean Architecture Refactoring

**Appetite**: 6 weeks (big batch)
**Scope**: Install command only (MVP)
**Circuit Breaker**: Ship partial or kill if not done in 6 weeks

### With 40-Minute Sessions

- Each session: focused work on cycle commitment
- Don't plan sessions individually
- Track progress weekly with hill charts
- Use sessions for both uphill (discovery) and downhill (execution)

### With TDD

- TDD fits naturally within Shape Up
- Red-Green-Refactor within sessions
- Tests help stay downhill (known execution)

### With Probe-Sense-Respond (Cynefin Complex)

- **Probe**: Uphill phase (exploring unknowns)
- **Sense**: Hill peak (understanding emerges)
- **Respond**: Downhill phase (executing known solution)

## Resources

- **Official Book**: https://basecamp.com/shapeup (free)
- **Key Chapters**: 1, 3-6, 8, 11, 14
- **Templates**: See `templates/` directory

## Quick Reference

| Question | Answer |
|----------|--------|
| How long should a cycle be? | 6 weeks (big batch) or 1-2 weeks (small batch) |
| What if I don't finish? | Circuit breaker: ship partial or kill (no extensions) |
| When do I shape work? | During 2-week cooldown after each cycle |
| Do I need a backlog? | No - ideas tracked loosely, resurface naturally |
| How do I track progress? | Hill charts (weekly updates) |
| What's the time budget? | Appetite (set before shaping, not an estimate) |
| When do I cut scope? | During shaping (primary) and building (if needed) |
