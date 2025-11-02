package runtime

import (
	"fmt"
	"os/exec"
	"strings"
)

// Manager handles Uberspace runtime version management
type Manager struct {
	dryRun  bool
	verbose bool
}

// NewManager creates a new runtime manager
func NewManager(dryRun, verbose bool) *Manager {
	return &Manager{
		dryRun:  dryRun,
		verbose: verbose,
	}
}

// SetVersion sets the runtime version using uberspace tools
func (m *Manager) SetVersion(language, version string) error {
	if m.verbose {
		fmt.Printf("Setting %s version to %s\n", language, version)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: uberspace tools version use %s %s\n", language, version)
		return nil
	}

	cmd := exec.Command("uberspace", "tools", "version", "use", language, version)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set %s version: %w\nOutput: %s", language, err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// GetVersion retrieves the current runtime version
func (m *Manager) GetVersion(language string) (string, error) {
	cmd := exec.Command("uberspace", "tools", "version", "show", language)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get %s version: %w", language, err)
	}

	version := strings.TrimSpace(string(output))
	return version, nil
}

// ListVersions lists available versions for a runtime
func (m *Manager) ListVersions(language string) ([]string, error) {
	cmd := exec.Command("uberspace", "tools", "version", "list", language)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list %s versions: %w", language, err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var versions []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && line != "Supported versions:" {
			// Remove leading "- " if present
			version := strings.TrimPrefix(line, "- ")
			versions = append(versions, version)
		}
	}

	return versions, nil
}

// RestartPHP restarts the PHP-FPM service
func (m *Manager) RestartPHP() error {
	if m.verbose {
		fmt.Println("Restarting PHP-FPM")
	}

	if m.dryRun {
		fmt.Println("[DRY RUN] Would execute: uberspace tools restart php")
		return nil
	}

	cmd := exec.Command("uberspace", "tools", "restart", "php")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to restart PHP: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}
