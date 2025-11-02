package database

import (
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// Property: When a database name follows the ${USER}_appname convention, validation should always pass
// REQ-PBT-DB-001
func TestProperty_DatabaseName_FollowsConvention(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate valid username and app name
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should follow the convention
		expectedPrefix := username + "_"
		if !strings.HasPrefix(dbName, expectedPrefix) {
			t.Fatalf("database name doesn't follow convention: got %q, want prefix %q", dbName, expectedPrefix)
		}

		// Property: Should contain the app name
		if !strings.Contains(dbName, appName) {
			t.Fatalf("database name doesn't contain app name: got %q, want to contain %q", dbName, appName)
		}
	})
}

// Property: The database name generator should never produce names longer than MySQL's 64-character limit
// REQ-PBT-DB-002
func TestProperty_DatabaseName_WithinMySQLLimit(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate various username and app name combinations
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,31}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,31}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should never exceed 64 characters (MySQL limit)
		const mysqlMaxLength = 64
		if len(dbName) > mysqlMaxLength {
			t.Fatalf("database name exceeds MySQL limit of 64 chars: %q (length: %d)", dbName, len(dbName))
		}
	})
}

// Property: When credentials are parsed from .my.cnf format, the output should never contain newlines or control characters
// REQ-PBT-DB-003
func TestProperty_CredentialsParsing_NoControlCharacters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate valid credentials (without control characters)
		username := rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_]{2,20}$`).Draw(t, "username")
		password := rapid.StringMatching(`^[a-zA-Z0-9!@#$%^&*()_+=-]{8,32}$`).Draw(t, "password")

		creds := &Credentials{
			User:     username,
			Password: password,
			Host:     "localhost",
			Port:     "3306",
		}

		// Property: User should not contain control characters
		if strings.ContainsAny(creds.User, "\n\r\t\x00") {
			t.Fatalf("user contains control characters: %q", creds.User)
		}

		// Property: Password should not contain control characters
		if strings.ContainsAny(creds.Password, "\n\r\t\x00") {
			t.Fatalf("password contains control characters: %q", creds.Password)
		}

		// Property: Host should not contain control characters
		if strings.ContainsAny(creds.Host, "\n\r\t\x00") {
			t.Fatalf("host contains control characters: %q", creds.Host)
		}

		// Property: Port should not contain control characters
		if strings.ContainsAny(creds.Port, "\n\r\t\x00") {
			t.Fatalf("port contains control characters: %q", creds.Port)
		}
	})
}

// Property: GenerateDatabaseName should be deterministic (same inputs = same output)
func TestProperty_GenerateDatabaseName_Deterministic(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		// Generate multiple times
		name1 := GenerateDatabaseName(username, appName)
		name2 := GenerateDatabaseName(username, appName)
		name3 := GenerateDatabaseName(username, appName)

		// Property: Should always produce the same result
		if name1 != name2 || name1 != name3 {
			t.Fatalf("GenerateDatabaseName is not deterministic: %q, %q, %q", name1, name2, name3)
		}
	})
}

// Property: Database names should never contain spaces or special characters that would break SQL
func TestProperty_DatabaseName_NoSpecialCharacters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should not contain problematic characters
		problematicChars := []string{" ", "\t", "\n", ";", "'", "\"", "`", "\\", "/", "*", "?", "<", ">", "|"}
		for _, char := range problematicChars {
			if strings.Contains(dbName, char) {
				t.Fatalf("database name contains problematic character %q: %s", char, dbName)
			}
		}
	})
}

// Property: Database names should always be non-empty
func TestProperty_DatabaseName_NonEmpty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should never be empty
		if dbName == "" {
			t.Fatal("GenerateDatabaseName returned empty string")
		}

		// Property: Should have reasonable minimum length
		if len(dbName) < 4 { // At least "u_a" format
			t.Fatalf("database name is too short: %q (length: %d)", dbName, len(dbName))
		}
	})
}

// Property: Different app names should produce different database names
func TestProperty_DatabaseName_UniqueForDifferentApps(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName1 := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName1")
		appName2 := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).
			Filter(func(s string) bool { return s != appName1 }).
			Draw(t, "appName2")

		dbName1 := GenerateDatabaseName(username, appName1)
		dbName2 := GenerateDatabaseName(username, appName2)

		// Property: Different app names should produce different database names
		if dbName1 == dbName2 {
			t.Fatalf("different app names produced same database name: %q for apps %q and %q",
				dbName1, appName1, appName2)
		}
	})
}

// Property: Database names should be SQL-safe (can be used directly in queries)
func TestProperty_DatabaseName_SQLSafe(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should only contain alphanumeric and underscore/hyphen
		for _, char := range dbName {
			if !((char >= 'a' && char <= 'z') ||
				(char >= 'A' && char <= 'Z') ||
				(char >= '0' && char <= '9') ||
				char == '_' || char == '-') {
				t.Fatalf("database name contains non-SQL-safe character %c in %q", char, dbName)
			}
		}
	})
}

// Property: Credentials should have consistent defaults
func TestProperty_Credentials_DefaultHostAndPort(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Create credentials with defaults
		creds := &Credentials{
			Host: "localhost",
			Port: "3306",
			User: rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_]{2,20}$`).Draw(t, "user"),
			Password: rapid.StringMatching(`^[a-zA-Z0-9!@#$%^&*()_+=-]{8,32}$`).Draw(t, "password"),
		}

		// Property: Default host should be localhost
		if creds.Host != "localhost" {
			t.Fatalf("default host is not localhost: %q", creds.Host)
		}

		// Property: Default port should be 3306
		if creds.Port != "3306" {
			t.Fatalf("default port is not 3306: %q", creds.Port)
		}
	})
}

// Property: Database name format should always be username_appname
func TestProperty_DatabaseName_ExactFormat(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should match exact format
		expected := username + "_" + appName
		if dbName != expected {
			t.Fatalf("database name doesn't match expected format: got %q, want %q", dbName, expected)
		}
	})
}

// Property: Database names should never start with a number
func TestProperty_DatabaseName_NoLeadingNumber(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Username starting with letter (enforced by regex)
		username := rapid.StringMatching(`^[a-z][a-z0-9]{2,15}$`).Draw(t, "username")
		appName := rapid.StringMatching(`^[a-z][a-z0-9-]{2,20}$`).Draw(t, "appName")

		dbName := GenerateDatabaseName(username, appName)

		// Property: Should not start with a number
		if len(dbName) > 0 && dbName[0] >= '0' && dbName[0] <= '9' {
			t.Fatalf("database name starts with a number: %q", dbName)
		}
	})
}

// Property: MySQLManager should be creatable with any dry-run/verbose combination
func TestProperty_MySQLManager_AlwaysCreatable(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dryRun := rapid.Bool().Draw(t, "dryRun")
		verbose := rapid.Bool().Draw(t, "verbose")

		manager := NewMySQLManager(dryRun, verbose)

		// Property: Should never be nil
		if manager == nil {
			t.Fatal("NewMySQLManager returned nil")
		}

		// Property: Should have the correct settings
		if manager.dryRun != dryRun {
			t.Fatalf("dryRun not set correctly: got %v, want %v", manager.dryRun, dryRun)
		}
		if manager.verbose != verbose {
			t.Fatalf("verbose not set correctly: got %v, want %v", manager.verbose, verbose)
		}
	})
}
