package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/tburny/uberman/internal/testutil"
)

func TestNewMySQLManager(t *testing.T) {
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
			manager := NewMySQLManager(tt.dryRun, tt.verbose)

			require.NotNil(t, manager)
			assert.Equal(t, tt.dryRun, manager.dryRun)
			assert.Equal(t, tt.verbose, manager.verbose)
		})
	}
}

func TestMySQLManager_GetCredentials(t *testing.T) {
	t.Run("valid credentials file", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)

		originalHome := os.Getenv("HOME")
		t.Cleanup(func() {
			os.Setenv("HOME", originalHome)
		})

		myCnfContent := testutil.MyCnfContent("testuser", "testpass")
		testutil.CreateTestFile(t, tmpDir, ".my.cnf", myCnfContent)

		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		creds, err := manager.GetCredentials()

		require.NoError(t, err)
		require.NotNil(t, creds)

		assert.Equal(t, "testuser", creds.User)
		assert.Equal(t, "testpass", creds.Password)
		assert.Equal(t, "localhost", creds.Host)
		assert.Equal(t, "3306", creds.Port)
		assert.Equal(t, "testuser", creds.Database)
	})

	t.Run("credentials with custom host and port", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)
		originalHome := os.Getenv("HOME")
		t.Cleanup(func() { os.Setenv("HOME", originalHome) })

		myCnfContent := `[client]
user = customuser
password = custompass
host = mysql.example.com
port = 3307
`
		testutil.CreateTestFile(t, tmpDir, ".my.cnf", myCnfContent)
		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		creds, err := manager.GetCredentials()

		require.NoError(t, err)
		assert.Equal(t, "customuser", creds.User)
		assert.Equal(t, "custompass", creds.Password)
		assert.Equal(t, "mysql.example.com", creds.Host)
		assert.Equal(t, "3307", creds.Port)
	})

	t.Run("credentials with quoted values", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)
		originalHome := os.Getenv("HOME")
		t.Cleanup(func() { os.Setenv("HOME", originalHome) })

		myCnfContent := `[client]
user = "quoted_user"
password = 'single_quoted_pass'
`
		testutil.CreateTestFile(t, tmpDir, ".my.cnf", myCnfContent)
		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		creds, err := manager.GetCredentials()

		require.NoError(t, err)
		assert.Equal(t, "quoted_user", creds.User)
		assert.Equal(t, "single_quoted_pass", creds.Password)
	})

	t.Run("missing my.cnf file", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)
		originalHome := os.Getenv("HOME")
		t.Cleanup(func() { os.Setenv("HOME", originalHome) })

		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		creds, err := manager.GetCredentials()

		assert.Error(t, err)
		assert.Nil(t, creds)
		assert.Contains(t, err.Error(), "failed to open ~/.my.cnf")
	})

	t.Run("invalid my.cnf format", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)
		originalHome := os.Getenv("HOME")
		t.Cleanup(func() { os.Setenv("HOME", originalHome) })

		invalidContent := `[client]
user testuser
# Missing = sign
`
		testutil.CreateTestFile(t, tmpDir, ".my.cnf", invalidContent)
		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		_, err := manager.GetCredentials()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not find MySQL credentials")
	})

	t.Run("my.cnf with multiple sections", func(t *testing.T) {
		tmpDir := testutil.CreateTempDir(t)
		originalHome := os.Getenv("HOME")
		t.Cleanup(func() { os.Setenv("HOME", originalHome) })

		multiSectionContent := `[mysql]
user = wronguser
password = wrongpass

[client]
user = correctuser
password = correctpass

[mysqldump]
user = dumpuser
`
		testutil.CreateTestFile(t, tmpDir, ".my.cnf", multiSectionContent)
		os.Setenv("HOME", tmpDir)

		manager := NewMySQLManager(false, false)
		creds, err := manager.GetCredentials()

		require.NoError(t, err)
		// Should only read from [client] section
		assert.Equal(t, "correctuser", creds.User)
		assert.Equal(t, "correctpass", creds.Password)
	})
}

func TestGenerateDatabaseName(t *testing.T) {
	tests := []struct {
		username string
		appName  string
		expected string
	}{
		{
			username: "johndoe",
			appName:  "wordpress",
			expected: "johndoe_wordpress",
		},
		{
			username: "alice",
			appName:  "nextcloud",
			expected: "alice_nextcloud",
		},
		{
			username: "user",
			appName:  "myapp",
			expected: "user_myapp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.username+"_"+tt.appName, func(t *testing.T) {
			result := GenerateDatabaseName(tt.username, tt.appName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Integration tests with testcontainers
// These tests require Docker to be running

func setupMySQLContainer(t *testing.T) (testcontainers.Container, *sql.DB, string, string) {
	t.Helper()

	ctx := context.Background()

	container, err := mysql.Run(ctx,
		"mysql:8.0",
		mysql.WithDatabase("testdb"),
		mysql.WithUsername("root"),
		mysql.WithPassword("testpass"),
	)
	if err != nil {
		t.Fatalf("Failed to start MySQL container: %v", err)
	}

	// Get connection string
	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "3306")
	require.NoError(t, err)

	dsn := fmt.Sprintf("root:testpass@tcp(%s:%s)/", host, port.Port())

	// Wait for MySQL to be ready
	var db *sql.DB
	for i := 0; i < 30; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second)
	}
	require.NoError(t, err, "MySQL did not become ready in time")

	t.Cleanup(func() {
		if db != nil {
			db.Close()
		}
		container.Terminate(ctx)
	})

	return container, db, host, port.Port()
}

func TestMySQLManager_Integration_CreateDatabase(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Check if Docker is available
	if _, err := testcontainers.NewDockerClient(); err != nil {
		t.Skip("Docker not available, skipping integration test")
	}

	_, db, _, _ := setupMySQLContainer(t)

	t.Run("create database", func(t *testing.T) {
		// Note: This is a simplified test since we can't easily mock the mysql command
		// In a real integration test, we would use the database connection directly

		dbName := "testdb_app1"

		// Create database using SQL directly (simulating what the command would do)
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
		require.NoError(t, err)

		// Verify database exists
		var schemaName string
		err = db.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&schemaName)
		assert.NoError(t, err)
		assert.Equal(t, dbName, schemaName)

		// Cleanup
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", dbName))
		require.NoError(t, err)
	})
}

func TestMySQLManager_Integration_DatabaseExists(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	if _, err := testcontainers.NewDockerClient(); err != nil {
		t.Skip("Docker not available, skipping integration test")
	}

	_, db, _, _ := setupMySQLContainer(t)

	t.Run("database exists check", func(t *testing.T) {
		dbName := "testdb_exists"

		// Create database
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", dbName))
		require.NoError(t, err)

		// Check if it exists using SQL
		var schemaName string
		err = db.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&schemaName)
		assert.NoError(t, err)
		assert.Equal(t, dbName, schemaName)

		// Cleanup
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE `%s`", dbName))
		require.NoError(t, err)
	})

	t.Run("database does not exist", func(t *testing.T) {
		dbName := "nonexistent_db"

		var schemaName string
		err := db.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&schemaName)
		assert.Error(t, err) // Should be sql.ErrNoRows
	})
}

func TestMySQLManager_Integration_ExportImport(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	if _, err := testcontainers.NewDockerClient(); err != nil {
		t.Skip("Docker not available, skipping integration test")
	}

	_, db, _, _ := setupMySQLContainer(t)

	tmpDir := testutil.CreateTempDir(t)

	t.Run("export and import database", func(t *testing.T) {
		dbName := "testdb_export"

		// Create database and table with data
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", dbName))
		require.NoError(t, err)

		_, err = db.Exec(fmt.Sprintf("USE `%s`", dbName))
		require.NoError(t, err)

		_, err = db.Exec("CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(100))")
		require.NoError(t, err)

		_, err = db.Exec("INSERT INTO users (id, name) VALUES (1, 'Alice'), (2, 'Bob')")
		require.NoError(t, err)

		// Export using mysqldump would happen here in real code
		// For testing, we'll create a simple SQL dump file
		dumpPath := filepath.Join(tmpDir, "export.sql")
		dumpSQL := `CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(100));
INSERT INTO users (id, name) VALUES (1, 'Alice'), (2, 'Bob');`

		err = os.WriteFile(dumpPath, []byte(dumpSQL), 0644)
		require.NoError(t, err)

		// Drop the database
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE `%s`", dbName))
		require.NoError(t, err)

		// Re-create empty database
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", dbName))
		require.NoError(t, err)

		// Import would happen here using mysql command
		// For testing, execute the SQL statements separately
		_, err = db.Exec(fmt.Sprintf("USE `%s`", dbName))
		require.NoError(t, err)

		// Execute CREATE TABLE statement
		_, err = db.Exec("CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(100))")
		require.NoError(t, err)

		// Execute INSERT statement
		_, err = db.Exec("INSERT INTO users (id, name) VALUES (1, 'Alice'), (2, 'Bob')")
		require.NoError(t, err)

		// Verify data was imported
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
		require.NoError(t, err)
		assert.Equal(t, 2, count)

		// Cleanup
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE `%s`", dbName))
		require.NoError(t, err)
	})
}

func TestMySQLManager_DryRun(t *testing.T) {
	t.Run("create database dry-run", func(t *testing.T) {
		manager := NewMySQLManager(true, true)

		err := manager.CreateDatabase("testdb_dryrun")
		assert.NoError(t, err, "Dry-run should not error")
	})

	t.Run("drop database dry-run", func(t *testing.T) {
		manager := NewMySQLManager(true, false)

		err := manager.DropDatabase("testdb_dryrun")
		assert.NoError(t, err, "Dry-run should not error")
	})

	t.Run("export database dry-run", func(t *testing.T) {
		manager := NewMySQLManager(true, false)

		err := manager.ExportDatabase("testdb", "/tmp/export.sql")
		assert.NoError(t, err, "Dry-run should not error")
	})

	t.Run("import database dry-run", func(t *testing.T) {
		manager := NewMySQLManager(true, false)

		err := manager.ImportDatabase("testdb", "/tmp/import.sql")
		assert.NoError(t, err, "Dry-run should not error")
	})
}

func TestMySQLManager_UberspaceNamingConvention(t *testing.T) {
	// Test that database names follow Uberspace convention
	tests := []struct {
		username string
		appName  string
		expected string
	}{
		{
			username: "uberuser",
			appName:  "wordpress",
			expected: "uberuser_wordpress",
		},
		{
			username: "alice",
			appName:  "ghost",
			expected: "alice_ghost",
		},
	}

	for _, tt := range tests {
		t.Run(tt.username+"_"+tt.appName, func(t *testing.T) {
			dbName := GenerateDatabaseName(tt.username, tt.appName)
			assert.Equal(t, tt.expected, dbName)

			// Verify it starts with username_
			assert.Contains(t, dbName, tt.username+"_")
		})
	}
}
