# Uberman Quick Start Guide

This guide will help you get started with Uberman in minutes.

## Prerequisites

- Uberspace hosting account
- SSH access to your Uberspace
- Basic familiarity with command line

## Installation (On Uberspace)

### Quick Install (Recommended)

Install with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
```

This automatically:
- Detects your platform (Linux amd64 on Uberspace)
- Downloads the latest release
- Installs to `~/bin/uberman`
- Verifies the installation

### Alternative: Build from Source

If you prefer to build from source:

```bash
# Install Go if needed
uberspace tools version use go 1.21

# Clone and build
cd ~/projekte
git clone https://github.com/tburny/uberman.git
cd uberman
make install-local

# Verify installation
uberman --version
```

## Your First App: WordPress

Let's install WordPress as a first example.

### Step 1: Preview the Installation (Dry Run)

```bash
# See what would happen without making changes
uberman --dry-run --verbose install wordpress
```

This shows you:
- Directory structure that will be created
- Runtime version that will be set
- Database that will be created
- All commands that would be executed

### Step 2: Install WordPress

```bash
# Install WordPress for real
uberman install wordpress
```

This will:
1. Create `~/apps/wordpress/` directory structure
2. Set PHP version to 8.3
3. Create MySQL database `${USER}_wordpress`
4. Download WordPress files
5. Generate `wp-config.php` with database credentials
6. Configure web backend
7. Set up cron jobs for background tasks

### Step 3: Complete WordPress Setup

```bash
# Access your WordPress site
# Open: https://yourusername.uber.space

# Follow the WordPress installation wizard
# Create admin user, set site title, etc.
```

### Step 4: Check Status

```bash
# Check WordPress status
uberman status wordpress

# View logs
tail -f ~/logs/wordpress/
```

## Installing Other Apps

### Ghost (Node.js Blog)

```bash
# Install Ghost blog platform
uberman install ghost

# Check service status
supervisorctl status ghost

# Access Ghost
# Open: https://yourusername.uber.space
```

### Nextcloud (File Sync)

```bash
# Install Nextcloud
uberman install nextcloud

# Complete setup via web interface
# Open: https://yourusername.uber.space
```

## Creating Custom Apps

### Step 1: Generate Manifest Template

```bash
# Create a new app manifest
uberman init myapp
```

This creates `apps/custom/myapp.toml` with a template.

### Step 2: Edit the Manifest

```bash
# Edit the manifest
nano apps/custom/myapp.toml
```

Example minimal Node.js app:

```toml
[app]
name = "myapp"
type = "nodejs"
version = "1.0.0"
description = "My custom Node.js app"

[runtime]
language = "node"
version = "20"

[database]
type = "mysql"
required = true

[install]
method = "git"
source = "https://github.com/yourusername/myapp.git"
location = "app/"

[web]
backend = "http"
port = 8080

[[services]]
name = "myapp"
command = "npm start"
directory = "app/"
port = 8080

[backup]
include = ["data/", "uploads/"]
exclude = ["node_modules/", "tmp/"]
```

### Step 3: Install Your App

```bash
# Test with dry-run first
uberman --dry-run install myapp

# Install for real
uberman install myapp
```

## Common Operations

### Backup

```bash
# Create a manual backup
uberman backup wordpress

# List available backups
ls ~/apps/wordpress/backups/
```

### Restore

```bash
# Restore from a specific backup
uberman restore wordpress 2025-11-02-14-30-00

# Restore from latest backup
uberman restore wordpress latest
```

### Upgrade

```bash
# Upgrade to latest version (with automatic backup)
uberman upgrade wordpress

# If upgrade fails, automatic rollback occurs
```

### Status Check

```bash
# Check app status
uberman status wordpress

# Check all installed apps
uberman list
```

### Uninstall

```bash
# Remove an app (creates final backup first)
uberman uninstall wordpress
```

## Directory Structure

After installing an app, you'll have:

```
~/apps/wordpress/
├── .uberman.toml          # App configuration
├── app/                   # WordPress files
│   └── wordpress/
├── data/                  # Persistent data
├── logs/                  # Application logs
├── backups/               # Local backups
└── tmp/                   # Temporary files
```

Logs are also symlinked from `~/logs/wordpress/` for easy access.

## Useful Commands

```bash
# Show help
uberman --help
uberman install --help

# Verbose output for debugging
uberman --verbose install myapp

# Dry-run to preview changes
uberman --dry-run upgrade myapp

# Check app health
uberman status myapp

# View logs
tail -f ~/logs/myapp/

# Check service status
supervisorctl status

# Check web backends
uberspace web backend list

# Check databases
mysql -e "SHOW DATABASES"
```

## Troubleshooting

### Port Already in Use

```bash
# Check used ports
uberspace web backend list

# Edit manifest to use different port
nano apps/custom/myapp.toml
```

### Service Won't Start

```bash
# Check service logs
supervisorctl tail myapp stderr

# Check service status
supervisorctl status myapp

# Restart service
supervisorctl restart myapp
```

### Database Connection Issues

```bash
# Check database exists
mysql -e "SHOW DATABASES" | grep myapp

# Check credentials
cat ~/.my.cnf

# Test connection
mysql yourusername_myapp -e "SELECT 1"
```

### App Won't Load in Browser

```bash
# Check web backend
uberspace web backend list

# Check service is running
supervisorctl status

# Check logs for errors
tail -f ~/logs/myapp/
```

## Next Steps

1. **Explore Examples**: Check `apps/examples/` for more app manifests
2. **Read Documentation**: See `CLAUDE.md` for detailed architecture
3. **Contribute**: Add your own app manifests! See `CONTRIBUTING.md`
4. **Get Help**: Open issues on GitHub or ask in Uberspace forum

## Tips

- Always test with `--dry-run` first
- Keep backups before upgrades
- Use verbose mode for debugging
- Check logs when things go wrong
- Follow Uberspace best practices
- Read app-specific documentation

## Resources

- **Uberman Docs**: [README.md](../README.md)
- **Uberspace Manual**: https://manual.uberspace.de/
- **Uberspace Lab**: https://lab.uberspace.de/
- **GitHub Issues**: https://github.com/tburny/uberman/issues

Happy deploying! 🚀
