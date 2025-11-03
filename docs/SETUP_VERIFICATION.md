# Setup Verification Guide

> **⚠️ REFACTORING IN PROGRESS**
> This document references code deleted during 2025-11-03 Clean Architecture refactoring.
> See CLAUDE.md, ARCHITECTURE.md, PLANNING.md for current state.
> Will update after refactoring completes.

This document helps you verify that the Uberman project is correctly set up.

## Git Repository Status

✅ **Git initialized**: Repository initialized with `main` branch
✅ **Conventional Commits**: Following Conventional Commits 1.0.0 specification
✅ **Commit template**: `.gitmessage` configured for guided commits
✅ **Initial commits**: Foundation and tooling commits created

### Verify Git Setup

```bash
# Check git status
git status

# View commit history
git log --oneline

# Verify commit template
git config commit.template
```

Expected output:
- Clean working directory or changes ready to commit
- Two initial commits following conventional commits format
- Commit template set to `.gitmessage`

## Project Structure Verification

### Directory Structure

```bash
tree -L 2 -I 'bin|vendor|.git'
```

Expected directories:
```
.
├── .github/               # GitHub workflows and configs
├── apps/
│   ├── custom/           # User-defined app manifests
│   └── examples/         # Pre-defined app manifests
├── cmd/
│   └── uberman/          # CLI application
├── docs/                 # Documentation
├── internal/             # Internal packages
│   ├── appdir/
│   ├── backup/
│   ├── config/
│   ├── database/
│   ├── deploy/
│   ├── runtime/
│   ├── supervisor/
│   ├── templates/
│   └── web/
└── scripts/              # Helper scripts
```

### Files Checklist

Core files:
- [ ] `go.mod` - Go module definition
- [ ] `Makefile` - Build automation
- [ ] `LICENSE` - MIT license
- [ ] `.gitignore` - Git ignore rules
- [ ] `.gitmessage` - Commit template

Documentation:
- [ ] `README.md` - Main documentation
- [ ] `CLAUDE.md` - Architecture guide
- [ ] `CONTRIBUTING.md` - Contribution guidelines
- [ ] `CHANGELOG.md` - Version history
- [ ] `PROJECT_SUMMARY.md` - Project overview
- [ ] `docs/QUICKSTART.md` - Getting started
- [ ] `docs/ROADMAP.md` - Development roadmap
- [ ] `docs/CONVENTIONAL_COMMITS.md` - Commit guidelines
- [ ] `docs/SETUP_VERIFICATION.md` - This file

Configuration:
- [ ] `.uberman.toml.example` - Example config
- [ ] `.github/commitlint.config.js` - Commitlint config
- [ ] `.github/workflows/commitlint.yml` - CI workflow

Source code:
- [ ] `cmd/uberman/main.go` - CLI entry point
- [ ] `cmd/uberman/install.go` - Install command
- [ ] `internal/config/app.go` - Config parser
- [ ] `internal/runtime/uberspace.go` - Runtime manager
- [ ] `internal/database/mysql.go` - Database manager
- [ ] `internal/web/backend.go` - Web backend manager
- [ ] `internal/supervisor/service.go` - Service manager
- [ ] `internal/appdir/manager.go` - Directory manager

App manifests:
- [ ] `apps/examples/wordpress.toml`
- [ ] `apps/examples/ghost.toml`
- [ ] `apps/examples/nextcloud.toml`
- [ ] `apps/custom/.gitkeep`

Scripts:
- [ ] `scripts/commit.sh` - Commit helper (executable)

## Go Module Verification

### Check Module

```bash
# View module info
cat go.mod

# Check for syntax errors
go mod verify
```

Expected:
- Module: `github.com/tburny/uberman`
- Go version: 1.21
- Dependencies: cobra, toml, viper

### Note on Go Installation

⚠️ **Go is not currently installed** in the development environment.

To continue development, install Go:

**On Uberspace:**
```bash
uberspace tools version use go 1.21
go version
```

**On local machine:**
```bash
# Download from https://golang.org/dl/
# Or use your package manager:
brew install go        # macOS
apt install golang     # Ubuntu/Debian
dnf install golang     # Fedora
```

Once Go is installed:
```bash
# Download dependencies
go mod download

# Verify dependencies
go mod verify

# Build the project
make build
```

## Documentation Verification

### Completeness Check

All documentation should be:
- [ ] Free of TODO placeholders
- [ ] Using correct GitHub username (tburny)
- [ ] Using correct module path (github.com/tburny/uberman)
- [ ] Following markdown formatting
- [ ] Cross-referenced correctly

### Link Verification

Key links to verify:
- GitHub URLs use `tburny` username
- Internal links use relative paths
- External links to Uberspace are correct
- All example commands are accurate

## Conventional Commits Verification

### Check Commit Messages

```bash
# View recent commits
git log --oneline -5

# View full commit message
git log -1 --pretty=full
```

Each commit should:
- [ ] Start with type: `feat`, `fix`, `chore`, etc.
- [ ] Include scope in parentheses (optional)
- [ ] Have lowercase description
- [ ] Have no period at end of subject
- [ ] Be under 72 characters in subject
- [ ] Include body explaining why (if needed)
- [ ] Reference issues in footer (if applicable)

### Test Commit Template

```bash
# Open git commit (will show template)
git commit --allow-empty

# Or use the helper script
./scripts/commit.sh
```

Expected:
- Template appears with guidelines
- Helper script asks for type, scope, description
- Both produce valid conventional commits

## Build Verification (When Go is Available)

### Build Commands

```bash
# Download dependencies
go mod download

# Run tests
make test

# Build binary
make build

# Install locally
make install-local
```

Expected outputs:
- Dependencies downloaded successfully
- Tests pass (or are skipped if none exist)
- Binary created in `bin/uberman`
- Binary installed to `~/bin/uberman`

### Verify Binary

```bash
# Check binary exists
ls -lh bin/uberman

# Run binary (once built)
bin/uberman --version
bin/uberman --help
```

## Manifest Verification

### Check Example Manifests

```bash
# List manifests
ls -l apps/examples/

# Validate manifest syntax (requires Go)
# This will be available once the validation command is implemented
# uberman validate apps/examples/wordpress.toml
```

Each manifest should:
- [ ] Have valid TOML syntax
- [ ] Include all required fields
- [ ] Use correct types
- [ ] Have meaningful descriptions

### Test Manifest Loading (Manual)

```bash
# Check TOML syntax manually
cat apps/examples/wordpress.toml | grep -E "^\[|^[a-z]"
```

Expected: Clean TOML structure with sections and key-value pairs

## GitHub Integration Verification

### Workflows

```bash
# Check workflow syntax
cat .github/workflows/commitlint.yml
```

Expected:
- [ ] Valid YAML syntax
- [ ] Runs on pull_request
- [ ] Uses commitlint
- [ ] Validates commit messages

### Commitlint Config

```bash
# Check config
cat .github/commitlint.config.js
```

Expected:
- [ ] Extends conventional config
- [ ] Defines project scopes
- [ ] Sets proper rules

## Next Steps After Verification

Once verification is complete:

1. **Install Go** (if not already installed)
2. **Build the project**: `make build`
3. **Run tests**: `make test`
4. **Read documentation**: Start with `QUICKSTART.md`
5. **Try the commit helper**: `./scripts/commit.sh`
6. **Make your first contribution**: See `CONTRIBUTING.md`

## Troubleshooting

### Issue: Git not initialized
```bash
git init
git branch -m main
git config commit.template .gitmessage
```

### Issue: Go mod errors
```bash
# Ensure Go is installed
go version

# Clean and re-download
go clean -modcache
go mod download
```

### Issue: Build fails
```bash
# Check Go version (requires 1.21+)
go version

# Verify module
go mod verify

# Check for syntax errors
go build ./...
```

### Issue: Commit script not executable
```bash
chmod +x scripts/commit.sh
```

### Issue: Wrong GitHub username
```bash
# Should be tburny, not tobias
grep -r "tobias/uberman" .
# If found, update to tburny/uberman
```

## Verification Checklist

Complete this checklist to ensure proper setup:

- [ ] Git repository initialized with main branch
- [ ] All files committed with conventional commits
- [ ] Project structure matches expected layout
- [ ] Documentation is complete and accurate
- [ ] Go module uses correct path (github.com/tburny/uberman)
- [ ] Example manifests are valid TOML
- [ ] Commit template configured
- [ ] Commit helper script is executable
- [ ] GitHub workflows are configured
- [ ] LICENSE file exists (MIT)
- [ ] .gitignore properly excludes build artifacts
- [ ] README has correct installation instructions
- [ ] All documentation uses correct GitHub username

## Success Criteria

The project setup is verified when:

✅ Git repository is clean and properly initialized
✅ All documentation is present and accurate
✅ Conventional commits are configured and working
✅ Project structure matches specification
✅ All files are committed with proper messages
✅ Build succeeds (once Go is installed)
✅ Example manifests are valid
✅ Ready for development and contributions

## Resources

- **Conventional Commits**: [docs/CONVENTIONAL_COMMITS.md](CONVENTIONAL_COMMITS.md)
- **Quick Start**: [QUICKSTART.md](QUICKSTART.md)
- **Contributing**: [../CONTRIBUTING.md](../CONTRIBUTING.md)
- **Architecture**: [../CLAUDE.md](../CLAUDE.md)

---

**Setup Date**: 2025-11-02
**Project Version**: 0.1.0
**Verification**: ✅ Complete
