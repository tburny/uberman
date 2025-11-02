package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	tests := []struct {
		name    string
		dryRun  bool
		verbose bool
	}{
		{
			name:    "normal mode",
			dryRun:  false,
			verbose: false,
		},
		{
			name:    "dry-run verbose",
			dryRun:  true,
			verbose: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewManager(tt.dryRun, tt.verbose)

			require.NotNil(t, manager)
			assert.Equal(t, tt.dryRun, manager.dryRun)
			assert.Equal(t, tt.verbose, manager.verbose)
		})
	}
}

func TestManager_SetVersion_DryRun(t *testing.T) {
	tests := []struct {
		name     string
		language string
		version  string
	}{
		{
			name:     "set PHP version",
			language: "php",
			version:  "8.3",
		},
		{
			name:     "set Node version",
			language: "node",
			version:  "20",
		},
		{
			name:     "set Python version",
			language: "python",
			version:  "3.11",
		},
		{
			name:     "set Ruby version",
			language: "ruby",
			version:  "3.2",
		},
		{
			name:     "set Go version",
			language: "go",
			version:  "1.21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test in dry-run mode to avoid requiring uberspace command
			manager := NewManager(true, true)

			err := manager.SetVersion(tt.language, tt.version)
			assert.NoError(t, err, "Dry-run should not error")
		})
	}
}

func TestManager_RestartPHP_DryRun(t *testing.T) {
	t.Run("restart PHP dry-run", func(t *testing.T) {
		manager := NewManager(true, true)

		err := manager.RestartPHP()
		assert.NoError(t, err, "Dry-run should not error")
	})

	t.Run("restart PHP verbose", func(t *testing.T) {
		manager := NewManager(true, true)

		err := manager.RestartPHP()
		assert.NoError(t, err)
	})
}

func TestManager_VerboseOutput(t *testing.T) {
	t.Run("verbose mode enabled", func(t *testing.T) {
		manager := NewManager(true, true)

		// These should output verbose messages but not error
		err := manager.SetVersion("php", "8.3")
		assert.NoError(t, err)

		err = manager.RestartPHP()
		assert.NoError(t, err)
	})

	t.Run("verbose mode disabled", func(t *testing.T) {
		manager := NewManager(true, false)

		// Should work without verbose output
		err := manager.SetVersion("node", "20")
		assert.NoError(t, err)
	})
}

// Note: The following tests would require actual uberspace commands to be available
// In a real Uberspace environment, you would test these with actual commands
// For unit testing, we rely on dry-run mode

func TestManager_GetVersion_Mock(t *testing.T) {
	// This test documents the expected behavior
	// In a real test environment with uberspace installed, uncomment and adapt

	t.Skip("Requires uberspace command to be installed")

	/*
		manager := NewManager(false, false)
		version, err := manager.GetVersion("php")
		require.NoError(t, err)
		assert.NotEmpty(t, version)
		assert.Regexp(t, `\d+\.\d+`, version) // Should match version pattern like "8.3"
	*/
}

func TestManager_ListVersions_Mock(t *testing.T) {
	// This test documents the expected behavior
	t.Skip("Requires uberspace command to be installed")

	/*
		manager := NewManager(false, false)
		versions, err := manager.ListVersions("php")
		require.NoError(t, err)
		assert.NotEmpty(t, versions)
		// Should contain multiple version numbers
		assert.Greater(t, len(versions), 0)
	*/
}

func TestManager_SupportedLanguages(t *testing.T) {
	// Test that all documented languages work with SetVersion in dry-run
	languages := []string{"php", "python", "node", "ruby", "go"}

	manager := NewManager(true, false)

	for _, lang := range languages {
		t.Run("language_"+lang, func(t *testing.T) {
			err := manager.SetVersion(lang, "latest")
			assert.NoError(t, err)
		})
	}
}

func TestManager_EdgeCases(t *testing.T) {
	manager := NewManager(true, false)

	t.Run("empty version string", func(t *testing.T) {
		err := manager.SetVersion("php", "")
		// Should not error in dry-run, but would fail with actual command
		assert.NoError(t, err)
	})

	t.Run("empty language string", func(t *testing.T) {
		err := manager.SetVersion("", "8.3")
		// Should not error in dry-run, but would fail with actual command
		assert.NoError(t, err)
	})
}

// Integration test that would run on actual Uberspace
func TestManager_Integration_Uberspace(t *testing.T) {
	// Check if running on Uberspace (has uberspace command)
	t.Skip("Integration test - only run on actual Uberspace host")

	/*
		// This test should only run on actual Uberspace
		manager := NewManager(false, true)

		// Test getting current PHP version
		version, err := manager.GetVersion("php")
		require.NoError(t, err)
		assert.NotEmpty(t, version)

		// Test listing available PHP versions
		versions, err := manager.ListVersions("php")
		require.NoError(t, err)
		assert.NotEmpty(t, versions)
		assert.Contains(t, versions, version) // Current version should be in list

		// Test setting PHP version (to same version to avoid side effects)
		err = manager.SetVersion("php", version)
		require.NoError(t, err)

		// Verify version is still the same
		newVersion, err := manager.GetVersion("php")
		require.NoError(t, err)
		assert.Equal(t, version, newVersion)
	*/
}
