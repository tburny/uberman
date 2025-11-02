# Release Process

This document describes the automated release process for Uberman using semantic-release and GitHub Actions.

## Overview

Uberman uses **semantic versioning** and **automated releases** based on **Conventional Commits**. Every commit to the `main` branch triggers the release workflow, which:

1. Analyzes commits since the last release
2. Determines the next version number (major, minor, or patch)
3. Generates release notes from commit messages
4. Updates CHANGELOG.md
5. Creates a GitHub release with version tag
6. Builds binaries for multiple platforms
7. Uploads binaries to the release

## Versioning

Uberman follows [Semantic Versioning 2.0.0](https://semver.org/):

```
MAJOR.MINOR.PATCH
```

- **MAJOR** (1.0.0): Breaking changes (incompatible API changes)
- **MINOR** (0.1.0): New features (backwards-compatible)
- **PATCH** (0.0.1): Bug fixes (backwards-compatible)

### Version Bumping Rules

Based on [Conventional Commits](https://www.conventionalcommits.org/):

| Commit Type | Version Bump | Example |
|-------------|--------------|---------|
| `feat:` | MINOR | `feat(cli): add list command` → 0.2.0 |
| `fix:` | PATCH | `fix(database): handle timeout` → 0.1.1 |
| `perf:` | PATCH | `perf(backup): use parallel compression` → 0.1.1 |
| `refactor:` | PATCH | `refactor(web): simplify logic` → 0.1.1 |
| `feat!:` or `BREAKING CHANGE:` | MAJOR | `feat(cli)!: change API` → 1.0.0 |
| `docs:`, `style:`, `test:`, `chore:` | None | No release created |

## Automated Release Workflow

### Trigger

The release workflow runs on every push to `main`:

```yaml
on:
  push:
    branches:
      - main
```

### Steps

1. **Commit Analysis** (`@semantic-release/commit-analyzer`)
   - Parses all commits since last release
   - Determines if a release is needed
   - Calculates next version number

2. **Release Notes** (`@semantic-release/release-notes-generator`)
   - Generates release notes from commits
   - Groups by type (Features, Bug Fixes, etc.)
   - Includes breaking changes prominently

3. **Changelog Update** (`@semantic-release/changelog`)
   - Updates CHANGELOG.md with new version
   - Maintains Keep a Changelog format
   - Commits back to repository

4. **GitHub Release** (`@semantic-release/github`)
   - Creates Git tag (e.g., `v0.2.0`)
   - Creates GitHub release with notes
   - Marks release as published

5. **Binary Building** (Custom job)
   - Builds for multiple platforms:
     - Linux (amd64, arm64)
     - macOS (amd64, arm64)
     - Windows (amd64)
     - FreeBSD (amd64)
   - Creates tarballs (.tar.gz) and zips (.zip)
   - Generates SHA256 checksums

6. **Asset Upload**
   - Uploads all binaries to GitHub release
   - Includes checksums for verification

## Supported Platforms

Binaries are built for:

| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Linux | amd64 | `uberman-linux-amd64.tar.gz` |
| Linux | arm64 | `uberman-linux-arm64.tar.gz` |
| macOS | amd64 (Intel) | `uberman-darwin-amd64.tar.gz` |
| macOS | arm64 (Apple Silicon) | `uberman-darwin-arm64.tar.gz` |
| Windows | amd64 | `uberman-windows-amd64.zip` |
| FreeBSD | amd64 | `uberman-freebsd-amd64.tar.gz` |

Each archive includes:
- Compiled binary
- SHA256 checksum file

## How to Trigger a Release

### Option 1: Feature Release (Minor Version)

```bash
# Make your changes
git add .

# Commit with feat: type
git commit -m "feat(cli): add status command for health monitoring"

# Push to main
git push origin main
```

Result: Version bumps from 0.1.0 → **0.2.0**

### Option 2: Bug Fix Release (Patch Version)

```bash
# Fix the bug
git add .

# Commit with fix: type
git commit -m "fix(database): handle connection timeout gracefully

Previously, database connection timeouts would cause the
application to crash. This adds proper error handling and retry logic.

Closes #42"

# Push to main
git push origin main
```

Result: Version bumps from 0.1.0 → **0.1.1**

### Option 3: Breaking Change Release (Major Version)

```bash
# Make breaking changes
git add .

# Option A: Use ! notation
git commit -m "feat(api)!: change database naming convention

BREAKING CHANGE: Database names now use underscore separator
instead of dash. Existing installations must migrate databases.

Migration guide: docs/MIGRATION.md"

# Option B: Use BREAKING CHANGE footer
git commit -m "feat(api): change database naming convention

Database names now use underscore for consistency with MySQL
naming conventions and improved readability.

BREAKING CHANGE: Database names now use underscore separator
instead of dash. Existing installations must run migration:
uberman migrate-database <app-name>"

# Push to main
git push origin main
```

Result: Version bumps from 0.1.0 → **1.0.0**

### Option 4: No Release (Documentation, etc.)

```bash
# Update documentation
git add .

# Commit with docs: type
git commit -m "docs: update installation instructions for Go 1.22"

# Push to main
git push origin main
```

Result: **No release created** (CHANGELOG not updated)

## Release Notes Format

Generated release notes follow this structure:

```markdown
# [0.2.0](https://github.com/tburny/uberman/compare/v0.1.0...v0.2.0) (2025-11-02)

## Features

* **cli:** add status command for health monitoring ([abc1234](commit-link))
* **backup:** implement incremental backup strategy ([def5678](commit-link))

## Bug Fixes

* **database:** handle connection timeout gracefully ([ghi9012](commit-link))

## Documentation

* update installation instructions for Go 1.22 ([jkl3456](commit-link))

## BREAKING CHANGES

* **api:** Database names now use underscore separator instead of dash.
  Migration required for existing installations.
```

## Manual Release (Emergency)

If automated release fails, you can create a manual release:

### 1. Create Tag Locally

```bash
# Determine next version
git log --oneline

# Create and push tag
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```

### 2. Build Binaries

```bash
# Build for all platforms
make release

# Or build individual platforms
GOOS=linux GOARCH=amd64 go build -o uberman-linux-amd64 ./cmd/uberman
GOOS=darwin GOARCH=arm64 go build -o uberman-darwin-arm64 ./cmd/uberman
```

### 3. Create GitHub Release

1. Go to GitHub → Releases → Draft a new release
2. Choose the tag (v0.2.0)
3. Generate release notes or write manually
4. Upload binaries
5. Publish release

### 4. Update CHANGELOG

```bash
# Manually edit CHANGELOG.md
vim CHANGELOG.md

# Add version and changes
git add CHANGELOG.md
git commit -m "chore(release): 0.2.0 [skip ci]"
git push origin main
```

## Skipping CI

To push commits without triggering release:

```bash
git commit -m "chore: update README [skip ci]"
```

Or in the commit body:

```bash
git commit -m "chore: update README

[skip ci]"
```

## Pre-release Versions

For alpha/beta releases (future feature):

```bash
# Create pre-release branch
git checkout -b beta

# Make changes and commit
git commit -m "feat(cli): add experimental feature"

# Push
git push origin beta
```

This would create: `v0.2.0-beta.1`

## Release Checklist

Before releasing a major version:

- [ ] Update documentation for breaking changes
- [ ] Create migration guide if needed
- [ ] Update CLAUDE.md with new architecture
- [ ] Test on multiple platforms
- [ ] Update example manifests
- [ ] Review CHANGELOG.md
- [ ] Announce in GitHub Discussions
- [ ] Update README.md with new features

## Monitoring Releases

### Check Release Status

1. **GitHub Actions**: Check workflow runs
   - Go to: https://github.com/tburny/uberman/actions
   - Look for "Release" workflow

2. **GitHub Releases**: View published releases
   - Go to: https://github.com/tburny/uberman/releases
   - Verify binaries uploaded

3. **CHANGELOG**: Check automated updates
   - View: https://github.com/tburny/uberman/blob/main/CHANGELOG.md

### Release Notifications

Releases trigger:
- GitHub release creation (subscribers notified)
- Git tag creation
- CHANGELOG commit to main

## Troubleshooting

### Release Workflow Failed

1. Check GitHub Actions logs
2. Verify GITHUB_TOKEN permissions
3. Ensure commits follow Conventional Commits
4. Check for merge conflicts in CHANGELOG

### No Release Created

Possible reasons:
- No `feat:`, `fix:`, or `perf:` commits since last release
- Only `docs:`, `chore:`, or `test:` commits
- Commits don't follow Conventional Commits format

Solution:
```bash
# View recent commits
git log --oneline

# Verify commit format
npx commitlint --from=HEAD~5
```

### Binary Build Failed

1. Check Go version compatibility
2. Verify imports are valid
3. Test local build: `make build`
4. Check for platform-specific issues

### Changelog Conflicts

If CHANGELOG.md has merge conflicts:

```bash
# Reset to remote
git fetch origin
git checkout origin/main -- CHANGELOG.md

# Re-run semantic-release (will skip)
git push origin main --force
```

## Configuration Files

### `.releaserc.json`

Main semantic-release configuration:
- Defines release rules
- Configures plugins
- Sets changelog format

### `.github/workflows/release.yml`

GitHub Actions workflow:
- Runs semantic-release
- Builds binaries
- Uploads assets

### `.github/workflows/test.yml`

CI workflow (runs before release):
- Tests code
- Verifies formatting
- Checks compilation

## Best Practices

1. **Always use Conventional Commits** for main branch
2. **Test locally** before pushing to main
3. **Write descriptive commit bodies** (appear in changelog)
4. **Reference issues** with `Closes #123`
5. **Group related changes** in single commits
6. **Use breaking changes sparingly**
7. **Update docs** with new features
8. **Review release notes** after automated creation

## Examples

### Good Release Commit

```bash
git commit -m "feat(backup): implement incremental backup strategy

Incremental backups reduce backup time by 80% and storage by 60%.
Especially beneficial for large Nextcloud installations with 100GB+ data.

Implementation uses file modification timestamps and maintains a
manifest of backed-up files. Full backups still created monthly.

Closes #45
Refs #42, #38"
```

This creates:
- Minor version bump (0.1.0 → 0.2.0)
- Well-formatted release notes
- Links to issues
- Detailed explanation for users

### Bad Release Commit

```bash
git commit -m "added stuff"
```

This creates:
- **No release** (doesn't follow convention)
- Rejected by commitlint on PR
- Must be fixed before merge

## Resources

- **Semantic Release**: https://semantic-release.gitbook.io/
- **Conventional Commits**: https://www.conventionalcommits.org/
- **Semantic Versioning**: https://semver.org/
- **Keep a Changelog**: https://keepachangelog.com/

## Support

Questions about releases?
- Open issue: https://github.com/tburny/uberman/issues
- GitHub Discussions: https://github.com/tburny/uberman/discussions
- Check workflow logs: https://github.com/tburny/uberman/actions

---

*This release process is fully automated using semantic-release and GitHub Actions.*
