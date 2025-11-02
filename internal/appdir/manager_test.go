package appdir

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tburny/uberman/internal/testutil"
)

func TestNewManager(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	// Set HOME to temp directory
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	tests := []struct {
		name    string
		appName string
		dryRun  bool
		verbose bool
	}{
		{
			name:    "create manager with normal mode",
			appName: "testapp",
			dryRun:  false,
			verbose: false,
		},
		{
			name:    "create manager with dry-run",
			appName: "wordpress",
			dryRun:  true,
			verbose: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager, err := NewManager(tt.appName, tt.dryRun, tt.verbose)

			require.NoError(t, err)
			require.NotNil(t, manager)

			assert.Equal(t, tt.dryRun, manager.dryRun)
			assert.Equal(t, tt.verbose, manager.verbose)
			assert.Equal(t, tt.appName, manager.appName)
			assert.Contains(t, manager.appRoot, filepath.Join("apps", tt.appName))
		})
	}
}

func TestManager_PathGetters(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	manager, err := NewManager("myapp", false, false)
	require.NoError(t, err)

	expectedRoot := filepath.Join(tmpDir, "apps", "myapp")

	tests := []struct {
		name     string
		getter   func() string
		expected string
	}{
		{
			name:     "AppRoot",
			getter:   manager.AppRoot,
			expected: expectedRoot,
		},
		{
			name:     "AppDir",
			getter:   manager.AppDir,
			expected: filepath.Join(expectedRoot, "app"),
		},
		{
			name:     "DataDir",
			getter:   manager.DataDir,
			expected: filepath.Join(expectedRoot, "data"),
		},
		{
			name:     "LogsDir",
			getter:   manager.LogsDir,
			expected: filepath.Join(expectedRoot, "logs"),
		},
		{
			name:     "BackupsDir",
			getter:   manager.BackupsDir,
			expected: filepath.Join(expectedRoot, "backups"),
		},
		{
			name:     "TmpDir",
			getter:   manager.TmpDir,
			expected: filepath.Join(expectedRoot, "tmp"),
		},
		{
			name:     "ConfigFile",
			getter:   manager.ConfigFile,
			expected: filepath.Join(expectedRoot, ".uberman.toml"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.getter()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestManager_Create(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("create directory structure", func(t *testing.T) {
		manager, err := NewManager("testapp", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		// Verify all directories were created
		dirs := []string{
			manager.AppRoot(),
			manager.AppDir(),
			manager.DataDir(),
			manager.LogsDir(),
			manager.BackupsDir(),
			manager.TmpDir(),
		}

		for _, dir := range dirs {
			info, err := os.Stat(dir)
			assert.NoError(t, err, "Directory should exist: %s", dir)
			assert.True(t, info.IsDir(), "Path should be a directory: %s", dir)
		}

		// Verify log symlink was created
		symlinkPath := filepath.Join(tmpDir, "logs", "testapp")
		info, err := os.Lstat(symlinkPath)
		assert.NoError(t, err, "Symlink should exist")
		assert.NotNil(t, info)

		// Verify symlink points to correct location
		target, err := os.Readlink(symlinkPath)
		assert.NoError(t, err)
		assert.Equal(t, manager.LogsDir(), target)
	})

	t.Run("create with dry-run", func(t *testing.T) {
		manager, err := NewManager("dryrunapp", true, true)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		// Verify directories were NOT created
		_, err = os.Stat(manager.AppRoot())
		assert.True(t, os.IsNotExist(err), "Directory should not exist in dry-run mode")
	})

	t.Run("create idempotent", func(t *testing.T) {
		manager, err := NewManager("idempotent", false, false)
		require.NoError(t, err)

		// Create once
		err = manager.Create()
		require.NoError(t, err)

		// Create again - should not error
		err = manager.Create()
		assert.NoError(t, err, "Creating existing structure should be idempotent")
	})
}

func TestManager_Exists(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("does not exist initially", func(t *testing.T) {
		manager, err := NewManager("nonexistent", false, false)
		require.NoError(t, err)

		exists := manager.Exists()
		assert.False(t, exists)
	})

	t.Run("exists after creation", func(t *testing.T) {
		manager, err := NewManager("existing", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		exists := manager.Exists()
		assert.True(t, exists)
	})
}

func TestManager_IsEmpty(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("empty after creation", func(t *testing.T) {
		manager, err := NewManager("emptyapp", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		isEmpty, err := manager.IsEmpty()
		assert.NoError(t, err)
		assert.True(t, isEmpty)
	})

	t.Run("not empty after adding files", func(t *testing.T) {
		manager, err := NewManager("nonempty", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		// Add a file to app directory
		testFile := filepath.Join(manager.AppDir(), "test.txt")
		err = os.WriteFile(testFile, []byte("test"), 0644)
		require.NoError(t, err)

		isEmpty, err := manager.IsEmpty()
		assert.NoError(t, err)
		assert.False(t, isEmpty)
	})

	t.Run("non-existent directory", func(t *testing.T) {
		manager, err := NewManager("notcreated", false, false)
		require.NoError(t, err)

		isEmpty, err := manager.IsEmpty()
		assert.NoError(t, err)
		assert.True(t, isEmpty, "Non-existent directory should be considered empty")
	})
}

func TestManager_Remove(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("remove app directory", func(t *testing.T) {
		manager, err := NewManager("toremove", false, false)
		require.NoError(t, err)

		// Create the structure
		err = manager.Create()
		require.NoError(t, err)

		// Verify it exists
		assert.True(t, manager.Exists())

		// Remove it
		err = manager.Remove()
		require.NoError(t, err)

		// Verify it's gone
		assert.False(t, manager.Exists())

		// Verify symlink is also gone
		symlinkPath := filepath.Join(tmpDir, "logs", "toremove")
		_, err = os.Lstat(symlinkPath)
		assert.True(t, os.IsNotExist(err), "Symlink should be removed")
	})

	t.Run("remove with dry-run", func(t *testing.T) {
		manager, err := NewManager("dryremove", false, false)
		require.NoError(t, err)

		// Create normally
		err = manager.Create()
		require.NoError(t, err)

		// Switch to dry-run mode
		manager.dryRun = true

		// Try to remove
		err = manager.Remove()
		require.NoError(t, err)

		// Verify it still exists
		assert.True(t, manager.Exists(), "Directory should still exist in dry-run mode")
	})

	t.Run("remove non-existent directory", func(t *testing.T) {
		manager, err := NewManager("neverexisted", false, false)
		require.NoError(t, err)

		// Remove non-existent directory - should not error
		err = manager.Remove()
		assert.NoError(t, err, "Removing non-existent directory should not error")
	})
}

func TestManager_Validate(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("validate complete structure", func(t *testing.T) {
		manager, err := NewManager("validapp", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		// Create config file
		configContent := testutil.MinimalManifest()
		err = os.WriteFile(manager.ConfigFile(), []byte(configContent), 0644)
		require.NoError(t, err)

		// Validate
		err = manager.Validate()
		assert.NoError(t, err)
	})

	t.Run("validate without config file", func(t *testing.T) {
		manager, err := NewManager("noconfig", false, false)
		require.NoError(t, err)

		err = manager.Create()
		require.NoError(t, err)

		// Don't create config file

		// Validate - should fail
		err = manager.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "config file does not exist")
	})

	t.Run("validate missing directories", func(t *testing.T) {
		manager, err := NewManager("incomplete", false, false)
		require.NoError(t, err)

		// Only create root directory
		err = os.MkdirAll(manager.AppRoot(), 0755)
		require.NoError(t, err)

		// Validate - should fail
		err = manager.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory does not exist")
	})

	t.Run("validate non-existent app", func(t *testing.T) {
		manager, err := NewManager("notexist", false, false)
		require.NoError(t, err)

		// Validate without creating
		err = manager.Validate()
		assert.Error(t, err)
	})
}

func TestManager_CreateLogSymlink(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	t.Run("replace existing symlink", func(t *testing.T) {
		manager, err := NewManager("replacelink", false, false)
		require.NoError(t, err)

		// Create structure first time
		err = manager.Create()
		require.NoError(t, err)

		symlinkPath := filepath.Join(tmpDir, "logs", "replacelink")

		// Verify initial symlink
		target1, err := os.Readlink(symlinkPath)
		require.NoError(t, err)

		// Create again (should replace symlink)
		err = manager.Create()
		require.NoError(t, err)

		// Verify symlink still works
		target2, err := os.Readlink(symlinkPath)
		require.NoError(t, err)

		assert.Equal(t, target1, target2)
	})
}

func TestManager_PermissionsCheck(t *testing.T) {
	tmpDir := testutil.CreateTempDir(t)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() {
		os.Setenv("HOME", originalHome)
	})

	manager, err := NewManager("permcheck", false, false)
	require.NoError(t, err)

	err = manager.Create()
	require.NoError(t, err)

	// Check that directories have correct permissions (0755)
	dirs := []string{
		manager.AppRoot(),
		manager.AppDir(),
		manager.DataDir(),
	}

	for _, dir := range dirs {
		info, err := os.Stat(dir)
		require.NoError(t, err)

		// Check directory has read/write/execute for owner
		mode := info.Mode().Perm()
		assert.Equal(t, os.FileMode(0755), mode, "Directory %s should have 0755 permissions", dir)
	}
}
