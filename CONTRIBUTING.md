# Contributing to Uberman

Thank you for your interest in contributing to Uberman! This document provides guidelines for contributing to the project.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version, Uberspace details if applicable)
- Relevant logs or error messages

### Suggesting Features

Feature suggestions are welcome! Please open an issue describing:
- The problem you're trying to solve
- Your proposed solution
- Any alternative solutions you've considered
- How this fits with Uberman's goals

### Adding New App Manifests

We encourage contributions of new app manifests! To add a new app:

1. **Create the manifest**:
   ```bash
   # Create a new manifest file
   cp apps/examples/wordpress.toml apps/examples/myapp.toml
   # Edit with your app's configuration
   nano apps/examples/myapp.toml
   ```

2. **Test thoroughly**:
   - Test installation on a real Uberspace instance
   - Verify all services start correctly
   - Test the backup/restore functionality
   - Document any manual steps required

3. **Add configuration templates** (if needed):
   - If your app requires custom config files, add templates to `internal/templates/`
   - Document template variables

4. **Update documentation**:
   - Add your app to the README.md supported apps list
   - Document any special requirements or post-install steps
   - Include example usage

5. **Submit a pull request** with:
   - The manifest file
   - Any templates or additional code
   - Documentation updates
   - Test results from your Uberspace installation

### Code Contributions

#### Setting Up Development Environment

```bash
# Clone the repository
git clone https://github.com/tburny/uberman.git
cd uberman

# Install dependencies
go mod download

# Build
make build

# Run tests
make test
```

#### Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and testable

#### Writing Tests

- Write tests for new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Mock external dependencies (Uberspace commands)

Example test structure:
```go
func TestFeature(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        // Test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

#### Pull Request Process

1. **Fork and create a branch**:
   ```bash
   git checkout -b feature/my-feature
   # or
   git checkout -b fix/my-bugfix
   ```

2. **Make your changes**:
   - Write clear, focused commits
   - Follow code style guidelines
   - Add tests for new functionality
   - Update documentation

3. **Test your changes**:
   ```bash
   make test
   make test-coverage
   ```

4. **Format and lint**:
   ```bash
   make fmt
   make lint
   ```

5. **Commit with clear messages**:
   ```bash
   git commit -m "Add feature: description"
   # or
   git commit -m "Fix: description of bug fix"
   ```

6. **Push and create PR**:
   ```bash
   git push origin feature/my-feature
   ```
   Then create a pull request on GitHub with:
   - Clear description of changes
   - Reference to related issues
   - Test results
   - Screenshots (if UI changes)

## Commit Message Guidelines

This project follows the [Conventional Commits 1.0.0 Specification](https://www.conventionalcommits.org/en/v1.0.0/).

### Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

- `feat:` - A new feature (correlates with MINOR in SemVer)
- `fix:` - A bug fix (correlates with PATCH in SemVer)
- `docs:` - Documentation only changes
- `style:` - Changes that don't affect code meaning (formatting, whitespace)
- `refactor:` - Code change that neither fixes a bug nor adds a feature
- `perf:` - Performance improvement
- `test:` - Adding or correcting tests
- `build:` - Changes to build system or dependencies
- `ci:` - Changes to CI configuration files and scripts
- `chore:` - Other changes that don't modify src or test files
- `revert:` - Reverts a previous commit

### Breaking Changes

Breaking changes MUST be indicated by:
- `!` after the type/scope: `feat!:` or `feat(api)!:`
- OR a `BREAKING CHANGE:` footer

### Scopes

Common scopes for this project:
- `config` - Configuration parsing and validation
- `runtime` - Runtime version management
- `database` - Database operations
- `web` - Web backend configuration
- `supervisor` - Service management
- `appdir` - Directory management
- `backup` - Backup/restore functionality
- `deploy` - Deployment features
- `cli` - CLI commands and flags
- `manifest` - App manifest handling

### Examples

```
feat(database): add PostgreSQL support

Implements PostgreSQL database creation and management
alongside existing MySQL support.

Closes #123

---

fix(supervisor): correct service file template syntax

The supervisord configuration template had incorrect
environment variable syntax that caused services to fail.

---

feat(cli)!: change install command argument format

BREAKING CHANGE: install command now requires explicit
app name instead of interactive selection.

Before: uberman install
After: uberman install wordpress

---

docs: add PostgreSQL setup instructions

---

test(database): add integration tests for MySQL operations

---

chore(deps): update cobra to v1.8.0

---

refactor(config): simplify manifest validation logic
```

### Additional Rules

1. Use imperative mood in subject line ("add" not "added" or "adds")
2. Don't capitalize first letter of description
3. No period at the end of subject line
4. Limit subject line to 72 characters
5. Separate subject from body with blank line
6. Wrap body at 72 characters
7. Use body to explain what and why, not how
8. Reference issues and pull requests in footer

## App Manifest Guidelines

When creating app manifests:

### Required Fields

All manifests must include:
```toml
[app]
name = "appname"
type = "php|python|nodejs|ruby|go|static"
version = "latest"
description = "Clear description"

[runtime]
language = "language"
version = "version"
```

### Best Practices

1. **Use stable versions**: Prefer stable releases over `latest`
2. **Document dependencies**: List all required system tools
3. **Secure by default**: Follow security best practices
4. **Test backup/restore**: Verify backup includes all critical data
5. **Minimal permissions**: Use least privilege principle
6. **Clear errors**: Provide helpful error messages

### Example Structure

```toml
[app]
name = "myapp"
type = "nodejs"
version = "2.0.0"
description = "My awesome application"

[runtime]
language = "node"
version = "20"

[database]
type = "mysql"
required = true

[install]
method = "git"
source = "https://github.com/user/repo.git"
location = "app/"

[web]
backend = "http"
port = 8080

[[services]]
name = "myapp"
command = "node server.js"
directory = "app/"
port = 8080
startsecs = 15
autorestart = true

[services.environment]
NODE_ENV = "production"

[backup]
include = ["data/", "uploads/"]
exclude = ["cache/", "tmp/", "node_modules/"]
```

## Testing on Uberspace

Before submitting app manifests:

1. **Test fresh installation**:
   ```bash
   uberman install myapp
   ```

2. **Verify services**:
   ```bash
   supervisorctl status
   uberman status myapp
   ```

3. **Test backup/restore**:
   ```bash
   uberman backup myapp
   uberman restore myapp <backup-id>
   ```

4. **Test upgrade path**:
   ```bash
   uberman upgrade myapp
   ```

5. **Check cleanup**:
   ```bash
   # Ensure no orphaned files or services
   ```

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Accept constructive criticism
- Focus on what's best for the project
- Show empathy towards others

## Questions?

- Open an issue for questions
- Join discussions on GitHub
- Check existing documentation

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
