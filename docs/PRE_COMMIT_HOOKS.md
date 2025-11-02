# Pre-Commit Hooks

This project uses [pre-commit](https://pre-commit.com/) to automatically run checks before commits, ensuring code quality and consistency.

## What Gets Checked

The pre-commit hooks run the following checks automatically:

### Go Checks
- **gofmt**: Formats Go code with `gofmt -s -w` (simplify + write)
- **go vet**: Checks for common Go errors
- **go mod tidy**: Ensures `go.mod` and `go.sum` are up to date

### General Checks
- **Trailing whitespace**: Removes trailing whitespace
- **End of file**: Ensures files end with a newline
- **YAML syntax**: Validates YAML files
- **TOML syntax**: Validates TOML files (app manifests)
- **Large files**: Prevents accidental commits of large files (>1MB)
- **Merge conflicts**: Detects unresolved merge conflict markers

## Installation

### Install pre-commit

**macOS/Linux:**
```bash
pip install pre-commit
# or
brew install pre-commit
```

**Windows:**
```bash
pip install pre-commit
```

### Enable hooks for this repository

```bash
pre-commit install
```

This creates git hooks in `.git/hooks/` that run automatically before commits.

## Usage

### Automatic (Recommended)

Once installed, hooks run automatically when you commit:

```bash
git commit -m "feat: add new feature"
```

If any hook fails:
1. The commit is aborted
2. Files are automatically fixed (when possible)
3. You review the changes
4. Stage the fixed files: `git add .`
5. Commit again

### Manual

Run hooks on all files:
```bash
pre-commit run --all-files
```

Run specific hook:
```bash
pre-commit run go-fmt --all-files
```

Update hook versions:
```bash
pre-commit autoupdate
```

## Example Workflow

```bash
# Make changes to Go files
vim cmd/uberman/main.go

# Stage changes
git add cmd/uberman/main.go

# Commit (hooks run automatically)
git commit -m "feat: add version command"

# If gofmt makes changes:
# - Review the formatting changes
# - Stage them: git add cmd/uberman/main.go
# - Commit again: git commit -m "feat: add version command"
```

## Bypassing Hooks

**Not recommended**, but if you need to skip hooks:

```bash
git commit --no-verify -m "emergency fix"
```

Only use this in emergencies (e.g., critical hotfix).

## CI/CD Integration

The release workflow runs the same checks:
- Tests must pass
- Code must be properly formatted
- `go vet` must pass

Pre-commit hooks ensure your code passes these checks **before** pushing to GitHub.

## Troubleshooting

### Hooks not running

```bash
# Reinstall hooks
pre-commit uninstall
pre-commit install
```

### Hook fails with permission error

```bash
# Update pre-commit
pip install --upgrade pre-commit

# Clear cache and reinstall
pre-commit clean
pre-commit install-hooks
```

### Want to disable a specific hook

Edit `.pre-commit-config.yaml` and comment out the hook:

```yaml
# - id: go-vet  # Temporarily disabled
```

## Benefits

✅ **Catch errors early**: Before pushing to GitHub
✅ **Consistent formatting**: All code follows the same style
✅ **Faster CI/CD**: Fewer failures in GitHub Actions
✅ **Better code quality**: Automatic checks enforce best practices
✅ **Time savings**: No need to manually run formatters

## Resources

- [pre-commit documentation](https://pre-commit.com/)
- [Available hooks](https://pre-commit.com/hooks.html)
- [Go hooks](https://github.com/dnephin/pre-commit-golang)
