# Commit Message Format Reference

Commit message format for session work, following Conventional Commits 1.0.

**Audience**: Developers committing session work or querying commit history.

**See also**:
- [Session Commands](session-commands.md) - /session command reference
- [Run a Session](../how-to/run-a-session.md) - Complete workflow

---

## Format Template

```
<type>(<scope>): <description>

Session: YYYYMMDD HH:MM-HH:MM (NN min actual, NN min paused)

[Optional body - discoveries, decisions, learnings]

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Components

**Header** (required):
- `<type>`: Commit type (see below)
- `<scope>`: Architecture layer or area
- `<description>`: Imperative mood, lowercase, no period

**Session Line** (required for session work):
- Date: `YYYYMMDD` (session start date)
- Time range: `HH:MM-HH:MM` (start-end in 24h format)
- Actual time: `NN min actual` (working time, excludes pauses)
- Paused time: `NN min paused` (total pause duration)

**Body** (optional):
- Multi-line explanation
- Discoveries made during session
- Decisions and rationale
- Learnings (especially for complex work)

**Footer** (always):
- Claude Code attribution
- Co-authored-by Claude

---

## Commit Types

Based on Conventional Commits 1.0:

| Type | Purpose | Example |
|------|---------|---------|
| `feat` | New feature | Add DatabaseName value object |
| `fix` | Bug fix | Fix port validation off-by-one |
| `docs` | Documentation only | Add session command reference |
| `refactor` | Code restructuring | Extract validation to helper |
| `test` | Test additions | Add property tests for AppName |
| `chore` | Maintenance | Update dependencies |

**Breaking changes**: Add `!` after type (e.g., `feat(domain)!: change Installation interface`)

---

## Scope Values

**Architecture layers** (Clean Architecture):

- `domain` - Domain layer (entities, value objects, aggregates)
- `application` - Application layer (use cases, services)
- `infrastructure` - Infrastructure layer (adapters, repositories)
- `cli` - CLI layer (commands, flags, output)

**Other areas**:

- `build` - Build system, dependencies
- `ci` - CI/CD configuration
- `docs` - Documentation (methodology or user-facing)

**Multiple scopes**: Use primary layer only (e.g., `domain` even if test touches infrastructure)

---

## Examples

### Value Object Creation

```
feat(domain): add DatabaseName value object

Session: 20251104 10:23-10:58 (35 min actual, 2 min paused)

Created DatabaseName value object with validation:
- Length: 3-64 characters
- Pattern: alphanumeric + underscores
- Prefix: ${USER}_ enforced

Property-based tests verify invariants using rapid framework.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `feat` (new capability)
- Scope: `domain` (domain layer)
- Session: 35 min working (target met), 2 min paused (within limits)
- Body: Details validation rules and testing approach

### Bug Fix

```
fix(infrastructure): correct directory permission handling

Session: 20251105 14:15-14:52 (37 min actual, 0 min paused)

Fixed directory creation to use 0755 permissions instead of 0644.
Directories need execute bit for traversal.

Discovery: os.MkdirAll defaults to exact permissions provided,
unlike touch which applies umask.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `fix` (corrects defect)
- Scope: `infrastructure` (adapter layer)
- Body: Includes discovery made during debugging

### Integration Test

```
test(infrastructure): add filesystem adapter integration tests

Session: 20251106 09:30-10:18 (48 min actual, 5 min paused)

Added testcontainers-based integration tests:
- Directory provisioning with actual filesystem
- Permission verification
- Cleanup on test completion

Tests run in Docker container for isolation.
Working time exceeded target due to testcontainers setup learning curve.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `test` (adds tests)
- Scope: `infrastructure`
- Working time: 48 min (8 min over target, acknowledged in body)
- Body: Explains test approach and time variance

### Refactoring

```
refactor(domain): extract validation to shared helpers

Session: 20251107 11:05-11:37 (32 min actual, 3 min paused)

Extracted common validation patterns into domain/validation package:
- AlphanumericUnderscore
- LengthBetween
- PrefixEnforced

Reduces duplication across value objects.
No behavior changes - pure refactoring.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `refactor` (restructure without behavior change)
- Body: Explicitly states "no behavior changes"

### Documentation

```
docs(methodology): add session command reference

Session: 20251108 15:20-15:58 (38 min actual, 1 min paused)

Created complete /session command reference:
- All 5 commands with syntax
- Auto-pause behavior
- Reminder system
- Examples for each command

Part of Diataxis reference documentation.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `docs`
- Scope: `methodology` (process docs, not user-facing)

### Multiple Concerns (Breaking Change)

```
feat(domain)!: change Installation aggregate interface

Session: 20251109 13:45-14:32 (47 min actual, 8 min paused)

BREAKING CHANGE: Installation.Execute() now returns (InstallationResult, error)
instead of error only.

Added InstallationResult value object to capture:
- Success status
- Installation path
- Provisioned resources list

Requires update to all adapters and tests.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

**Explanation**:
- Type: `feat` with `!` (breaking change marker)
- Body: `BREAKING CHANGE:` footer (required for breaking changes)
- Working time: 47 min (7 min over due to API design complexity)

### Discarded Session (Learning Capture)

When session is discarded, update cycle file instead of creating commit:

**Cycle file update**:
```markdown
## Session Metrics

### Week 1
- **Sessions discarded**: 1

**Discarded sessions**:
1. 20251110-domain-state-machine: 52 min (too complex)
   - Learning: Explicit state machine adds cognitive overhead
   - Alternative: Try procedural approach with clear phases
   - Decision: Simplify Installation aggregate
```

**No commit** - branch deleted without merge.

---

## Session Line Details

### Date Format

`YYYYMMDD` - ISO 8601 compact date

**Examples**:
- 2025-11-04 → `20251104`
- 2025-12-25 → `20251225`

### Time Range Format

`HH:MM-HH:MM` - 24-hour format, start to end

**Examples**:
- Started 10:23, ended 10:58 → `10:23-10:58`
- Started 14:15, ended 15:03 → `14:15-15:03`

**Note**: Use session start/end times, not clock time (excludes pauses).

### Actual Time

`NN min actual` - Working time only (excludes pauses)

**Calculation**:
```
actual = (endTime - startTime) - totalPausedTime
```

**Examples**:
- Started 10:00, ended 10:40, no pauses → `40 min actual`
- Started 10:00, ended 10:50, paused 8 min → `42 min actual`

**Target**: 40 minutes

### Paused Time

`NN min paused` - Total duration of all pauses

**Includes**:
- Auto-pause for Task agent (>10 min)
- Auto-pause for idle (>5 min)
- Manual pause (`/session pause`)

**Examples**:
- No pauses → `0 min paused`
- Single 12-min Task agent pause → `12 min paused`
- Two pauses (5 min + 3 min) → `8 min paused`

---

## Body Guidelines

### When to Include Body

**Always include** for:
- Complex implementations (>50 LOC changed)
- Design decisions (why this approach?)
- Discoveries (unexpected learnings)
- Breaking changes (impact explanation)
- Time variances (why over 40 min?)

**Optional** for:
- Simple value objects
- Straightforward bug fixes
- Trivial refactorings

### Body Content

**Good body content**:
- **Discoveries**: "Learned that os.MkdirAll doesn't apply umask"
- **Decisions**: "Chose rapid over gopter for simpler API"
- **Trade-offs**: "Sacrificed performance for readability in validation"
- **Warnings**: "This approach works but may need revisiting for multiple databases"

**Avoid**:
- Repeating what code shows ("Added field X, added method Y")
- Implementation minutiae ("Changed i to index in loop")
- Future planning ("TODO: Add caching later")

---

## Querying Session History

### List All Sessions

```bash
git log --grep="Session:" --oneline
```

Output:
```
a1b2c3d feat(domain): add DatabaseName value object
e4f5g6h fix(infrastructure): correct directory permissions
i7j8k9l test(infrastructure): add filesystem integration tests
```

### Show Full Session Details

```bash
git log --grep="Session:" --format="%h %s%n%b" -n 5
```

Output:
```
a1b2c3d feat(domain): add DatabaseName value object
Session: 20251104 10:23-10:58 (35 min actual, 2 min paused)

Created DatabaseName value object with validation...
```

### Count Sessions This Week

```bash
git log --grep="Session:" --since="1 week ago" --oneline | wc -l
```

Output: `12` (sessions completed this week)

### Extract Session Metrics

Bash script to calculate average session duration:

```bash
#!/bin/bash
# extract-session-metrics.sh

git log --grep="Session:" --format="%b" | \
  grep -oP '\d+ min actual' | \
  grep -oP '\d+' | \
  awk '{sum+=$1; count++} END {print "Average:", sum/count, "min"}'
```

Output: `Average: 38.5 min`

### Sessions by Scope

```bash
git log --grep="Session:" --format="%s" | \
  grep -oP '\(\w+\)' | \
  sort | uniq -c | sort -rn
```

Output:
```
  8 (domain)
  4 (infrastructure)
  3 (test)
  2 (cli)
  1 (docs)
```

### Sessions Over Target

```bash
git log --grep="Session:" --format="%h %s%n%b" | \
  grep -B1 "actual" | \
  awk '/min actual/ {if ($1 > 40) print}'
```

Shows sessions exceeding 40 min working time.

---

## Automation

### Git Alias for Session Log

Add to `.gitconfig`:

```ini
[alias]
  sessions = log --grep="Session:" --format="%C(yellow)%h%Creset %s%n%C(dim)%b%Creset" --reverse
  session-stats = !git log --grep='Session:' --format='%b' | grep -oP '\\d+ min actual' | grep -oP '\\d+' | awk '{sum+=$1; count++} END {print \"Total:\", count, \"sessions, Avg:\", sum/count, \"min\"}'
```

**Usage**:
```bash
git sessions              # List all sessions
git session-stats         # Calculate statistics
```

### Commit Message Template

Create `.git/commit-template.txt`:

```
<type>(<scope>): <description>

Session: YYYYMMDD HH:MM-HH:MM (NN min actual, NN min paused)



🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

Configure git:
```bash
git config commit.template .git/commit-template.txt
```

**Note**: `/session stop` automates this - template useful for manual commits only.

---

## Commit Message Validation

### Pre-commit Hook

Create `.git/hooks/commit-msg`:

```bash
#!/bin/bash
# Validate commit message format

commit_msg=$(cat "$1")

# Check for session line
if ! echo "$commit_msg" | grep -qP "^Session: \d{8} \d{2}:\d{2}-\d{2}:\d{2} \(\d+ min actual, \d+ min paused\)"; then
  echo "Error: Missing or invalid Session line"
  echo "Format: Session: YYYYMMDD HH:MM-HH:MM (NN min actual, NN min paused)"
  exit 1
fi

# Check for conventional commit header
if ! echo "$commit_msg" | head -n1 | grep -qP "^(feat|fix|docs|refactor|test|chore)(\(.+\))?!?: .+"; then
  echo "Error: Invalid commit header"
  echo "Format: <type>(<scope>): <description>"
  exit 1
fi

exit 0
```

Make executable:
```bash
chmod +x .git/hooks/commit-msg
```

**Note**: Hook only applies to manual commits, not `/session stop` automation.

---

## Related Documentation

- [Session Commands](session-commands.md) - /session command syntax
- [Run a Session](../how-to/run-a-session.md) - Complete workflow
- [Branch Naming](branch-naming.md) - Session branch conventions
- [Shape Up and Sessions](../explanation/shape-up-and-sessions.md) - Methodology overview
