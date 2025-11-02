#!/bin/bash
# Verification script for Docker testing setup
# This script checks that all Docker testing infrastructure is properly configured

set -e  # Exit on error

echo "========================================"
echo "Verifying Docker Testing Setup"
echo "========================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check Docker availability
echo -n "1. Checking Docker installation... "
if command -v docker &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
    docker --version
else
    echo -e "${RED}✗${NC}"
    echo "   ERROR: Docker not found. Please install Docker."
    exit 1
fi
echo ""

# Check Docker Compose
echo -n "2. Checking Docker Compose... "
if command -v docker compose &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
    docker compose version
else
    echo -e "${RED}✗${NC}"
    echo "   ERROR: Docker Compose not found."
    exit 1
fi
echo ""

# Check Docker daemon
echo -n "3. Checking Docker daemon... "
if docker ps &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC}"
    echo "   ERROR: Docker daemon not running. Start Docker and try again."
    exit 1
fi
echo ""

# Check required files
echo "4. Checking required files..."
FILES=(
    "Dockerfile.test"
    "docker-compose.test.yml"
    ".dockerignore"
    "internal/testutil/docker.go"
    "docs/TESTING.md"
    "docs/DOCKER_TESTING.md"
)

for file in "${FILES[@]}"; do
    echo -n "   - $file... "
    if [ -f "$file" ]; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗${NC}"
        echo "   ERROR: File not found: $file"
        exit 1
    fi
done
echo ""

# Validate docker-compose.yml syntax
echo -n "5. Validating docker-compose.test.yml... "
if docker compose -f docker-compose.test.yml config --quiet; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC}"
    echo "   ERROR: docker-compose.test.yml has syntax errors"
    exit 1
fi
echo ""

# Check Makefile targets
echo "6. Checking Makefile targets..."
TARGETS=(
    "test-docker-build"
    "test-docker"
    "test-docker-properties"
    "test-docker-integration"
    "test-docker-coverage"
    "clean-docker"
)

for target in "${TARGETS[@]}"; do
    echo -n "   - $target... "
    if make -n "$target" &> /dev/null; then
        echo -e "${GREEN}✓${NC}"
    else
        echo -e "${RED}✗${NC}"
        echo "   ERROR: Makefile target not found: $target"
        exit 1
    fi
done
echo ""

# Check Go modules
echo -n "7. Checking Go modules... "
if go mod verify &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${YELLOW}⚠${NC} Warning: go mod verify failed (run 'go mod tidy')"
fi
echo ""

# Check testcontainers-go dependency
echo -n "8. Checking testcontainers-go dependency... "
if go list -m github.com/testcontainers/testcontainers-go &> /dev/null; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC}"
    echo "   ERROR: testcontainers-go not found in go.mod"
    exit 1
fi
echo ""

# Summary
echo "========================================"
echo -e "${GREEN}All checks passed!${NC}"
echo "========================================"
echo ""
echo "Next steps:"
echo "  1. Build the Docker test image:"
echo "     make test-docker-build"
echo ""
echo "  2. Run property-based tests:"
echo "     make test-docker-properties"
echo ""
echo "  3. View documentation:"
echo "     cat docs/TESTING_CHEATSHEET.md"
echo ""
echo "For more information, see:"
echo "  - docs/TESTING.md (comprehensive guide)"
echo "  - docs/DOCKER_TESTING.md (Docker details)"
echo "  - docs/TESTING_CHEATSHEET.md (quick reference)"
echo ""
