package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

// CreateTempDir creates a temporary directory for testing
func CreateTempDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "uberman-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})
	return dir
}

// CreateTestFile creates a test file with the given content
func CreateTestFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	filePath := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		t.Fatalf("failed to create parent dir: %v", err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	return filePath
}

// MinimalManifest returns a minimal valid app manifest in TOML format
func MinimalManifest() string {
	return `[app]
name = "testapp"
type = "php"
version = "1.0"
description = "Test application"

[runtime]
language = "php"
version = "8.3"

[database]
type = "mysql"
required = false

[install]
method = "download"
source = "https://example.com/app.tar.gz"
location = "app/"

[web]
backend = "apache"
document_root = "app/"

[backup]
include = ["app/"]
exclude = ["tmp/"]
`
}

// NodeJSManifest returns a Node.js app manifest in TOML format
func NodeJSManifest() string {
	return `[app]
name = "nodeapp"
type = "nodejs"
version = "1.0"
description = "Node.js application"

[runtime]
language = "node"
version = "20"

[database]
type = "mysql"
required = true
name = "custom_db"

[install]
method = "npm"
source = "express"
location = "app/"

[web]
backend = "http"
port = 8080
document_root = "app/"

[[services]]
name = "nodeapp"
command = "node server.js"
directory = "app/"
port = 8080
startsecs = 10
autorestart = true

[services.environment]
NODE_ENV = "production"
PORT = "8080"

[backup]
include = ["app/", "data/"]
exclude = ["node_modules/", "tmp/"]
`
}

// InvalidManifest returns an invalid manifest (missing required fields)
func InvalidManifest() string {
	return `[app]
name = "testapp"
# Missing type - should fail validation

[runtime]
# Missing language - should fail validation
version = "8.3"
`
}

// MyCnfContent returns a test ~/.my.cnf file content
func MyCnfContent(user, password string) string {
	return `[client]
user = ` + user + `
password = ` + password + `
host = localhost
port = 3306
`
}
