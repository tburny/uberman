# Requirements Engineering Agent - Usage Guide

**Comprehensive guide to using the Uberman Requirements Engineering toolkit**

## Overview

The Requirements Engineering Agent helps you create, analyze, and improve requirements using EARS (Easy Approach to Requirements Syntax). It provides:

- **Interactive skill** for comprehensive requirements work
- **Quick commands** for common tasks
- **Documentation** and templates for reference
- **Quality analysis** based on INCOSE/IEEE standards

## Getting Started

### Invoking the Agent

**Method 1: Use the Skill (Recommended)**

```
@requirements-engineer

I need help writing requirements for [your feature/system]
```

The agent will greet you and ask which mode you want to use.

**Method 2: Use Quick Commands**

For specific tasks, use slash commands:
- `/requirements-convert` - Convert informal to EARS format
- `/requirements-analyze` - Analyze quality and suggest improvements
- `/requirements-elicit` - Start interactive elicitation session
- `/requirements-batch` - Process multiple requirements

### Choosing the Right Approach

| Task | Use | Example |
|------|-----|---------|
| Write new requirements from scratch | `/requirements-elicit` | "I need requirements for password reset" |
| Convert informal description | `/requirements-convert` | "Users should be able to export data" |
| Improve existing requirement | `/requirements-analyze` | "The system must be fast" |
| Review many requirements | `/requirements-batch` | Paste list of 20+ requirements |
| Learn EARS | `@requirements-engineer` Mode 5: Teach | "Explain state-driven requirements" |

## Usage Examples

### Example 1: Creating New Requirements

**Scenario**: You need requirements for a login feature

**Step 1**: Start elicitation
```
/requirements-elicit
```

**Step 2**: Provide context
```
I need requirements for a login feature in a web application.
Users should be able to log in with email and password.
```

**Agent Response**:
```
**Current Context**:
System: Web Application
Feature: User Login
Stakeholders: End users, system administrators

**Requirements Drafted**:
[REQ-LOGIN-001] When a user enters valid credentials and clicks Login,
the authentication service shall verify the credentials and shall grant
access to the user's account within 1 second.

**Clarifying Questions**:
1. How should the system handle invalid credentials?
2. Are there account lockout requirements after failed attempts?
3. Should there be session timeout requirements?
4. Do you need "Remember Me" functionality?
```

**Step 3**: Answer questions, iterate until complete

### Example 2: Improving Existing Requirements

**Input**:
```
/requirements-analyze

The app should respond quickly when users click buttons.
```

**Agent Response**:
```
**Original**: "The app should respond quickly when users click buttons."

**Quality Scores** (1-5 scale):
- Clarity: 1/5 - Vague "quickly", unclear which app, which buttons
- Testability: 1/5 - Cannot measure "quickly" objectively
- Completeness: 2/5 - Missing specific trigger, measurable criteria

**EARS Pattern**: None

**Issues Detected**:
1. Weak modal: Uses "should" instead of "shall"
2. Vague term: "quickly" (unmeasurable)
3. Vague nouns: "app" (which app?), "buttons" (which buttons?)
4. Missing response specification

**Improved Version:**
[REQ-UI-001] When a user clicks any button in the mobile application,
the application shall provide visual feedback within 200 milliseconds to
maintain perceived responsiveness.

**Improvements Made**:
- Replaced "should" with "shall"
- Added specific trigger ("user clicks any button")
- Specified system ("mobile application")
- Replaced "quickly" with measurable "200 milliseconds"
- Added rationale for timing ("perceived responsiveness")
```

### Example 3: Batch Processing

**Input**:
```
/requirements-batch

1. System should validate input
2. Users must be able to download reports
3. The API shall return JSON responses
4. Handle errors appropriately
5. System should log all transactions
```

**Agent Response**: [Comprehensive batch analysis with pattern distribution, common issues, detailed analysis per requirement, and prioritized improvements]

### Example 4: Learning EARS

**Input**:
```
@requirements-engineer

Teach me about event-driven requirements with examples
```

**Agent Response**: [Detailed explanation of "When" pattern with multiple examples, when to use it, how to write test cases, etc.]

## Operating Modes

The main skill has 5 modes:

### Mode 1: Generate
**When to use**: Creating new requirements from descriptions

**What it does**:
- Parses your description
- Identifies appropriate EARS pattern
- Generates well-formed requirement
- Explains reasoning
- Provides alternatives

**Example**:
```
Input: "Users need to export their data"

Output:
[REQ-EXPORT-001] When a user clicks the Export Data button,
the application shall generate a CSV file containing all user
data and shall initiate download within 3 seconds.

Pattern: Event-Driven (When)
Rationale: User action triggers system response
```

### Mode 2: Analyze
**When to use**: Improving existing requirements

**What it does**:
- Assesses quality (1-5 scores)
- Detects ambiguities and weak terms
- Identifies EARS pattern (if present)
- Suggests specific improvements
- Generates EARS-compliant version

### Mode 3: Elicit
**When to use**: Interactive requirements gathering

**What it does**:
- Asks targeted questions
- Extracts requirements from answers
- Identifies gaps in coverage
- Drafts requirements iteratively
- Ensures completeness

### Mode 4: Batch
**When to use**: Processing many requirements

**What it does**:
- Analyzes entire requirement set
- Checks cross-requirement consistency
- Identifies patterns and issues
- Generates summary statistics
- Prioritizes improvements

### Mode 5: Teach
**When to use**: Learning EARS and best practices

**What it does**:
- Explains EARS patterns
- Provides examples
- Shows before/after transformations
- References standards
- Answers questions about requirements engineering

## Best Practices

### 1. Start with Context

Always provide:
- System/feature name
- Domain/industry
- Key stakeholders
- Any constraints

**Good**:
```
I'm writing requirements for an e-commerce checkout process.
Users are customers making purchases. We need PCI compliance
and sub-3-second page loads.
```

**Not helpful**:
```
I need requirements for my project.
```

### 2. Be Specific About Your Needs

Tell the agent:
- What you want to accomplish
- What you already have (if anything)
- What you're uncertain about
- Any special constraints

### 3. Iterate

Requirements rarely emerge perfect on first try:
1. Get initial draft from agent
2. Review and identify gaps
3. Ask clarifying questions
4. Refine iteratively
5. Validate with stakeholders

### 4. Use the Right Tool

- **Quick task** → Use slash command
- **Complex feature** → Use full skill with elicitation
- **Learning** → Use teach mode
- **Review** → Use batch mode

### 5. Maintain Requirements in Version Control

After using the agent:
- Copy requirements to your requirements document
- Commit to version control
- Link to issues/stories
- Update as needed

## Integration with Development

### Linking to User Stories (Agile)

EARS requirements complement user stories:

**User Story**:
```
As a user, I want to reset my password so that I can regain access if I forget it.
```

**EARS Requirements** (from agent):
```
[REQ-PWD-001] When a user clicks "Forgot Password" on the login page,
the system shall send a password reset link to the user's registered
email address within 30 seconds.

[REQ-PWD-002] When a user clicks the password reset link, the system
shall display the password reset form and shall ensure the link expires
after 1 hour.

[REQ-PWD-003] When a user submits a new password, the system shall
validate the password meets complexity requirements (minimum 8
characters, 1 uppercase, 1 number, 1 special character) and shall
update the password if valid.

[REQ-PWD-004] If the password reset link has expired, then the system
shall display "This link has expired. Please request a new password
reset."
```

### Creating Acceptance Criteria

Ask the agent to generate acceptance criteria:

```
For requirement [REQ-PWD-001], generate acceptance criteria
```

**Agent provides**:
```
Acceptance Criteria for REQ-PWD-001:
- [ ] "Forgot Password" link is visible on login page
- [ ] Clicking link triggers email send
- [ ] Email contains valid reset link
- [ ] Email arrives within 30 seconds (measure with test)
- [ ] Email goes to correct registered address
- [ ] System handles missing/invalid email addresses
- [ ] Appropriate error messages shown if send fails
```

### Test Case Generation

Use requirements to drive test cases:

```
Help me write test cases for [REQ-PWD-003]
```

**Agent provides**:
```
Test Cases for REQ-PWD-003:

TC-PWD-003-001: Valid password accepted
1. Navigate to password reset form
2. Enter password: "SecureP@ss1"
3. Submit form
4. Expected: Password updated successfully

TC-PWD-003-002: Password too short rejected
1. Navigate to password reset form
2. Enter password: "Short1!"
3. Submit form
4. Expected: Error "Password must be at least 8 characters"

TC-PWD-003-003: Missing uppercase rejected
1. Navigate to password reset form
2. Enter password: "securep@ss1"
3. Submit form
4. Expected: Error "Password must contain at least 1 uppercase letter"

[Additional test cases for each validation rule...]
```

## Troubleshooting

### Agent doesn't understand my description

**Solution**: Provide more context
- Specify the system/component
- Clarify the trigger/event
- Describe the desired outcome
- Mention any constraints

### Generated requirement doesn't match my intent

**Solution**: Clarify and iterate
```
That's close, but the system should [your correction].
Can you revise the requirement?
```

### Too many requirements generated

**Solution**: Focus the scope
```
Let's focus on just the happy path for now.
We'll handle error cases in a separate session.
```

### Requirement seems too rigid

**Solution**: Discuss with agent
```
This seems overly specific. Can we make it more flexible
while staying testable?
```

The agent will explain trade-offs and offer alternatives.

## Quick Reference

### Slash Commands

| Command | Purpose | Input |
|---------|---------|-------|
| `/requirements-convert` | Convert to EARS | Informal description |
| `/requirements-analyze` | Quality analysis | Existing requirement |
| `/requirements-elicit` | Interactive session | System/feature description |
| `/requirements-batch` | Bulk processing | List of requirements |

### Skill Invocation

```
@requirements-engineer [your request or question]
```

### Quality Scoring

| Score | Meaning |
|-------|---------|
| 5 | Excellent - ready for implementation |
| 4 | Good - minor refinements needed |
| 3 | Acceptable - some improvements recommended |
| 2 | Poor - significant issues to address |
| 1 | Unacceptable - major rework required |

### EARS Patterns Quick Reference

| Keyword | Pattern | Example |
|---------|---------|---------|
| *(none)* | Ubiquitous | "The API **shall** use HTTPS" |
| **While** | State-Driven | "**While** offline, the app **shall** queue requests" |
| **When** | Event-Driven | "**When** user clicks Save, system **shall** save file" |
| **Where** | Optional | "**Where** GPS enabled, app **shall** show location" |
| **If...Then** | Unwanted | "**If** timeout occurs, **then** system **shall** retry" |

## Additional Resources

- **EARS Guide**: See [EARS_GUIDE.md](EARS_GUIDE.md) for complete EARS reference
- **Quality Framework**: See [REQUIREMENTS_QUALITY.md](REQUIREMENTS_QUALITY.md) for quality criteria
- **Examples**: See [REQUIREMENTS_EXAMPLES.md](REQUIREMENTS_EXAMPLES.md) for domain-specific examples
- **Templates**: See `templates/requirements/` for worksheets and checklists

## Support

For questions or issues:
1. Ask the agent itself (teach mode)
2. Consult the documentation
3. Review examples
4. Check INCOSE/IEEE standards

---

*Part of the Uberman Requirements Engineering toolkit*
