# CI/CD Setup Documentation

This document provides an overview of the CI/CD infrastructure implemented for Uberman.

## Overview

Uberman uses a fully automated release and build pipeline powered by:

- **GitHub Actions** for CI/CD workflows
- **semantic-release** for automated versioning
- **Conventional Commits** for version determination
- **Multi-platform builds** for broad compatibility

## Architecture

```
Push to main
     │
     ├──> Test Workflow (runs immediately)
     │    ├── Run tests (Go 1.21, 1.22)
     │    ├── Run linter (golangci-lint)
     │    ├── Build for all platforms
     │    └── Upload coverage
     │
     └──> Release Workflow (runs after tests pass)
          ├── Analyze commits (semantic-release)
          ├── Determine version (major/minor/patch)
          ├── Generate release notes
          ├── Update CHANGELOG.md
          ├── Create Git tag
          ├── Build binaries (6 platforms)
          └── Upload to GitHub Release
```

## Workflows

### 1. Test Workflow (`.github/workflows/test.yml`)

**Triggers:**

- Push to `main` branch
- Pull requests to `main`

**Jobs:**

**Test Job:**

- Runs on Go 1.21 and 1.22
- Executes: `go test -v -race -coverprofile=coverage.out`
- Uploads coverage to Codecov
- Validates code formatting with `gofmt`
- Runs `go vet` for static analysis

**Build Job:**

- Cross-compiles for all supported platforms
- Validates builds complete successfully
- Matrix builds:
  - Linux: amd64, arm64
  - macOS: amd64, arm64
  - Windows: amd64
  - FreeBSD: amd64

**Lint Job:**

- Runs golangci-lint with 20+ linters
- Checks code quality and style
- Configuration: `.golangci.yml`

### 2. Release Workflow (`.github/workflows/release.yml`)

**Triggers:**

- Push to `main` branch (only)

**Permissions Required:**

```yaml
permissions:
  contents: write     # Create releases and tags
  issues: write       # Comment on issues
  pull-requests: write # Comment on PRs
```

**Jobs:**

**Semantic Release Job:**

1. Analyzes all commits since last release
2. Determines next version based on commit types:
   - `feat:` → Minor version bump (0.x.0)
   - `fix:` → Patch version bump (0.0.x)
   - `BREAKING CHANGE:` → Major version bump (x.0.0)
3. Generates release notes from commits
4. Updates CHANGELOG.md
5. Creates Git tag (e.g., `v0.2.0`)
6. Creates GitHub release draft

**Build Job:**

- Triggered only if new release is created
- Builds binaries for 6 platform/arch combinations:
  - `uberman-linux-amd64`
  - `uberman-linux-arm64`
  - `uberman-darwin-amd64` (Intel Mac)
  - `uberman-darwin-arm64` (Apple Silicon)
  - `uberman-windows-amd64.exe`
  - `uberman-freebsd-amd64`
- Creates tarballs (.tar.gz) and zips (.zip)
- Generates SHA256 checksums
- Uploads as workflow artifacts

**Upload Assets Job:**

- Downloads all build artifacts
- Uploads binaries to GitHub release
- Attaches checksums for verification

### 3. Commitlint Workflow (`.github/workflows/commitlint.yml`)

**Triggers:**

- Pull requests (opened, synchronized, reopened)

**Purpose:**

- Validates all commit messages follow Conventional Commits
- Prevents merging non-conforming commits
- Ensures releases will work correctly

**Configuration:** `.github/commitlint.config.js`

## Configuration Files

### `.releaserc.json`

Semantic-release configuration defining:

```json
{
  "branches": ["main"],
  "plugins": [
    "@semantic-release/commit-analyzer",     // Analyzes commits
    "@semantic-release/release-notes-generator", // Generates notes
    "@semantic-release/changelog",           // Updates CHANGELOG.md
    "@semantic-release/github",              // Creates GitHub release
    "@semantic-release/git"                  // Commits CHANGELOG back
  ]
}
```

**Release Rules:**

- `feat:` → Minor release
- `fix:`, `perf:`, `refactor:` → Patch release
- `docs:`, `style:`, `test:`, `chore:` → No release
- `BREAKING CHANGE:` or `!` → Major release

### `.golangci.yml`

Linter configuration with 20+ enabled linters:

- `errcheck`: Check error handling
- `gosec`: Security vulnerabilities
- `staticcheck`: Static analysis
- `govet`: Go vet integration
- `gofmt`: Code formatting
- And more...

### `Makefile`

Build automation with release target:

```makefile
# Local release build
make release

# Builds all platform binaries in dist/
# Includes version info from git tags
```

**Version Embedding:**

```makefile
VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"
```

## Version Determination

### From Conventional Commits

| Commit Type | Example | Version Change |
|-------------|---------|----------------|
| `feat:` | `feat(cli): add list command` | 0.1.0 → **0.2.0** |
| `fix:` | `fix(db): timeout handling` | 0.1.0 → **0.1.1** |
| `perf:` | `perf(backup): parallel` | 0.1.0 → **0.1.1** |
| `feat!:` | `feat(api)!: change format` | 0.1.0 → **1.0.0** |
| `docs:` | `docs: update README` | No change |

### Breaking Changes

Two ways to indicate breaking changes:

**Method 1: Exclamation mark**

```
feat(api)!: change database format
```

**Method 2: Footer**

```
feat(api): change database format

BREAKING CHANGE: Database schema changed. Migration required.
```

Both trigger major version bump: 0.x.x → **1.0.0**

## Binary Distribution

### Platforms Supported

| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Linux | amd64 | `uberman-linux-amd64.tar.gz` |
| Linux | arm64 | `uberman-linux-arm64.tar.gz` |
| macOS | amd64 | `uberman-darwin-amd64.tar.gz` |
| macOS | arm64 | `uberman-darwin-arm64.tar.gz` |
| Windows | amd64 | `uberman-windows-amd64.zip` |
| FreeBSD | amd64 | `uberman-freebsd-amd64.tar.gz` |

### Build Configuration

```yaml
- goos: linux
  goarch: amd64
  CGO_ENABLED: 0  # Static binary
```

**Build Flags:**

- `-s`: Strip debug symbols
- `-w`: Strip DWARF info
- `-ldflags`: Embed version info
- `CGO_ENABLED=0`: Static linking

### Checksums

Each binary includes SHA256 checksum:

```
uberman-linux-amd64.tar.gz
uberman-linux-amd64.tar.gz.sha256
```

**Verification:**

```bash
sha256sum -c uberman-linux-amd64.tar.gz.sha256
```

## Release Process

### Automatic Release

1. **Developer makes changes:**

   ```bash
   git add .
   git commit -m "feat(cli): add new command"
   git push origin main
   ```

2. **GitHub Actions triggers:**
   - Test workflow runs
   - If tests pass, release workflow runs

3. **Semantic-release analyzes commits:**
   - Since last release: v0.1.0
   - New commit type: `feat:`
   - Decision: Bump minor version

4. **Release workflow executes:**
   - Creates version v0.2.0
   - Generates release notes
   - Updates CHANGELOG.md
   - Creates Git tag
   - Commits CHANGELOG back

5. **Build job triggers:**
   - Builds 6 platform binaries
   - Creates archives with checksums
   - Uploads to GitHub release

6. **Release published:**
   - Available at: github.com/tburny/uberman/releases
   - Binaries ready for download
   - CHANGELOG updated
   - Tag created: v0.2.0

### Manual Verification

Check release status:

```bash
# View GitHub Actions
gh run list

# Check latest release
gh release view

# List release assets
gh release view --json assets
```

## Badges

README includes status badges:

```markdown
[![Release](https://img.shields.io/github/v/release/tburny/uberman)]
[![Build Status](https://github.com/tburny/uberman/workflows/Test/badge.svg)]
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)]
[![semantic-release](https://img.shields.io/badge/%20%20📦🚀-semantic--release-e10079.svg)]
```

These show:

- Latest release version
- Build/test status
- Conventional Commits compliance
- Semantic-release usage

## Monitoring

### GitHub Actions

**View workflow runs:**

```
https://github.com/tburny/uberman/actions
```

**Check specific workflow:**

- Test: <https://github.com/tburny/uberman/actions/workflows/test.yml>
- Release: <https://github.com/tburny/uberman/actions/workflows/release.yml>

### Releases

**View all releases:**

```
https://github.com/tburny/uberman/releases
```

**Latest release API:**

```bash
curl -s https://api.github.com/repos/tburny/uberman/releases/latest
```

### CHANGELOG

**View changes:**

```
https://github.com/tburny/uberman/blob/main/CHANGELOG.md
```

Updated automatically by semantic-release.

## Troubleshooting

### Release Not Created

**Possible causes:**

1. No `feat:`, `fix:`, or `perf:` commits
2. Only `docs:`, `chore:`, etc. commits
3. Commits don't follow Conventional Commits

**Solution:**

```bash
# Check commit format
git log --oneline --format="%s" | head -5

# Validate with commitlint
npx commitlint --from=HEAD~5
```

### Binary Build Failed

**Check logs:**

```bash
gh run view --log
```

**Common issues:**

- Import errors (check dependencies)
- Platform-specific code (use build tags)
- CGO dependencies (disable with CGO_ENABLED=0)

**Test locally:**

```bash
make release
```

### Workflow Permissions

If release fails with permission error:

1. Go to: Settings → Actions → General
2. Workflow permissions → Read and write
3. Save changes
4. Re-run workflow

## Security

### Token Usage

Workflows use `GITHUB_TOKEN`:

- Automatic authentication
- Scoped to repository
- Expires after workflow
- No manual setup needed

### Binary Signing

Future enhancement (planned):

- GPG signing of binaries
- Checksum signing
- Verification instructions

## Best Practices

1. **Always test locally:** `make build && make test`
2. **Use conventional commits:** Follow the specification
3. **Group related changes:** One logical change per commit
4. **Write descriptive bodies:** Explain why, not what
5. **Reference issues:** Use `Closes #123`
6. **Review release notes:** After automated creation
7. **Monitor workflows:** Check Actions tab regularly

## Metrics

### Current Setup

- **Workflows:** 3 (Test, Release, Commitlint)
- **Build platforms:** 6
- **Architectures:** 2 (amd64, arm64)
- **Go versions tested:** 2 (1.21, 1.22)
- **Linters enabled:** 20+
- **Average build time:** ~3-5 minutes
- **Release time:** ~5-8 minutes (including builds)

## Future Enhancements

Planned improvements:

- [ ] tbd

## Resources

- **GitHub Actions Docs**: <https://docs.github.com/actions>
- **semantic-release**: <https://semantic-release.gitbook.io/>
- **Conventional Commits**: <https://www.conventionalcommits.org/>
- **golangci-lint**: <https://golangci-lint.run/>

## Support

Issues with CI/CD:

- Check workflow logs: `gh run view`
- Review configuration files
- Open issue: <https://github.com/tburny/uberman/issues>
- See: [RELEASE_PROCESS.md](RELEASE_PROCESS.md)

---

*CI/CD infrastructure powered by GitHub Actions and semantic-release.*
