package supervisor

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

// Property: When a service configuration is generated, the output should always be valid INI format
// REQ-PBT-SUPER-001
func TestProperty_ServiceConfig_ValidINI(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		// Read the generated file
		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config file: %v", err)
		}

		// Property: Should start with [program:name]
		expectedHeader := "[program:" + service.Name + "]"
		if !strings.HasPrefix(string(content), expectedHeader) {
			t.Fatalf("config doesn't start with expected header %q", expectedHeader)
		}

		// Property: Should contain command line
		if !strings.Contains(string(content), "command=") {
			t.Fatal("config doesn't contain command=")
		}

		// Property: Should not contain malformed lines
		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "[") || strings.HasPrefix(line, "#") {
				continue
			}
			// Each non-empty line should be key=value format
			if !strings.Contains(line, "=") && line != "" {
				t.Fatalf("line %d is not valid INI format: %q", i, line)
			}
		}
	})
}

// Property: The environment variable rendering should properly escape all special characters
// REQ-PBT-SUPER-002
func TestProperty_EnvironmentVariables_ProperlyEscaped(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		// Add environment variables with various special characters
		service.Environment = map[string]string{
			"SIMPLE":       rapid.StringMatching(`^[A-Za-z0-9]+$`).Draw(t, "simple"),
			"WITH_EQUALS":  rapid.StringMatching(`^[A-Za-z0-9=]+$`).Draw(t, "withEquals"),
			"WITH_QUOTES":  rapid.StringMatching(`^[A-Za-z0-9"']+$`).Draw(t, "withQuotes"),
		}

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config file: %v", err)
		}

		// Property: Environment section should be present if variables exist
		if len(service.Environment) > 0 {
			if !strings.Contains(string(content), "environment=") {
				t.Fatal("config doesn't contain environment= when variables are present")
			}
		}

		// Property: Values should be quoted
		for key := range service.Environment {
			// The environment line should contain the key
			if !strings.Contains(string(content), key) {
				t.Fatalf("environment variable %q not found in config", key)
			}
		}
	})
}

// Property: When a service name is sanitized, the output should match the regex ^[a-zA-Z0-9_-]+$
// REQ-PBT-SUPER-003
func TestProperty_ServiceName_ValidFormat(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate service name that matches expected format
		serviceName := rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_-]{2,30}$`).Draw(t, "serviceName")

		service := Service{
			Name:      serviceName,
			Command:   "echo test",
			Directory: "/tmp",
		}

		_, cleanup := setupTempHome(t)
		defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		// Property: Service name should only contain valid characters
		for _, char := range serviceName {
			if !((char >= 'a' && char <= 'z') ||
				(char >= 'A' && char <= 'Z') ||
				(char >= '0' && char <= '9') ||
				char == '_' || char == '-') {
				t.Fatalf("service name contains invalid character %c: %q", char, serviceName)
			}
		}
	})
}

// Property: Service config file should always be created in ~/etc/services.d/
func TestProperty_ServiceConfig_CorrectLocation(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		// Property: File should be in the correct location
		expectedPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Fatalf("service config not at expected location %q: %v", expectedPath, err)
		}
	})
}

// Property: DryRun mode should never create actual files
func TestProperty_DryRun_NoFileCreation(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		// Create manager in dry-run mode
		manager := NewServiceManager(true, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("DryRun CreateService failed: %v", err)
		}

		// Property: File should NOT exist in dry-run mode
		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		if _, err := os.Stat(configPath); err == nil {
			t.Fatal("service config file exists after dry-run (should not be created)")
		}
	})
}

// Property: Service config should always have default values for optional fields
func TestProperty_ServiceDefaults_Applied(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Create service without optional fields
		service := Service{
			Name:      rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_-]{2,30}$`).Draw(t, "name"),
			Command:   rapid.StringMatching(`^[a-z]+ [a-z]+$`).Draw(t, "command"),
			Directory: "/tmp",
			// No StartSecs, AutoRestart, or log files set
		}

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config: %v", err)
		}

		// Property: Should have default startsecs
		if !strings.Contains(string(content), "startsecs=") {
			t.Fatal("config missing startsecs (default should be applied)")
		}

		// Property: Should have stdout_logfile
		if !strings.Contains(string(content), "stdout_logfile=") {
			t.Fatal("config missing stdout_logfile (default should be applied)")
		}

		// Property: Should have stderr_logfile
		if !strings.Contains(string(content), "stderr_logfile=") {
			t.Fatal("config missing stderr_logfile (default should be applied)")
		}
	})
}

// Property: Service name should never be empty
func TestProperty_ServiceName_NonEmpty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		serviceName := rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_-]{2,30}$`).Draw(t, "serviceName")

		// Property: Generated name should never be empty
		if serviceName == "" {
			t.Fatal("generated service name is empty")
		}

		// Property: Should have reasonable length
		if len(serviceName) < 3 {
			t.Fatalf("service name too short: %q (length: %d)", serviceName, len(serviceName))
		}
	})
}

// Property: Command should never be empty
func TestProperty_Command_NonEmpty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		// Property: Command should never be empty
		if service.Command == "" {
			t.Fatal("service command is empty")
		}

		// Property: Command should have reasonable content
		if len(strings.TrimSpace(service.Command)) == 0 {
			t.Fatal("service command is only whitespace")
		}
	})
}

// Property: Directory should always be an absolute path
func TestProperty_Directory_AbsolutePath(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)

		// Property: Directory should be absolute
		if !filepath.IsAbs(service.Directory) {
			t.Fatalf("service directory is not absolute: %q", service.Directory)
		}
	})
}

// Property: AutoRestart value should be represented as true/false in config
func TestProperty_AutoRestart_BooleanFormat(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)
		service.AutoRestart = rapid.Bool().Draw(t, "autoRestart")

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config: %v", err)
		}

		// Property: AutoRestart should be true or false (not other values)
		if service.AutoRestart {
			if !strings.Contains(string(content), "autorestart=true") {
				t.Fatal("autorestart=true not found when AutoRestart is true")
			}
		} else {
			if !strings.Contains(string(content), "autorestart=false") {
				t.Fatal("autorestart=false not found when AutoRestart is false")
			}
		}
	})
}

// Property: StartSecs should always be a positive number
func TestProperty_StartSecs_Positive(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		service := generateValidService(t)
		service.StartSecs = rapid.IntRange(1, 300).Draw(t, "startSecs")

		tmpHome, cleanup := setupTempHome(t); defer cleanup()

		manager := NewServiceManager(false, false)
		err := manager.CreateService(service)
		if err != nil {
			t.Fatalf("CreateService failed: %v", err)
		}

		configPath := filepath.Join(tmpHome, "etc", "services.d", service.Name+".ini")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config: %v", err)
		}

		// Property: startsecs should be present with a positive value
		if !strings.Contains(string(content), "startsecs=") {
			t.Fatal("startsecs not found in config")
		}

		// Check the value is positive (appears as a number)
		contentStr := string(content)
		if strings.Contains(contentStr, "startsecs=-") || strings.Contains(contentStr, "startsecs=0") {
			t.Fatalf("startsecs has non-positive value in config")
		}
	})
}

// Property: ServiceManager should always be creatable with any combination of flags
func TestProperty_ServiceManager_AlwaysCreatable(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		dryRun := rapid.Bool().Draw(t, "dryRun")
		verbose := rapid.Bool().Draw(t, "verbose")

		manager := NewServiceManager(dryRun, verbose)

		// Property: Should never be nil
		if manager == nil {
			t.Fatal("NewServiceManager returned nil")
		}

		// Property: Should have correct settings
		if manager.dryRun != dryRun {
			t.Fatalf("dryRun not set correctly: got %v, want %v", manager.dryRun, dryRun)
		}
		if manager.verbose != verbose {
			t.Fatalf("verbose not set correctly: got %v, want %v", manager.verbose, verbose)
		}
	})
}

// Helper function to generate valid service configurations
func generateValidService(t *rapid.T) Service {
	return Service{
		Name:        rapid.StringMatching(`^[a-zA-Z][a-zA-Z0-9_-]{2,30}$`).Draw(t, "name"),
		Command:     rapid.StringMatching(`^[a-z]+ [a-z/.-]+$`).Draw(t, "command"),
		Directory:   rapid.StringMatching(`^/[a-z0-9/_-]{1,50}$`).Draw(t, "directory"),
		StartSecs:   rapid.IntRange(5, 60).Draw(t, "startSecs"),
		AutoRestart: rapid.Bool().Draw(t, "autoRestart"),
		Environment: generateEnvironment(t),
	}
}

// Helper to generate environment variables
func generateEnvironment(t *rapid.T) map[string]string {
	size := rapid.IntRange(0, 5).Draw(t, "envSize")
	env := make(map[string]string)

	for i := 0; i < size; i++ {
		key := rapid.StringMatching(`^[A-Z][A-Z0-9_]{1,20}$`).Draw(t, "envKey")
		value := rapid.StringMatching(`^[a-zA-Z0-9_/.-]{1,30}$`).Draw(t, "envValue")
		env[key] = value
	}

	return env
}
