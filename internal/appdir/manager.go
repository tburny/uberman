package appdir

import (
	"fmt"
	"os"
	"path/filepath"
)

// Manager handles app directory structure
type Manager struct {
	appRoot string
	appName string
	dryRun  bool
	verbose bool
}

// NewManager creates a new directory structure manager
func NewManager(appName string, dryRun, verbose bool) (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	appRoot := filepath.Join(homeDir, "apps", appName)

	return &Manager{
		appRoot: appRoot,
		appName: appName,
		dryRun:  dryRun,
		verbose: verbose,
	}, nil
}

// AppRoot returns the root directory path for the app
func (m *Manager) AppRoot() string {
	return m.appRoot
}

// AppDir returns the path for the application code directory
func (m *Manager) AppDir() string {
	return filepath.Join(m.appRoot, "app")
}

// DataDir returns the path for persistent data
func (m *Manager) DataDir() string {
	return filepath.Join(m.appRoot, "data")
}

// LogsDir returns the path for logs
func (m *Manager) LogsDir() string {
	return filepath.Join(m.appRoot, "logs")
}

// BackupsDir returns the path for local backups
func (m *Manager) BackupsDir() string {
	return filepath.Join(m.appRoot, "backups")
}

// TmpDir returns the path for temporary files
func (m *Manager) TmpDir() string {
	return filepath.Join(m.appRoot, "tmp")
}

// ConfigFile returns the path to the app's uberman config
func (m *Manager) ConfigFile() string {
	return filepath.Join(m.appRoot, ".uberman.toml")
}

// Create creates the complete app directory structure
func (m *Manager) Create() error {
	dirs := []string{
		m.appRoot,
		m.AppDir(),
		m.DataDir(),
		m.LogsDir(),
		m.BackupsDir(),
		m.TmpDir(),
	}

	if m.verbose {
		fmt.Printf("Creating app directory structure for: %s\n", m.appName)
	}

	for _, dir := range dirs {
		if m.verbose {
			fmt.Printf("Creating directory: %s\n", dir)
		}

		if m.dryRun {
			fmt.Printf("[DRY RUN] Would create directory: %s\n", dir)
			continue
		}

		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create symlink for logs to central location
	if err := m.createLogSymlink(); err != nil {
		return err
	}

	if m.verbose {
		fmt.Printf("App directory structure created successfully\n")
	}

	return nil
}

// createLogSymlink creates a symlink from ~/logs/<app-name> to app logs
func (m *Manager) createLogSymlink() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	centralLogsDir := filepath.Join(homeDir, "logs")
	if err := os.MkdirAll(centralLogsDir, 0755); err != nil {
		return fmt.Errorf("failed to create central logs directory: %w", err)
	}

	symlinkPath := filepath.Join(centralLogsDir, m.appName)
	targetPath := m.LogsDir()

	if m.verbose {
		fmt.Printf("Creating log symlink: %s -> %s\n", symlinkPath, targetPath)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would create symlink: %s -> %s\n", symlinkPath, targetPath)
		return nil
	}

	// Remove existing symlink if it exists
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			return fmt.Errorf("failed to remove existing symlink: %w", err)
		}
	}

	if err := os.Symlink(targetPath, symlinkPath); err != nil {
		return fmt.Errorf("failed to create log symlink: %w", err)
	}

	return nil
}

// Exists checks if the app directory structure exists
func (m *Manager) Exists() bool {
	info, err := os.Stat(m.appRoot)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsEmpty checks if the app directory is empty
func (m *Manager) IsEmpty() (bool, error) {
	entries, err := os.ReadDir(m.AppDir())
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}
	return len(entries) == 0, nil
}

// Remove removes the entire app directory structure
func (m *Manager) Remove() error {
	if m.verbose {
		fmt.Printf("Removing app directory: %s\n", m.appRoot)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would remove directory: %s\n", m.appRoot)
		return nil
	}

	// Remove log symlink
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	symlinkPath := filepath.Join(homeDir, "logs", m.appName)
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			if m.verbose {
				fmt.Printf("Warning: failed to remove log symlink: %v\n", err)
			}
		}
	}

	// Remove app directory
	if err := os.RemoveAll(m.appRoot); err != nil {
		return fmt.Errorf("failed to remove app directory: %w", err)
	}

	if m.verbose {
		fmt.Printf("App directory removed successfully\n")
	}

	return nil
}

// Validate checks if the directory structure is valid
func (m *Manager) Validate() error {
	requiredDirs := []string{
		m.appRoot,
		m.AppDir(),
		m.DataDir(),
		m.LogsDir(),
	}

	for _, dir := range requiredDirs {
		info, err := os.Stat(dir)
		if err != nil {
			return fmt.Errorf("directory does not exist: %s", dir)
		}
		if !info.IsDir() {
			return fmt.Errorf("path is not a directory: %s", dir)
		}
	}

	// Check if config file exists
	if _, err := os.Stat(m.ConfigFile()); err != nil {
		return fmt.Errorf("config file does not exist: %s", m.ConfigFile())
	}

	return nil
}
