#!/bin/bash
# Helper script for creating conventional commits

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Conventional Commit Helper${NC}"
echo "Following the Conventional Commits 1.0.0 specification"
echo ""

# Ask for commit type
echo "Select commit type:"
echo "  1) feat     - New feature"
echo "  2) fix      - Bug fix"
echo "  3) docs     - Documentation changes"
echo "  4) style    - Code style changes (formatting, etc.)"
echo "  5) refactor - Code refactoring"
echo "  6) perf     - Performance improvement"
echo "  7) test     - Adding or updating tests"
echo "  8) build    - Build system or dependencies"
echo "  9) ci       - CI/CD configuration"
echo " 10) chore    - Other changes"
echo " 11) revert   - Revert a previous commit"
read -p "Enter number [1-11]: " type_num

case $type_num in
  1) TYPE="feat" ;;
  2) TYPE="fix" ;;
  3) TYPE="docs" ;;
  4) TYPE="style" ;;
  5) TYPE="refactor" ;;
  6) TYPE="perf" ;;
  7) TYPE="test" ;;
  8) TYPE="build" ;;
  9) TYPE="ci" ;;
  10) TYPE="chore" ;;
  11) TYPE="revert" ;;
  *) echo -e "${RED}Invalid selection${NC}"; exit 1 ;;
esac

# Ask for scope
echo ""
echo "Common scopes: config, runtime, database, web, supervisor, appdir, backup, deploy, cli, manifest"
read -p "Enter scope (optional, press enter to skip): " SCOPE

# Ask for description
echo ""
read -p "Enter short description (max 72 chars, no period at end): " DESCRIPTION

# Validate description length
if [ ${#DESCRIPTION} -gt 72 ]; then
  echo -e "${YELLOW}Warning: Description is ${#DESCRIPTION} characters (recommended max: 72)${NC}"
  read -p "Continue anyway? [y/N]: " continue
  if [[ ! $continue =~ ^[Yy]$ ]]; then
    exit 1
  fi
fi

# Ask for breaking change
read -p "Is this a breaking change? [y/N]: " BREAKING

# Ask for body
echo ""
echo "Enter commit body (optional, explain why this change is being made)"
echo "Press Ctrl+D when done, or just Ctrl+D to skip:"
BODY=$(cat)

# Ask for footer
echo ""
read -p "Enter footer (e.g., 'Closes #123', optional): " FOOTER

# Build commit message
if [ -n "$SCOPE" ]; then
  if [[ $BREAKING =~ ^[Yy]$ ]]; then
    HEADER="${TYPE}(${SCOPE})!: ${DESCRIPTION}"
  else
    HEADER="${TYPE}(${SCOPE}): ${DESCRIPTION}"
  fi
else
  if [[ $BREAKING =~ ^[Yy]$ ]]; then
    HEADER="${TYPE}!: ${DESCRIPTION}"
  else
    HEADER="${TYPE}: ${DESCRIPTION}"
  fi
fi

COMMIT_MSG="$HEADER"

if [ -n "$BODY" ]; then
  COMMIT_MSG="$COMMIT_MSG

$BODY"
fi

if [[ $BREAKING =~ ^[Yy]$ ]]; then
  read -p "Describe the breaking change: " BREAKING_DESC
  COMMIT_MSG="$COMMIT_MSG

BREAKING CHANGE: $BREAKING_DESC"
fi

if [ -n "$FOOTER" ]; then
  COMMIT_MSG="$COMMIT_MSG

$FOOTER"
fi

# Show preview
echo ""
echo -e "${GREEN}Commit message preview:${NC}"
echo "----------------------------------------"
echo "$COMMIT_MSG"
echo "----------------------------------------"
echo ""

# Confirm
read -p "Create this commit? [Y/n]: " confirm
if [[ $confirm =~ ^[Nn]$ ]]; then
  echo -e "${YELLOW}Commit cancelled${NC}"
  exit 0
fi

# Create commit
git commit -m "$COMMIT_MSG"

echo -e "${GREEN}Commit created successfully!${NC}"
