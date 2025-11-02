package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// AppManifest represents the complete app configuration
type AppManifest struct {
	App      AppConfig       `toml:"app"`
	Runtime  RuntimeConfig   `toml:"runtime"`
	Database DatabaseConfig  `toml:"database"`
	Install  InstallConfig   `toml:"install"`
	Web      WebConfig       `toml:"web"`
	Services []ServiceConfig `toml:"services"`
	Cron     CronConfig      `toml:"cron"`
	Backup   BackupConfig    `toml:"backup"`
}

// AppConfig defines basic app metadata
type AppConfig struct {
	Name        string `toml:"name"`
	Type        string `toml:"type"` // php, python, nodejs, ruby, go
	Version     string `toml:"version"`
	Description string `toml:"description"`
}

// RuntimeConfig defines language runtime requirements
type RuntimeConfig struct {
	Language string `toml:"language"`
	Version  string `toml:"version"`
}

// DatabaseConfig defines database requirements
type DatabaseConfig struct {
	Type     string `toml:"type"` // mysql, postgresql, mongodb, redis
	Required bool   `toml:"required"`
	Name     string `toml:"name"` // optional custom name, defaults to ${USER}_${APP}
}

// InstallConfig defines installation method
type InstallConfig struct {
	Method   string            `toml:"method"` // download, git, cli_tool, composer, pip, npm
	Source   string            `toml:"source"` // URL, git repo, or package name
	Command  string            `toml:"command"`
	Location string            `toml:"location"` // relative to app root
	Env      map[string]string `toml:"env"`
}

// WebConfig defines web server configuration
type WebConfig struct {
	Backend      string   `toml:"backend"` // apache, http
	Port         int      `toml:"port"`
	DocumentRoot string   `toml:"document_root"`
	StaticPaths  []string `toml:"static_paths"`
}

// ServiceConfig defines a supervisord service
type ServiceConfig struct {
	Name         string            `toml:"name"`
	Command      string            `toml:"command"`
	Directory    string            `toml:"directory"`
	Port         int               `toml:"port"`
	Environment  map[string]string `toml:"environment"`
	StartSecs    int               `toml:"startsecs"`
	AutoRestart  bool              `toml:"autorestart"`
	StdoutLogfile string           `toml:"stdout_logfile"`
	StderrLogfile string           `toml:"stderr_logfile"`
}

// CronConfig defines scheduled tasks
type CronConfig struct {
	Jobs []CronJob `toml:"jobs"`
}

// CronJob represents a single cron entry
type CronJob struct {
	Schedule string `toml:"schedule"`
	Command  string `toml:"command"`
	Log      string `toml:"log"`
}

// BackupConfig defines backup strategy
type BackupConfig struct {
	Include []string `toml:"include"`
	Exclude []string `toml:"exclude"`
}

// LoadManifest loads an app manifest from a TOML file
func LoadManifest(path string) (*AppManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	var manifest AppManifest
	if err := toml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	return &manifest, nil
}

// FindManifest searches for an app manifest in standard locations
func FindManifest(appName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Search order:
	// 1. ~/apps/<app-name>/.uberman.toml (installed app)
	// 2. ./apps/<app-name>.toml (project directory)
	// 3. ./apps/custom/<app-name>.toml (user-defined)
	// 4. ~/.uberman/apps/<app-name>.toml (user home)

	searchPaths := []string{
		filepath.Join(homeDir, "apps", appName, ".uberman.toml"),
		filepath.Join("apps", appName+".toml"),
		filepath.Join("apps", "custom", appName+".toml"),
		filepath.Join(homeDir, ".uberman", "apps", appName+".toml"),
	}

	for _, path := range searchPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("manifest not found for app: %s", appName)
}

// Validate checks if the manifest is valid
func (m *AppManifest) Validate() error {
	if m.App.Name == "" {
		return fmt.Errorf("app name is required")
	}
	if m.App.Type == "" {
		return fmt.Errorf("app type is required")
	}
	if m.Runtime.Language == "" {
		return fmt.Errorf("runtime language is required")
	}

	// Validate app type
	validTypes := map[string]bool{
		"php": true, "python": true, "nodejs": true,
		"ruby": true, "go": true, "static": true,
	}
	if !validTypes[m.App.Type] {
		return fmt.Errorf("invalid app type: %s", m.App.Type)
	}

	// Validate web backend for non-PHP apps
	if m.App.Type != "php" && m.Web.Backend == "http" && m.Web.Port == 0 {
		return fmt.Errorf("port is required for http backend")
	}

	return nil
}
