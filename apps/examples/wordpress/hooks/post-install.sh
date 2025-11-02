#!/bin/bash
# WordPress Post-Install Hook
# This script runs after WordPress has been installed

set -e

APP_ROOT="$1"
APP_NAME="$2"

echo "Running WordPress post-install hook..."

# Configure WordPress
cd "${APP_ROOT}/app/wordpress"

# Generate secret keys
SALTS=$(curl -s https://api.wordpress.org/secret-key/1.1/salt/)

# Create wp-config.php if not exists
if [ ! -f wp-config.php ]; then
    echo "Creating wp-config.php..."

    # Get database credentials
    DB_NAME="${USER}_${APP_NAME}"
    DB_USER="${USER}"
    DB_PASSWORD=$(grep "^password" ~/.my.cnf | cut -d= -f2 | tr -d ' ')

    # Create config from sample
    cp wp-config-sample.php wp-config.php

    # Replace database settings
    sed -i "s/database_name_here/${DB_NAME}/" wp-config.php
    sed -i "s/username_here/${DB_USER}/" wp-config.php
    sed -i "s/password_here/${DB_PASSWORD}/" wp-config.php

    # Replace salt placeholders
    sed -i "/AUTH_KEY/,/NONCE_SALT/d" wp-config.php
    echo "$SALTS" >> wp-config.php
fi

# Set proper permissions
chmod 644 wp-config.php

# Create uploads directory
mkdir -p wp-content/uploads
chmod 755 wp-content/uploads

echo "WordPress post-install hook completed!"
echo ""
echo "Next steps:"
echo "1. Visit your domain to complete WordPress installation"
echo "2. Create an admin user"
echo "3. Configure permalinks"
