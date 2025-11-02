package web

import (
	"fmt"
	"os/exec"
	"strings"
)

// BackendManager handles Uberspace web backend configuration
type BackendManager struct {
	dryRun  bool
	verbose bool
}

// NewBackendManager creates a new web backend manager
func NewBackendManager(dryRun, verbose bool) *BackendManager {
	return &BackendManager{
		dryRun:  dryRun,
		verbose: verbose,
	}
}

// Backend represents a web backend configuration
type Backend struct {
	Path   string
	Type   string // "apache" or "http"
	Port   int    // only for http type
	Domain string // optional, for domain-specific backends
}

// SetBackend configures a web backend
func (m *BackendManager) SetBackend(backend Backend) error {
	var args []string

	if backend.Domain != "" {
		args = append(args, "web", "backend", "set", backend.Path, "--domain", backend.Domain)
	} else {
		args = append(args, "web", "backend", "set", backend.Path)
	}

	if backend.Type == "apache" {
		args = append(args, "--apache")
	} else if backend.Type == "http" {
		args = append(args, "--http", "--port", fmt.Sprintf("%d", backend.Port))
	} else {
		return fmt.Errorf("invalid backend type: %s (must be 'apache' or 'http')", backend.Type)
	}

	if m.verbose {
		fmt.Printf("Setting backend: %s\n", strings.Join(args, " "))
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: uberspace %s\n", strings.Join(args, " "))
		return nil
	}

	cmd := exec.Command("uberspace", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set backend: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// DeleteBackend removes a web backend configuration
func (m *BackendManager) DeleteBackend(path string, domain string) error {
	args := []string{"web", "backend", "del", path}

	if domain != "" {
		args = append(args, "--domain", domain)
	}

	if m.verbose {
		fmt.Printf("Deleting backend: %s\n", strings.Join(args, " "))
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: uberspace %s\n", strings.Join(args, " "))
		return nil
	}

	cmd := exec.Command("uberspace", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to delete backend: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// ListBackends lists all configured backends
func (m *BackendManager) ListBackends() ([]string, error) {
	cmd := exec.Command("uberspace", "web", "backend", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list backends: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var backends []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			backends = append(backends, line)
		}
	}

	return backends, nil
}

// ListDomains lists all configured domains
func (m *BackendManager) ListDomains() ([]string, error) {
	cmd := exec.Command("uberspace", "web", "domain", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list domains: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var domains []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			domains = append(domains, line)
		}
	}

	return domains, nil
}

// FindAvailablePort finds an available port in the Uberspace range
func (m *BackendManager) FindAvailablePort(start, end int) (int, error) {
	if start < 1024 {
		start = 1024
	}
	if end > 65535 {
		end = 65535
	}

	backends, err := m.ListBackends()
	if err != nil {
		return 0, err
	}

	usedPorts := make(map[int]bool)
	for _, backend := range backends {
		// Parse port from backend string (format varies)
		// This is a simplified implementation
		var port int
		if _, err := fmt.Sscanf(backend, "%*s => http:%d", &port); err == nil {
			usedPorts[port] = true
		}
	}

	// Find first available port
	for port := start; port <= end; port++ {
		if !usedPorts[port] {
			return port, nil
		}
	}

	return 0, fmt.Errorf("no available ports in range %d-%d", start, end)
}
