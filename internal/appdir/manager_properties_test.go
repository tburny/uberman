package appdir

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// Helper to set up temp home directory
func setupTempHome(t *rapid.T) (tmpHome string, cleanup func()) {
	tmpHome, err := os.MkdirTemp("", "uberman-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpHome)

	cleanup = func() {
		os.Setenv("HOME", originalHome)
		os.RemoveAll(tmpHome)
	}
	return tmpHome, cleanup
}

// Property: When a valid directory structure is created, all required subdirectories should exist
// REQ-PBT-APPDIR-001
func TestProperty_DirectoryCreation_AllSubdirectoriesExist(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate valid app name
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		// Set up temporary home directory
		_, cleanup := setupTempHome(t)
		defer cleanup()

		// Create manager
		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// Create directory structure
		err = manager.Create()
		if err != nil {
			t.Fatalf("failed to create directory structure: %v", err)
		}

		// Property: All required directories should exist
		requiredDirs := []string{
			manager.AppRoot(),
			manager.AppDir(),
			manager.DataDir(),
			manager.LogsDir(),
			manager.BackupsDir(),
			manager.TmpDir(),
		}

		for _, dir := range requiredDirs {
			info, err := os.Stat(dir)
			if err != nil {
				t.Fatalf("required directory does not exist: %s (error: %v)", dir, err)
			}
			if !info.IsDir() {
				t.Fatalf("path is not a directory: %s", dir)
			}
		}
	})
}

// Property: When ValidateStructure is called on a created structure, validation should always pass
// REQ-PBT-APPDIR-002
func TestProperty_CreatedStructure_AlwaysValidates(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		// Set up temporary home directory
		_, cleanup := setupTempHome(t); defer cleanup()

		// Create manager and structure
		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		err = manager.Create()
		if err != nil {
			t.Fatalf("failed to create structure: %v", err)
		}

		// Create config file (required for validation)
		configContent := []byte("[app]\nname = \"" + appName + "\"\n")
		err = os.WriteFile(manager.ConfigFile(), configContent, 0644)
		if err != nil {
			t.Fatalf("failed to create config file: %v", err)
		}

		// Property: Validation should always pass for created structure
		err = manager.Validate()
		if err != nil {
			t.Fatalf("validation failed for created structure: %v", err)
		}
	})
}

// Property: The symlink path resolution should always return an absolute path
// REQ-PBT-APPDIR-003
func TestProperty_PathResolution_AlwaysAbsolute(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		// Set up temporary home directory
		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// Property: All path methods should return absolute paths
		paths := []struct {
			name string
			path string
		}{
			{"AppRoot", manager.AppRoot()},
			{"AppDir", manager.AppDir()},
			{"DataDir", manager.DataDir()},
			{"LogsDir", manager.LogsDir()},
			{"BackupsDir", manager.BackupsDir()},
			{"TmpDir", manager.TmpDir()},
			{"ConfigFile", manager.ConfigFile()},
		}

		for _, p := range paths {
			if !filepath.IsAbs(p.path) {
				t.Fatalf("%s returned relative path: %s", p.name, p.path)
			}
		}
	})
}

// Property: AppRoot should always follow the pattern ~/apps/<app-name>
func TestProperty_AppRoot_FollowsConvention(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		tmpHome, cleanup := setupTempHome(t)
		defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		appRoot := manager.AppRoot()
		expectedRoot := filepath.Join(tmpHome, "apps", appName)

		if appRoot != expectedRoot {
			t.Fatalf("AppRoot doesn't follow convention: got %s, want %s", appRoot, expectedRoot)
		}
	})
}

// Property: Manager methods should always be consistent with AppRoot
func TestProperty_ManagerPaths_ConsistentWithRoot(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		appRoot := manager.AppRoot()

		// All subdirectories should be children of AppRoot
		subdirs := []struct {
			name string
			path string
		}{
			{"AppDir", manager.AppDir()},
			{"DataDir", manager.DataDir()},
			{"LogsDir", manager.LogsDir()},
			{"BackupsDir", manager.BackupsDir()},
			{"TmpDir", manager.TmpDir()},
		}

		for _, subdir := range subdirs {
			if !strings.HasPrefix(subdir.path, appRoot) {
				t.Fatalf("%s (%s) is not a child of AppRoot (%s)", subdir.name, subdir.path, appRoot)
			}
		}
	})
}

// Property: Exists should return false before Create and true after
func TestProperty_Exists_BeforeAndAfterCreate(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// Property: Should not exist before creation
		if manager.Exists() {
			t.Fatal("directory exists before Create was called")
		}

		// Create structure
		err = manager.Create()
		if err != nil {
			t.Fatalf("failed to create structure: %v", err)
		}

		// Property: Should exist after creation
		if !manager.Exists() {
			t.Fatal("directory does not exist after Create was called")
		}
	})
}

// Property: IsEmpty should return true for freshly created app directory
func TestProperty_IsEmpty_AfterCreation(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		err = manager.Create()
		if err != nil {
			t.Fatalf("failed to create structure: %v", err)
		}

		// Property: App directory should be empty after creation
		isEmpty, err := manager.IsEmpty()
		if err != nil {
			t.Fatalf("IsEmpty returned error: %v", err)
		}
		if !isEmpty {
			t.Fatal("app directory is not empty after creation")
		}
	})
}

// Property: DryRun mode should never create actual directories
func TestProperty_DryRun_NoActualCreation(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		// Create manager in dry-run mode
		manager, err := NewManager(appName, true, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// Call Create in dry-run mode
		err = manager.Create()
		if err != nil {
			t.Fatalf("dry-run Create failed: %v", err)
		}

		// Property: Directory should not exist after dry-run
		if manager.Exists() {
			t.Fatal("directory exists after dry-run Create (should not have been created)")
		}
	})
}

// Property: Remove should make Exists return false
func TestProperty_Remove_MakesNotExist(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// Create and verify exists
		err = manager.Create()
		if err != nil {
			t.Fatalf("failed to create structure: %v", err)
		}

		if !manager.Exists() {
			t.Fatal("directory does not exist after creation")
		}

		// Remove and verify does not exist
		err = manager.Remove()
		if err != nil {
			t.Fatalf("failed to remove structure: %v", err)
		}

		// Property: Should not exist after removal
		if manager.Exists() {
			t.Fatal("directory still exists after Remove was called")
		}
	})
}

// Property: ConfigFile should always end with .uberman.toml
func TestProperty_ConfigFile_CorrectExtension(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		configFile := manager.ConfigFile()

		// Property: Should end with .uberman.toml
		if !strings.HasSuffix(configFile, ".uberman.toml") {
			t.Fatalf("ConfigFile doesn't end with .uberman.toml: %s", configFile)
		}

		// Property: Should be in AppRoot
		if !strings.HasPrefix(configFile, manager.AppRoot()) {
			t.Fatalf("ConfigFile is not in AppRoot: %s", configFile)
		}
	})
}

// Property: Multiple Create calls should be idempotent (not fail)
func TestProperty_Create_Idempotent(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		_, cleanup := setupTempHome(t); defer cleanup()

		manager, err := NewManager(appName, false, false)
		if err != nil {
			t.Fatalf("failed to create manager: %v", err)
		}

		// First create
		err = manager.Create()
		if err != nil {
			t.Fatalf("first Create failed: %v", err)
		}

		// Second create should not fail (idempotent)
		err = manager.Create()
		if err != nil {
			t.Fatalf("second Create failed (not idempotent): %v", err)
		}

		// Property: Should still exist
		if !manager.Exists() {
			t.Fatal("directory does not exist after second Create")
		}
	})
}
