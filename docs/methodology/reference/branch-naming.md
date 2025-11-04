# Session Branch Naming Reference

Branch naming conventions for ephemeral session branches.

**Audience**: Developers looking up branch naming rules or cleanup procedures.

**See also**:
- [Session Commands](session-commands.md) - /session start creates branches
- [Run a Session](../how-to/run-a-session.md) - Complete workflow

---

## Convention

```
session/YYYYMMDD-scope-task
```

### Components

| Component | Description | Format | Example |
|-----------|-------------|--------|---------|
| `session/` | Fixed prefix | Literal | `session/` |
| `YYYYMMDD` | Session start date | ISO 8601 compact | `20251104` |
| `scope` | Hill chart scope/layer | Lowercase, hyphenated | `domain`, `infrastructure` |
| `task` | Specific work | Lowercase, hyphenated, 2-3 words | `databasename`, `port-validation` |

### Format Rules

1. **Lowercase only** - No uppercase letters
2. **Hyphens separate words** - Use `-` not `_` or spaces
3. **No special characters** - Alphanumeric and hyphens only
4. **Task brevity** - Maximum 3 words for task component
5. **Date format** - Exactly 8 digits (YYYYMMDD)

---

## Examples

### Domain Layer Work

**Value object creation**:
```
session/20251104-domain-databasename
session/20251104-domain-appname
session/20251104-domain-directorypath
```

**Aggregate implementation**:
```
session/20251105-domain-installation-aggregate
session/20251106-domain-manifest-aggregate
```

**Domain service**:
```
session/20251107-domain-prerequisite-validator
```

### Infrastructure Layer Work

**Repository adapter**:
```
session/20251108-infrastructure-inmemory-adapter
session/20251109-infrastructure-filesystem-adapter
```

**Specific adapter functionality**:
```
session/20251110-infrastructure-directory-provisioning
session/20251111-infrastructure-database-creation
```

**Dry-run adapter**:
```
session/20251112-infrastructure-dryrun-adapter
```

### Application Layer Work

**Use case**:
```
session/20251113-application-install-usecase
```

**Service**:
```
session/20251114-application-validation-service
```

### CLI Layer Work

**Command implementation**:
```
session/20251115-cli-install-command
session/20251116-cli-flags-parsing
```

**Wiring/integration**:
```
session/20251117-cli-install-wiring
session/20251118-cli-dependency-injection
```

### Testing Work

**Unit tests**:
```
session/20251119-test-databasename-unit
session/20251120-test-installation-unit
```

**Property-based tests**:
```
session/20251121-test-appname-properties
session/20251122-test-port-properties
```

**Integration tests**:
```
session/20251123-test-filesystem-integration
session/20251124-test-installation-workflow
```

### Documentation Work

**Methodology docs**:
```
session/20251125-docs-session-commands
session/20251126-docs-commit-format
```

**User-facing docs**:
```
session/20251127-docs-installation-guide
session/20251128-docs-cli-reference
```

---

## Rationale

### Date Component (YYYYMMDD)

**Purpose**: Indicates branch staleness at a glance.

**Benefits**:
- Quickly identify old branches
- Automated cleanup based on age
- Debugging: correlate with commit history
- Metrics: sessions per day/week

**Example**: `session/20251001-domain-task` in December 2025 → Obviously stale (2 months old)

### Scope Component

**Purpose**: Links session to hill chart scope or architecture layer.

**Benefits**:
- Know which scope/layer work belongs to
- Group sessions by scope for metrics
- Understand architecture location at glance

**Common scopes**:
- Architecture layers: `domain`, `application`, `infrastructure`, `cli`
- Cross-cutting: `test`, `docs`, `build`

### Task Component

**Purpose**: Brief description of specific work.

**Benefits**:
- Understand what session does
- Avoid duplicate sessions
- Readable in git logs

**Keep short**: 2-3 words maximum
- ✅ `databasename` (value object name)
- ✅ `port-validation` (what + where)
- ✅ `install-wiring` (action + target)
- ❌ `create-database-name-value-object-with-validation` (too verbose)

---

## Branch Lifecycle

### Creation

**Trigger**: `/session start <task>`

**Process**:
1. Generate branch name: `session/YYYYMMDD-scope-task`
2. Create branch from current HEAD (usually `main`)
3. Checkout new branch
4. Begin session work

**Example**:
```bash
/session start domain-databasename
```

Creates and checks out: `session/20251104-domain-databasename`

### Active Work

**All commits** happen on session branch during session.

**Freedom**: Experiment freely, branch is ephemeral.

**Examples**:
```bash
# Multiple commits allowed on session branch
git commit -m "wip: add validation tests"
git commit -m "wip: implement validation logic"
git commit -m "wip: refactor for clarity"
```

**Note**: Commit messages during session can be informal ("wip:", "temp:", etc.) - squashed on merge.

### Completion (Keep)

**Trigger**: `/session stop` → Keep

**Process**:
1. Squash merge to `main`
2. Single commit with session metadata (see [Commit Format](commit-format.md))
3. Delete session branch
4. Cleanup session state

**Example**:
```bash
# After /session stop → Keep
# Squash merged: session/20251104-domain-databasename
# Result: One commit on main with proper format
# Branch deleted
```

**Git log shows**:
```
feat(domain): add DatabaseName value object
Session: 20251104 10:23-10:58 (35 min actual, 2 min paused)
...
```

### Completion (Discard)

**Trigger**: `/session stop` → Discard

**Process**:
1. Record learning in cycle file
2. Delete session branch (no merge)
3. Cleanup session state
4. Return to `main`

**Example**:
```bash
# After /session stop → Discard
# Branch deleted: session/20251104-domain-state-machine
# No commits merged to main
# Learning recorded in cycle file
```

**Git history**: No trace of session branch (intentional).

---

## Branch Management

### List Session Branches

**All session branches**:
```bash
git branch | grep 'session/'
```

Output:
```
  session/20251104-domain-databasename
  session/20251105-infrastructure-inmemory
* session/20251106-test-installation
```

**Count active sessions**:
```bash
git branch | grep -c 'session/'
```

Output: `3`

**Should be**: 0-1 (only current session, if any)

### Identify Stale Branches

**Manual inspection**:
```bash
git branch | grep 'session/' | sort
```

Look for dates >7 days old.

**Automated check** (requires date parsing):
```bash
#!/bin/bash
# find-stale-sessions.sh

TODAY=$(date +%Y%m%d)
THRESHOLD=7  # days

git branch | grep 'session/' | while read -r branch; do
  # Extract date from branch name
  branch_date=$(echo "$branch" | grep -oP '\d{8}')

  if [ -n "$branch_date" ]; then
    # Calculate age (simplified - doesn't handle month boundaries)
    age=$((TODAY - branch_date))

    if [ $age -gt $THRESHOLD ]; then
      echo "Stale: $branch (age: $age days)"
    fi
  fi
done
```

**Usage**:
```bash
bash find-stale-sessions.sh
```

Output:
```
Stale: session/20251028-domain-task (age: 10 days)
Stale: session/20251030-infrastructure-test (age: 8 days)
```

### Cleanup Stale Branches

**Manual deletion** (safe, one-by-one):
```bash
git branch -D session/20251028-domain-task
```

**Bulk deletion** (requires confirmation):
```bash
# Delete all session branches older than 7 days
git branch | grep 'session/' | while read -r branch; do
  branch_date=$(echo "$branch" | grep -oP '\d{8}')
  age=$(($(date +%Y%m%d) - branch_date))

  if [ $age -gt 7 ]; then
    echo "Delete $branch? (y/n)"
    read -r confirm
    if [ "$confirm" = "y" ]; then
      git branch -D "$branch"
    fi
  fi
done
```

**Warning**: Use `-D` (force delete) only if certain branch is abandoned. Use `-d` to fail if branch has unmerged work.

---

## Troubleshooting

### Branch Name Validation Errors

**Error**: "Invalid task name"

**Cause**: Task name contains uppercase, underscores, or special characters.

**Fix**:
```bash
# ❌ Invalid
/session start Domain-DatabaseName
/session start domain_database_name
/session start domain/database-name

# ✅ Valid
/session start domain-databasename
/session start domain-database-name
```

### Stale Branch Warning

**Scenario**: Starting session when old session branch exists.

**Warning**:
```
⚠️ Warning: Stale session branch detected
  session/20251028-domain-task (10 days old)

Delete stale branch? [y/n]
```

**Action**:
- `y` - Delete stale branch, proceed with new session
- `n` - Keep stale branch, abort new session start

### Multiple Active Sessions

**Scenario**: Interrupted session, started new one.

**Problem**:
```
✗ Error: Session already in progress (domain-databasename)
  Use `/session stop` first, then start new session.
```

**Fix**:
1. Stop current session: `/session stop`
2. Choose: Keep or discard
3. Start new session: `/session start <task>`

### Branch Name Conflicts

**Scenario**: Two sessions same day, same task.

**Problem**:
```
✗ Error: Branch already exists: session/20251104-domain-databasename
```

**Fix**: Make task name more specific:
```bash
# First session
/session start domain-databasename

# Second session (more specific)
/session start domain-databasename-refactor
/session start domain-databasename-tests
```

---

## Naming Anti-Patterns

### Too Verbose

❌ **Bad**:
```
session/20251104-domain-create-database-name-value-object-with-full-validation-and-property-tests
```

✅ **Good**:
```
session/20251104-domain-databasename
```

**Reason**: Task name should be brief. Details go in commit message, not branch name.

### Too Generic

❌ **Bad**:
```
session/20251104-domain-work
session/20251105-infrastructure-stuff
session/20251106-test-tests
```

✅ **Good**:
```
session/20251104-domain-databasename
session/20251105-infrastructure-inmemory
session/20251106-test-installation
```

**Reason**: Generic names lose information. Be specific about what you're working on.

### Wrong Separators

❌ **Bad**:
```
session/20251104_domain_databasename    # Underscores
session/20251104-Domain-DatabaseName    # CamelCase
session/20251104-domain/databasename    # Slash in task
```

✅ **Good**:
```
session/20251104-domain-databasename    # Hyphens only
```

**Reason**: Consistency. Always lowercase, always hyphens.

### Missing Components

❌ **Bad**:
```
session/domain-databasename             # Missing date
session/20251104-databasename           # Missing scope
session/20251104-domain                 # Missing task
```

✅ **Good**:
```
session/20251104-domain-databasename    # All three components
```

**Reason**: Each component serves a purpose. All three required.

---

## Integration with Tooling

### Git Aliases

Add to `.gitconfig`:

```ini
[alias]
  # List session branches
  sessions = !git branch | grep 'session/'

  # Count session branches
  session-count = !git branch | grep -c 'session/' || echo 0

  # Show session with commit count
  session-status = !git branch -v | grep 'session/'
```

**Usage**:
```bash
git sessions           # List all session branches
git session-count      # Count active sessions
git session-status     # Show branches with last commit
```

### Shell Prompt Integration

Show current session branch in shell prompt:

**Bash** (`.bashrc`):
```bash
# Add to PS1
parse_git_branch() {
  git branch 2>/dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/'
}

PS1='[\u@\h \W$(parse_git_branch)]\$ '
```

**Zsh** (`.zshrc`):
```zsh
# Enable git prompt
autoload -Uz vcs_info
precmd() { vcs_info }
setopt prompt_subst
PROMPT='%F{green}%n@%m%f %F{blue}%~%f %F{red}${vcs_info_msg_0_}%f$ '
```

**Result**:
```
user@host ~/uberman session/20251104-domain-databasename$
```

### Editor Integration

**VS Code** (shows branch in status bar):
```json
{
  "git.decorations.enabled": true,
  "git.showBranchName": true
}
```

**Vim** (vim-fugitive):
```vim
" Shows branch in statusline
set statusline+=%{FugitiveStatusline()}
```

---

## Related Documentation

- [Session Commands](session-commands.md) - /session start creates branches
- [Commit Format](commit-format.md) - Session metadata in commits
- [Run a Session](../how-to/run-a-session.md) - Complete workflow
- [Shape Up and Sessions](../explanation/shape-up-and-sessions.md) - Methodology overview
