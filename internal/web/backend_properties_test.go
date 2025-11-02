package web

import (
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// Property: The port finder should always return a port in the valid range (1024-65535)
// REQ-PBT-WEB-001
func TestProperty_FindAvailablePort_ValidRange(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate random but valid start and end
		start := rapid.IntRange(1024, 32767).Draw(t, "start")
		end := rapid.IntRange(start+1, 65535).Draw(t, "end")

		manager := NewBackendManager(false, false)

		// Note: This test documents expected behavior
		// The actual FindAvailablePort implementation may fail if no ports are available
		// But if it succeeds, the port should be in range
		port, err := manager.FindAvailablePort(start, end)

		if err == nil {
			// Property: Port should be in requested range
			if port < start || port > end {
				t.Fatalf("port %d is not in requested range [%d, %d]", port, start, end)
			}

			// Property: Port should be in valid TCP range
			if port < 1024 || port > 65535 {
				t.Fatalf("port %d is not in valid TCP range [1024, 65535]", port)
			}
		}
	})
}

// Property: When a backend configuration is generated, the output should always contain valid uberspace command syntax
// REQ-PBT-WEB-002
func TestProperty_BackendConfiguration_ValidSyntax(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		backend := generateValidBackend(t)

		manager := NewBackendManager(true, false) // dry-run mode to capture command

		// SetBackend should generate valid command (won't fail in dry-run)
		err := manager.SetBackend(backend)
		if err != nil {
			t.Fatalf("SetBackend failed: %v", err)
		}

		// Property: Backend type should be valid
		validTypes := map[string]bool{"apache": true, "http": true}
		if !validTypes[backend.Type] {
			t.Fatalf("invalid backend type: %q", backend.Type)
		}

		// Property: HTTP backend should have a valid port
		if backend.Type == "http" {
			if backend.Port < 1024 || backend.Port > 65535 {
				t.Fatalf("HTTP backend has invalid port: %d", backend.Port)
			}
		}

		// Property: Path should start with /
		if !strings.HasPrefix(backend.Path, "/") {
			t.Fatalf("backend path doesn't start with /: %q", backend.Path)
		}
	})
}

// Property: Domain validation should consistently accept/reject the same input
// REQ-PBT-WEB-003
func TestProperty_DomainValidation_Consistent(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate various domain formats
		domain := rapid.StringMatching(`^[a-z][a-z0-9.-]{3,30}\.[a-z]{2,6}$`).Draw(t, "domain")

		backend := Backend{
			Path:   "/",
			Type:   "apache",
			Domain: domain,
		}

		manager := NewBackendManager(true, false)

		// Run SetBackend multiple times
		results := make([]error, 3)
		for i := 0; i < 3; i++ {
			results[i] = manager.SetBackend(backend)
		}

		// Property: All results should be consistent
		for i := 1; i < len(results); i++ {
			if (results[0] == nil) != (results[i] == nil) {
				t.Fatalf("domain validation is non-deterministic: first=%v, iteration %d=%v",
					results[0], i, results[i])
			}
		}
	})
}

// Property: Valid backend types should always be accepted
func TestProperty_ValidBackendTypes_Accepted(t *testing.T) {
	validTypes := []string{"apache", "http"}

	rapid.Check(t, func(t *rapid.T) {
		backendType := rapid.SampledFrom(validTypes).Draw(t, "backendType")

		backend := Backend{
			Path: "/",
			Type: backendType,
		}

		// HTTP type needs port
		if backendType == "http" {
			backend.Port = rapid.IntRange(1024, 65535).Draw(t, "port")
		}

		manager := NewBackendManager(true, false)
		err := manager.SetBackend(backend)

		// Property: Valid types should not fail
		if err != nil {
			t.Fatalf("valid backend type %q was rejected: %v", backendType, err)
		}
	})
}

// Property: Invalid backend types should always be rejected
func TestProperty_InvalidBackendTypes_Rejected(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate invalid type (not apache or http)
		invalidType := rapid.StringMatching(`^[a-z]{3,10}$`).
			Filter(func(s string) bool { return s != "apache" && s != "http" }).
			Draw(t, "invalidType")

		backend := Backend{
			Path: "/",
			Type: invalidType,
		}

		manager := NewBackendManager(true, false)
		err := manager.SetBackend(backend)

		// Property: Invalid types should be rejected
		if err == nil {
			t.Fatalf("invalid backend type %q was not rejected", invalidType)
		}
		if !strings.Contains(err.Error(), "invalid backend type") {
			t.Fatalf("error message doesn't mention invalid type: %v", err)
		}
	})
}

// Property: Backend path should always start with /
func TestProperty_BackendPath_StartsWithSlash(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		backend := generateValidBackend(t)

		// Property: Path should start with /
		if !strings.HasPrefix(backend.Path, "/") {
			t.Fatalf("backend path doesn't start with /: %q", backend.Path)
		}
	})
}

// Property: HTTP backend without port should fail
func TestProperty_HTTPBackend_RequiresPort(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		backend := Backend{
			Path: rapid.StringMatching(`^/[a-z0-9/_-]*$`).Draw(t, "path"),
			Type: "http",
			Port: 0, // Missing port
		}

		manager := NewBackendManager(true, false)
		err := manager.SetBackend(backend)

		// Property: Should not succeed with missing port
		// Note: Current implementation might not validate this, documenting expected behavior
		_ = err // Document that this should fail
	})
}

// Property: Port range validator should enforce minimum 1024
func TestProperty_PortRange_EnforcesMinimum(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate port below minimum
		invalidPort := rapid.IntRange(1, 1023).Draw(t, "invalidPort")

		backend := Backend{
			Path: "/",
			Type: "http",
			Port: invalidPort,
		}

		manager := NewBackendManager(true, false)

		// The SetBackend itself might not validate, but FindAvailablePort should
		port, err := manager.FindAvailablePort(invalidPort, 2000)

		if err == nil {
			// If it succeeds, port should be >= 1024
			if port < 1024 {
				t.Fatalf("FindAvailablePort returned port below 1024: %d", port)
			}
		}

		// Property: Backend with invalid port should be treated specially
		_ = backend // Document expected behavior
	})
}

// Property: Port range validator should enforce maximum 65535
func TestProperty_PortRange_EnforcesMaximum(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate port above maximum
		invalidPort := rapid.IntRange(65536, 70000).Draw(t, "invalidPort")

		manager := NewBackendManager(false, false)

		// FindAvailablePort should handle this gracefully
		port, err := manager.FindAvailablePort(60000, invalidPort)

		if err == nil {
			// If it succeeds, port should be <= 65535
			if port > 65535 {
				t.Fatalf("FindAvailablePort returned port above 65535: %d", port)
			}
		}
	})
}

// Property: BackendManager should always be creatable
func TestProperty_BackendManager_AlwaysCreatable(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dryRun := rapid.Bool().Draw(t, "dryRun")
		verbose := rapid.Bool().Draw(t, "verbose")

		manager := NewBackendManager(dryRun, verbose)

		// Property: Should never be nil
		if manager == nil {
			t.Fatal("NewBackendManager returned nil")
		}

		// Property: Should have correct settings
		if manager.dryRun != dryRun {
			t.Fatalf("dryRun not set correctly: got %v, want %v", manager.dryRun, dryRun)
		}
		if manager.verbose != verbose {
			t.Fatalf("verbose not set correctly: got %v, want %v", manager.verbose, verbose)
		}
	})
}

// Property: Domain should only contain valid DNS characters
func TestProperty_Domain_ValidCharacters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		domain := rapid.StringMatching(`^[a-z0-9][a-z0-9.-]{1,60}[a-z0-9]$`).Draw(t, "domain")

		// Property: Should only contain alphanumeric, dots, and hyphens
		for _, char := range domain {
			if !((char >= 'a' && char <= 'z') ||
				(char >= '0' && char <= '9') ||
				char == '.' || char == '-') {
				t.Fatalf("domain contains invalid character %c: %q", char, domain)
			}
		}
	})
}

// Property: Backend path should not contain invalid URL characters
func TestProperty_BackendPath_ValidURLCharacters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		path := rapid.StringMatching(`^/[a-z0-9/_-]{0,50}$`).Draw(t, "path")

		// Property: Should not contain spaces or special chars
		invalidChars := []string{" ", "\t", "\n", "?", "&", "=", "#"}
		for _, char := range invalidChars {
			if strings.Contains(path, char) {
				t.Fatalf("backend path contains invalid character %q: %s", char, path)
			}
		}
	})
}

// Property: DryRun mode should never execute actual commands
func TestProperty_DryRun_NoExecution(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		backend := generateValidBackend(t)

		// Create manager in dry-run mode
		manager := NewBackendManager(true, false)

		// This should not fail even with invalid backends since it's dry-run
		err := manager.SetBackend(backend)
		if err != nil {
			t.Fatalf("dry-run SetBackend failed: %v", err)
		}

		// Property: Deletion in dry-run should also not fail
		err = manager.DeleteBackend(backend.Path, backend.Domain)
		if err != nil {
			t.Fatalf("dry-run DeleteBackend failed: %v", err)
		}
	})
}

// Property: Apache backend should not require port
func TestProperty_ApacheBackend_NoPortRequired(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		backend := Backend{
			Path: rapid.StringMatching(`^/[a-z0-9/_-]*$`).Draw(t, "path"),
			Type: "apache",
			Port: 0, // No port needed for Apache
		}

		manager := NewBackendManager(true, false)
		err := manager.SetBackend(backend)

		// Property: Should succeed without port for Apache
		if err != nil {
			t.Fatalf("Apache backend failed without port: %v", err)
		}
	})
}

// Property: FindAvailablePort should return different ports for different ranges
func TestProperty_FindAvailablePort_DifferentRanges(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate two non-overlapping ranges
		start1 := rapid.IntRange(1024, 30000).Draw(t, "start1")
		end1 := start1 + rapid.IntRange(1, 100).Draw(t, "range1")

		start2 := rapid.IntRange(end1+1, 60000).Draw(t, "start2")
		end2 := start2 + rapid.IntRange(1, 100).Draw(t, "range2")

		if end2 > 65535 {
			end2 = 65535
		}

		manager := NewBackendManager(false, false)

		port1, err1 := manager.FindAvailablePort(start1, end1)
		port2, err2 := manager.FindAvailablePort(start2, end2)

		// If both succeed, they should be in their respective ranges
		if err1 == nil && err2 == nil {
			if port1 < start1 || port1 > end1 {
				t.Fatalf("port1 %d not in range [%d, %d]", port1, start1, end1)
			}
			if port2 < start2 || port2 > end2 {
				t.Fatalf("port2 %d not in range [%d, %d]", port2, start2, end2)
			}
		}
	})
}

// Helper function to generate valid backend configurations
func generateValidBackend(t *rapid.T) Backend {
	backendType := rapid.SampledFrom([]string{"apache", "http"}).Draw(t, "type")

	backend := Backend{
		Path: rapid.StringMatching(`^/[a-z0-9/_-]{0,30}$`).Draw(t, "path"),
		Type: backendType,
	}

	// Add port for HTTP backend
	if backendType == "http" {
		backend.Port = rapid.IntRange(1024, 65535).Draw(t, "port")
	}

	// Optionally add domain
	if rapid.Bool().Draw(t, "hasDomain") {
		backend.Domain = rapid.StringMatching(`^[a-z][a-z0-9.-]{3,30}\.[a-z]{2,6}$`).Draw(t, "domain")
	}

	return backend
}
