package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tburny/uberman/internal/config"
)

var (
	version = "0.1.0"
	cfgFile string
	dryRun  bool
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "uberman",
	Short: "Uberspace App Management System",
	Long: `uberman provides reproducible installation, upgrades, backups,
and deployment strategies for applications on Uberspace hosting.

Each app is installed in a self-contained directory following the
convention: ~/apps/<app-name>`,
	Version: version,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uberman.toml)")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "show what would be done without making changes")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
