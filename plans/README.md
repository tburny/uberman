# Plans Directory - Shape Up Methodology

This directory contains shaped work, active cycles, and betting decisions following Shape Up methodology.

## Directory Structure

```
plans/
├── templates/              # Pitch, hill chart, betting table templates
├── pitches/                # Shaped work ready for betting (YYYY-MM-DD-title.md)
├── cycles/                 # Active cycle tracking (cycle-NN-name.md)
└── cooldown/               # Betting decisions (YYYY-MM-DD-betting-table.md)
```

## Core Principles

### Fixed Time, Variable Scope

- Set appetite FIRST (1-2 weeks or 6 weeks)
- Design solution to fit appetite
- Cut scope aggressively if needed
- **Never extend deadlines by default**

### Cycles and Cooldown

**6-week cycle (big batch)**: Full focus on one major project
**1-2 week cycle (small batch)**: Quick wins, bug fixes, small features
**2-week cooldown**: Bug fixes, exploration, shaping, betting table

### Appetite

**Appetite** = time budget you're willing to spend (NOT an estimate).

- **Small batch**: 1-2 weeks
- **Big batch**: 6 weeks

A creative constraint that forces hard choices.

## Shaping (During Cooldown)

Four steps:
1. **Set Boundaries**: Define appetite and frame problem
2. **Rough Out Elements**: Use breadboarding (workflows) or fat marker sketches (structure)
3. **Address Rabbit Holes**: Identify technical risks, de-risk or exclude
4. **Write Pitch**: Document in `pitches/YYYY-MM-DD-title.md`

## Scope Hammering

Forcefully question every feature to fit time box.

**Techniques**:
- Must-haves vs Nice-to-haves (~)
- Core vs Edge cases (focus on 80%)
- Good enough > perfect

**When**: During shaping (primary) and building (when hitting constraints).

## Circuit Breaker

Hard deadline that kills projects not shipped within one cycle.

**Options**:
1. **Ship What's Done** (most common) - scope hammer, ship working feature
2. **Kill and Reshape** - discard work, learn, reshape if still valuable
3. **Extend** (RARE) - only if work is downhill and nearly complete

## Hill Charts

Visual progress tracking: **Uphill** (figuring out) vs **Downhill** (executing).

```
      /\
     /  \
    /    \
Uphill  Downhill
(Unknown) (Known)
```

**Weekly check-in**: Place each scope on hill.
**Warning sign**: Stuck uphill for 2+ weeks.

## Betting Table (Solo)

Self-reflection during cooldown to decide next cycle.

**Process**:
1. Review shaped pitches
2. Evaluate: problem importance, appetite fit, solution quality, timing
3. Make bet: one big batch (6 weeks) OR 3-4 small batches (1-2 weeks)
4. Let other ideas die

Document in `cooldown/YYYY-MM-DD-betting-table.md`.

## No Backlogs

- Ideas tracked loosely (IDEAS.md)
- Important ideas resurface naturally
- Unimportant ideas die intentionally
- No constant grooming

## Quick Reference

| Question | Answer |
|----------|--------|
| Cycle length? | 6 weeks (big batch) or 1-2 weeks (small batch) |
| Don't finish? | Ship partial or kill (circuit breaker) |
| When to shape? | During 2-week cooldown |
| Need backlog? | No - ideas resurface naturally |
| Track progress? | Hill charts (weekly) |
| Time budget? | Appetite (not an estimate) |
| Cut scope? | During shaping and building |

## Resources

- **Official Book**: https://basecamp.com/shapeup (free)
- **Key Chapters**: 1, 3-6, 8, 11, 14
- **Templates**: See `templates/` directory
