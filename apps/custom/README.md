# Custom App Definitions

This directory is for your custom app definitions. Apps defined here are user-specific and not committed to the repository.

## Creating a Custom App

### Method 1: Using uberman init (Coming Soon)

```bash
uberman init myapp
```

This will create:
```
apps/custom/myapp/
├── app.toml          # App manifest template
└── hooks/            # Hook scripts directory
    └── README.md
```

### Method 2: Copy from Example

```bash
cp -r apps/examples/wordpress apps/custom/myapp
cd apps/custom/myapp
nano app.toml  # Edit manifest
```

### Method 3: Create from Scratch

```bash
mkdir -p apps/custom/myapp/hooks
touch apps/custom/myapp/app.toml
```

Then edit `app.toml` with your app configuration.

## App Directory Structure

Each custom app should follow this structure:

```
apps/custom/myapp/
├── app.toml                   # Required: App manifest
├── hooks/                     # Optional: Lifecycle hooks
│   ├── pre-install.sh
│   ├── post-install.sh
│   ├── pre-upgrade.sh
│   ├── post-upgrade.sh
│   ├── pre-backup.sh
│   ├── post-backup.sh
│   ├── pre-start.sh
│   └── post-start.sh
└── templates/                 # Optional: Config templates
    └── config.json.tmpl
```

## Minimal app.toml Example

```toml
[app]
name = "myapp"
type = "nodejs"
version = "1.0.0"
description = "My custom app"

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
```

## Adding Hooks

Create executable bash scripts in the `hooks/` directory:

```bash
mkdir -p apps/custom/myapp/hooks

cat > apps/custom/myapp/hooks/post-install.sh << 'EOF'
#!/bin/bash
set -e

APP_ROOT="$1"
APP_NAME="$2"

echo "Running post-install hook..."
cd "${APP_ROOT}/app"

# Your post-install logic here
npm run setup

echo "Post-install complete!"
EOF

chmod +x apps/custom/myapp/hooks/post-install.sh
```

## Testing Your App

### Dry Run

Test without making changes:

```bash
uberman --dry-run install myapp
```

### Verbose Mode

See detailed output:

```bash
uberman --verbose install myapp
```

### Install

Install for real:

```bash
uberman install myapp
```

## Best Practices

1. **Version Control**: Keep your custom apps in a separate git repository
2. **Testing**: Always test with `--dry-run` first
3. **Hooks**: Use hooks for app-specific setup
4. **Documentation**: Add comments to your manifest
5. **Templates**: Use templates for configuration files
6. **Backup**: Test backup/restore before relying on it

## Examples

See `apps/examples/` for complete, tested examples:
- `wordpress/` - PHP application with hooks
- `ghost/` - Node.js application
- `nextcloud/` - PHP with complex setup

## Support

- **Documentation**: See [apps/README.md](../README.md)
- **Schema**: See [CLAUDE.md](../../CLAUDE.md)
- **Issues**: https://github.com/tburny/uberman/issues
