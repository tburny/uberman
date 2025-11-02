package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBackendManager(t *testing.T) {
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
			manager := NewBackendManager(tt.dryRun, tt.verbose)

			require.NotNil(t, manager)
			assert.Equal(t, tt.dryRun, manager.dryRun)
			assert.Equal(t, tt.verbose, manager.verbose)
		})
	}
}

func TestBackendManager_SetBackend_DryRun(t *testing.T) {
	tests := []struct {
		name    string
		backend Backend
		wantErr bool
		errMsg  string
	}{
		{
			name: "apache backend for root path",
			backend: Backend{
				Path: "/",
				Type: "apache",
			},
			wantErr: false,
		},
		{
			name: "http backend with port",
			backend: Backend{
				Path: "/",
				Type: "http",
				Port: 8080,
			},
			wantErr: false,
		},
		{
			name: "apache backend for subpath",
			backend: Backend{
				Path: "/blog",
				Type: "apache",
			},
			wantErr: false,
		},
		{
			name: "http backend with domain",
			backend: Backend{
				Path:   "/",
				Type:   "http",
				Port:   3000,
				Domain: "example.com",
			},
			wantErr: false,
		},
		{
			name: "static path backend",
			backend: Backend{
				Path: "/static",
				Type: "apache",
			},
			wantErr: false,
		},
		{
			name: "invalid backend type",
			backend: Backend{
				Path: "/",
				Type: "invalid",
			},
			wantErr: true,
			errMsg:  "invalid backend type",
		},
		{
			name: "http backend without port should error in real execution",
			backend: Backend{
				Path: "/",
				Type: "http",
				Port: 0,
			},
			wantErr: false, // In dry-run, this won't error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewBackendManager(true, true)

			err := manager.SetBackend(tt.backend)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBackendManager_DeleteBackend_DryRun(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		domain string
	}{
		{
			name:   "delete root backend",
			path:   "/",
			domain: "",
		},
		{
			name:   "delete subpath backend",
			path:   "/api",
			domain: "",
		},
		{
			name:   "delete backend for specific domain",
			path:   "/",
			domain: "example.com",
		},
		{
			name:   "delete static path",
			path:   "/static",
			domain: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewBackendManager(true, false)

			err := manager.DeleteBackend(tt.path, tt.domain)
			assert.NoError(t, err, "Dry-run should not error")
		})
	}
}

func TestBackendManager_BackendTypes(t *testing.T) {
	manager := NewBackendManager(true, false)

	t.Run("apache backend", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "apache",
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("http backend", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "http",
			Port: 8080,
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})
}

func TestBackendManager_FindAvailablePort(t *testing.T) {
	t.Run("find port in valid range", func(t *testing.T) {
		manager := NewBackendManager(true, false)

		// This test would require mocking ListBackends
		// In dry-run mode, we can test the edge cases

		port, err := manager.FindAvailablePort(8000, 9000)
		if err == nil {
			// If no error, port should be in range
			assert.GreaterOrEqual(t, port, 8000)
			assert.LessOrEqual(t, port, 9000)
		}
	})

	t.Run("adjust range below minimum", func(t *testing.T) {
		manager := NewBackendManager(true, false)

		port, err := manager.FindAvailablePort(500, 2000)
		if err == nil {
			// Should adjust to start at 1024
			assert.GreaterOrEqual(t, port, 1024)
		}
	})

	t.Run("adjust range above maximum", func(t *testing.T) {
		manager := NewBackendManager(true, false)

		port, err := manager.FindAvailablePort(60000, 70000)
		if err == nil {
			// Should adjust to end at 65535
			assert.LessOrEqual(t, port, 65535)
		}
	})
}

func TestBackend_Validation(t *testing.T) {
	tests := []struct {
		name    string
		backend Backend
		isValid bool
	}{
		{
			name: "valid apache backend",
			backend: Backend{
				Path: "/",
				Type: "apache",
			},
			isValid: true,
		},
		{
			name: "valid http backend",
			backend: Backend{
				Path: "/",
				Type: "http",
				Port: 8080,
			},
			isValid: true,
		},
		{
			name: "invalid type",
			backend: Backend{
				Path: "/",
				Type: "nginx",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewBackendManager(true, false)
			err := manager.SetBackend(tt.backend)

			if tt.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBackendManager_CommonUseCases(t *testing.T) {
	manager := NewBackendManager(true, true)

	t.Run("PHP app with Apache", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "apache",
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("Node.js app with HTTP backend", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "http",
			Port: 8080,
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("Python Flask app", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "http",
			Port: 5000,
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("Static files via Apache", func(t *testing.T) {
		backend := Backend{
			Path: "/static",
			Type: "apache",
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("Multiple backends - app on HTTP, static on Apache", func(t *testing.T) {
		// Main app backend
		appBackend := Backend{
			Path: "/",
			Type: "http",
			Port: 3000,
		}
		err := manager.SetBackend(appBackend)
		assert.NoError(t, err)

		// Static files backend
		staticBackend := Backend{
			Path: "/static",
			Type: "apache",
		}
		err = manager.SetBackend(staticBackend)
		assert.NoError(t, err)
	})
}

func TestBackendManager_DomainSpecific(t *testing.T) {
	manager := NewBackendManager(true, false)

	t.Run("backend for specific domain", func(t *testing.T) {
		backend := Backend{
			Path:   "/",
			Type:   "http",
			Port:   8080,
			Domain: "blog.example.com",
		}

		err := manager.SetBackend(backend)
		assert.NoError(t, err)
	})

	t.Run("delete domain-specific backend", func(t *testing.T) {
		err := manager.DeleteBackend("/", "blog.example.com")
		assert.NoError(t, err)
	})
}

func TestBackendManager_PortRanges(t *testing.T) {
	// Test various port configurations
	manager := NewBackendManager(true, false)

	validPorts := []int{1024, 3000, 5000, 8000, 8080, 9000, 65535}

	for _, port := range validPorts {
		t.Run("port_"+string(rune(port)), func(t *testing.T) {
			backend := Backend{
				Path: "/",
				Type: "http",
				Port: port,
			}

			err := manager.SetBackend(backend)
			assert.NoError(t, err)
		})
	}
}

func TestBackendManager_ErrorMessages(t *testing.T) {
	manager := NewBackendManager(true, false)

	t.Run("invalid backend type error message", func(t *testing.T) {
		backend := Backend{
			Path: "/",
			Type: "invalid",
		}

		err := manager.SetBackend(backend)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid backend type")
		assert.Contains(t, err.Error(), "must be 'apache' or 'http'")
	})
}

// Integration tests that would run on actual Uberspace
func TestBackendManager_Integration_Uberspace(t *testing.T) {
	t.Skip("Integration test - only run on actual Uberspace host")

	/*
		// This test should only run on actual Uberspace
		manager := NewBackendManager(false, true)

		// Set a backend
		backend := Backend{
			Path: "/test",
			Type: "apache",
		}
		err := manager.SetBackend(backend)
		require.NoError(t, err)

		// List backends to verify
		backends, err := manager.ListBackends()
		require.NoError(t, err)
		assert.NotEmpty(t, backends)

		// Clean up
		err = manager.DeleteBackend("/test", "")
		require.NoError(t, err)
	*/
}
