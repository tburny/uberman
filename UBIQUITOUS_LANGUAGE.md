# Ubiquitous Language
# Bounded Context: App Installation

All terms used consistently across code, documentation, and conversation within the "App Installation" bounded context.

## Bounded Context Name

**Context**: "App Installation" (noun + verb pattern per DDD)
**Go Directory**: `appinstallation` (lowercase, no separators per Go conventions)

## Core Workflow Concepts

### Installation
**Definition**: The workflow/process of setting up an app on Uberspace hosting.

**States**: NotStarted → ManifestLoaded → PrerequisitesValidated → ProvisioningInProgress → Provisioned → Configured → Installed

**Responsibilities**:
- Coordinate all provisioning activities
- Maintain workflow state
- Validate prerequisites
- Emit domain events

**Code**: `internal/appinstallation/domain/workflow/installation.go` (aggregate root)

**Examples**:
- "The installation transitions to ManifestLoaded state"
- "Installation failed during database provisioning"
- "Check installation status before proceeding"

### App
**Definition**: Installable software package (WordPress, Ghost, Nextcloud, etc.)

**Properties**:
- Has a manifest defining requirements
- Becomes an Instance once installed
- Identified by AppName

**Code**: Referenced via AppName value object

**Examples**:
- "Install the wordpress app"
- "The app requires a MySQL database"
- "App manifest not found"

### Manifest
**Definition**: TOML file defining app requirements and configuration.

**Properties**:
- Immutable once loaded
- Contains: runtime, database, web, services, backup config
- Located at apps/<appname>/app.toml

**Code**: `internal/appinstallation/domain/manifest/manifest.go` (aggregate)

**Examples**:
- "Load manifest from disk"
- "Manifest specifies PHP 8.3"
- "Validate manifest structure"

### Provisioning
**Definition**: The act of creating infrastructure for an app.

**Includes**:
- Directories (app/, data/, logs/, backups/, tmp/)
- Database (MySQL with ${USER}_<appname>)
- Web backend (Apache or HTTP proxy)
- Service (supervisord configuration)
- Runtime (language version)

**Code**: Commands on Installation aggregate

**Examples**:
- "Provision database for the app"
- "Provisioning failed: port unavailable"
- "All provisioning steps completed"

### Instance
**Definition**: A specific installed app at ~/apps/<appname>/.

**Properties**:
- Has physical directory structure
- Has associated database (if required)
- Has web backend configuration
- Has supervisord service (if required)
- Running or stopped state

**Examples**:
- "The wordpress instance is at ~/apps/wordpress"
- "Instance already exists at this location"
- "Remove the instance directory"

## Workflow Terms

### State
**Definition**: Current phase of the installation workflow.

**Properties**:
- Immutable - creates new Installation on transition
- One of: NotStarted, ManifestLoaded, PrerequisitesValidated, ProvisioningInProgress, Provisioned, Configured, Installed, Failed

**Code**: `internal/appinstallation/domain/workflow/states.go`

**Examples**:
- "Installation is in ManifestLoaded state"
- "Transition to PrerequisitesValidated state"
- "Cannot provision from NotStarted state"

### Command
**Definition**: Action that drives state transition in Installation aggregate.

**Examples**:
- LoadManifest(appName)
- ValidatePrerequisites()
- ProvisionDirectories()
- ProvisionDatabase()
- ProvisionWebBackend()
- ProvisionService()
- SetRuntimeVersion()
- FinalizeInstallation()

**Code**: Methods on Installation aggregate

**Usage**:
- "Execute LoadManifest command"
- "ProvisionDatabase command failed"
- "Commands validate preconditions before executing"

### Event
**Definition**: Outcome of successful command execution (domain event).

**Properties**:
- Past tense naming
- Emitted after state change
- Immutable
- Can be used for logging, monitoring, event sourcing

**Examples**:
- ManifestLoaded
- PrerequisitesValidated
- DirectoriesProvisioned
- DatabaseProvisioned
- WebBackendProvisioned
- ServiceProvisioned
- RuntimeVersionSet
- InstallationCompleted
- InstallationFailed(error)

**Code**: `internal/appinstallation/domain/workflow/events.go`

**Usage**:
- "Emit ManifestLoaded event"
- "Listen for InstallationCompleted"
- "Log all domain events"

## Value Objects

### AppName
**Definition**: Validated app name following Uberspace conventions.

**Rules**:
- Lowercase only
- 3-21 characters
- Pattern: ^[a-z][a-z0-9-]{2,20}$
- Must start with letter

**Examples**: "wordpress", "ghost", "my-app-123"
**Invalid**: "WordPress", "ab", "123app", "app_name"

### Port
**Definition**: Valid port number for HTTP backends.

**Rules**:
- Range: 1024-65535 (no privileged ports)
- Must be available (not in use)

**Examples**: 8080, 3000, 5432
**Invalid**: 80 (privileged), 99999 (out of range)

### DatabaseName
**Definition**: MySQL database name following Uberspace convention.

**Rules**:
- Pattern: ${USER}_<appname>
- Example: If USER=johndoe and app=wordpress, then database=johndoe_wordpress

**Code**: Constructed from username + AppName

### DirectoryPath
**Definition**: Absolute directory path.

**Rules**:
- Must be absolute (starts with /)
- Validated for safety (no ../ traversal)

**Examples**: "/home/user/apps/wordpress", "/home/user/apps/ghost/data"
**Invalid**: "apps/wordpress" (relative), "../etc/passwd" (traversal)

## Infrastructure Terms

### Backend (Web)
**Definition**: Uberspace web server configuration.

**Types**:
- **Apache**: Directly serves PHP applications
- **HTTP**: Reverse proxy to application server on specific port

**Configuration**:
- Path: URL path (e.g., "/", "/blog")
- Port: Required for HTTP backend
- Document root: Optional for Apache

**Examples**:
- "Configure Apache backend for WordPress"
- "Set HTTP backend on port 8080"
- "Backend type determines supervisord service need"

### Service
**Definition**: Supervisord process definition for running long-lived application processes.

**Properties**:
- Config location: ~/etc/services.d/<appname>.ini
- INI format with [program:name] section
- Auto-restart, logging, environment variables

**Examples**:
- "Create supervisord service for Ghost"
- "Service failed to start"
- "Reload supervisorctl configuration"

### Runtime
**Definition**: Language runtime version (PHP, Python, Node.js, Ruby, Go).

**Management**: Via uberspace tools version command

**Examples**:
- "Set PHP runtime to 8.3"
- "Node.js runtime version 20"
- "Runtime version not available"

## Anti-Patterns (Don't Use These Terms)

### ❌ "Setup"
**Problem**: Too vague
**Use Instead**: "Installation" (the workflow) or "Provisioning" (creating infrastructure)

### ❌ "Deploy"
**Problem**: Ambiguous (could mean install, upgrade, or push code)
**Use Instead**: "Install" (first-time setup)

### ❌ "Config" (as noun)
**Problem**: Conflicts with package name
**Use Instead**: "Configuration" or "Manifest"

### ❌ "App Management"
**Problem**: Not workflow-focused, too generic
**Use Instead**: "App Installation" (our bounded context name)

### ❌ "Create"
**Problem**: Ambiguous (create what? manifest? instance? database?)
**Use Instead**: Specific verb - "Install app", "Provision database", "Configure backend"

## Consistency Rules

1. **Always use "Installation" (capital I)** when referring to the aggregate/workflow
2. **Always use "Manifest" (capital M)** when referring to the aggregate/TOML file
3. **Use "app" (lowercase)** when referring generically to software packages
4. **Use "Instance" (capital I)** when referring to installed app at specific location
5. **Event names are past tense**: ManifestLoaded, not LoadManifest (that's the command)
6. **Command names are imperative**: LoadManifest, not LoadingManifest
7. **State names are adjectives/nouns**: ManifestLoaded, ProvisioningInProgress, not Loading

## Usage in Code

### Package Names
```go
internal/appinstallation/domain/workflow     // Installation aggregate
internal/appinstallation/domain/manifest     // Manifest aggregate
internal/appinstallation/domain/valueobjects // AppName, Port, etc.
internal/appinstallation/application/installapp  // Use case
internal/appinstallation/infrastructure/filesystem // Adapters
```

### Type Names
```go
type Installation struct { ... }      // Aggregate root
type Manifest struct { ... }          // Aggregate
type AppName string                   // Value object
type Port int                         // Value object
type State int                        // Enum
type ManifestLoaded struct { ... }    // Domain event
```

### Method Names
```go
func (i *Installation) LoadManifest(...) error        // Command
func (i *Installation) ProvisionDatabase(...) error   // Command
func (i *Installation) State() State                  // Query
func (i *Installation) IsComplete() bool              // Query
```

## References

- PRD.md - EARS requirements using this language
- ARCHITECTURE.md - Architecture enforcing these terms
- Code - All types and methods use this vocabulary
