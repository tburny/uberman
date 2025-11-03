# Uberman - Uberspace App Management System

> **⚠️ REFACTORING IN PROGRESS**
> This document references code deleted during 2025-11-03 Clean Architecture refactoring.
> See CLAUDE.md, ARCHITECTURE.md, PLANNING.md for current state.
> Will update after refactoring completes.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/tburny/uberman)](https://github.com/tburny/uberman/releases)
[![Build Status](https://github.com/tburny/uberman/workflows/Test/badge.svg)](https://github.com/tburny/uberman/actions)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

**Uberman** is a CLI tool for reproducible installation, upgrades, backups, and deployment of applications on [Uberspace](https://uberspace.de) hosting. It automates the manual installation patterns from [Uberspace Lab](https://lab.uberspace.de/) while following a "convention over configuration" approach.

## Features

- **Reproducible Installations**: Install apps with a single command using TOML manifests
- **Safe Upgrades**: Automatic backups before upgrades with rollback capability
- **Convention Over Configuration**: Self-contained app directories at `~/apps/<app-name>`
- **Lifecycle Hooks**: Customize installation, upgrades, and backups with bash scripts
- **Multi-Language Support**: PHP, Python, Node.js, Ruby, Go applications
- **Uberspace Integration**: Native support for `uberspace` commands, supervisord, MySQL
- **Dry-Run Mode**: Preview changes before executing
- **User-Extensible**: Create custom app definitions or use pre-defined ones
- **Requirements Engineering Agent**: AI-powered EARS requirements toolkit (see below)

## Requirements Engineering Toolkit

Uberman includes a comprehensive **Requirements Engineering Agent** that helps you write, analyze, and improve requirements using **EARS (Easy Approach to Requirements Syntax)** - an industry-standard notation developed at Rolls-Royce and adopted by Airbus, NASA, Bosch, and others.

### Features

- **EARS Format Conversion**: Transform informal descriptions into properly formatted EARS requirements
- **Quality Analysis**: Assess requirements against INCOSE and IEEE 29148 standards
- **Interactive Elicitation**: Guided requirements gathering through conversation
- **Batch Processing**: Analyze multiple requirements with comprehensive reports
- **Standards-Based**: Follows INCOSE 42-rule framework and IEEE 29148 guidelines

### Quick Start

**Interactive skill** (comprehensive mode):
```
@requirements-engineer

I need help writing requirements for [your feature]
```

**Quick commands** (specific tasks):
```bash
/requirements-convert    # Convert informal to EARS format
/requirements-analyze    # Analyze quality and suggest improvements
/requirements-elicit     # Start interactive elicitation session
/requirements-batch      # Process multiple requirements
```

### Documentation

- **[EARS Guide](docs/EARS_GUIDE.md)** - Complete EARS reference with all 5 patterns
- **[Quality Framework](docs/REQUIREMENTS_QUALITY.md)** - INCOSE/IEEE standards and scoring
- **[Usage Guide](docs/REQUIREMENTS_ENGINEERING_AGENT.md)** - How to use the agent
- **[Examples](docs/REQUIREMENTS_EXAMPLES.md)** - Real-world requirements across domains
- **[Templates](templates/requirements/)** - Worksheets and checklists

### Example

**Input (informal)**:
```
Users should be able to export their data quickly
```

**Output (EARS-formatted)**:
```
[REQ-EXPORT-001] When a user clicks the Export Data button, the application
shall generate a CSV file containing all user data and shall initiate download
within 3 seconds.

Pattern: Event-Driven (When)
Quality: Clarity 5/5, Testability 5/5, Completeness 5/5
```

## Supported Apps

Pre-defined app manifests are available for:

- **WordPress** - Popular CMS (PHP)
- **Ghost** - Modern publishing platform (Node.js)
- **Nextcloud** - Self-hosted file sync (PHP)

Additional apps can be defined using TOML manifests.

## Installation

### One-Line Install (Recommended)

Install the latest version with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
```

This will automatically:
- Detect your OS and architecture
- Download the latest release
- Install to `~/bin/uberman`
- Verify the installation

**Custom install location:**
```bash
# Install to /usr/local/bin (requires sudo)
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | sudo INSTALL_DIR=/usr/local/bin bash

# Install to custom directory
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | INSTALL_DIR=~/.local/bin bash
```

### Manual Install

Download and install manually:

```bash
# On Linux (Uberspace)
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-linux-amd64.tar.gz | tar xz
mv uberman-linux-amd64 ~/bin/uberman
chmod +x ~/bin/uberman

# On macOS (Intel)
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-darwin-amd64.tar.gz | tar xz
mv uberman-darwin-amd64 ~/bin/uberman
chmod +x ~/bin/uberman

# On macOS (Apple Silicon)
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-darwin-arm64.tar.gz | tar xz
mv uberman-darwin-arm64 ~/bin/uberman
chmod +x ~/bin/uberman

# Verify installation
uberman --version
```

**Available Platforms:**
- Linux (amd64, arm64)
- macOS (amd64/Intel, arm64/Apple Silicon)
- Windows (amd64)
- FreeBSD (amd64)

See [Releases](https://github.com/tburny/uberman/releases) for all versions and platforms.

### Build from Source

```bash
# Prerequisites: Go 1.21 or higher
# On Uberspace: uberspace tools version use go 1.21

# Clone the repository
git clone https://github.com/tburny/uberman.git
cd uberman

# Build and install
make install-local

# Verify installation
uberman --version
```

### Local Development

```bash
# Clone the repository
git clone https://github.com/tburny/uberman.git
cd uberman

# Install dependencies
go mod download

# Build
go build -o bin/uberman cmd/uberman/main.go

# Run
./bin/uberman --help
```

## Quick Start

### Install WordPress

```bash
# Install WordPress with all dependencies
uberman install wordpress

# Check status
uberman status wordpress
```

### Install Ghost Blog

```bash
# Install Ghost with Node.js and MySQL
uberman install ghost

# View logs
tail -f ~/logs/ghost/ghost.log
```

### Create Custom App

```bash
# Generate app manifest template
uberman init myapp

# Edit the manifest
nano apps/custom/myapp.toml

# Install your app
uberman install myapp
```

## Usage

### Core Commands

```bash
# Install an app
uberman install <app-name>

# Upgrade an app (with automatic backup)
uberman upgrade <app-name>

# Create manual backup
uberman backup <app-name>

# Restore from backup
uberman restore <app-name> <backup-id>

# Check app status and health
uberman status <app-name>

# List installed apps
uberman list

# Deploy from git repository
uberman deploy <app-name>

# Create new app manifest
uberman init <app-name>
```

### Global Flags

```bash
# Dry-run mode (show what would be done)
uberman --dry-run install wordpress

# Verbose output
uberman --verbose upgrade ghost

# Custom config file
uberman --config ~/.uberman.toml status myapp
```

## App Directory Structure

Each app is installed in a self-contained directory:

```
~/apps/<app-name>/
├── .uberman.toml          # App instance configuration
├── app/                   # Application code/files
├── data/                  # Persistent application data
├── logs/                  # Application logs (symlinked to ~/logs/)
├── backups/               # Local backup storage
└── tmp/                   # Temporary files (excluded from backups)
```

## App Manifest Format

App definitions use TOML format. Example:

```toml
[app]
name = "myapp"
type = "nodejs"
version = "latest"
description = "My awesome app"

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

[backup]
include = ["data/", "uploads/"]
exclude = ["cache/", "tmp/"]
```

See `apps/examples/` for complete examples.

## How It Works

Uberman automates the typical Uberspace installation workflow:

1. **Runtime Setup**: Sets language version using `uberspace tools version`
2. **Database Creation**: Creates MySQL database with proper naming (`${USER}_<app>`)
3. **App Download**: Fetches source code via download, git, or package manager
4. **Configuration**: Generates config files with injected credentials
5. **Service Setup**: Creates supervisord services for long-running processes
6. **Web Backend**: Configures web routing via `uberspace web backend`
7. **Verification**: Checks service status and connectivity

All operations respect Uberspace constraints:
- No Docker or system packages
- User-space operations only
- Proper database naming conventions
- Correct port binding (`0.0.0.0`, not `localhost`)

## Safety Features

- **Automatic Backups**: Before upgrades or destructive operations
- **Dry-Run Mode**: Preview changes without execution
- **Rollback Capability**: Restore previous version on failure
- **Lock Files**: Prevent concurrent operations on same app
- **Validation**: Verify manifests before execution
- **Pre-Flight Checks**: Disk space, dependencies, port availability

## Development

### Project Structure

```
uberman/
├── cmd/uberman/           # CLI entry point
├── internal/              # Internal packages
│   ├── config/           # TOML parsing and validation
│   ├── runtime/          # Runtime version management
│   ├── database/         # MySQL/PostgreSQL operations
│   ├── web/              # Web backend configuration
│   ├── supervisor/       # Supervisord service management
│   ├── appdir/           # Directory structure management
│   ├── backup/           # Backup/restore logic
│   ├── deploy/           # Deployment strategies
│   └── templates/        # Configuration templates
├── apps/                 # App manifests
│   ├── examples/         # Pre-defined apps
│   └── custom/           # User-defined apps
└── docs/                 # Documentation
```

### Building

```bash
# Build for current platform
go build -o bin/uberman cmd/uberman/main.go

# Build for Linux (Uberspace)
GOOS=linux GOARCH=amd64 go build -o bin/uberman-linux cmd/uberman/main.go

# Install to $GOPATH/bin
go install ./cmd/uberman
```

### Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run integration tests (requires Uberspace)
go test -tags=integration ./...
```

## Contributing

Contributions are welcome! To add a new app manifest:

1. Create `apps/examples/<appname>.toml`
2. Test installation on Uberspace
3. Document special requirements
4. Submit pull request

See [CLAUDE.md](CLAUDE.md) for detailed development guidelines.

## Uberspace Resources

- **Uberspace Manual**: https://manual.uberspace.de/
- **Uberspace Lab** (installation guides): https://lab.uberspace.de/
- **Lab Git Repository**: https://github.com/Uberspace/lab

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Documentation

### User Documentation

- **Reference**: [CLI Commands](docs/cli-reference.md) | [Manifest Schema](docs/manifest-schema.md) | [Error Codes](docs/error-codes.md)
- **Guides**: [Creating Manifests](docs/creating-manifests.md) | [Troubleshooting](docs/troubleshooting.md)
- **Apps**: See `apps/<app_name>/README.md` for installation guides
- **Requirements**: [EARS Guide](docs/EARS_GUIDE.md) | [Quality Framework](docs/REQUIREMENTS_QUALITY.md) | [Agent Usage](docs/REQUIREMENTS_ENGINEERING_AGENT.md)

### Developer Documentation

- **Architecture**: [ARCHITECTURE.md](ARCHITECTURE.md) | [Clean Architecture](CLAUDE.md) | [Planning](PLANNING.md)
- **Decisions**: See [decisions/](decisions/) for architectural decision records (ADRs)
- **AI Collaboration**: See [docs/ai-copilot/](docs/ai-copilot/) for lessons learned working with AI

### Project Management

- **Shape Up**: [Methodology](plans/README.md) | [Current Cycle](plans/cycles/) | [Pitches](plans/pitches/)
- **Ubiquitous Language**: [Domain Glossary](UBIQUITOUS_LANGUAGE.md)

## Acknowledgments

- Inspired by [Uberspace Lab](https://lab.uberspace.de/) installation guides
- Built for the Uberspace hosting community
- Thanks to all Uberspace Lab contributors

## Support

- **Issues**: [GitHub Issues](https://github.com/tburny/uberman/issues)
- **Discussions**: [GitHub Discussions](https://github.com/tburny/uberman/discussions)
- **Uberspace Forum**: https://forum.uberspace.de/

---

**Note**: This is an unofficial tool and is not affiliated with Uberspace. Always maintain your own backups and test thoroughly before using in production.
