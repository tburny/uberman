# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## [Unreleased]

### Added
- Initial project structure with Go modules
- CLI framework with cobra
- TOML configuration parser for app manifests
- Runtime version management (Uberspace integration)
- MySQL database operations (create, export, import)
- Web backend configuration for Apache and HTTP
- Supervisord service management
- Self-contained app directory structure manager
- Example app manifests (WordPress, Ghost, Nextcloud)
- Comprehensive documentation (CLAUDE.md, README.md, CONTRIBUTING.md)
- Quick start guide
- Development roadmap
- Makefile for build automation
- Conventional Commits configuration and guidelines

### Changed
- N/A

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

### Security
- N/A

## [0.1.0] - 2025-11-02

### Added
- Initial foundation release
- Core internal packages (config, runtime, database, web, supervisor, appdir)
- Basic install command structure
- Dry-run and verbose modes
- App manifest schema and validation
- Example manifests for WordPress, Ghost, and Nextcloud

[Unreleased]: https://github.com/tburny/uberman/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/tburny/uberman/releases/tag/v0.1.0
