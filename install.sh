#!/bin/bash
# Uberman Installation Script
# Usage: curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
GITHUB_REPO="tburny/uberman"
BINARY_NAME="uberman"
INSTALL_DIR="${INSTALL_DIR:-$HOME/bin}"

# Helper functions
log_info() {
    echo -e "${BLUE}==>${NC} $1"
}

log_success() {
    echo -e "${GREEN}==>${NC} $1"
}

log_error() {
    echo -e "${RED}Error:${NC} $1" >&2
}

log_warning() {
    echo -e "${YELLOW}Warning:${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    local os arch

    # Detect OS
    case "$(uname -s)" in
        Linux*)     os="linux" ;;
        Darwin*)    os="darwin" ;;
        FreeBSD*)   os="freebsd" ;;
        CYGWIN*|MINGW*|MSYS*) os="windows" ;;
        *)
            log_error "Unsupported operating system: $(uname -s)"
            exit 1
            ;;
    esac

    # Detect architecture
    case "$(uname -m)" in
        x86_64|amd64)  arch="amd64" ;;
        aarch64|arm64) arch="arm64" ;;
        *)
            log_error "Unsupported architecture: $(uname -m)"
            exit 1
            ;;
    esac

    # Windows uses .exe extension
    if [ "$os" = "windows" ]; then
        BINARY_NAME="${BINARY_NAME}.exe"
    fi

    echo "${os}-${arch}"
}

# Get latest release version from GitHub
get_latest_version() {
    local version
    version=$(curl -fsSL "https://api.github.com/repos/${GITHUB_REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$version" ]; then
        log_error "Failed to fetch latest version"
        exit 1
    fi

    echo "$version"
}

# Download and extract binary
download_binary() {
    local platform=$1
    local version=$2
    local tmpdir

    tmpdir=$(mktemp -d)
    trap "rm -rf $tmpdir" EXIT

    local archive_name="${BINARY_NAME}-${platform}"
    local download_url="https://github.com/${GITHUB_REPO}/releases/download/${version}/${archive_name}.tar.gz"

    # Windows uses .zip
    if [[ "$platform" == *"windows"* ]]; then
        download_url="https://github.com/${GITHUB_REPO}/releases/download/${version}/${archive_name}.zip"
    fi

    log_info "Downloading ${BINARY_NAME} ${version} for ${platform}..."

    if ! curl -fsSL "$download_url" -o "${tmpdir}/archive"; then
        log_error "Failed to download from ${download_url}"
        exit 1
    fi

    log_info "Extracting archive..."

    cd "$tmpdir"
    if [[ "$platform" == *"windows"* ]]; then
        if command -v unzip >/dev/null 2>&1; then
            unzip -q archive
        else
            log_error "unzip not found. Please install unzip and try again."
            exit 1
        fi
    else
        tar -xzf archive
    fi

    echo "${tmpdir}/${archive_name}"
}

# Install binary
install_binary() {
    local binary_path=$1

    # Create install directory if it doesn't exist
    if [ ! -d "$INSTALL_DIR" ]; then
        log_info "Creating installation directory: ${INSTALL_DIR}"
        mkdir -p "$INSTALL_DIR"
    fi

    # Move binary to install directory
    log_info "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."

    if [ -f "${INSTALL_DIR}/${BINARY_NAME}" ]; then
        log_warning "Existing installation found. Backing up to ${INSTALL_DIR}/${BINARY_NAME}.backup"
        mv "${INSTALL_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}.backup"
    fi

    mv "$binary_path" "${INSTALL_DIR}/${BINARY_NAME}"
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

    log_success "Installation complete!"
}

# Verify installation
verify_installation() {
    if [ ! -x "${INSTALL_DIR}/${BINARY_NAME}" ]; then
        log_error "Binary is not executable"
        exit 1
    fi

    # Check if install directory is in PATH
    if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
        log_warning "Installation directory ${INSTALL_DIR} is not in your PATH"
        log_info "Add it to your PATH by adding this line to your shell profile:"
        echo ""
        echo "    export PATH=\"\$PATH:${INSTALL_DIR}\""
        echo ""
    fi

    # Try to run the binary
    log_info "Verifying installation..."
    if "${INSTALL_DIR}/${BINARY_NAME}" --version >/dev/null 2>&1; then
        local version=$("${INSTALL_DIR}/${BINARY_NAME}" --version 2>/dev/null || echo "unknown")
        log_success "Uberman installed successfully! Version: ${version}"
    else
        log_warning "Binary installed but unable to verify version"
        log_success "Uberman installed to ${INSTALL_DIR}/${BINARY_NAME}"
    fi
}

# Show usage instructions
show_usage() {
    cat << EOF

${GREEN}Next steps:${NC}

1. Verify installation:
   ${BLUE}\$ uberman --version${NC}

2. View available commands:
   ${BLUE}\$ uberman --help${NC}

3. Install your first app (e.g., WordPress):
   ${BLUE}\$ uberman install wordpress${NC}

4. Read the documentation:
   https://github.com/${GITHUB_REPO}

${YELLOW}Note:${NC} If you see "command not found", make sure ${INSTALL_DIR} is in your PATH.

EOF
}

# Main installation flow
main() {
    echo ""
    echo "${BLUE}╔════════════════════════════════════════╗${NC}"
    echo "${BLUE}║                                        ║${NC}"
    echo "${BLUE}║        Uberman Installer               ║${NC}"
    echo "${BLUE}║                                        ║${NC}"
    echo "${BLUE}╚════════════════════════════════════════╝${NC}"
    echo ""

    # Check prerequisites
    if ! command -v curl >/dev/null 2>&1; then
        log_error "curl is required but not installed. Please install curl and try again."
        exit 1
    fi

    if ! command -v tar >/dev/null 2>&1; then
        log_error "tar is required but not installed. Please install tar and try again."
        exit 1
    fi

    # Detect platform
    log_info "Detecting platform..."
    platform=$(detect_platform)
    log_info "Platform: ${platform}"

    # Get latest version
    log_info "Fetching latest version..."
    version=$(get_latest_version)
    log_info "Latest version: ${version}"

    # Download binary
    binary_path=$(download_binary "$platform" "$version")

    # Install binary
    install_binary "$binary_path"

    # Verify installation
    verify_installation

    # Show usage
    show_usage
}

# Handle command-line arguments
case "${1:-}" in
    -h|--help)
        cat << EOF
Uberman Installation Script

Usage:
  curl -fsSL https://raw.githubusercontent.com/${GITHUB_REPO}/main/install.sh | bash

  Or with custom install directory:
  curl -fsSL https://raw.githubusercontent.com/${GITHUB_REPO}/main/install.sh | INSTALL_DIR=/usr/local/bin bash

Options:
  -h, --help     Show this help message

Environment Variables:
  INSTALL_DIR    Installation directory (default: \$HOME/bin)

Examples:
  # Install to default location (\$HOME/bin)
  curl -fsSL https://raw.githubusercontent.com/${GITHUB_REPO}/main/install.sh | bash

  # Install to /usr/local/bin (requires sudo)
  curl -fsSL https://raw.githubusercontent.com/${GITHUB_REPO}/main/install.sh | sudo INSTALL_DIR=/usr/local/bin bash

EOF
        exit 0
        ;;
    *)
        main "$@"
        ;;
esac
