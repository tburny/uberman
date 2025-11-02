package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tburny/uberman/internal/testutil"
)

func TestLoadManifest(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectError bool
		validate    func(t *testing.T, manifest *AppManifest)
	}{
		{
			name:        "valid minimal manifest",
			content:     testutil.MinimalManifest(),
			expectError: false,
			validate: func(t *testing.T, manifest *AppManifest) {
				assert.Equal(t, "testapp", manifest.App.Name)
				assert.Equal(t, "php", manifest.App.Type)
				assert.Equal(t, "php", manifest.Runtime.Language)
				assert.Equal(t, "8.3", manifest.Runtime.Version)
				assert.Equal(t, "mysql", manifest.Database.Type)
				assert.False(t, manifest.Database.Required)
				assert.Equal(t, "apache", manifest.Web.Backend)
			},
		},
		{
			name:        "valid nodejs manifest with services",
			content:     testutil.NodeJSManifest(),
			expectError: false,
			validate: func(t *testing.T, manifest *AppManifest) {
				assert.Equal(t, "nodeapp", manifest.App.Name)
				assert.Equal(t, "nodejs", manifest.App.Type)
				assert.Equal(t, "http", manifest.Web.Backend)
				assert.Equal(t, 8080, manifest.Web.Port)
				assert.Len(t, manifest.Services, 1)
				assert.Equal(t, "nodeapp", manifest.Services[0].Name)
				assert.Equal(t, "node server.js", manifest.Services[0].Command)
				assert.Equal(t, 8080, manifest.Services[0].Port)
				assert.Equal(t, "production", manifest.Services[0].Environment["NODE_ENV"])
			},
		},
		{
			name:        "invalid TOML syntax",
			content:     "[app\nname = invalid",
			expectError: true,
		},
		{
			name:        "empty file",
			content:     "",
			expectError: false, // TOML parser accepts empty input
			validate: func(t *testing.T, manifest *AppManifest) {
				// Empty manifest should have zero values
				assert.Empty(t, manifest.App.Name)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := testutil.CreateTempDir(t)
			manifestPath := testutil.CreateTestFile(t, tmpDir, "app.toml", tt.content)

			// Load manifest
			manifest, err := LoadManifest(manifestPath)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, manifest)

			if tt.validate != nil {
				tt.validate(t, manifest)
			}
		})
	}
}

func TestLoadManifest_FileErrors(t *testing.T) {
	t.Run("non-existent file", func(t *testing.T) {
		manifest, err := LoadManifest("/nonexistent/path/app.toml")
		assert.Error(t, err)
		assert.Nil(t, manifest)
		assert.Contains(t, err.Error(), "failed to read manifest")
	})
}

func TestFindManifest(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	// Set HOME to temp directory for testing
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	tests := []struct {
		name         string
		appName      string
		setupFiles   func(t *testing.T)
		expectError  bool
		expectedPath string // relative to tmpDir or current dir
	}{
		{
			name:    "installed app instance",
			appName: "wordpress",
			setupFiles: func(t *testing.T) {
				testutil.CreateTestFile(t, tmpDir, "apps/wordpress/.uberman.toml", testutil.MinimalManifest())
			},
			expectError:  false,
			expectedPath: "apps/wordpress/.uberman.toml",
		},
		{
			name:    "project examples directory (new structure)",
			appName: "ghost",
			setupFiles: func(t *testing.T) {
				// Create in current directory
				testutil.CreateTestFile(t, ".", "apps/examples/ghost/app.toml", testutil.NodeJSManifest())
			},
			expectError:  false,
			expectedPath: "apps/examples/ghost/app.toml",
		},
		{
			name:    "custom app directory",
			appName: "myapp",
			setupFiles: func(t *testing.T) {
				testutil.CreateTestFile(t, ".", "apps/custom/myapp/app.toml", testutil.MinimalManifest())
			},
			expectError:  false,
			expectedPath: "apps/custom/myapp/app.toml",
		},
		{
			name:    "legacy flat structure",
			appName: "legacy",
			setupFiles: func(t *testing.T) {
				testutil.CreateTestFile(t, ".", "apps/legacy.toml", testutil.MinimalManifest())
			},
			expectError:  false,
			expectedPath: "apps/legacy.toml",
		},
		{
			name:        "app not found",
			appName:     "nonexistent",
			setupFiles:  func(t *testing.T) {},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up files from previous test
			os.RemoveAll(filepath.Join(tmpDir, "apps"))
			os.RemoveAll("apps")

			tt.setupFiles(t)

			// Clean up after test
			t.Cleanup(func() {
				os.RemoveAll("apps")
			})

			path, err := FindManifest(tt.appName)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "manifest not found")
				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, path)

			// Verify file exists
			_, err = os.Stat(path)
			assert.NoError(t, err)
		})
	}
}

func TestGetAppDirectory(t *testing.T) {
	tests := []struct {
		name         string
		manifestPath string
		expected     string
	}{
		{
			name:         "app.toml in directory",
			manifestPath: "/home/user/apps/examples/wordpress/app.toml",
			expected:     "/home/user/apps/examples/wordpress",
		},
		{
			name:         "direct .toml file",
			manifestPath: "/home/user/apps/wordpress.toml",
			expected:     "/home/user/apps",
		},
		{
			name:         "instance config file",
			manifestPath: "/home/user/apps/wordpress/.uberman.toml",
			expected:     "/home/user/apps/wordpress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAppDirectory(tt.manifestPath)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindHookScript(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	tests := []struct {
		name         string
		manifestPath string
		hookName     string
		setupFiles   func(t *testing.T) string
		expectError  bool
	}{
		{
			name:     "hook script exists",
			hookName: "post-install",
			setupFiles: func(t *testing.T) string {
				manifestPath := testutil.CreateTestFile(t, tmpDir, "wordpress/app.toml", testutil.MinimalManifest())
				testutil.CreateTestFile(t, tmpDir, "wordpress/hooks/post-install.sh", "#!/bin/bash\necho 'post-install'")
				return manifestPath
			},
			expectError: false,
		},
		{
			name:     "hook script does not exist",
			hookName: "pre-upgrade",
			setupFiles: func(t *testing.T) string {
				return testutil.CreateTestFile(t, tmpDir, "ghost/app.toml", testutil.NodeJSManifest())
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manifestPath := tt.setupFiles(t)
			hookPath, err := FindHookScript(manifestPath, tt.hookName)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "hook script")
				assert.Contains(t, err.Error(), "not found")
				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, hookPath)

			// Verify hook file exists
			_, err = os.Stat(hookPath)
			assert.NoError(t, err)
		})
	}
}

func TestAppManifest_Validate(t *testing.T) {
	tests := []struct {
		name        string
		manifest    *AppManifest
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid PHP app",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "wordpress",
					Type: "php",
				},
				Runtime: RuntimeConfig{
					Language: "php",
					Version:  "8.3",
				},
				Web: WebConfig{
					Backend: "apache",
				},
			},
			expectError: false,
		},
		{
			name: "valid Node.js app with HTTP backend",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "ghost",
					Type: "nodejs",
				},
				Runtime: RuntimeConfig{
					Language: "node",
					Version:  "20",
				},
				Web: WebConfig{
					Backend: "http",
					Port:    8080,
				},
			},
			expectError: false,
		},
		{
			name: "missing app name",
			manifest: &AppManifest{
				App: AppConfig{
					Type: "php",
				},
				Runtime: RuntimeConfig{
					Language: "php",
				},
			},
			expectError: true,
			errorMsg:    "app name is required",
		},
		{
			name: "missing app type",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "testapp",
				},
				Runtime: RuntimeConfig{
					Language: "php",
				},
			},
			expectError: true,
			errorMsg:    "app type is required",
		},
		{
			name: "missing runtime language",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "testapp",
					Type: "php",
				},
				Runtime: RuntimeConfig{},
			},
			expectError: true,
			errorMsg:    "runtime language is required",
		},
		{
			name: "invalid app type",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "testapp",
					Type: "invalid",
				},
				Runtime: RuntimeConfig{
					Language: "php",
				},
			},
			expectError: true,
			errorMsg:    "invalid app type",
		},
		{
			name: "HTTP backend missing port",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "nodeapp",
					Type: "nodejs",
				},
				Runtime: RuntimeConfig{
					Language: "node",
					Version:  "20",
				},
				Web: WebConfig{
					Backend: "http",
					Port:    0, // Missing port
				},
			},
			expectError: true,
			errorMsg:    "port is required for http backend",
		},
		{
			name: "valid static app",
			manifest: &AppManifest{
				App: AppConfig{
					Name: "staticsite",
					Type: "static",
				},
				Runtime: RuntimeConfig{
					Language: "static",
				},
				Web: WebConfig{
					Backend: "apache",
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.manifest.Validate()

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestAppManifest_ValidateAllTypes(t *testing.T) {
	// Test all valid app types
	validTypes := []string{"php", "python", "nodejs", "ruby", "go", "static"}

	for _, appType := range validTypes {
		t.Run("valid_type_"+appType, func(t *testing.T) {
			manifest := &AppManifest{
				App: AppConfig{
					Name: "testapp",
					Type: appType,
				},
				Runtime: RuntimeConfig{
					Language: "php", // Just needs to be non-empty
				},
				Web: WebConfig{
					Backend: "apache",
				},
			}

			err := manifest.Validate()
			assert.NoError(t, err)
		})
	}
}

// TestLoadManifest_Integration loads a real WordPress manifest
func TestLoadManifest_Integration(t *testing.T) {
	// This test checks if we can load the actual WordPress manifest from the repo
	wordpressManifest := filepath.Join("..", "..", "apps", "examples", "wordpress", "app.toml")

	// Skip if file doesn't exist (e.g., in isolated test environment)
	if _, err := os.Stat(wordpressManifest); os.IsNotExist(err) {
		t.Skip("WordPress manifest not found, skipping integration test")
	}

	manifest, err := LoadManifest(wordpressManifest)
	require.NoError(t, err)
	require.NotNil(t, manifest)

	// Validate WordPress-specific fields
	assert.Equal(t, "wordpress", manifest.App.Name)
	assert.Equal(t, "php", manifest.App.Type)
	assert.Equal(t, "php", manifest.Runtime.Language)
	assert.Equal(t, "8.3", manifest.Runtime.Version)
	assert.Equal(t, "mysql", manifest.Database.Type)
	assert.True(t, manifest.Database.Required)
	assert.Equal(t, "apache", manifest.Web.Backend)

	// Validate cron jobs
	assert.NotEmpty(t, manifest.Cron.Jobs)

	// Validate backup config
	assert.NotEmpty(t, manifest.Backup.Include)
	assert.NotEmpty(t, manifest.Backup.Exclude)

	// Run validation
	err = manifest.Validate()
	assert.NoError(t, err)
}
