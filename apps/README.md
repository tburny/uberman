# App Definitions

This directory contains app definitions for Uberman. Each app is defined in its own subdirectory with a manifest file and optional hook scripts.

## Directory Structure

```
apps/
├── examples/           # Pre-defined, tested apps
│   ├── wordpress/
│   │   ├── app.toml           # App manifest
│   │   ├── hooks/             # Hook scripts (optional)
│   │   │   ├── pre-install.sh
│   │   │   ├── post-install.sh
│   │   │   ├── pre-upgrade.sh
│   │   │   ├── post-upgrade.sh
│   │   │   ├── pre-backup.sh
│   │   │   ├── post-backup.sh
│   │   │   ├── pre-start.sh
│   │   │   └── post-start.sh
│   │   └── templates/         # Config file templates (optional)
│   │       └── wp-config.php.tmpl
│   ├── ghost/
│   └── nextcloud/
└── custom/            # User-defined apps
    └── myapp/
        ├── app.toml
        └── hooks/
```

## App Manifest (`app.toml`)

Each app requires an `app.toml` manifest file that defines:

- App metadata (name, version, description)
- Runtime requirements (PHP, Node.js, Python, etc.)
- Database requirements
- Installation method
- Web backend configuration
- Services (for non-PHP apps)
- Cron jobs
- Backup configuration

**Example:**

```toml
[app]
name = "myapp"
type = "nodejs"
version = "1.0.0"
description = "My awesome application"

[runtime]
language = "node"
version = "20"

[database]
type = "mysql"
required = true

[install]
method = "git"
source = "https://github.com/user/myapp.git"
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

See example apps in `examples/` for complete manifests.

## Hook Scripts

Hook scripts are optional bash scripts that run at specific lifecycle events. They allow customization of the installation, upgrade, and maintenance processes.

### Available Hooks

| Hook | When it Runs | Use Cases |
|------|--------------|-----------|
| `pre-install.sh` | Before app installation | Check prerequisites, prepare environment |
| `post-install.sh` | After app installation | Configure app, create admin user, set permissions |
| `pre-upgrade.sh` | Before app upgrade | Backup custom configs, check compatibility |
| `post-upgrade.sh` | After app upgrade | Run migrations, update database, clear caches |
| `pre-backup.sh` | Before creating backup | Flush caches, export data, prepare for backup |
| `post-backup.sh` | After creating backup | Verify backup, upload to remote storage |
| `pre-start.sh` | Before starting service | Check config, run migrations, verify database |
| `post-start.sh` | After starting service | Warm up caches, verify service health |
| `pre-stop.sh` | Before stopping service | Drain connections, save state |
| `post-stop.sh` | After stopping service | Cleanup temp files |

### Hook Script Template

```bash
#!/bin/bash
# Hook: <hook-name>
# Description: What this hook does

set -e  # Exit on error

# Arguments provided by uberman
APP_ROOT="$1"      # ~/apps/<app-name>
APP_NAME="$2"      # <app-name>
ACTION="$3"        # install|upgrade|backup|start|stop

echo "Running <hook-name> hook for ${APP_NAME}..."

# Your hook logic here

echo "Hook completed successfully!"
```

### Hook Requirements

1. **Executable**: Hooks must be executable (`chmod +x`)
2. **Bash**: Use `#!/bin/bash` shebang
3. **Error handling**: Use `set -e` to exit on errors
4. **Arguments**: Accept APP_ROOT, APP_NAME, ACTION
5. **Output**: Provide clear messages for logging
6. **Idempotent**: Should be safe to run multiple times

### Hook Examples

**Post-Install: Create Admin User**

```bash
#!/bin/bash
set -e

APP_ROOT="$1"
APP_NAME="$2"

cd "${APP_ROOT}/app"

# Create admin user if not exists
if ! php artisan user:exists admin@example.com; then
    php artisan user:create admin@example.com --admin
fi
```

**Post-Upgrade: Run Migrations**

```bash
#!/bin/bash
set -e

APP_ROOT="$1"
cd "${APP_ROOT}/app"

# Run database migrations
php artisan migrate --force

# Clear and rebuild caches
php artisan cache:clear
php artisan config:cache
php artisan route:cache
```

**Pre-Backup: Export Database**

```bash
#!/bin/bash
set -e

APP_ROOT="$1"
APP_NAME="$2"

DB_NAME="${USER}_${APP_NAME}"
BACKUP_DIR="${APP_ROOT}/backups"

mkdir -p "${BACKUP_DIR}"
mysqldump "${DB_NAME}" | gzip > "${BACKUP_DIR}/db-$(date +%Y%m%d-%H%M%S).sql.gz"
```

### Environment Variables

Hooks have access to:

- `$APP_ROOT` - App installation directory (`~/apps/<app-name>`)
- `$APP_NAME` - App name
- `$USER` - Current username
- `$HOME` - User home directory
- All standard shell environment variables

### Hook Execution Order

**Installation:**
1. Create app directory structure
2. Set runtime version
3. Create database
4. Run `pre-install.sh` (if exists)
5. Download/install app code
6. Install dependencies
7. Generate configuration
8. Set up web backend
9. Create services
10. Run `post-install.sh` (if exists)
11. Start services

**Upgrade:**
1. Run `pre-upgrade.sh` (if exists)
2. Create automatic backup
3. Stop services
4. Download new version
5. Install dependencies
6. Run `post-upgrade.sh` (if exists)
7. Start services
8. Verify upgrade

**Backup:**
1. Run `pre-backup.sh` (if exists)
2. Create file backup
3. Export database
4. Run `post-backup.sh` (if exists)

### Best Practices

1. **Logging**: Echo clear messages for debugging
2. **Error handling**: Always use `set -e`
3. **Checks**: Verify prerequisites before acting
4. **Idempotency**: Make scripts safe to re-run
5. **Documentation**: Comment complex operations
6. **Testing**: Test hooks before committing
7. **Security**: Never hardcode passwords or secrets
8. **Permissions**: Set appropriate file permissions

### Debugging Hooks

Enable verbose mode to see hook output:

```bash
uberman --verbose install myapp
```

View hook execution logs:

```bash
tail -f ~/logs/myapp/install.log
```

Test hooks manually:

```bash
cd ~/apps/myapp
./hooks/post-install.sh ~/apps/myapp myapp install
```

## Configuration Templates

Store configuration file templates in `templates/` directory. Templates use Go's text/template syntax.

**Example: `templates/config.json.tmpl`**

```json
{
  "database": {
    "host": "localhost",
    "name": "{{.DatabaseName}}",
    "user": "{{.DatabaseUser}}",
    "password": "{{.DatabasePassword}}"
  },
  "domain": "{{.Domain}}",
  "secret": "{{.SecretKey}}"
}
```

Available template variables:
- `{{.AppName}}` - App name
- `{{.AppRoot}}` - App directory
- `{{.DatabaseName}}` - Database name
- `{{.DatabaseUser}}` - Database user
- `{{.DatabasePassword}}` - Database password
- `{{.Domain}}` - Primary domain
- `{{.SecretKey}}` - Generated secret key

## Creating Custom Apps

1. **Create app directory:**
   ```bash
   mkdir -p apps/custom/myapp
   ```

2. **Create manifest:**
   ```bash
   uberman init myapp
   # Or copy from example:
   cp -r apps/examples/wordpress apps/custom/myapp
   ```

3. **Edit `app.toml`:**
   ```bash
   nano apps/custom/myapp/app.toml
   ```

4. **Add hooks (optional):**
   ```bash
   mkdir -p apps/custom/myapp/hooks
   nano apps/custom/myapp/hooks/post-install.sh
   chmod +x apps/custom/myapp/hooks/post-install.sh
   ```

5. **Test installation:**
   ```bash
   uberman --dry-run install myapp
   uberman install myapp
   ```

## Contributing Apps

To contribute a new app definition to `examples/`:

1. Create app directory with manifest and hooks
2. Test thoroughly on Uberspace
3. Document any special requirements
4. Add templates if needed
5. Submit pull request

See [CONTRIBUTING.md](../CONTRIBUTING.md) for details.

## App Examples

Browse `examples/` for complete, tested app definitions:

- **WordPress** - PHP CMS with wp-cli integration
- **Ghost** - Node.js blog platform
- **Nextcloud** - PHP file sync and sharing

Each example includes:
- Complete `app.toml` manifest
- Lifecycle hooks
- Configuration templates (where needed)
- Comments and documentation

## Support

- **Documentation**: See [CLAUDE.md](../CLAUDE.md)
- **Issues**: https://github.com/tburny/uberman/issues
- **Examples**: Check `apps/examples/`

---

*For more details on manifest schema, see [CLAUDE.md](../CLAUDE.md#app-manifest-schema-toml)*
