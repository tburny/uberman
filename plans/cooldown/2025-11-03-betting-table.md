# Betting Table - 2025-11-03 Cooldown

## Context

**Previous Cycle:** N/A (first formal cycle)
**Cooldown Period:** Shaped Clean Architecture refactoring pitch, deleted existing `internal/` code
**Capacity:** 6 weeks (big batch)
**Date:** 2025-11-03

## Pitches Evaluated

### Pitch 1: Rebuild Install Command with Clean Architecture
- **Problem:** Existing codebase violates DDD/Clean Architecture principles - domain layer polluted with I/O
- **Appetite:** 6 weeks
- **Solution:** Delete existing code, rebuild from scratch with single bounded context ("App Installation")
- **Rabbit Holes:** Over-engineering workflow state machine, TOML parsing location, application layer necessity
- **Decision:** ✓ BET
- **Reasoning:**
  - Core functionality needs solid foundation for future features
  - MVP scope (install command only) fits 6 weeks with scope hammering
  - Probe-Sense-Respond approach handles complexity (Cynefin: COMPLEX domain)
  - Circuit breaker acceptable - can ship partial (domain + 1 adapter) or kill and document learnings
  - No backlog debt - clean slate allows correct architecture from start

## Bets Made (Commitments)

### Cycle 1: Clean Architecture Refactoring
- **Commitment:** 6 weeks, solo developer, full focus
- **Scope:** `uberman install <app>` command ONLY (MVP)
  - Must-haves: Domain layer (workflow + manifest aggregates), filesystem adapter, CLI integration
  - Nice-to-haves: In-memory adapter, dry-run adapter, application layer (can be cut)
- **No-gos:** upgrade, backup, restore, deploy, list, status commands (defer to future cycles)
- **Circuit Breaker:**
  - **Ideal:** Ship working install command with all adapters
  - **Acceptable:** Ship domain layer + filesystem adapter (CLI calls domain directly)
  - **Kill:** If domain layer doesn't feel right after week 2, reshape or abandon
- **Start Date:** 2025-11-03
- **End Date:** 2025-12-15 (6 weeks)

## Ideas Not Bet On

None evaluated - focused on single refactoring effort.

## Reflection

**What I learned from code deletion:**
- Existing architecture had no clear boundaries
- Domain logic mixed with infrastructure concerns
- Tests were slow because domain wasn't pure
- TOML parsing in wrong layer (infrastructure instead of domain)

**Shaping quality:**
- Pitch well-shaped with clear rabbit holes identified
- MVP scope explicit (install only)
- Probe-Sense-Respond approach appropriate for COMPLEX domain
- Scope hammer applied: cut all non-install functionality

**Appetite accuracy:**
- 6 weeks feels right for foundational refactoring
- Risk: state machine might be too complex (TBD week 1-2)
- Mitigation: Can simplify to procedural workflow if needed

**Circuit breaker effectiveness:**
- Clear kill criteria: domain layer feels wrong after 2 weeks
- Ship partial acceptable: domain + filesystem adapter + CLI
- No shame in killing: would document learnings in ADRs

**Build order (Probe-Sense-Respond):**
1. **Weeks 1-2 (PROBE):** Build domain layer, validate approach
2. **End of Week 2 (SENSE):** Evaluate state machine complexity, test speed, domain purity
3. **Weeks 3-4 (RESPOND):** Build infrastructure based on learnings
4. **Week 5 (PROBE):** Build application layer (or skip if not needed)
5. **Week 6 (RESPOND):** CLI integration, ship working command

---

**Next Betting Table:** 2025-12-22 (after 6-week cycle + 1-week cooldown)
**Next Cooldown:** 2025-12-16 to 2025-12-22
