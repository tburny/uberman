package supervisor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// ServiceManager handles supervisord service management
type ServiceManager struct {
	dryRun  bool
	verbose bool
}

// NewServiceManager creates a new service manager
func NewServiceManager(dryRun, verbose bool) *ServiceManager {
	return &ServiceManager{
		dryRun:  dryRun,
		verbose: verbose,
	}
}

// Service represents a supervisord service configuration
type Service struct {
	Name          string
	Command       string
	Directory     string
	Environment   map[string]string
	StartSecs     int
	AutoRestart   bool
	StdoutLogfile string
	StderrLogfile string
	User          string
}

const serviceTemplate = `[program:{{.Name}}]
command={{.Command}}
directory={{.Directory}}
{{- if .Environment}}
environment={{range $key, $value := .Environment}}{{$key}}="{{$value}}",{{end}}
{{- end}}
startsecs={{.StartSecs}}
autorestart={{if .AutoRestart}}true{{else}}false{{end}}
{{- if .StdoutLogfile}}
stdout_logfile={{.StdoutLogfile}}
{{- end}}
{{- if .StderrLogfile}}
stderr_logfile={{.StderrLogfile}}
{{- end}}
`

// CreateService creates a supervisord service configuration file
func (m *ServiceManager) CreateService(service Service) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Default values
	if service.StartSecs == 0 {
		service.StartSecs = 15
	}
	if service.AutoRestart {
		service.AutoRestart = true
	}
	if service.StdoutLogfile == "" {
		service.StdoutLogfile = filepath.Join(homeDir, "logs", service.Name+".log")
	}
	if service.StderrLogfile == "" {
		service.StderrLogfile = filepath.Join(homeDir, "logs", service.Name+"-error.log")
	}

	servicesDir := filepath.Join(homeDir, "etc", "services.d")
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return fmt.Errorf("failed to create services directory: %w", err)
	}

	configPath := filepath.Join(servicesDir, service.Name+".ini")

	if m.verbose {
		fmt.Printf("Creating service configuration: %s\n", configPath)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would create service file: %s\n", configPath)
		fmt.Println("[DRY RUN] Service configuration:")
		tmpl, _ := template.New("service").Parse(serviceTemplate)
		_ = tmpl.Execute(os.Stdout, service)
		return nil
	}

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse service template: %w", err)
	}

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, service); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	if m.verbose {
		fmt.Printf("Service configuration created successfully\n")
	}

	return nil
}

// ReloadServices reloads supervisord configuration
func (m *ServiceManager) ReloadServices() error {
	if m.verbose {
		fmt.Println("Reloading supervisord configuration")
	}

	if m.dryRun {
		fmt.Println("[DRY RUN] Would execute: supervisorctl reread")
		fmt.Println("[DRY RUN] Would execute: supervisorctl update")
		return nil
	}

	// Reread configuration
	cmd := exec.Command("supervisorctl", "reread")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to reread services: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Reread output: %s\n", string(output))
	}

	// Update services
	cmd = exec.Command("supervisorctl", "update")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update services: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Update output: %s\n", string(output))
	}

	return nil
}

// StartService starts a supervisord service
func (m *ServiceManager) StartService(name string) error {
	if m.verbose {
		fmt.Printf("Starting service: %s\n", name)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: supervisorctl start %s\n", name)
		return nil
	}

	cmd := exec.Command("supervisorctl", "start", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start service: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// StopService stops a supervisord service
func (m *ServiceManager) StopService(name string) error {
	if m.verbose {
		fmt.Printf("Stopping service: %s\n", name)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: supervisorctl stop %s\n", name)
		return nil
	}

	cmd := exec.Command("supervisorctl", "stop", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop service: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// RestartService restarts a supervisord service
func (m *ServiceManager) RestartService(name string) error {
	if m.verbose {
		fmt.Printf("Restarting service: %s\n", name)
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: supervisorctl restart %s\n", name)
		return nil
	}

	cmd := exec.Command("supervisorctl", "restart", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to restart service: %w\nOutput: %s", err, string(output))
	}

	if m.verbose {
		fmt.Printf("Output: %s\n", string(output))
	}

	return nil
}

// ServiceStatus gets the status of a service
func (m *ServiceManager) ServiceStatus(name string) (string, error) {
	cmd := exec.Command("supervisorctl", "status", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get service status: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// RemoveService removes a service configuration and stops the service
func (m *ServiceManager) RemoveService(name string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Stop the service first
	if err := m.StopService(name); err != nil {
		// Continue even if stop fails (service might not be running)
		if m.verbose {
			fmt.Printf("Warning: failed to stop service: %v\n", err)
		}
	}

	if m.dryRun {
		fmt.Printf("[DRY RUN] Would execute: supervisorctl remove %s\n", name)
		fmt.Printf("[DRY RUN] Would delete: %s\n", filepath.Join(homeDir, "etc", "services.d", name+".ini"))
		return nil
	}

	// Remove from supervisord
	cmd := exec.Command("supervisorctl", "remove", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if m.verbose {
			fmt.Printf("Warning: failed to remove service from supervisord: %v\nOutput: %s\n", err, string(output))
		}
	}

	// Delete the configuration file
	configPath := filepath.Join(homeDir, "etc", "services.d", name+".ini")
	if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete service file: %w", err)
	}

	if m.verbose {
		fmt.Printf("Service %s removed successfully\n", name)
	}

	return nil
}
