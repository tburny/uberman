# Session Commands Reference

Quick reference for `/session` slash command syntax and behavior.

**Audience**: Developers looking up command syntax during development.

**See also**:
- [Run a Session](../how-to/run-a-session.md) - Complete workflow guide
- [Session State Schema](session-state-schema.md) - JSON state structure

---

## Command Overview

| Command | Purpose | Creates Branch | Updates State |
|---------|---------|----------------|---------------|
| `/session start <task>` | Begin new session | ✓ | ✓ |
| `/session stop` | End session | - | Deletes |
| `/session pause [reason]` | Pause timer | - | ✓ |
| `/session resume` | Resume timer | - | ✓ |
| `/session status` | Show details | - | - |

---

## `/session start <task>`

Begin a new session, create branch, and start automatic time tracking.

### Syntax

```bash
/session start <task>
```

**Parameters**:
- `<task>`: Brief task description (2-4 words, lowercase, hyphen-separated)

### Behavior

1. Creates branch: `session/YYYYMMDD-scope-task`
2. Initializes session state in `.claude/session-state.json`
3. Starts timer (tracks actual working time)
4. Updates cycle file with active session marker
5. Begins automatic monitoring (auto-pause, reminders)

### Examples

**Starting work on domain layer value object**:
```bash
/session start domain-databasename
```
Branch created: `session/20251104-domain-databasename`

**Starting work on infrastructure adapter**:
```bash
/session start infrastructure-inmemory
```
Branch created: `session/20251104-infrastructure-inmemory`

**Starting work on integration test**:
```bash
/session start test-installation-workflow
```
Branch created: `session/20251104-test-installation-workflow`

### Output

```
✓ Session started: domain-databasename
✓ Branch created: session/20251104-domain-databasename
✓ Timer started: 00:00

Active session:
- Scope: Domain Layer Foundation
- Task: Create DatabaseName value object
- Started: 2025-11-04 10:23

Working on session/20251104-domain-databasename
```

### Errors

**Session already active**:
```
✗ Error: Session already in progress (domain-appname)
  Use `/session stop` first, then start new session.
```

**Invalid task name**:
```
✗ Error: Task name must be lowercase with hyphens (e.g., domain-databasename)
  Got: domain_databaseName
```

---

## `/session stop`

End current session and decide whether to keep or discard work.

### Syntax

```bash
/session stop
```

**Parameters**: None

### Behavior

1. Stops timer, calculates final metrics
2. Prompts: "Keep or discard this session?"
3. **If Keep**:
   - Squash merge branch to main
   - Commit with session metadata (see [Commit Format](commit-format.md))
   - Updates cycle metrics
   - Deletes session state
   - Deletes session branch
4. **If Discard**:
   - Deletes branch without merging
   - Records learning in cycle file
   - Updates cycle metrics (discarded count)
   - Deletes session state

### Examples

**Keeping successful work**:
```bash
/session stop
```

Output:
```
Session complete: domain-databasename
- Working time: 35 min
- Paused time: 2 min (1 auto-pause)
- Total time: 37 min

Keep or discard this session?
[K]eep / [D]iscard: K

✓ Squash merging session/20251104-domain-databasename to main
✓ Committed with session metadata
✓ Updated cycle metrics
✓ Session state cleaned up

Ready for 10-minute break.
```

**Discarding failed experiment**:
```bash
/session stop
```

Output:
```
Session complete: domain-state-machine
- Working time: 52 min
- Paused time: 8 min (2 auto-pauses)
- Total time: 60 min ⚠️

Keep or discard this session?
[K]eep / [D]iscard: D

Reason for discarding (optional): State machine too complex, try simpler approach

✓ Branch deleted (no merge)
✓ Learning recorded in cycle file
✓ Updated cycle metrics
✓ Session state cleaned up

Ready for 10-minute break.
```

### Output (Keep)

Commit message created:
```
feat(domain): add DatabaseName value object

Session: 20251104 10:23-10:58 (35 min actual, 2 min paused)

Created DatabaseName value object with validation:
- Length: 3-64 characters
- Pattern: alphanumeric + underscores
- Prefix: ${USER}_ enforced

Property-based tests verify invariants.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Output (Discard)

Cycle file updated:
```markdown
## Session Metrics

### Week 1
- **Sessions completed**: 5
- **Sessions discarded**: 1 (domain-state-machine: too complex)
```

### Errors

**No active session**:
```
✗ Error: No active session to stop.
```

---

## `/session pause [reason]`

Manually pause the session timer (overrides automatic pause detection).

### Syntax

```bash
/session pause [reason]
```

**Parameters**:
- `[reason]` (optional): Why pausing (e.g., "lunch", "meeting", "research")

### Behavior

1. Pauses timer (working time stops accumulating)
2. Records pause start time and reason
3. Status changes: `IN_PROGRESS` → `PAUSED`
4. Pause type: `MANUAL`

**Note**: Auto-pause still applies (Task agent >10 min, idle >5 min), but manual pause takes precedence.

### Examples

**Taking lunch break**:
```bash
/session pause lunch
```

Output:
```
⏸ Session paused: lunch
- Working time so far: 23 min
- Started: 2025-11-04 10:23
- Paused: 2025-11-04 10:46

Use `/session resume` to continue.
```

**Taking meeting**:
```bash
/session pause meeting
```

**No reason specified**:
```bash
/session pause
```

Output:
```
⏸ Session paused
- Working time so far: 18 min
```

### When to Use

Use manual pause for:
- Scheduled breaks (lunch, meetings)
- Switching contexts temporarily
- Research that doesn't count as session work

**Don't use** for:
- Task agent research (auto-pauses automatically)
- Brief idle periods (auto-pauses automatically)

### Errors

**No active session**:
```
✗ Error: No active session to pause.
```

**Session already paused**:
```
✗ Error: Session already paused.
  Use `/session resume` to continue.
```

---

## `/session resume`

Resume session from paused state.

### Syntax

```bash
/session resume
```

**Parameters**: None

### Behavior

1. Calculates pause duration
2. Records pause duration in state
3. Resumes timer (working time resumes)
4. Status changes: `PAUSED` → `IN_PROGRESS`

### Examples

**Resuming after lunch**:
```bash
/session resume
```

Output:
```
▶ Session resumed
- Paused for: 32 min (lunch)
- Working time so far: 23 min
- Total time: 55 min

Continuing session/20251104-domain-databasename
```

**Resuming after auto-pause**:
```bash
/session resume
```

Output:
```
▶ Session resumed
- Auto-paused for: 12 min (Task agent: codebase exploration)
- Working time so far: 18 min
- Total time: 30 min
```

### Errors

**No active session**:
```
✗ Error: No active session to resume.
```

**Session not paused**:
```
✗ Error: Session is not paused (status: IN_PROGRESS).
```

---

## `/session status`

Display current session details (read-only, no state changes).

### Syntax

```bash
/session status
```

**Parameters**: None

### Behavior

Shows:
- Session ID and branch name
- Scope and task
- Start time
- Current status (IN_PROGRESS or PAUSED)
- Working time, paused time, total time
- Pause history (if any)
- Reminders triggered (if any)

### Examples

**Active session, no pauses**:
```bash
/session status
```

Output:
```
Session: 20251104-domain-databasename
Branch: session/20251104-domain-databasename
Status: IN_PROGRESS ✓

Scope: Domain Layer Foundation
Task: Create DatabaseName value object
Started: 2025-11-04 10:23

Time:
- Working: 18 min
- Paused: 0 min
- Total: 18 min

Reminders: None
```

**Paused session with history**:
```bash
/session status
```

Output:
```
Session: 20251104-infrastructure-filesystem
Branch: session/20251104-infrastructure-filesystem
Status: PAUSED ⏸

Scope: Filesystem Adapter
Task: Implement directory provisioning
Started: 2025-11-04 14:15
Paused: 2025-11-04 15:03

Time:
- Working: 42 min ⚠️ (target: 40 min)
- Paused: 6 min (2 pauses)
- Total: 48 min

Pause history:
1. AUTO_TASK: 12 min (Task agent: codebase exploration)
2. MANUAL: 5 min (ongoing - meeting)

Reminders:
- Break reminder (40 min): Triggered ✓
```

**Over time, timeout triggered**:
```bash
/session status
```

Output:
```
Session: 20251104-cli-install-wiring
Branch: session/20251104-cli-install-wiring
Status: PAUSED ⏸ (TIMEOUT)

Scope: CLI Layer
Task: Wire install command to use cases
Started: 2025-11-04 16:20
Paused: 2025-11-04 17:53 (auto - timeout)

Time:
- Working: 90 min 🚨 (200% over target)
- Paused: 3 min
- Total: 93 min

Reminders:
- Break reminder (40 min): Triggered ✓
- Warning (60 min): Triggered ✓
- Timeout (90 min): Triggered ✓ - Hard timeout enforced

⚠️ Session automatically paused at 90 min working time.
   Take a break before resuming. Consider stopping session.
```

### Errors

**No active session**:
```
No active session.
```

---

## Auto-Pause Behavior

Sessions automatically pause in two scenarios:

### 1. Task Agent Running (>10 minutes)

When Task agent runs for exploratory work:
- Pause type: `AUTO_TASK`
- Reason: "Task agent: [agent description]"
- Working time stops accumulating
- Resumes automatically when agent completes OR manually via `/session resume`

Example:
```
⏸ Auto-pause: Task agent exploring codebase (12 min)
  Working time paused. Will auto-resume when agent completes.
```

### 2. User Idle (>5 minutes)

When >5 minutes pass between messages:
- Pause type: `AUTO_IDLE`
- Reason: "User idle"
- Working time stops accumulating
- Resumes automatically on next message OR manually via `/session resume`

Example:
```
⏸ Auto-pause: User idle for 7 min
  Working time paused. Send a message to auto-resume.
```

**Note**: Auto-pause is transparent - you don't need to manually resume unless you want to.

---

## Reminder System

Progressive reminders enforce healthy session lengths:

| Reminder | Working Time | Behavior | Severity |
|----------|--------------|----------|----------|
| Break | 40 min | Gentle suggestion | ℹ️ Info |
| Warning | 60 min | Strong warning | ⚠️ Warning |
| Timeout | 90 min | Hard pause | 🚨 Critical |

### Break Reminder (40 min)

```
ℹ️ Break reminder: 40 min working time reached (target)
   Consider stopping session with `/session stop`.
   Take a 10-minute break.
```

**Action**: Finish current task, stop session soon.

### Warning (60 min)

```
⚠️ Warning: 60 min working time (50% over target)
   Strongly consider stopping session.
   Prolonged focus increases error risk.
```

**Action**: Stop session within next 10 minutes.

### Timeout (90 min)

```
🚨 Timeout: 90 min working time reached
   Session automatically paused.
   REQUIRED break before resuming.
   Consider discarding if stuck.
```

**Action**: Session paused automatically. Take break. Evaluate if task needs replanning.

---

## State File Location

Session state stored in: `.claude/session-state.json`

**Lifecycle**:
- Created: `/session start`
- Updated: Every message (time tracking, pauses, reminders)
- Deleted: `/session stop`

**Ephemeral**: Never committed to git. See [Session State Schema](session-state-schema.md) for structure.

---

## Related Documentation

- [Run a Session](../how-to/run-a-session.md) - Complete workflow guide
- [Commit Format](commit-format.md) - Session metadata in commits
- [Branch Naming](branch-naming.md) - Session branch conventions
- [Session State Schema](session-state-schema.md) - JSON structure
- [Shape Up and Sessions](../explanation/shape-up-and-sessions.md) - Methodology overview
