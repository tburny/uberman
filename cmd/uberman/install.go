package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tburny/uberman/internal/appdir"
	"github.com/tburny/uberman/internal/config"
	"github.com/tburny/uberman/internal/database"
	"github.com/tburny/uberman/internal/runtime"
)

var installCmd = &cobra.Command{
	Use:   "install <app-name>",
	Short: "Install an application",
	Long: `Install an application from a manifest file.

The manifest will be searched in the following locations:
  1. ./apps/<app-name>.toml
  2. ./apps/custom/<app-name>.toml
  3. ~/.uberman/apps/<app-name>.toml

The app will be installed to ~/apps/<app-name>/`,
	Args: cobra.ExactArgs(1),
	RunE: runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) error {
	appName := args[0]

	if verbose {
		fmt.Printf("Installing app: %s\n", appName)
	}

	// Find and load manifest
	manifestPath, err := config.FindManifest(appName)
	if err != nil {
		return fmt.Errorf("failed to find manifest: %w", err)
	}

	if verbose {
		fmt.Printf("Loading manifest from: %s\n", manifestPath)
	}

	manifest, err := config.LoadManifest(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to load manifest: %w", err)
	}

	// Validate manifest
	if err := manifest.Validate(); err != nil {
		return fmt.Errorf("invalid manifest: %w", err)
	}

	if verbose {
		fmt.Printf("Manifest validated successfully\n")
		fmt.Printf("App: %s v%s (%s)\n", manifest.App.Name, manifest.App.Version, manifest.App.Type)
		fmt.Printf("Runtime: %s %s\n", manifest.Runtime.Language, manifest.Runtime.Version)
	}

	// Create app directory structure
	dirManager, err := appdir.NewManager(appName, dryRun, verbose)
	if err != nil {
		return fmt.Errorf("failed to create directory manager: %w", err)
	}

	// Check if app already exists
	if dirManager.Exists() {
		return fmt.Errorf("app already exists at: %s", dirManager.AppRoot())
	}

	// Create directories
	if err := dirManager.Create(); err != nil {
		return fmt.Errorf("failed to create app directories: %w", err)
	}

	fmt.Printf("✓ Created app directory structure at: %s\n", dirManager.AppRoot())

	// Set runtime version
	runtimeMgr := runtime.NewManager(dryRun, verbose)
	if err := runtimeMgr.SetVersion(manifest.Runtime.Language, manifest.Runtime.Version); err != nil {
		return fmt.Errorf("failed to set runtime version: %w", err)
	}

	fmt.Printf("✓ Set %s version to %s\n", manifest.Runtime.Language, manifest.Runtime.Version)

	// Create database if required
	if manifest.Database.Required {
		dbMgr := database.NewMySQLManager(dryRun, verbose)

		username := os.Getenv("USER")
		dbName := manifest.Database.Name
		if dbName == "" {
			dbName = database.GenerateDatabaseName(username, appName)
		}

		if err := dbMgr.CreateDatabase(dbName); err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}

		fmt.Printf("✓ Created database: %s\n", dbName)
	}

	// TODO: Implement remaining installation steps:
	// - Download/clone application code
	// - Install dependencies (composer, pip, npm)
	// - Generate configuration files
	// - Set up web backend
	// - Create supervisord services
	// - Set up cron jobs
	// - Run post-install hooks

	fmt.Printf("\n✓ App '%s' installed successfully!\n", appName)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  - Check status: uberman status %s\n", appName)
	fmt.Printf("  - View logs: tail -f ~/logs/%s/\n", appName)
	fmt.Printf("  - App directory: %s\n", dirManager.AppRoot())

	return nil
}
