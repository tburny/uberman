#!/bin/bash
# Ghost Pre-Start Hook
# This script runs before Ghost service starts

set -e

APP_ROOT="$1"
APP_NAME="$2"

echo "Running Ghost pre-start hook..."

cd "${APP_ROOT}/app"

# Ensure config exists
if [ ! -f config.production.json ]; then
    echo "Error: config.production.json not found!"
    exit 1
fi

# Check database connection
DB_NAME="${USER}_${APP_NAME}"
if ! mysql -e "USE ${DB_NAME}" 2>/dev/null; then
    echo "Error: Cannot connect to database ${DB_NAME}"
    exit 1
fi

# Run database migrations
echo "Running database migrations..."
NODE_ENV=production node current/index.js migrate || true

echo "Ghost pre-start hook completed!"
