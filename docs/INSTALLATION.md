# Installation Guide

Complete installation instructions for Uberman across all platforms.

## Quick Start

### One-Line Install

The fastest way to install Uberman:

```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
```

This automatically detects your platform and installs the latest version.

## Installation Methods

### 1. Automated Installation Script (Recommended)

**Features:**
- Automatic platform detection
- Downloads latest release
- Verifies installation
- Supports custom install locations

**Default Installation:**
```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
```

Installs to: `~/bin/uberman`

**Custom Installation Directory:**

```bash
# Install to /usr/local/bin (system-wide, requires sudo)
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | sudo INSTALL_DIR=/usr/local/bin bash

# Install to ~/.local/bin
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | INSTALL_DIR=~/.local/bin bash

# Install to any custom directory
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | INSTALL_DIR=/custom/path bash
```

**What the script does:**
1. Detects your OS (Linux, macOS, FreeBSD, Windows)
2. Detects your architecture (amd64, arm64)
3. Downloads the appropriate binary from GitHub releases
4. Extracts and installs to specified directory
5. Makes the binary executable
6. Verifies the installation
7. Shows next steps

**Prerequisites:**
- `curl` (for downloading)
- `tar` (for extracting, pre-installed on most systems)
- `unzip` (only for Windows)

### 2. Manual Binary Installation

Download pre-built binaries from [GitHub Releases](https://github.com/tburny/uberman/releases).

#### Linux (Uberspace)

```bash
# Download
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-linux-amd64.tar.gz -o uberman.tar.gz

# Extract
tar -xzf uberman.tar.gz

# Install
mkdir -p ~/bin
mv uberman-linux-amd64 ~/bin/uberman
chmod +x ~/bin/uberman

# Verify
uberman --version
```

#### Linux (ARM64 - Raspberry Pi, etc.)

```bash
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-linux-arm64.tar.gz -o uberman.tar.gz
tar -xzf uberman.tar.gz
mv uberman-linux-arm64 ~/bin/uberman
chmod +x ~/bin/uberman
```

#### macOS (Intel)

```bash
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-darwin-amd64.tar.gz -o uberman.tar.gz
tar -xzf uberman.tar.gz
mv uberman-darwin-amd64 /usr/local/bin/uberman
chmod +x /usr/local/bin/uberman
```

#### macOS (Apple Silicon - M1/M2/M3)

```bash
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-darwin-arm64.tar.gz -o uberman.tar.gz
tar -xzf uberman.tar.gz
mv uberman-darwin-arm64 /usr/local/bin/uberman
chmod +x /usr/local/bin/uberman
```

#### Windows

```bash
# PowerShell
Invoke-WebRequest -Uri "https://github.com/tburny/uberman/releases/latest/download/uberman-windows-amd64.zip" -OutFile "uberman.zip"
Expand-Archive -Path uberman.zip -DestinationPath .
Move-Item uberman-windows-amd64.exe C:\Windows\System32\uberman.exe
```

Or download and extract manually from the [releases page](https://github.com/tburny/uberman/releases).

#### FreeBSD

```bash
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-freebsd-amd64.tar.gz -o uberman.tar.gz
tar -xzf uberman.tar.gz
mv uberman-freebsd-amd64 /usr/local/bin/uberman
chmod +x /usr/local/bin/uberman
```

### 3. Build from Source

**Prerequisites:**
- Go 1.21 or higher
- Git

**On Uberspace:**

```bash
# Install Go
uberspace tools version use go 1.21

# Clone repository
cd ~/projekte
git clone https://github.com/tburny/uberman.git
cd uberman

# Build and install
make install-local

# Verify
uberman --version
```

**On other systems:**

```bash
# Install Go from https://golang.org/dl/

# Clone repository
git clone https://github.com/tburny/uberman.git
cd uberman

# Download dependencies
go mod download

# Build
go build -o uberman ./cmd/uberman

# Install
sudo mv uberman /usr/local/bin/
# Or for user-only install:
mv uberman ~/bin/

# Verify
uberman --version
```

**Build all platforms:**

```bash
make release
```

This creates binaries for all supported platforms in the `dist/` directory.

### 4. Package Managers

**Coming soon:**
- Homebrew (macOS/Linux)
- APT repository (Debian/Ubuntu)
- Snapcraft (Ubuntu)
- AUR (Arch Linux)

## Verifying Installation

### Check Version

```bash
uberman --version
```

Expected output:
```
uberman version v1.0.0
```

### Check Help

```bash
uberman --help
```

Should display available commands and options.

### Test Installation

```bash
# Check if binary is executable
which uberman

# Expected: /home/username/bin/uberman (or your install path)
```

## Verifying Downloads

### Checksums

Each release includes SHA256 checksums for verification:

```bash
# Download binary and checksum
curl -LO https://github.com/tburny/uberman/releases/latest/download/uberman-linux-amd64.tar.gz
curl -LO https://github.com/tburny/uberman/releases/latest/download/uberman-linux-amd64.tar.gz.sha256

# Verify checksum
sha256sum -c uberman-linux-amd64.tar.gz.sha256
```

Expected output:
```
uberman-linux-amd64.tar.gz: OK
```

### GPG Signatures (Coming Soon)

Future releases will include GPG signatures for enhanced security verification.

## Post-Installation Setup

### 1. Verify PATH

Ensure the installation directory is in your PATH:

```bash
echo $PATH | grep -o ~/bin
```

If not found, add to your shell profile:

**For bash (`~/.bashrc` or `~/.bash_profile`):**
```bash
export PATH="$PATH:$HOME/bin"
```

**For zsh (`~/.zshrc`):**
```bash
export PATH="$PATH:$HOME/bin"
```

**For fish (`~/.config/fish/config.fish`):**
```fish
set -gx PATH $PATH $HOME/bin
```

Then reload your shell:
```bash
source ~/.bashrc  # or ~/.zshrc, etc.
```

### 2. Shell Completion (Coming Soon)

Future versions will support shell completion for bash, zsh, and fish.

### 3. Configuration

Create optional configuration file:

```bash
# Copy example configuration
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/.uberman.toml.example -o ~/.uberman.toml

# Edit configuration
nano ~/.uberman.toml
```

See [Configuration Guide](../CLAUDE.md) for details.

## Updating Uberman

### Using Installation Script

Simply re-run the installation script:

```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
```

The script automatically backs up the existing installation.

### Manual Update

1. Download latest release
2. Replace existing binary
3. Verify new version

```bash
# Example for Linux
curl -L https://github.com/tburny/uberman/releases/latest/download/uberman-linux-amd64.tar.gz | tar xz
mv ~/bin/uberman ~/bin/uberman.backup
mv uberman-linux-amd64 ~/bin/uberman
chmod +x ~/bin/uberman
uberman --version
```

### Check for Updates

```bash
# View current version
uberman --version

# Check latest release on GitHub
curl -s https://api.github.com/repos/tburny/uberman/releases/latest | grep '"tag_name"'
```

## Uninstallation

### Remove Binary

```bash
# If installed to ~/bin
rm ~/bin/uberman

# If installed to /usr/local/bin
sudo rm /usr/local/bin/uberman
```

### Remove Configuration

```bash
rm ~/.uberman.toml
```

### Remove Installed Apps

```bash
# Remove all installed apps (optional)
rm -rf ~/apps/*

# Or remove specific apps
rm -rf ~/apps/wordpress
```

## Troubleshooting

### "Command not found"

**Problem:** Shell can't find `uberman` command.

**Solutions:**

1. Check installation location:
   ```bash
   ls -l ~/bin/uberman
   ```

2. Verify PATH includes install directory:
   ```bash
   echo $PATH
   ```

3. Add to PATH if needed (see Post-Installation Setup above)

4. Try with full path:
   ```bash
   ~/bin/uberman --version
   ```

### "Permission denied"

**Problem:** Binary is not executable.

**Solution:**
```bash
chmod +x ~/bin/uberman
```

### "Cannot execute binary file"

**Problem:** Downloaded wrong platform binary.

**Solutions:**

1. Check your platform:
   ```bash
   uname -s   # OS
   uname -m   # Architecture
   ```

2. Download correct binary for your platform

3. Or use installation script (auto-detects):
   ```bash
   curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh | bash
   ```

### Installation Script Fails

**Check prerequisites:**
```bash
# Check curl
curl --version

# Check tar
tar --version
```

**Download script and run manually:**
```bash
curl -fsSL https://raw.githubusercontent.com/tburny/uberman/main/install.sh -o install-uberman.sh
chmod +x install-uberman.sh
./install-uberman.sh
```

### macOS "cannot be opened" error

**Problem:** macOS Gatekeeper blocks unsigned binary.

**Solution:**
```bash
# Remove quarantine attribute
xattr -d com.apple.quarantine ~/bin/uberman

# Or allow in System Preferences → Security & Privacy
```

## Platform-Specific Notes

### Uberspace

- Default install location: `~/bin/` (already in PATH)
- No sudo required
- Use Linux amd64 binaries
- Compatible with Uberspace 7

### macOS

- Binaries are not signed (Gatekeeper may complain)
- Use Homebrew when available (coming soon)
- Two architectures: Intel (amd64) and Apple Silicon (arm64)

### Windows

- Use PowerShell or WSL
- Native Windows binary available (.exe)
- Add to PATH manually if needed

### FreeBSD

- Limited testing (community-maintained)
- Report issues on GitHub

## Getting Help

### Documentation

- **Quick Start**: [docs/QUICKSTART.md](QUICKSTART.md)
- **Full Guide**: [README.md](../README.md)
- **Architecture**: [CLAUDE.md](../CLAUDE.md)

### Support Channels

- **GitHub Issues**: https://github.com/tburny/uberman/issues
- **Discussions**: https://github.com/tburny/uberman/discussions
- **Uberspace Forum**: https://forum.uberspace.de/

### Reporting Installation Issues

When reporting installation problems, include:
1. Your OS and architecture (`uname -a`)
2. Installation method used
3. Complete error messages
4. Steps to reproduce

---

*For more information, see the [README](../README.md) and [QUICKSTART](QUICKSTART.md) guides.*
