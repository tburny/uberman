# Conventional Commits Quick Reference

This project follows the [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/) specification.

## Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

## Quick Examples

### Simple Feature
```
feat(database): add PostgreSQL support
```

### Bug Fix
```
fix(supervisor): correct service template syntax
```

### Breaking Change (Option 1)
```
feat(cli)!: change install command argument format

BREAKING CHANGE: install command now requires explicit app name
```

### Breaking Change (Option 2)
```
feat(cli): change install command argument format

BREAKING CHANGE: The install command now requires an explicit app name
instead of interactive selection. This improves scriptability but
requires updating existing scripts.

Before: uberman install
After: uberman install wordpress
```

### With Body and Footer
```
feat(backup): implement incremental backup strategy

Incremental backups significantly reduce backup time and storage
requirements by only backing up changed files since the last backup.
Full backups are still created on the first backup of each month.

Closes #45
See also: #42, #38
```

## Commit Types

### Production Code Changes

| Type | Description | SemVer Impact | Example |
|------|-------------|---------------|---------|
| `feat` | New feature | MINOR (0.x.0) | `feat(cli): add list command` |
| `fix` | Bug fix | PATCH (0.0.x) | `fix(database): handle connection timeout` |
| `perf` | Performance improvement | PATCH | `perf(backup): use parallel compression` |

### Development Changes

| Type | Description | Example |
|------|-------------|---------|
| `docs` | Documentation only | `docs: update installation guide` |
| `style` | Code style (no logic change) | `style(config): fix indentation` |
| `refactor` | Code restructuring | `refactor(web): simplify backend logic` |
| `test` | Add or update tests | `test(database): add MySQL integration tests` |
| `build` | Build system or dependencies | `build(deps): update cobra to v1.8.0` |
| `ci` | CI/CD configuration | `ci: add automated testing workflow` |
| `chore` | Other maintenance | `chore: update .gitignore` |
| `revert` | Revert previous commit | `revert: feat(database): add PostgreSQL` |

## Scopes for This Project

| Scope | Description | Example |
|-------|-------------|---------|
| `config` | Configuration parsing | `feat(config): add YAML manifest support` |
| `runtime` | Runtime management | `fix(runtime): correct PHP version detection` |
| `database` | Database operations | `feat(database): add PostgreSQL support` |
| `web` | Web backend config | `fix(web): handle port conflicts` |
| `supervisor` | Service management | `feat(supervisor): add service restart` |
| `appdir` | Directory structure | `refactor(appdir): simplify path handling` |
| `backup` | Backup/restore | `feat(backup): add incremental backups` |
| `deploy` | Deployment | `feat(deploy): add git-based deployment` |
| `cli` | CLI commands/flags | `feat(cli): add verbose flag` |
| `manifest` | App manifests | `feat(manifest): validate dependencies` |
| `templates` | Config templates | `feat(templates): add Django template` |
| `deps` | Dependencies | `build(deps): update all dependencies` |

## Breaking Changes

### Indication Methods

1. **Exclamation mark after type/scope:**
   ```
   feat(api)!: change database name format
   ```

2. **BREAKING CHANGE footer:**
   ```
   feat(api): change database name format

   BREAKING CHANGE: Database names now use underscore separator
   instead of dash. Existing installations must migrate.
   ```

### When to Mark as Breaking

- Changes that require user action
- Changes to public APIs
- Changes to configuration format
- Changes to command arguments/flags
- Removal of features
- Changes to default behavior

## Tools and Automation

### Interactive Commit Helper

Use the provided script for guided commit creation:
```bash
./scripts/commit.sh
```

### Git Commit Template

The repository includes a commit template. Use it with:
```bash
git commit
# Editor opens with template and guidelines
```

### Commitlint (CI/CD)

Pull requests are automatically validated using commitlint. Local validation:
```bash
npm install
npx commitlint --from=HEAD~1
```

## Common Patterns

### Multiple Changes

If your commit includes multiple unrelated changes, split into separate commits:

❌ Bad:
```
feat: add backup and fix database bug
```

✅ Good:
```
# First commit
feat(backup): add incremental backup support

# Second commit
fix(database): handle connection timeout
```

### Squashing Commits

When squashing commits in a PR, use a proper conventional commit message:

```
feat(backup): implement incremental backup strategy

- Add incremental backup logic
- Implement changed file detection
- Add backup metadata storage
- Update backup restoration to handle incremental

Closes #45
```

### Reverting Commits

```
revert: feat(backup): implement incremental backup

This reverts commit abc123def456.

The incremental backup feature caused data corruption in
edge cases and needs to be redesigned.

Refs #47
```

## Message Quality

### Subject Line

✅ Good:
- `feat(cli): add status command for health checks`
- `fix(database): prevent SQL injection in queries`
- `docs: add troubleshooting section to README`

❌ Bad:
- `Added stuff` (no type, vague)
- `feat(cli): Added status command.` (capitalized, period)
- `fix: Fixed a bug` (not descriptive)
- `WIP` (not conventional format)

### Body

Use the body to explain:
- **What** changed (details beyond the subject)
- **Why** the change was necessary
- **How** it solves the problem
- Any side effects or considerations

```
feat(backup): implement incremental backup strategy

Incremental backups reduce backup time by 80% and storage usage
by 60% compared to full backups. This is especially important
for large Nextcloud installations with 100GB+ of data.

Implementation uses file modification timestamps and maintains
a manifest of backed-up files. Full backups are still created
monthly to ensure restore reliability.

Benchmarks:
- Full backup: 45 minutes, 80GB
- Incremental: 9 minutes, 16GB (average)

Closes #45
Related: #42 (backup compression), #38 (backup scheduling)
```

## Integration with Workflow

### Before Committing

1. Stage your changes: `git add <files>`
2. Review changes: `git diff --staged`
3. Use commit helper: `./scripts/commit.sh`
   OR manually craft message following the spec
4. Verify message: Check for typos, length, format

### During Pull Request

- PR title should be a conventional commit
- Commits in PR will be validated by CI
- Squash commits into conventional format before merge

### Release Process

Conventional commits enable automated versioning:
- `feat:` → bump minor version
- `fix:` → bump patch version
- `BREAKING CHANGE:` → bump major version

## Resources

- **Specification**: https://www.conventionalcommits.org/en/v1.0.0/
- **Commitlint**: https://commitlint.js.org/
- **Semantic Versioning**: https://semver.org/
- **Keep a Changelog**: https://keepachangelog.com/

## Tips

1. **Write in imperative mood**: "add feature" not "added feature"
2. **Be specific**: "fix typo in README" not "fix bug"
3. **One logical change per commit**: Easy to review and revert
4. **Reference issues**: Use `Closes #123` or `Refs #456`
5. **Think about changelog**: Your commit will appear there
6. **Use scope consistently**: Pick from defined scopes
7. **When in doubt, use the script**: `./scripts/commit.sh`

## Enforcement

- **Local**: Use commit template and helper script
- **CI**: Commitlint validates all PR commits
- **Review**: Maintainers check commit quality
- **Education**: This guide and CONTRIBUTING.md

---

Questions? See [CONTRIBUTING.md](../CONTRIBUTING.md) or open an issue.
