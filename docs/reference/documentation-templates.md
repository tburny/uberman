---
title: Documentation Templates Reference
type: reference
date: 2025-11-03
---

# Documentation Templates Reference

*Copy-paste ready templates for all documentation types*

## Frontmatter Standard

All documentation files should include YAML frontmatter:

```yaml
---
title: [Human-friendly title]
type: tutorial|howto|reference|explanation
date: YYYY-MM-DD
---
```

**For Uberman**: Minimal frontmatter (personal tool, no tracking needed)
- No `owner` field (it's you)
- No `status` field (if it exists, it's current)
- No `last_reviewed` field (you'll notice when outdated)

---

## Reference Document Template

**Use for**: API reference, CLI commands, configuration options, error codes, schemas

```markdown
---
title: [Thing] Reference
type: reference
date: YYYY-MM-DD
---

# [Thing] Reference

[One-sentence description of what this reference covers]

## Syntax

\`\`\`
[Formal syntax definition]
\`\`\`

## Options / Parameters

| Name | Type | Required | Default | Description |
|------|------|----------|---------|-------------|
| [name] | [string/int/bool] | [yes/no] | [value] | [What it does] |
| [name] | [string/int/bool] | [yes/no] | [value] | [What it does] |

## Exit Codes / Error Codes

| Code | Meaning | Description |
|------|---------|-------------|
| 0 | Success | [When this occurs] |
| 1 | [Error type] | [What went wrong] |
| 2 | [Error type] | [What went wrong] |

## Examples

### Example 1: [Common Use Case]

\`\`\`bash
# Brief comment explaining what this does
[command with options]
\`\`\`

**Output**:
\`\`\`
[Expected output shown here]
\`\`\`

### Example 2: [Another Use Case]

\`\`\`bash
[command with different options]
\`\`\`

**Output**:
\`\`\`
[Expected output]
\`\`\`

### Example 3: [Edge Case or Advanced]

\`\`\`bash
[command]
\`\`\`

## See Also

- [Related reference document]
- [How-to guide for usage]
- [Explanation for design rationale]
```

---

## How-to Guide Template

**Use for**: Task-oriented guides, troubleshooting, configuration tasks

```markdown
---
title: How to [Accomplish Specific Task]
type: howto
date: YYYY-MM-DD
---

# How to [Accomplish Specific Task]

[One sentence describing what problem this solves or task this accomplishes]

## Prerequisites

- [ ] [Assumed knowledge or skill]
- [ ] [Required software or tools]
- [ ] [Necessary access or permissions]

## Steps

### 1. [First Action]

[Brief explanation of what this step does and why]

\`\`\`bash
# Command with clear comments
[command]
\`\`\`

**Expected result**: [What you should see after this step]

### 2. [Second Action]

[Continue with clear, numbered steps]

\`\`\`bash
[command]
\`\`\`

### 3. [Final Action]

[Last step]

\`\`\`bash
[command]
\`\`\`

## Verification

Verify the task completed successfully:

\`\`\`bash
# Test command
[verification command]
\`\`\`

**Expected output**:
\`\`\`
[What success looks like]
\`\`\`

## Troubleshooting

| Issue | Symptoms | Solution |
|-------|----------|----------|
| [Common problem 1] | [How it manifests] | [How to fix] |
| [Common problem 2] | [How it manifests] | [How to fix] |
| [Common problem 3] | [How it manifests] | [How to fix] |

## Next Steps

- [Link to related how-to guide]
- [Link to reference for more details]
- [Link to explanation for background]
```

---

## Tutorial Template

**Use for**: Learning-oriented, step-by-step guides for beginners

```markdown
---
title: [Learning Goal - e.g., "Installing Your First App"]
type: tutorial
date: YYYY-MM-DD
---

# [Learning Goal]

[What you'll accomplish in this tutorial and why it's valuable]

## What You'll Learn

- [Learning objective 1]
- [Learning objective 2]
- [Learning objective 3]

## Prerequisites

- [ ] [Required knowledge - be explicit, assume nothing]
- [ ] [Required software or tools]
- [ ] [Required access or setup]

## What You'll Build

[Brief description of the end result - what they'll have when done]

## Step 1: [First Action]

[Detailed, explicit instructions. Assume beginner - don't skip "obvious" steps]

\`\`\`bash
# Clear comment explaining what this command does
[exact command to run]
\`\`\`

**You should see**:
\`\`\`
[Exact expected output]
\`\`\`

✓ **Check**: [Verification step - how to know it worked]

## Step 2: [Second Action]

[Continue with detailed steps]

\`\`\`bash
[command]
\`\`\`

**You should see**:
\`\`\`
[Expected output]
\`\`\`

## Step 3: [Continue through completion]

[Keep going until tutorial is complete]

## Verification

Let's verify everything worked:

\`\`\`bash
# Test the result
[verification commands]
\`\`\`

**Expected result**: [Describe success]

## What You Learned

[Summary of what was accomplished]

- [Key takeaway 1]
- [Key takeaway 2]
- [Key takeaway 3]

## Next Steps

Now that you've completed this tutorial:

- [Link to next tutorial in series]
- [Link to how-to guides for common tasks]
- [Link to reference for complete options]

## Troubleshooting

If something went wrong:

| Issue | Solution |
|-------|----------|
| [Problem] | [Fix] |
```

---

## Explanation / ADR Template

**Use for**: Architectural decisions, design rationale, understanding concepts

```markdown
---
title: ADR-NNNN: [Decision Title]
type: explanation
date: YYYY-MM-DD
---

# ADR-NNNN: [Decision Title]

**Status**: Proposed | Accepted | Deprecated | Superseded
**Date**: YYYY-MM-DD
**Deciders**: [Name(s)]
**Supersedes**: [ADR-NNNN] (if applicable)
**Superseded by**: [ADR-NNNN] (if deprecated)

## Context

[What's the issue or decision to be made? What's the background? What factors are driving this decision?]

## Decision

[What did you decide? State the decision clearly and concisely.]

## Rationale

[Why did you make this decision? What alternatives did you consider? What are the trade-offs?]

### Alternatives Considered

1. **[Option A]**
   - **Pros**: [Benefits]
   - **Cons**: [Drawbacks]
   - **Why rejected**: [Reason]

2. **[Option B]**
   - **Pros**: [Benefits]
   - **Cons**: [Drawbacks]
   - **Why rejected**: [Reason]

### Trade-offs

- [Trade-off 1: What you gain vs what you sacrifice]
- [Trade-off 2: Another consideration]

## Consequences

### Positive

- [What this decision enables]
- [Benefits gained]
- [Problems solved]

### Negative

- [What this decision constrains]
- [Limitations accepted]
- [New challenges introduced]

### Risks

- [What could go wrong]
- [Mitigation strategies]

## Related Decisions

- Supersedes: [ADR-NNNN - Previous Decision]
- Related to: [ADR-NNNN - Connected Decision]
- See also: [ADR-NNNN - Similar Context]

## Notes

[Additional observations, follow-up actions, or future refinements to consider]
```

---

## App Installation Guide Template

**Use for**: App-specific installation documentation (Uberman apps)

```markdown
---
title: Installing [App Name] on Uberspace
type: tutorial
date: YYYY-MM-DD
---

# Installing [App Name] on Uberspace

[One-sentence description of what this app is and why you'd install it]

## Using Uberman (Automated)

Install [App] with a single command:

\`\`\`bash
uberman install [app-name]
\`\`\`

The automated installation handles:
- [Key task 1]
- [Key task 2]
- [Key task 3]

## Manual Process (What Uberman Automates)

[Document the step-by-step manual installation. This preserves knowledge about what the automation does and serves as a reference if troubleshooting is needed.]

### Prerequisites

- Uberspace account (v7+)
- [Required software or dependencies]
- [Any specific Uberspace configuration]

### Step 1: [First Manual Step]

\`\`\`bash
# Manual commands
[command]
\`\`\`

[Explanation of what this does]

### Step 2: [Second Manual Step]

\`\`\`bash
[command]
\`\`\`

[Continue documenting manual process]

### Step 3: [Continue through all steps]

[Keep going until manual installation is complete]

## Post-Installation

[Any manual configuration that Uberman cannot automate]

1. [Manual config step 1]
2. [Manual config step 2]

## Manifest Design

[Explain the choices made in the app's manifest.toml file]

**Key decisions**:
- **Port**: [Why this port was chosen]
- **Database**: [Database configuration reasoning]
- **Directory structure**: [Why organized this way]
- **Services**: [Supervisord service configuration]

**Trade-offs**:
- [What was prioritized]
- [What was deferred]

## Verification

Verify the installation worked:

\`\`\`bash
# Check service status
supervisorctl status [app-name]

# Test web access
curl https://[youraccount].uber.space/[app-path]
\`\`\`

**Expected results**:
- Service status: `RUNNING`
- HTTP response: `200 OK`

## Troubleshooting

Common issues and solutions:

| Issue | Symptoms | Solution |
|-------|----------|----------|
| Database connection fails | Service won't start, logs show connection error | Check database credentials in config file |
| Port already in use | Service fails with "address already in use" | Run `ss -tulpn` to find conflicting service |
| Web backend not working | 502 Bad Gateway | Verify `uberspace web backend list` shows correct configuration |

## Configuration

[Post-install configuration options, if relevant]

### [Configuration Topic 1]

[How to configure this aspect]

\`\`\`bash
[config commands]
\`\`\`

### [Configuration Topic 2]

[Continue with configuration options]

## References

- **Uberspace Lab**: [Link to lab guide if it exists]
- **Official Documentation**: [Link to app's official docs]
- **Manifest**: [Link to apps/[app-name]/manifest.toml]
- **Uberman Docs**: [Link to relevant Uberman documentation]

## Notes

[Any additional context, gotchas, future improvements planned, or lessons learned]
```

---

## README Template

**Use for**: Project overview, repository entry point

```markdown
---
title: [Project Name] README
type: reference
date: YYYY-MM-DD
---

# [Project Name]

[One-sentence description of what this project does]

[![Badge 1](url)](link)
[![Badge 2](url)](link)

[Brief paragraph expanding on the description - what problem does this solve?]

## Features

- **[Feature 1]**: [Brief description]
- **[Feature 2]**: [Brief description]
- **[Feature 3]**: [Brief description]
- **[Feature 4]**: [Brief description]

## Installation

### Quick Install

\`\`\`bash
[One-line install command if available]
\`\`\`

### Manual Install

\`\`\`bash
[Step-by-step installation commands]
\`\`\`

See [Installation Guide](docs/INSTALLATION.md) for detailed instructions.

## Quick Start

Get started quickly:

\`\`\`bash
# Most common usage
[basic command example]
\`\`\`

See [Getting Started Tutorial](docs/tutorials/getting-started.md) for complete walkthrough.

## Documentation

### User Documentation

- **Getting Started**: [Link to tutorial]
- **How-to Guides**: [Link to how-to directory]
- **Reference**: [CLI Commands](docs/reference/cli-reference.md) | [Configuration](docs/reference/config-reference.md)

### Developer Documentation

- **Architecture**: [ARCHITECTURE.md](ARCHITECTURE.md)
- **Contributing**: [CONTRIBUTING.md](CONTRIBUTING.md)
- **Decisions**: [decisions/](decisions/) (ADRs)

### Project Management

- **Methodology**: [Shape Up](plans/README.md)
- **Planning**: [PLANNING.md](PLANNING.md)

## Usage

### Basic Commands

\`\`\`bash
# Command 1
[command with example]

# Command 2
[command with example]
\`\`\`

See [CLI Reference](docs/reference/cli-reference.md) for all commands.

## Configuration

[Brief configuration overview]

See [Configuration Guide](docs/reference/config-reference.md) for details.

## Development

### Prerequisites

- [Requirement 1]
- [Requirement 2]

### Setup

\`\`\`bash
# Clone repository
git clone [url]

# Install dependencies
[install command]

# Build
[build command]

# Run tests
[test command]
\`\`\`

See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## License

[License type] - see [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Credit 1]
- [Credit 2]

## Support

- **Issues**: [GitHub Issues](url)
- **Discussions**: [GitHub Discussions](url)
- **Documentation**: [Full docs](url)
```

---

## Release Notes Template

**Use for**: Version release announcements, changelogs

```markdown
---
title: Release Notes - v[X.Y.Z]
type: reference
date: YYYY-MM-DD
---

# Release Notes - v[X.Y.Z]

**Release Date**: YYYY-MM-DD
**Status**: Stable | Beta | Release Candidate

[Brief 1-2 sentence summary of this release]

## New Features

### [Feature Name]

[Description of the feature and why it's valuable]

**Usage**:
\`\`\`bash
[Example of how to use new feature]
\`\`\`

See [Documentation link] for details.

### [Another Feature]

[Continue listing new features]

## Improvements

- **[Component]**: [What was improved]
- **[Component]**: [What was improved]
- **[Component]**: [What was improved]

## Bug Fixes

- **[Issue #123]**: [Description of bug fixed]
- **[Issue #456]**: [Description of bug fixed]
- **[Issue #789]**: [Description of bug fixed]

## Breaking Changes

⚠️ **[Breaking Change 1]**

[Description of what changed and why]

**Impact**: [Who this affects]

**Migration**: [How to update existing code/config]
- Step 1: [Action]
- Step 2: [Action]

See [Migration Guide](link) for complete instructions.

⚠️ **[Breaking Change 2]**

[Continue with additional breaking changes]

## Deprecations

- **[Feature/API]**: Deprecated in v[X.Y.Z], will be removed in v[X+1.0.0]
  - **Replacement**: Use [alternative] instead
  - **Migration**: See [guide link]

## Known Issues

- **[Issue description]**: [Workaround if available]
- **[Issue description]**: [Workaround if available]

## Upgrade Instructions

### From v[X.Y-1.Z]

\`\`\`bash
[Upgrade commands]
\`\`\`

### From earlier versions

[Special upgrade notes if coming from older versions]

## Dependencies

### Updated

- [Dependency 1]: v[old] → v[new]
- [Dependency 2]: v[old] → v[new]

### Added

- [New dependency]: v[X.Y.Z]

### Removed

- [Deprecated dependency]: No longer required

## Contributors

Thank you to all contributors for this release:

- @[username] - [Contributions]
- @[username] - [Contributions]

## Full Changelog

See [CHANGELOG.md](CHANGELOG.md) or [GitHub Releases](url) for complete changes.
```

---

## Frontmatter Variations

### With Version Applicability

```yaml
---
title: [Title]
type: reference
date: YYYY-MM-DD
applies_to:
  min_version: "1.0.0"
  max_version: "2.0.0"
---
```

### With Prerequisites

```yaml
---
title: [Title]
type: tutorial
date: YYYY-MM-DD
prerequisites:
  - "Understanding of Go basics"
  - "Uberspace account"
estimated_time: "30 minutes"
---
```

### With Status (If Tracking Needed)

```yaml
---
title: [Title]
type: howto
date: YYYY-MM-DD
status: draft | review | published | outdated
---
```

---

## Quick Selection Guide

**Need to document**:
- CLI command or API? → **Reference** template
- How to accomplish task? → **How-to** template
- Learning experience? → **Tutorial** template
- Design decision? → **ADR** (Explanation) template
- App installation? → **App Installation** template
- Project overview? → **README** template
- Version release? → **Release Notes** template

---

## Related Documentation

- **Diataxis Framework**: See `docs/reference/diataxis-framework.md`
- **Quality Workflow**: See `docs/reference/docs-for-developers-workflow.md`
- **Writing Documentation Skill**: See `.claude/skills/writing-documentation.md`
