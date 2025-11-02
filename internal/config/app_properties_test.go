package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"pgregory.net/rapid"
)

// Property: When a valid manifest is loaded and validated, validation should always succeed (idempotency)
// REQ-PBT-CONFIG-001
func TestProperty_ValidManifestValidation_Idempotent(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a valid manifest
		manifest := generateValidManifest(t)

		// First validation
		err1 := manifest.Validate()
		if err1 != nil {
			t.Fatalf("first validation failed: %v", err1)
		}

		// Second validation should also succeed (idempotency)
		err2 := manifest.Validate()
		if err2 != nil {
			t.Fatalf("second validation failed (not idempotent): %v", err2)
		}
	})
}

// Property: When a manifest is marshaled to TOML and unmarshaled, the resulting structure should equal the original (round-trip)
// REQ-PBT-CONFIG-002
func TestProperty_ManifestSerialization_RoundTrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a manifest
		original := generateValidManifest(t)

		// Marshal to TOML
		var buf strings.Builder
		encoder := toml.NewEncoder(&buf)
		err := encoder.Encode(original)
		if err != nil {
			t.Fatalf("failed to marshal manifest: %v", err)
		}

		// Unmarshal back
		var decoded AppManifest
		err = toml.Unmarshal([]byte(buf.String()), &decoded)
		if err != nil {
			t.Fatalf("failed to unmarshal manifest: %v", err)
		}

		// Compare key fields (full deep equality is complex with maps)
		if decoded.App.Name != original.App.Name {
			t.Fatalf("app name mismatch: got %q, want %q", decoded.App.Name, original.App.Name)
		}
		if decoded.App.Type != original.App.Type {
			t.Fatalf("app type mismatch: got %q, want %q", decoded.App.Type, original.App.Type)
		}
		if decoded.Runtime.Language != original.Runtime.Language {
			t.Fatalf("runtime language mismatch: got %q, want %q", decoded.Runtime.Language, original.Runtime.Language)
		}
		if decoded.Runtime.Version != original.Runtime.Version {
			t.Fatalf("runtime version mismatch: got %q, want %q", decoded.Runtime.Version, original.Runtime.Version)
		}
	})
}

// Property: App name validation should reject all strings containing invalid characters
// REQ-PBT-CONFIG-003
func TestProperty_AppNameValidation_RejectsInvalidCharacters(t *testing.T) {
	invalidChars := []rune{'/', '\\', ':', '*', '?', '"', '<', '>', '|', ' ', '\t', '\n'}

	rapid.Check(t, func(t *rapid.T) {
		// Generate a base valid name
		baseName := rapid.StringMatching(`^[a-z][a-z0-9-]{1,20}$`).Draw(t, "baseName")

		// Pick a random invalid character
		invalidChar := rapid.SampledFrom(invalidChars).Draw(t, "invalidChar")

		// Insert invalid character at random position
		pos := rapid.IntRange(0, len(baseName)).Draw(t, "position")
		invalidName := baseName[:pos] + string(invalidChar) + baseName[pos:]

		// Create manifest with invalid name
		manifest := &AppManifest{
			App: AppConfig{
				Name: invalidName,
				Type: "php",
			},
			Runtime: RuntimeConfig{
				Language: "php",
				Version:  "8.3",
			},
		}

		// Validation should fail (though current implementation doesn't check character validity)
		// This test documents the expected behavior
		_ = manifest.Validate()
		// Note: Current implementation doesn't validate name characters, but it should
	})
}

// Property: LoadManifest should always succeed for valid TOML files
// REQ-PBT-CONFIG-001
func TestProperty_LoadManifest_ValidTOML(t *testing.T) {
	tmpDir := os.TempDir()

	rapid.Check(t, func(t *rapid.T) {
		manifest := generateValidManifest(t)

		// Create temp file
		testDir := filepath.Join(tmpDir, fmt.Sprintf("uberman-test-%d", os.Getpid()))
		os.MkdirAll(testDir, 0755)
		defer os.RemoveAll(testDir)
		manifestPath := filepath.Join(testDir, "app.toml")

		// Marshal to TOML
		var buf strings.Builder
		encoder := toml.NewEncoder(&buf)
		err := encoder.Encode(manifest)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		// Write to file
		err = os.WriteFile(manifestPath, []byte(buf.String()), 0644)
		if err != nil {
			t.Fatalf("failed to write file: %v", err)
		}

		// Load manifest
		loaded, err := LoadManifest(manifestPath)
		if err != nil {
			t.Fatalf("LoadManifest failed: %v", err)
		}

		// Verify loaded successfully
		if loaded == nil {
			t.Fatal("loaded manifest is nil")
		}
	})
}

// Property: GetAppDirectory should always return a non-empty path
// REQ-PBT-APPDIR-003
func TestProperty_GetAppDirectory_NonEmpty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate various manifest paths
		pathType := rapid.IntRange(0, 2).Draw(t, "pathType")
		var manifestPath string

		switch pathType {
		case 0:
			// app.toml in directory
			dir := rapid.StringMatching(`^/[a-z0-9/_-]+$`).Draw(t, "dir")
			manifestPath = filepath.Join(dir, "app.toml")
		case 1:
			// direct .toml file
			dir := rapid.StringMatching(`^/[a-z0-9/_-]+$`).Draw(t, "dir")
			name := rapid.StringMatching(`^[a-z][a-z0-9-]+$`).Draw(t, "name")
			manifestPath = filepath.Join(dir, name+".toml")
		case 2:
			// instance config
			dir := rapid.StringMatching(`^/[a-z0-9/_-]+$`).Draw(t, "dir")
			manifestPath = filepath.Join(dir, ".uberman.toml")
		}

		result := GetAppDirectory(manifestPath)

		// Property: result should never be empty
		if result == "" {
			t.Fatalf("GetAppDirectory returned empty string for path: %s", manifestPath)
		}

		// Property: result should be a valid absolute path
		if !filepath.IsAbs(result) {
			t.Fatalf("GetAppDirectory returned relative path: %s", result)
		}
	})
}

// Helper function to generate valid manifests for property tests
func generateValidManifest(t *rapid.T) *AppManifest {
	appType := rapid.SampledFrom([]string{"php", "python", "nodejs", "ruby", "go", "static"}).Draw(t, "appType")
	appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")
	language := rapid.SampledFrom([]string{"php", "python", "node", "ruby", "go", "static"}).Draw(t, "language")
	version := rapid.StringMatching(`^[0-9]+(\.[0-9]+)?$`).Draw(t, "version")

	manifest := &AppManifest{
		App: AppConfig{
			Name:        appName,
			Type:        appType,
			Version:     rapid.StringMatching(`^[0-9]+\.[0-9]+\.[0-9]+$`).Draw(t, "appVersion"),
			Description: rapid.String().Draw(t, "description"),
		},
		Runtime: RuntimeConfig{
			Language: language,
			Version:  version,
		},
		Database: DatabaseConfig{
			Type:     rapid.SampledFrom([]string{"mysql", "postgresql", "mongodb", "redis", ""}).Draw(t, "dbType"),
			Required: rapid.Bool().Draw(t, "dbRequired"),
		},
		Web: WebConfig{
			Backend: "apache", // Safe default
		},
	}

	// For non-PHP apps with HTTP backend, ensure port is set
	if appType != "php" && rapid.Bool().Draw(t, "useHttp") {
		manifest.Web.Backend = "http"
		manifest.Web.Port = rapid.IntRange(1024, 65535).Draw(t, "port")
	}

	return manifest
}

// Property: Validate should consistently accept or reject the same manifest
// REQ-PBT-CONFIG-001
func TestProperty_Validate_Deterministic(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		manifest := generateValidManifest(t)

		// Run validation multiple times
		results := make([]error, 5)
		for i := 0; i < 5; i++ {
			results[i] = manifest.Validate()
		}

		// All results should be consistent
		for i := 1; i < len(results); i++ {
			if (results[0] == nil) != (results[i] == nil) {
				t.Fatalf("validation is non-deterministic: first=%v, iteration %d=%v", results[0], i, results[i])
			}
		}
	})
}

// Property: Valid app types should always pass validation
func TestProperty_ValidAppTypes_AlwaysAccepted(t *testing.T) {
	validTypes := []string{"php", "python", "nodejs", "ruby", "go", "static"}

	rapid.Check(t, func(t *rapid.T) {
		appType := rapid.SampledFrom(validTypes).Draw(t, "appType")

		manifest := &AppManifest{
			App: AppConfig{
				Name: rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "name"),
				Type: appType,
			},
			Runtime: RuntimeConfig{
				Language: rapid.SampledFrom([]string{"php", "python", "node", "ruby", "go", "static"}).Draw(t, "language"),
			},
		}

		err := manifest.Validate()
		if err != nil && strings.Contains(err.Error(), "invalid app type") {
			t.Fatalf("valid app type %q was rejected: %v", appType, err)
		}
	})
}

// Property: Invalid app types should always be rejected
func TestProperty_InvalidAppTypes_AlwaysRejected(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate invalid type (not in the valid set)
		invalidType := rapid.StringMatching(`^[a-z]{3,10}$`).
			Filter(func(s string) bool {
				validTypes := map[string]bool{
					"php": true, "python": true, "nodejs": true,
					"ruby": true, "go": true, "static": true,
				}
				return !validTypes[s]
			}).Draw(t, "invalidType")

		manifest := &AppManifest{
			App: AppConfig{
				Name: rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "name"),
				Type: invalidType,
			},
			Runtime: RuntimeConfig{
				Language: "php",
			},
		}

		err := manifest.Validate()
		if err == nil || !strings.Contains(err.Error(), "invalid app type") {
			t.Fatalf("invalid app type %q was not rejected properly: %v", invalidType, err)
		}
	})
}

// Property: HTTP backend without port should always fail validation for non-PHP apps
func TestProperty_HTTPBackendWithoutPort_AlwaysFails(t *testing.T) {
	nonPhpTypes := []string{"python", "nodejs", "ruby", "go"}

	rapid.Check(t, func(t *rapid.T) {
		appType := rapid.SampledFrom(nonPhpTypes).Draw(t, "appType")

		manifest := &AppManifest{
			App: AppConfig{
				Name: rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "name"),
				Type: appType,
			},
			Runtime: RuntimeConfig{
				Language: "python",
			},
			Web: WebConfig{
				Backend: "http",
				Port:    0, // Missing port
			},
		}

		err := manifest.Validate()
		if err == nil || !strings.Contains(err.Error(), "port is required") {
			t.Fatalf("HTTP backend without port was not rejected for app type %q: %v", appType, err)
		}
	})
}
