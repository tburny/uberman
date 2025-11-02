# Uberman Development Roadmap

This document outlines the current implementation status and planned features for Uberman.

## Current Status: v0.1.0 (Foundation)

### ✅ Completed

#### Core Infrastructure
- [x] Go project structure with modules
- [x] CLI framework with cobra
- [x] TOML configuration parser
- [x] Dry-run and verbose modes
- [x] Comprehensive documentation (CLAUDE.md, README.md)

#### Internal Packages
- [x] `config` - App manifest parser and validation
- [x] `runtime` - Uberspace runtime version management
- [x] `database` - MySQL database operations
- [x] `web` - Web backend configuration
- [x] `supervisor` - Supervisord service management
- [x] `appdir` - Directory structure management

#### Example Manifests
- [x] WordPress (PHP)
- [x] Ghost (Node.js)
- [x] Nextcloud (PHP)

#### Documentation
- [x] README with quick start
- [x] CONTRIBUTING guidelines
- [x] QUICKSTART guide
- [x] Example configuration file

### 🚧 In Progress

#### Commands (Partial Implementation)
- [x] `install` - Basic structure (needs completion)
- [ ] `upgrade` - Not implemented
- [ ] `backup` - Not implemented
- [ ] `restore` - Not implemented
- [ ] `status` - Not implemented
- [ ] `deploy` - Not implemented
- [ ] `list` - Not implemented
- [ ] `init` - Not implemented

#### Missing Components
- [ ] Configuration template engine
- [ ] Backup/restore logic
- [ ] Deployment strategies
- [ ] Cron job management
- [ ] PHP configuration generation

---

## Phase 1: Core Functionality (v0.2.0)

**Goal**: Complete basic install/upgrade/backup workflow

### Priority 1: Installation Pipeline
- [ ] Download strategies (curl, git, wget)
- [ ] Language-specific dependency installation
  - [ ] PHP: composer, CLI tools (wp-cli, occ)
  - [ ] Python: pip with virtualenv
  - [ ] Node.js: npm/yarn
- [ ] Configuration file generation from templates
- [ ] Web backend setup implementation
- [ ] Supervisord service creation (for non-PHP apps)
- [ ] Cron job installation
- [ ] Post-install verification

### Priority 2: Backup & Restore
- [ ] Implement backup package
  - [ ] File system backup with compression
  - [ ] Database dump integration
  - [ ] Versioning and retention policies
  - [ ] Metadata storage (manifest, timestamp)
- [ ] Implement restore functionality
  - [ ] File restoration
  - [ ] Database restoration
  - [ ] Service restart after restore
  - [ ] Validation after restore

### Priority 3: Status & Health Checks
- [ ] Service status checking
- [ ] Port connectivity tests
- [ ] Database connectivity tests
- [ ] Web backend verification
- [ ] Log file analysis
- [ ] Health score calculation

### Priority 4: Upgrade Workflow
- [ ] Pre-upgrade backup
- [ ] Version detection
- [ ] Safe upgrade execution
- [ ] Rollback on failure
- [ ] Post-upgrade verification

---

## Phase 2: User Experience (v0.3.0)

**Goal**: Improve usability and extensibility

### Enhanced Commands
- [ ] `init` - Scaffold new app manifests
- [ ] `list` - Show installed apps with status
- [ ] `logs` - Tail app logs
- [ ] `shell` - Open shell in app directory
- [ ] `config` - View/edit app configuration

### Template System
- [ ] Template engine for config files
- [ ] Variable substitution (database, domains, secrets)
- [ ] Pre-defined templates for common apps
- [ ] User-defined templates
- [ ] Template validation

### Interactive Mode
- [ ] Prompts for missing configuration
- [ ] Confirmation for destructive operations
- [ ] Progress bars for long operations
- [ ] Colored output for better readability
- [ ] Interactive app selection

### Validation & Error Handling
- [ ] Pre-flight checks (disk space, ports, dependencies)
- [ ] Better error messages with suggestions
- [ ] Manifest schema validation
- [ ] Dependency checking
- [ ] Conflict detection (port, database names)

---

## Phase 3: Advanced Features (v0.4.0)

**Goal**: Production-ready features and automation

### Multi-App Management
- [ ] Install multiple apps in one command
- [ ] App dependency resolution (app A requires app B)
- [ ] Batch operations (backup all, upgrade all)
- [ ] App tags and filtering

### Deployment Features
- [ ] Git-based deployment
- [ ] Blue-green deployments
- [ ] Deployment hooks (pre/post deploy)
- [ ] Environment-specific configurations
- [ ] Secrets management

### Monitoring & Alerts
- [ ] Health monitoring
- [ ] Log monitoring and alerts
- [ ] Resource usage tracking
- [ ] Email notifications
- [ ] Webhook notifications

### Security
- [ ] Secrets encryption
- [ ] SSL/TLS certificate monitoring
- [ ] Security updates detection
- [ ] Vulnerability scanning
- [ ] Access control (multi-user support)

---

## Phase 4: Ecosystem (v0.5.0)

**Goal**: Community and integration

### App Registry
- [ ] Central registry of app manifests
- [ ] Search and discover apps
- [ ] Community contributions
- [ ] Rating and reviews
- [ ] Version compatibility tracking

### Integrations
- [ ] Ansible playbooks
- [ ] Terraform provider
- [ ] CI/CD integration (GitHub Actions, GitLab CI)
- [ ] Backup to cloud storage (S3, Backblaze)
- [ ] Monitoring integration (Prometheus, Grafana)

### Migration Tools
- [ ] Import existing manual installations
- [ ] Migrate from other hosting platforms
- [ ] Export configurations for documentation
- [ ] Disaster recovery tools

### Web Interface (Optional)
- [ ] Web UI for app management
- [ ] Dashboard with app statuses
- [ ] Log viewer
- [ ] Backup management
- [ ] User management

---

## Testing Strategy

### Unit Tests
- [ ] Config package tests
- [ ] Runtime package tests
- [ ] Database package tests
- [ ] Web backend package tests
- [ ] Supervisor package tests
- [ ] Template engine tests

### Integration Tests
- [ ] Full install workflow tests
- [ ] Upgrade tests
- [ ] Backup/restore tests
- [ ] Multi-app scenarios
- [ ] Error recovery tests

### End-to-End Tests
- [ ] Test on real Uberspace instances
- [ ] Test all example manifests
- [ ] Performance testing
- [ ] Stress testing

---

## Documentation Improvements

### User Documentation
- [ ] Video tutorials
- [ ] Interactive getting started guide
- [ ] Common recipes and patterns
- [ ] Troubleshooting guide
- [ ] FAQ

### Developer Documentation
- [ ] API documentation (godoc)
- [ ] Architecture diagrams
- [ ] Contributing guide enhancements
- [ ] Code examples for extensions
- [ ] Manifest schema reference

### App-Specific Docs
- [ ] WordPress deployment guide
- [ ] Ghost blog setup guide
- [ ] Nextcloud configuration guide
- [ ] Django/Flask best practices
- [ ] Custom app development guide

---

## Community & Growth

### Community Building
- [ ] GitHub Discussions setup
- [ ] Discord/Slack community
- [ ] Regular contributor calls
- [ ] Showcase successful deployments
- [ ] User testimonials

### Outreach
- [ ] Blog posts about Uberman
- [ ] Presentations at meetups
- [ ] Integration with Uberspace Lab
- [ ] Collaboration with Uberspace team
- [ ] Social media presence

---

## Technical Debt & Refactoring

### Code Quality
- [ ] Increase test coverage to >80%
- [ ] Add linting configuration
- [ ] Set up CI/CD pipeline
- [ ] Add pre-commit hooks
- [ ] Code documentation improvements

### Performance
- [ ] Optimize file operations
- [ ] Parallel execution where possible
- [ ] Caching for repeated operations
- [ ] Memory usage optimization

### Maintenance
- [ ] Dependency updates strategy
- [ ] Security audit process
- [ ] Regular release schedule
- [ ] Deprecation policy
- [ ] Backwards compatibility strategy

---

## Open Questions

Questions to be resolved during development:

1. **PostgreSQL Support**: Should we support PostgreSQL in addition to MySQL?
2. **Docker Integration**: Any way to leverage containers despite Uberspace limitations?
3. **Multi-Domain Apps**: How to handle apps that need multiple domains?
4. **Resource Limits**: Should we track/enforce resource usage?
5. **Backup Storage**: Default to local only or integrate cloud storage?
6. **Update Frequency**: How often should we check for app updates?
7. **Security Scanning**: Should we include vulnerability scanning?
8. **App Isolation**: Do we need better isolation between apps?

---

## Contributing

Want to help? Pick an item from the roadmap and:

1. Comment on the related issue (or create one)
2. Discuss approach in the issue
3. Submit a PR referencing the issue
4. Update this roadmap when complete

See [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

---

## Version History

- **v0.1.0** (Current) - Foundation and core packages
- **v0.2.0** (Planned) - Complete install/upgrade/backup workflow
- **v0.3.0** (Planned) - Enhanced user experience
- **v0.4.0** (Planned) - Production features
- **v0.5.0** (Planned) - Ecosystem and integrations

---

*Last updated: 2025-11-02*
