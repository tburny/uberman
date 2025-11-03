---
title: Diataxis Framework Reference
type: reference
date: 2025-11-03
---

# Diataxis Framework Reference

*Content type taxonomy for technical documentation*

**Source**: <https://diataxis.fr/>
**Author**: Daniele Procida

## Overview

Diataxis organizes documentation into four distinct types based on two axes:

**Axis 1**: Learning ↔ Working (user's goal)
**Axis 2**: Practical ↔ Theoretical (content nature)

```
                    LEARNING
                       |
    TUTORIALS          |          EXPLANATION
    (Practical)        |          (Theoretical)
    Learning           |          Understanding
    -------------------|-------------------
    HOW-TO GUIDES      |          REFERENCE
    (Practical)        |          (Theoretical)
    Working            |          Information
                       |
                    WORKING
```

## The Four Types: Quick Comparison

| Type | Axis | Purpose | Metaphor | User Question | Voice |
|------|------|---------|----------|---------------|-------|
| Tutorial | Learning × Practical | Learn by doing | Driving lesson | "Can you teach me?" | "First, you will..." |
| How-to | Working × Practical | Solve problem | Recipe | "How do I...?" | "Do this..." |
| Reference | Working × Theoretical | Look up facts | Encyclopedia | "What is...?" | Neutral description |
| Explanation | Learning × Theoretical | Understand | Article | "Why...?" | Discursive |

---

## Tutorial

### Purpose

**Learning-oriented, practical**: Take a beginner by the hand through a learning experience.

### Metaphor

A driving lesson - guided practice with guaranteed success.

### Characteristics

- **Step-by-step guidance** with explicit instructions
- **Guaranteed successful outcome** (builds confidence)
- **Minimal explanation** (save for Explanation section)
- **One path through domain** (not comprehensive)
- **Focus on learning**, not production use
- **Assume no prior knowledge**
- **Safe environment** (nothing can break)

### What to Include

- Clear learning objectives (what you'll build)
- Prerequisites listed explicitly
- Each step numbered and detailed
- Expected results at each stage
- Verification steps ("You should see...")
- Immediate feedback (success/failure indicators)
- Next steps after completion

### What to Avoid

- Multiple options or branches (decision paralysis)
- Detailed explanations (defer to Explanation docs)
- Production-ready code (simplify for learning)
- Assumptions about knowledge
- Skipping "obvious" steps

### For Uberman Project

**Target**: Apps only, NOT tool itself

- You don't need tutorials about Uberman (experienced user)
- DO create tutorials for installing apps (Lauti, WordPress, etc.)
- Tutorial = "Installing Lauti on Uberspace" (shows what's manual)

**Location**: Co-located with manifest

```
apps/<app>/
├── manifest.toml
└── README.md (tutorial-style)
```

### Example Structure

```markdown
# Installing [App] on Uberspace

[What you'll accomplish]

## Prerequisites
- Uberspace account
- [Software needed]

## Step 1: [Action]
```bash
# Exact commands
```

You should see: [Expected output]

## Step 2: [Next Action]

[Continue step-by-step]

## Verification

[How to confirm success]

## What You Learned

[Summary]

## Next Steps

- [Link to how-to guides]
- [Link to reference]

```markdown

---

## How-to Guide

### Purpose
**Task-oriented, practical**: Help a competent user solve a specific problem.

### Metaphor
A recipe - assumes kitchen skills, focuses on making one dish.

### Characteristics
- **Goal-oriented** (accomplish specific task)
- **Assume existing knowledge** (not teaching basics)
- **Problem-focused** (starts with "you want to...")
- **Flexible and adaptable** to user's context
- **Omit unnecessary detail** (link to reference instead)
- **Multiple guides** for multiple problems

### What to Include
- Problem statement (what you'll accomplish)
- Prerequisites (assumed knowledge, required setup)
- Clear, numbered steps
- Variations (if applicable)
- Verification (how to check success)
- Troubleshooting (common issues)
- Next steps or related guides

### What to Avoid
- Teaching fundamentals (not a tutorial)
- Comprehensive coverage (focus on one problem)
- Reference-style information (link instead)
- Unnecessary background (save for Explanation)

### For Uberman Project
**Target**: Infrequent tasks you might forget
- How to create custom manifest
- How to troubleshoot failed installations
- How to configure database credentials
- How to add custom lifecycle hooks

**Distribution**: ~10% of documentation

### Example Structure
```markdown
# How to [Accomplish Task]

[One sentence: problem being solved]

## Prerequisites
- [Assumed knowledge]
- [Required setup]

## Steps

1. **[Action 1]**
   \`\`\`bash
   # Command
   \`\`\`

2. **[Action 2]**
   [Continue]

## Verification
[Check success]

## Troubleshooting

| Issue | Solution |
|-------|----------|
| [Problem] | [Fix] |

## Next Steps
- [Related how-to]
- [Reference for details]
```

---

## Reference

### Purpose

**Information-oriented, theoretical**: Provide accurate, complete technical facts for lookup.

### Metaphor

An encyclopedia entry - neutral, comprehensive, authoritative.

### Characteristics

- **Structure mirrors code/system** (API, commands, config)
- **Comprehensive and authoritative** (single source of truth)
- **Consistent formatting** (tables, definitions)
- **Minimal prose** (maximum information density)
- **Truth and certainty focused** (facts, not opinions)
- **Quick to scan** (users know what they're looking for)

### What to Include

- Syntax definitions
- Complete parameter/option lists (types, defaults, constraints)
- Return values or exit codes
- Examples (brief, illustrative)
- Version information
- See also links

### What to Avoid

- How-to instructions (defer to how-to guides)
- Tutorials (defer to tutorials)
- Design rationale (defer to explanations)
- Verbose descriptions (be concise)
- Opinion or recommendation

### For Uberman Project

**Target**: PRIMARY documentation type (60%)

- CLI command reference (`uberman install --help` formalized)
- Manifest schema (TOML structure, all fields)
- Error codes and meanings
- Exit codes
- Configuration file reference

**Why 60%**: Experienced user needs quick lookups when forgetting syntax

### Example Structure

```markdown
# [Thing] Reference

[One-sentence description]

## Syntax

\`\`\`
[Formal syntax]
\`\`\`

## Parameters / Options

| Name | Type | Required | Default | Description |
|------|------|----------|---------|-------------|
| [param] | [type] | [yes/no] | [value] | [what it does] |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | [Error type] |

## Examples

### Example 1: [Common case]
\`\`\`bash
[command]
\`\`\`

### Example 2: [Another case]
[2-3 more examples]

## See Also
- [Related reference]
- [How-to for usage]
```

---

## Explanation

### Purpose

**Understanding-oriented, theoretical**: Clarify and illuminate concepts, design decisions, and alternatives.

### Metaphor

An academic article - explores ideas, discusses trade-offs, provides context.

### Characteristics

- **Design decisions** and alternatives considered
- **Context and background** (why things are this way)
- **Trade-offs explored** (pros/cons of choices)
- **Not about tasks** (no step-by-step instructions)
- **Can include opinions** and perspectives
- **Discursive** (follow interesting tangents)

### What to Include

- Problem or decision context
- Alternatives considered
- Rationale for chosen approach
- Trade-offs and consequences
- Related concepts
- Historical context (if relevant)
- Future implications

### What to Avoid

- Step-by-step instructions (defer to how-to)
- Complete technical specs (defer to reference)
- Teaching basics (defer to tutorials)
- Assuming one right answer

### For Uberman Project

**Target**: ADRs + architecture rationale (15%)

- Why Clean Architecture for Uberman
- Why single bounded context (App Installation)
- Why Shape Up methodology
- Understanding Uberspace platform constraints
- AI copilot collaboration lessons

**Format**: Architecture Decision Records (ADRs)

### Example Structure (ADR)

```markdown
# ADR-NNNN: [Decision Title]

**Status**: Accepted | Proposed | Deprecated
**Date**: YYYY-MM-DD

## Context
[What decision needs to be made? What's driving this?]

## Decision
[What was decided?]

## Rationale
[Why? Alternatives considered? Trade-offs?]

**Alternatives**:
1. [Option A]: [Why rejected]
2. [Option B]: [Why rejected]

## Consequences

**Positive**:
- [What this enables]

**Negative**:
- [What this constrains]

## Related Decisions
- [Other ADRs]
```

---

## Separation Principle

### Critical Rule

**Each document = ONE type only**

Mixing types creates confusion because users come with different mindsets:

- **Learning mode**: Want hand-holding, explanations, practice
- **Working mode**: Want quick answers, assume knowledge, no fluff

### What NOT to Mix

| ❌ Don't Do This | ✓ Do This Instead |
|------------------|-------------------|
| Tutorial with reference tables | Tutorial links to reference for details |
| How-to with design rationale | How-to links to explanation for "why" |
| Reference with step-by-step | Reference links to how-to for usage |
| Explanation with procedures | Explanation links to how-to for implementation |

### How to Link Between Types

**From Tutorial**:

- → Reference: "See [CLI Reference] for all options"
- → How-to: "Next, learn [How to Customize Installation]"

**From How-to**:

- → Reference: "See [Manifest Schema] for complete field list"
- → Explanation: "For rationale, see [Why This Approach]"

**From Reference**:

- → How-to: "Usage examples: [How to Use This Command]"
- → Tutorial: "Learning guide: [Getting Started Tutorial]"

**From Explanation**:

- → How-to: "To implement this: [How to Apply Pattern]"
- → Tutorial: "Hands-on practice: [Build Example App]"

---

## Distribution Guidelines

### General Guidance

Balance depends on audience and product maturity:

**Beginner-heavy audience**:

- Tutorials: 30% (lots of learning)
- How-to: 30% (common tasks)
- Reference: 20% (lookup)
- Explanation: 20% (understanding)

**Expert audience**:

- Tutorials: 5% (minimal learning)
- How-to: 35% (task-focused)
- Reference: 50% (primary need)
- Explanation: 10% (novel concepts only)

### Uberman-Specific Distribution

**Target**: 60% reference, 25% apps (tutorials), 15% explanation, 10% how-to

**Rationale**:

- **Personal tool** for experienced user (you)
- **Quick reference primary need** (CLI syntax, manifest format)
- **App tutorials preserve knowledge** (what manual process looks like)
- **Explanation for decisions** (ADRs for future-you)
- **Minimal how-to** (you can figure most things out)

**No tool tutorials**: You don't need teaching about Uberman itself

---

## Decision Tree

Use this to determine content type:

```
START: What does the user want to accomplish?

├─ Learn something NEW (never done before)
│  └─ TUTORIAL
│     Example: First time installing app on Uberspace
│
├─ Solve SPECIFIC PROBLEM (knows basics)
│  └─ HOW-TO GUIDE
│     Example: How to troubleshoot database connection
│
├─ Look up TECHNICAL FACTS (forgot details)
│  └─ REFERENCE
│     Example: What are all the manifest.toml fields?
│
└─ Understand WHY or DESIGN (context/rationale)
   └─ EXPLANATION
      Example: Why did we choose Clean Architecture?
```

---

## Quality Checklist by Type

### Tutorial Checklist

- [ ] Clear learning objective stated upfront
- [ ] Prerequisites explicitly listed
- [ ] Every step numbered and detailed
- [ ] Expected outcomes shown for each step
- [ ] Verification steps included
- [ ] Guaranteed successful completion
- [ ] Minimal branching or options
- [ ] Links to how-to/reference for next steps

### How-to Checklist

- [ ] Problem statement clear (what you'll accomplish)
- [ ] Prerequisites listed (assumed knowledge)
- [ ] Steps are clear and actionable
- [ ] Verification included (check success)
- [ ] Troubleshooting table provided
- [ ] Focused on single problem
- [ ] No teaching of fundamentals
- [ ] Links to reference for details

### Reference Checklist

- [ ] Syntax formally defined
- [ ] All parameters documented (type, default, required)
- [ ] Information presented in tables
- [ ] Minimal prose, maximum density
- [ ] Examples brief and illustrative
- [ ] Neutral tone (no opinion)
- [ ] Comprehensive coverage
- [ ] Version information included

### Explanation Checklist

- [ ] Context and background provided
- [ ] Alternatives discussed
- [ ] Trade-offs explained
- [ ] Rationale clear
- [ ] Consequences identified (positive/negative)
- [ ] No step-by-step instructions
- [ ] Related concepts linked
- [ ] Can include opinion/perspective

---

## Further Reading

- **Official site**: <https://diataxis.fr/>
- **Book reference**: "Docs for Developers" (Bhatti et al.) uses similar taxonomy
- **Workflow framework**: See `docs/reference/docs-for-developers-workflow.md`
- **Templates**: See `docs/reference/documentation-templates.md`
