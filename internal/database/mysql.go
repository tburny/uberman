package database

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// MySQLManager handles MySQL database operations on Uberspace
type MySQLManager struct {
	dryRun  bool
	verbose bool
}

// NewMySQLManager creates a new MySQL manager
func NewMySQLManager(dryRun, verbose bool) *MySQLManager {
	return &MySQLManager{
		dryRun:  dryRun,
		verbose: verbose,
	}
}

// Credentials holds MySQL connection information
type Credentials struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// GetCredentials reads MySQL credentials from ~/.my.cnf
func (m *MySQLManager) GetCredentials() (*Credentials, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	myCnfPath := filepath.Join(homeDir, ".my.cnf")
	file, err := os.Open(myCnfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open ~/.my.cnf: %w", err)
	}
	defer file.Close()

	creds := &Credentials{
		Host: "localhost",
		Port: "3306",
	}

	scanner := bufio.NewScanner(file)
	inClientSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "[client]" {
			inClientSection = true
			continue
		}

		if strings.HasPrefix(line, "[") {
			inClientSection = false
			continue
		}

		if inClientSection && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), "\"'")

			switch key {
			case "user":
				creds.User = value
				creds.Database = value // Default database is username
			case "password":
				creds.Password = value
			case "host":
				creds.Host = value
			case "port":
				creds.Port = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading ~/.my.cnf: %w", err)
	}

	if creds.User == "" || creds.Password == "" {
		return nil, fmt.Errorf("could not find MySQL credentials in ~/.my.cnf")
	}

	return creds, nil
}

// CreateDatabase creates a new MySQL database
func (m *MySQLManager) CreateDatabase(name string) error {
	if m.verbose {
		fmt.Printf("Creating database: %s\n", name)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: mysql -e \"CREATE DATABASE IF NOT EXISTS %s\"\n", name)
		return nil
	}

	cmd := exec.Command("mysql", "-e", fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", name))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create database: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Database %s created successfully\n", name)
	}

	return nil
}

// DatabaseExists checks if a database exists
func (m *MySQLManager) DatabaseExists(name string) (bool, error) {
	cmd := exec.Command("mysql", "-e", "SHOW DATABASES")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to list databases: %w", err)
	}

	databases := strings.Split(string(output), "\n")
	for _, db := range databases {
		if strings.TrimSpace(db) == name {
			return true, nil
		}
	}

	return false, nil
}

// DropDatabase drops a database (with safety check)
func (m *MySQLManager) DropDatabase(name string) error {
	if m.verbose {
		fmt.Printf("Dropping database: %s\n", name)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: mysql -e \"DROP DATABASE IF EXISTS %s\"\n", name)
		return nil
	}

	cmd := exec.Command("mysql", "-e", fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", name))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to drop database: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Database %s dropped successfully\n", name)
	}

	return nil
}

// ExportDatabase exports a database to SQL file
func (m *MySQLManager) ExportDatabase(name, outputPath string) error {
	if m.verbose {
		fmt.Printf("Exporting database %s to %s\n", name, outputPath)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: mysqldump %s > %s\n", name, outputPath)
		return nil
	}

	cmd := exec.Command("mysqldump", name)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to export database: %w", err)
	}

	if err := os.WriteFile(outputPath, output, 0644); err != nil {
		return fmt.Errorf("failed to write database dump: %w", err)
	}

	if m.verbose {
		fmt.Printf("Database exported successfully\n")
	}

	return nil
}

// ImportDatabase imports a database from SQL file
func (m *MySQLManager) ImportDatabase(name, inputPath string) error {
	if m.verbose {
		fmt.Printf("Importing database %s from %s\n", name, inputPath)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: mysql %s < %s\n", name, inputPath)
		return nil
	}

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	cmd := exec.Command("mysql", name)
	cmd.Stdin = strings.NewReader(string(data))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to import database: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Database imported successfully\n")
	}

	return nil
}

// GenerateDatabaseName creates a database name following Uberspace convention
func GenerateDatabaseName(username, appName string) string {
	return fmt.Sprintf("%s_%s", username, appName)
}
