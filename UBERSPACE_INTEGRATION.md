# Uberspace Platform Integration

Technical reference for Uberspace-specific commands and constraints.

## Platform Constraints

**Critical limitations**:
- **No Docker/containers**: Platform does not support containerization
- **No root/system packages**: Cannot use apt/yum/pacman
- **User-space only**: All operations in user context
- **Supervisord**: Long-running processes managed via supervisord (not systemd)
- **Database naming**: MySQL databases must use `${USER}_<appname>` pattern
- **Port binding**: Applications must bind to `0.0.0.0` or `::` (NOT localhost/127.0.0.1)

## Runtime Version Management

### Commands
```bash
# Set runtime version
uberspace tools version use php 8.3
uberspace tools version use python 3.11
uberspace tools version use node 20

# Show current version
uberspace tools version show php

# List available versions
uberspace tools version list php

# Restart PHP-FPM (PHP only)
uberspace tools restart php
```

### Supported Runtimes
- PHP (multiple versions)
- Python (multiple versions)
- Node.js (multiple versions)
- Ruby (multiple versions)
- Go (multiple versions)

## Web Backend Configuration

### Commands
```bash
# Apache backend (for PHP apps)
uberspace web backend set / --apache
uberspace web backend set /subpath --apache

# HTTP backend (for app servers)
uberspace web backend set / --http --port 8080
uberspace web backend set /api --http --port 3000

# List configured backends
uberspace web backend list

# Delete backend
uberspace web backend del /
uberspace web backend del /subpath
```

### Backend Types

**Apache**:
- Direct PHP serving via PHP-FPM
- No supervisord service needed
- DocumentRoot configurable

**HTTP**:
- Reverse proxy to application server
- Requires port specification (1024-65535)
- Requires supervisord service for app server

## Database Management

### Credentials

Location: `~/.my.cnf`
```ini
[client]
user = username
password = secret_password
```

### MySQL Commands

```bash
# Create database (MUST use ${USER}_ prefix)
mysql -e "CREATE DATABASE ${USER}_appname"

# Check if database exists
mysql -e "SHOW DATABASES LIKE '${USER}_appname'" | grep -q appname
# Exit code 0 = exists, 1 = not exists

# List all user databases
mysql -e "SHOW DATABASES LIKE '${USER}_%'"

# Export database
mysqldump ${USER}_appname > backup.sql
mysqldump ${USER}_appname | gzip > backup.sql.gz

# Import database
mysql ${USER}_appname < backup.sql
gunzip < backup.sql.gz | mysql ${USER}_appname

# Drop database
mysql -e "DROP DATABASE ${USER}_appname"
```

### Database Naming Convention

**Pattern**: `${USER}_<appname>`

**Examples**:
- User: `johndoe`, App: `myapp` → Database: `johndoe_myapp`
- User: `alice`, App: `blog` → Database: `alice_blog`

**Invalid**:
- `myapp` (missing user prefix)
- `johndoe-myapp` (wrong separator)

## Process Management (Supervisord)

### Configuration

**Location**: `~/etc/services.d/<servicename>.ini`

**Format**:
```ini
[program:myapp]
command=/home/%(ENV_USER)s/.local/bin/gunicorn --bind 0.0.0.0:8000 app:application
directory=%(ENV_HOME)s/apps/myapp/app
startsecs=15
autorestart=true
stdout_logfile=%(ENV_HOME)s/logs/myapp.log
stderr_logfile=%(ENV_HOME)s/logs/myapp-error.log

[program:myapp]
environment=NODE_ENV=production,PORT=8080
```

### Management Commands

```bash
# Reload configuration (detect new/changed .ini files)
supervisorctl reread

# Apply changes (start new services)
supervisorctl update

# Check status
supervisorctl status
supervisorctl status myapp

# Start service
supervisorctl start myapp

# Stop service
supervisorctl stop myapp

# Restart service
supervisorctl restart myapp

# View logs
tail -f ~/logs/myapp.log
tail -f ~/logs/myapp-error.log
```

### Common Patterns

**Node.js app**:
```ini
[program:nodeapp]
command=node server.js
directory=%(ENV_HOME)s/apps/nodeapp/app
environment=NODE_ENV=production,PORT=8080
```

**Python app (Gunicorn)**:
```ini
[program:pythonapp]
command=gunicorn --bind 0.0.0.0:8000 app:application
directory=%(ENV_HOME)s/apps/pythonapp/app
environment=PYTHONPATH=%(ENV_HOME)s/apps/pythonapp/app
```

**Ruby app (Puma)**:
```ini
[program:rubyapp]
command=bundle exec puma -b tcp://0.0.0.0:9292
directory=%(ENV_HOME)s/apps/rubyapp/app
```

## Port Allocation

### Available Range
- **User ports**: 1024-65535
- **Privileged ports**: 1-1023 (NOT available)

### Finding Available Ports

```bash
# Check if port is in use
ss -tulpn | grep :8080
# Empty output = port available
# Output with process = port in use

# Find next available port in range
for port in {8000..8100}; do
    if ! ss -tulpn | grep -q ":$port "; then
        echo $port
        break
    fi
done
```

### Port Binding Rules

**MUST bind to**:
- `0.0.0.0` (all IPv4 interfaces)
- `::` (all IPv6 interfaces)

**DO NOT bind to**:
- `localhost` or `127.0.0.1` (will not be accessible via web backend)
- Specific IP addresses (may not work with platform routing)

## Backup System

### Automatic Backups

**File backups**: `/backup/`
- 7 daily backups
- 7 weekly backups

**Database backups**: `/mysql_backup/`
- 21 days of daily backups

### Excluded Patterns

Directories automatically excluded from backups:
- `no_backup/`
- `tmp/`
- `.cache/`
- `cache/`

### Manual Backup Access

```bash
# List available file backups
ls -lh /backup/

# List available database backups
ls -lh /mysql_backup/

# Restore from backup
cp /backup/YYYY-MM-DD/path/to/file ~/destination/
```

## Directory Structure Convention

**Standard app layout**:
```
~/apps/<appname>/
├── .uberman.toml      # App metadata (optional)
├── app/               # Application code
├── data/              # Persistent data
├── logs/              # Application logs (may be symlinked to ~/logs/)
├── backups/           # Local backups
└── tmp/               # Temporary files (excluded from backups)
```

## Language-Specific Patterns

### PHP Apps
1. Install to DocumentRoot or subdirectory
2. No supervisord service needed (PHP-FPM automatic)
3. Configure via `~/etc/php.d/<app>.ini` if needed
4. Use CLI tools for automation (e.g., WP-CLI)

### Python Apps
1. Install outside DocumentRoot (security)
2. Create virtual environment: `python3 -m venv venv`
3. Install dependencies: `pip install -r requirements.txt`
4. Create supervisord service with Gunicorn/uWSGI
5. Collect static files to accessible location
6. **MUST** bind to `0.0.0.0:<port>`

### Node.js Apps
1. Install outside DocumentRoot
2. Set Node version: `uberspace tools version use node 20`
3. Install dependencies: `npm install`
4. Create supervisord service
5. Set environment: `NODE_ENV=production`
6. **MUST** bind to `0.0.0.0:<port>`

## Resources

- **Uberspace Manual**: https://manual.uberspace.de/
- **Uberspace Lab**: https://lab.uberspace.de/ (manual installation guides)
- **Uberspace Lab Git**: https://github.com/Uberspace/lab
