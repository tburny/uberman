# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)
and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

## [1.2.1](https://github.com/tburny/uberman/compare/v1.2.0...v1.2.1) (2025-11-02)

### Bug Fixes

* **ci:** capture semantic-release outputs for build job ([4ceca7a](https://github.com/tburny/uberman/commit/4ceca7ad9a2fee437c6a240df66b75d81bf00540))

## [1.2.0](https://github.com/tburny/uberman/compare/v1.1.0...v1.2.0) (2025-11-02)

### Features

* **apps:** restructure to subdirectories with lifecycle hooks ([693ddf3](https://github.com/tburny/uberman/commit/693ddf31abf6649fcd796c2d6779011a42dc2441))

## [1.1.0](https://github.com/tburny/uberman/compare/v1.0.0...v1.1.0) (2025-11-02)

### Features

* **install:** add one-line installation script ([861cfe7](https://github.com/tburny/uberman/commit/861cfe7154927e36daa7176ef1dd65a052bf642d))

### Documentation

* add comprehensive CI/CD setup documentation ([0896f6d](https://github.com/tburny/uberman/commit/0896f6dcc6df4166703753f64716c4264e73557c))

## 1.0.0 (2025-11-02)

### Features

* **ci:** add semantic-release with automated binary builds ([e03c616](https://github.com/tburny/uberman/commit/e03c61626ce29701a4cfbca6a0995506ab625714))
* initial project foundation for uberman ([e765f51](https://github.com/tburny/uberman/commit/e765f51fc6018ccf7b2d724bb63ed6be540e8ec6))

### Documentation

* add setup verification guide ([ead6f5e](https://github.com/tburny/uberman/commit/ead6f5e07c25e344774b2c458724fb99e7092d56))

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
