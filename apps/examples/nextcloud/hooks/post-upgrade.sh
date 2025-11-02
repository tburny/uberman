#!/bin/bash
# Nextcloud Post-Upgrade Hook
# This script runs after Nextcloud has been upgraded

set -e

APP_ROOT="$1"
APP_NAME="$2"

echo "Running Nextcloud post-upgrade hook..."

cd "${APP_ROOT}/app/nextcloud"

# Run Nextcloud upgrade process
echo "Running Nextcloud upgrade..."
php occ upgrade --no-interaction

# Add missing indices
echo "Adding missing database indices..."
php occ db:add-missing-indices --no-interaction

# Convert filecache to bigint if needed
echo "Converting filecache to bigint..."
php occ db:convert-filecache-bigint --no-interaction || true

# Update all apps
echo "Updating apps..."
php occ app:update --all

# Clear caches
echo "Clearing caches..."
php occ files:scan-app-data
php occ files:cleanup

# Restart notify_push if it exists
if supervisorctl status nextcloud_notify_push >/dev/null 2>&1; then
    echo "Restarting notify_push..."
    supervisorctl restart nextcloud_notify_push
fi

echo "Nextcloud post-upgrade hook completed!"
