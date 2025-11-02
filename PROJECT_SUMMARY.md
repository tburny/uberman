# Uberman Project Summary

## Overview

**Uberman** is a Go-based CLI tool designed to automate application deployment and management on Uberspace hosting. It provides reproducible installations following a "convention over configuration" approach, where each app is self-contained in `~/apps/<app-name>`.

## Project Status

**Version**: 0.1.0 (Foundation Phase)
**Status**: Core infrastructure complete, command implementations in progress
**License**: MIT

## What Has Been Created

### Core Infrastructure ✅

1. **Go Project Structure**
   - Module: `github.com/tburny/uberman`
   - CLI framework using cobra
   - Global flags: `--dry-run`, `--verbose`, `--config`

2. **Internal Packages** (7 packages)
   - `config` - TOML manifest parser and validation
   - `runtime` - Uberspace runtime version management
   - `database` - MySQL database operations
   - `web` - Web backend configuration
   - `supervisor` - Supervisord service management
   - `appdir` - Directory structure management
   - Placeholders: `backup`, `deploy`, `templates`

3. **Example App Manifests** (3 apps)
   - WordPress (PHP/MySQL)
   - Ghost (Node.js/MySQL)
   - Nextcloud (PHP/MySQL)

4. **Documentation** (Complete)
   - `CLAUDE.md` - Comprehensive architecture and development guide
   - `README.md` - User-facing documentation with features and usage
   - `CONTRIBUTING.md` - Contribution guidelines
   - `QUICKSTART.md` - Getting started tutorial
   - `ROADMAP.md` - Development phases and future features

5. **Development Tools**
   - `Makefile` - Build, test, install targets
   - `.gitignore` - Proper exclusions
   - `.uberman.toml.example` - Configuration template
   - `LICENSE` - MIT license

### Project Structure

```
uberman/
├── cmd/uberman/              # CLI application
│   ├── main.go              # Entry point with cobra setup
│   └── install.go           # Install command (partial)
├── internal/                # Internal packages
│   ├── config/             # Manifest parsing
│   ├── runtime/            # Version management
│   ├── database/           # MySQL operations
│   ├── web/                # Backend configuration
│   ├── supervisor/         # Service management
│   ├── appdir/             # Directory management
│   ├── backup/             # (placeholder)
│   ├── deploy/             # (placeholder)
│   └── templates/          # (placeholder)
├── apps/                   # App manifests
│   ├── examples/           # Pre-defined apps
│   └── custom/             # User apps
├── docs/                   # Documentation
├── go.mod                  # Go dependencies
├── Makefile               # Build automation
└── [configuration files]   # .gitignore, LICENSE, etc.
```

## Key Design Decisions

### Technology Choices

1. **Go** - For single binary distribution, excellent stdlib, strong typing
2. **TOML** - For human-readable, strongly-typed configuration
3. **Cobra** - For CLI framework with subcommands and flags
4. **Convention over Configuration** - Standardized `~/apps/<app-name>` layout

### Architecture Patterns

1. **Self-Contained Apps** - Each app in isolated directory with all dependencies
2. **Manifest-Driven** - Declarative TOML manifests define app requirements
3. **Safety First** - Dry-run mode, automatic backups, validation before execution
4. **Uberspace Integration** - Direct use of `uberspace` commands, supervisord, MySQL

### Directory Convention

Every app follows this structure:
```
~/apps/<app-name>/
├── .uberman.toml          # Instance configuration
├── app/                   # Application code
├── data/                  # Persistent data
├── logs/                  # Logs (symlinked to ~/logs/)
├── backups/               # Local backups
└── tmp/                   # Temporary files
```

## What Works Now

### Implemented Functionality

1. **CLI Framework**
   - Main command with version and help
   - Global flags (dry-run, verbose, config)
   - Subcommand structure ready

2. **Manifest System**
   - TOML parsing with full schema support
   - Manifest validation
   - Multi-location manifest search
   - Example manifests for 3 popular apps

3. **Runtime Management**
   - Set/get language versions (PHP, Python, Node.js, etc.)
   - PHP-FPM restart capability
   - Version listing

4. **Database Operations**
   - MySQL credential parsing from `~/.my.cnf`
   - Database creation with proper naming (`${USER}_<app>`)
   - Database existence checks
   - Import/export functionality

5. **Web Backend**
   - Apache and HTTP backend configuration
   - Port allocation in range 1024-65535
   - Backend listing and deletion
   - Domain configuration

6. **Service Management**
   - Supervisord service file generation
   - Service lifecycle management (start/stop/restart)
   - Configuration reload
   - Service status checking

7. **Directory Management**
   - App directory structure creation
   - Log symlink to `~/logs/<app>`
   - Directory validation
   - Cleanup operations

8. **Basic Install Command**
   - Manifest loading and validation
   - Directory creation
   - Runtime version setting
   - Database creation
   - (Needs completion for full workflow)

## What's Missing

### Phase 1 Priorities (v0.2.0)

1. **Complete Install Command**
   - Download/clone app code
   - Install dependencies (composer, pip, npm)
   - Generate configuration files from templates
   - Configure web backends
   - Create supervisord services
   - Set up cron jobs
   - Post-install verification

2. **Template Engine**
   - Configuration file generation
   - Variable substitution
   - Pre-defined templates (wp-config.php, etc.)

3. **Backup/Restore Commands**
   - File system backup with compression
   - Database dump integration
   - Versioning and retention
   - Restore functionality

4. **Status Command**
   - Service health checks
   - Port connectivity
   - Database connectivity
   - Web backend verification

5. **Upgrade Command**
   - Pre-upgrade backup
   - Safe upgrade execution
   - Rollback on failure

6. **Additional Commands**
   - `list` - Show installed apps
   - `init` - Create manifest template
   - `deploy` - Git-based deployment

## How to Use (When Complete)

### Installation
```bash
cd ~/projekte
git clone https://github.com/tobias/uberman.git
cd uberman
make install-local
```

### Install WordPress
```bash
uberman install wordpress
```

### Create Custom App
```bash
uberman init myapp
nano apps/custom/myapp.toml
uberman install myapp
```

## Development Workflow

### Building
```bash
make build          # Build binary
make install-local  # Install to ~/bin
make test          # Run tests
make fmt           # Format code
```

### Adding Commands
1. Create `cmd/uberman/<command>.go`
2. Register with cobra in `init()`
3. Implement logic using `internal/` packages
4. Add tests
5. Update documentation

### Adding App Manifests
1. Create `apps/examples/<app>.toml`
2. Test on real Uberspace instance
3. Document in README.md

## Success Criteria

The project will be considered successful when:

1. ✅ Core infrastructure is complete
2. ⏳ WordPress can be installed with single command
3. ⏳ Upgrades work reliably with automatic backup
4. ⏳ Backup/restore workflow is functional
5. ⏳ Status checks provide accurate information
6. ⏳ User can create custom app manifests
7. ⏳ Documentation is comprehensive and accurate
8. ⏳ Tests cover critical functionality

## Next Steps

### Immediate (v0.2.0)
1. Complete the install command implementation
2. Build template engine for config files
3. Implement backup/restore commands
4. Add status and list commands
5. Create more app manifest examples

### Short Term (v0.3.0)
1. Add interactive mode with prompts
2. Improve error messages and validation
3. Implement `init` command for scaffolding
4. Add progress indicators
5. Create integration tests

### Long Term (v0.4.0+)
1. Multi-app orchestration
2. Monitoring and alerts
3. Web interface (optional)
4. CI/CD integration
5. App registry/marketplace

## Resources

- **Repository**: (to be published)
- **Documentation**: See `docs/` directory
- **Uberspace Lab**: https://lab.uberspace.de/
- **Uberspace Manual**: https://manual.uberspace.de/

## Contributors

- Initial development: Tobias (with Claude Code assistance)
- Open for contributions! See CONTRIBUTING.md

## Notes

- Go is not installed in the current development environment
- Once Go is available, run `go mod download` to fetch dependencies
- Test thoroughly on real Uberspace instance before production use
- Always maintain separate backups (don't rely solely on Uberspace backups)

---

*Project created: 2025-11-02*
*Last updated: 2025-11-02*
