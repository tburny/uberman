# Development Methodology

Documentation for Uberman's development process following the [Diataxis](https://diataxis.fr/) framework.

**Audience**: Contributors, developers, future maintainers

**Separation**: This directory documents HOW WE BUILD uberman. For user-facing product documentation, see parent `docs/` directory.

---

## Documentation Types

### Explanation (Understanding)
Understanding-oriented documentation explains concepts, context, and philosophy.

- [Shape Up and Session Management](explanation/shape-up-and-sessions.md) - Complete development methodology

### How-To Guides (Tasks)
Task-oriented guides for accomplishing specific development tasks.

- [Plan a Cycle](how-to/plan-a-cycle.md) - Shaping, pitching, betting
- [Run a Session](how-to/run-a-session.md) - Start/work/stop workflow
- [Update Hill Chart](how-to/update-hill-chart.md) - Weekly review ritual
- [Scope Hammer](how-to/scope-hammer.md) - Cut features during cycle
- [Circuit Breaker](how-to/circuit-breaker.md) - End cycle decisions
- [Archive Cycle](how-to/archive-cycle.md) - Post-mortem and archival

### Reference (Information)
Information-oriented documentation for looking up syntax, formats, and schemas.

- [Session Commands](reference/session-commands.md) - /session command syntax
- [Commit Format](reference/commit-format.md) - Git commit message template
- [Cycle File Format](reference/cycle-file-format.md) - Cycle file structure
- [Session State Schema](reference/session-state-schema.md) - JSON schema
- [Branch Naming](reference/branch-naming.md) - Session branch conventions

### Tutorials (Learning)
Learning-oriented step-by-step lessons for new contributors.

- [First Cycle Walkthrough](tutorial/first-cycle-walkthrough.md) - Complete cycle from start to finish

---

## Quick Start

**New contributor?** Start here:
1. Read [Shape Up and Session Management](explanation/shape-up-and-sessions.md) (understanding WHY)
2. Follow [First Cycle Walkthrough](tutorial/first-cycle-walkthrough.md) (hands-on learning)
3. Reference [How-To Guides](how-to/) as needed (doing specific tasks)
4. Look up syntax in [Reference](reference/) (quick lookup)

---

## Related Documentation

**Technical Architecture**:
- [ARCHITECTURE.md](../../ARCHITECTURE.md) - Clean Architecture design
- [UBIQUITOUS_LANGUAGE.md](../../UBIQUITOUS_LANGUAGE.md) - Domain glossary
- [PRD.md](../../PRD.md) - Product requirements (EARS format)

**Planning**:
- [plans/](../../plans/) - Pitches, cycles, cooldown decisions
- [CLAUDE.md](../../CLAUDE.md) - AI assistant guidance

**User-Facing Docs**:
- [docs/](../) - Installation, usage, CLI reference (for end users)
