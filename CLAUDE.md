# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Uberman** is a Go-based CLI tool for reproducible installation, upgrades, backups, and deployment of applications on Uberspace hosting (uberspace.de). It follows a "convention over configuration" approach where each app is installed in a self-contained directory structure at `~/apps/<app-name>`.

### Purpose

Uberspace provides manual installation guides at https://lab.uberspace.de/ (GitHub: https://github.com/Uberspace/lab) that intentionally include manual steps for educational purposes. Uberman automates these installation patterns while maintaining best practices and safety.

### Key Constraints

- **No Docker/Containers**: Uberspace does not support containerization
- **No System Package Managers**: Cannot use apt/yum/pacman; must rely on language-specific package managers (npm, pip, composer, gem, etc.)
- **User-Space Only**: All operations run in user context without root privileges
- **Supervisord for Services**: Long-running processes managed via supervisord (not systemd)
- **Database Naming**: MySQL databases must be prefixed with username: `${USER}_<appname>`
- **Port Binding**: Applications must bind to `0.0.0.0` or `::` (NOT localhost/127.0.0.1)

## Commands

### Development Setup

**Note**: Go is required but not currently installed in the development environment. Once Go is available:

```bash
# Prerequisites: Go 1.21+ must be installed
# On Uberspace: uberspace tools version use go 1.21
# On local machine: Install from https://golang.org/dl/

# Install dependencies
go mod download

# Build the binary
go build -o bin/uberman cmd/uberman/main.go

# Install locally
go install ./cmd/uberman

# Run with verbose output
uberman --verbose <command>

# Dry-run mode (show what would be done)
uberman --dry-run <command>
```

### Available Commands

```bash
# Install an app from manifest
uberman install wordpress

# Upgrade an app (with automatic backup)
uberman upgrade wordpress

# Create manual backup
uberman backup wordpress

# Restore from backup
uberman restore wordpress <backup-id>

# Check app health and status
uberman status wordpress

# Deploy app from git/source
uberman deploy myapp

# List installed apps
uberman list

# Create new app manifest template
uberman init myapp
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/config/

# Run integration tests (requires Uberspace environment)
go test -tags=integration ./...
```

## Architecture

### Directory Structure Convention

All apps are installed following this structure:

```
~/apps/<app-name>/
├── .uberman.toml           # App instance configuration
├── app/                    # Application code/files
│   ├── wordpress/         # (example: WordPress files)
│   └── ...
├── data/                   # Persistent application data
├── logs/                   # Application logs (symlinked to ~/logs/)
├── backups/                # Local backup storage
└── tmp/                    # Temporary files (excluded from backups)
```

### Core Packages

#### `internal/config` - App Manifest Parser
- Parses TOML app manifests
- Validates configuration schema
- Searches for manifests in multiple locations (project, user home, installed app)
- Key types: `AppManifest`, `AppConfig`, `RuntimeConfig`, `DatabaseConfig`, `InstallConfig`, `WebConfig`, `ServiceConfig`

#### `internal/runtime` - Runtime Version Management
- Manages language runtime versions via `uberspace tools version`
- Supports: PHP, Python, Node.js, Ruby, Go, and others
- Key methods: `SetVersion()`, `GetVersion()`, `ListVersions()`, `RestartPHP()`

#### `internal/database` - Database Provisioning
- MySQL database creation and management
- Reads credentials from `~/.my.cnf`
- Enforces `${USER}_<appname>` naming convention
- Key methods: `CreateDatabase()`, `DatabaseExists()`, `ExportDatabase()`, `ImportDatabase()`

#### `internal/web` - Web Backend Configuration
- Configures Uberspace web backends via `uberspace web backend`
- Supports Apache (for PHP) and HTTP (for application servers)
- Finds available ports in range 1024-65535
- Key methods: `SetBackend()`, `DeleteBackend()`, `FindAvailablePort()`

#### `internal/supervisor` - Service Management
- Generates supervisord configuration files (`~/etc/services.d/*.ini`)
- Manages service lifecycle (start/stop/restart)
- Key methods: `CreateService()`, `ReloadServices()`, `StartService()`, `StopService()`

#### `internal/backup` - Backup/Restore Logic
- Creates compressed backups with versioning
- Excludes cache and temporary directories
- Integrates with Uberspace automatic backups (`/backup/`, `/mysql_backup/`)

#### `internal/deploy` - Deployment Strategies
- Supports git-based deployments
- Handles dependency installation per language
- Runs migrations and post-deploy hooks

#### `internal/templates` - Configuration Templates
- Template engine for app-specific config files
- Injects database credentials, domains, secrets
- Generates: `wp-config.php`, `config.production.json`, `.env`, etc.

### App Manifest Schema (TOML)

App definitions live in `apps/` directory with each app in its own subdirectory:

```
apps/
├── examples/                  # Pre-defined, tested apps
│   ├── wordpress/
│   │   ├── app.toml          # App manifest
│   │   └── hooks/            # Lifecycle hook scripts
│   ├── ghost/
│   └── nextcloud/
└── custom/                    # User-defined apps
    └── myapp/
        ├── app.toml
        └── hooks/
```

This structure allows each app to include:
- **app.toml**: App manifest configuration
- **hooks/**: Lifecycle hook scripts (pre-install, post-install, etc.)
- **templates/**: Configuration file templates (optional)

Manifest structure:
```toml
[app]
name = "appname"
type = "php|python|nodejs|ruby|go|static"
version = "latest"
description = "App description"

[runtime]
language = "php|python|node|ruby|go"
version = "8.3"  # Language version

[database]
type = "mysql|postgresql|mongodb|redis"
required = true
name = "optional_custom_name"  # Defaults to ${USER}_<appname>

[install]
method = "download|git|cli_tool|composer|pip|npm"
source = "URL or package name"
command = "optional command to run"
location = "app/"  # Relative to app root

[web]
backend = "apache|http"
port = 8080  # Required for http backend
document_root = "app/subdir"
static_paths = ["/static", "/media"]

[[services]]  # For non-PHP apps
name = "service-name"
command = "gunicorn app:application"
directory = "app/"
port = 8080
startsecs = 15
autorestart = true

[services.environment]
NODE_ENV = "production"
PORT = "8080"

[cron]
[[cron.jobs]]
schedule = "*/5 * * * *"
command = "wp cron event run"
log = "logs/cron.log"

[backup]
include = ["data/", "uploads/"]
exclude = ["cache/", "tmp/"]
```

### Lifecycle Hooks

Apps can include optional hook scripts that run at specific lifecycle events:

**Available Hooks:**
- `pre-install.sh` - Before app installation
- `post-install.sh` - After app installation
- `pre-upgrade.sh` - Before app upgrade
- `post-upgrade.sh` - After app upgrade
- `pre-backup.sh` - Before creating backup
- `post-backup.sh` - After creating backup
- `pre-start.sh` - Before starting service
- `post-start.sh` - After starting service

**Hook Script Format:**
```bash
#!/bin/bash
set -e

APP_ROOT="$1"      # ~/apps/<app-name>
APP_NAME="$2"      # <app-name>
ACTION="$3"        # install|upgrade|backup|start

# Your hook logic here
cd "${APP_ROOT}/app"
npm run setup
```

Hooks must be:
- Executable (`chmod +x`)
- Located in `hooks/` directory
- Named `<hook-name>.sh`
- Handle errors appropriately

See `apps/README.md` for complete hook documentation.

## Uberspace-Specific Integration

### Runtime Version Management

```bash
# Managed via: uberspace tools version
uberspace tools version use php 8.3
uberspace tools version show php
uberspace tools version list php
uberspace tools restart php  # For PHP only
```

### Web Backend Configuration

```bash
# Apache backend (for PHP apps)
uberspace web backend set / --apache

# HTTP backend (for app servers)
uberspace web backend set / --http --port 8080

# Static files
uberspace web backend set /static --apache

# List configured backends
uberspace web backend list
```

### Database Management

Credentials stored in `~/.my.cnf`:
```ini
[client]
user = username
password = secret
```

Database operations:
```bash
# Create database (must prefix with username)
mysql -e "CREATE DATABASE ${USER}_appname"

# Export database
mysqldump ${USER}_appname > backup.sql

# Import database
mysql ${USER}_appname < backup.sql
```

### Process Management (Supervisord)

Service files in `~/etc/services.d/<service>.ini`:
```ini
[program:myapp]
command=gunicorn --bind 0.0.0.0:8000 app:application
directory=%(ENV_HOME)s/apps/myapp/app
startsecs=15
autorestart=true
stdout_logfile=%(ENV_HOME)s/logs/myapp.log
stderr_logfile=%(ENV_HOME)s/logs/myapp-error.log
```

Management commands:
```bash
supervisorctl reread    # Detect config changes
supervisorctl update    # Apply changes
supervisorctl status    # Show all services
supervisorctl restart <service>
```

### Backup System

Uberspace provides automatic backups:
- **Files**: `/backup/` (7 daily + 7 weekly)
- **Databases**: `/mysql_backup/` (21 days)
- Excluded: `no_backup/`, `tmp/`, `.cache/`, `cache/`

## Implementation Patterns

### Language-Specific Installation

**PHP Apps** (WordPress, Nextcloud):
1. Install to DocumentRoot or subdirectory
2. No supervisord service needed (PHP-FPM automatic)
3. Configure PHP via `~/etc/php.d/<app>.ini`
4. Use CLI tools for automation (wp-cli, occ)
5. Set up cron jobs for background tasks

**Python Apps** (Django, Flask):
1. Install outside DocumentRoot (security)
2. Create virtual environment: `python3 -m venv venv`
3. Install dependencies: `pip install -r requirements.txt`
4. Create supervisord service with Gunicorn/uWSGI
5. Collect static files to DocumentRoot
6. Bind to `0.0.0.0:<port>` (NOT localhost)

**Node.js Apps** (Ghost, Express):
1. Install outside DocumentRoot
2. Set Node version: `uberspace tools version use node 20`
3. Install dependencies: `npm install`
4. Create supervisord service
5. Set environment: `NODE_ENV=production`
6. Bind to `0.0.0.0:<port>`

### Safety Features

All operations should implement:
- **Dry-run mode**: Show what would be done without executing
- **Pre-flight checks**: Disk space, dependencies, port availability
- **Automatic backups**: Before upgrades or destructive operations
- **Rollback capability**: Restore previous version on failure
- **Lock files**: Prevent concurrent operations on same app
- **Validation**: Verify manifest schema before execution
- **Verbose logging**: Detailed output for debugging

### Error Handling

- Always return descriptive errors with context
- Include command output in error messages
- Check for Uberspace-specific constraints (database naming, ports)
- Validate prerequisites before starting operations
- Clean up partial installations on failure

## Adding New App Manifests

To add a new pre-defined app:

1. Create `apps/examples/<appname>.toml` with complete manifest
2. Test installation in Uberspace environment
3. Document any special requirements or post-install steps
4. Add configuration templates to `internal/templates/` if needed
5. Update README.md with app support

To create custom app (user-extensible):

1. Run `uberman init <appname>` to scaffold template
2. Edit generated manifest with app-specific details
3. Test with `uberman --dry-run install <appname>`
4. Install with `uberman install <appname>`

## Key Design Decisions

### Why Go?
- Single static binary (easy distribution)
- Excellent standard library for system operations
- Fast compilation and execution
- Strong typing prevents configuration errors
- Cross-compilation support

### Why TOML for Manifests?
- Human-readable and easy to edit
- Strong typing (vs YAML's ambiguity)
- Good Go library support (BurntSushi/toml)
- Clear structure for nested configuration

### Why Convention Over Configuration?
- Reduces cognitive load for Ubernauts
- Standardizes directory layouts across apps
- Makes automation more reliable
- Easier to troubleshoot and debug
- Consistent backup/restore workflows

### Self-Contained App Directories
Each app in `~/apps/<app-name>/` contains everything needed:
- Application code
- Configuration files
- Logs (symlinked to ~/logs/ for centralization)
- Data and uploads
- Local backups
- This enables easy migration, cleanup, and isolation

## Development Guidelines

### Adding New Commands

1. Create command file in `cmd/uberman/<command>.go`
2. Register with cobra in `init()` function
3. Implement core logic in appropriate `internal/` package
4. Add `--dry-run` support
5. Add `--verbose` logging
6. Write tests in `*_test.go`
7. Update README.md and this file

### Code Style

- Follow standard Go conventions (gofmt, golint)
- Use descriptive variable names
- Add comments for exported functions
- Return errors, don't panic
- Use structured logging
- Validate inputs early
- Keep functions focused and testable

### Testing Strategy

- Unit tests for all packages
- Integration tests marked with build tag: `// +build integration`
- Mock Uberspace commands in tests
- Test dry-run mode separately
- Test error conditions
- Verify backup creation before destructive ops

## Future Enhancements

Planned features:
- Multi-app orchestration (install multiple apps)
- App dependencies (app A requires app B)
- Health monitoring and alerts
- Automatic security updates
- Migration tools (import existing manual installations)
- Web UI for non-CLI users
- Ansible/Terraform integration
- Git-based app definitions (remote manifests)

## Related Resources

- Uberspace Manual: https://manual.uberspace.de/
- Uberspace Lab (installation guides): https://lab.uberspace.de/
- Uberspace Lab Git Repository: https://github.com/Uberspace/lab
- Project Issue Tracker: (to be added)
