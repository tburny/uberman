---
name: writing-documentation
description: Create and review technical documentation using Diataxis framework and Docs for Developers methodology. Use when writing new docs, reviewing documentation quality, determining content type, editing for clarity, or instructing Claude to generate documentation. Includes prompting patterns for AI-assisted documentation and incremental "good enough for now" approach.
version: 1.0.0
allowed-tools: [Read, Grep, Glob]
metadata:
  author: Uberman Project
  tags: [documentation, diataxis, docs-for-developers, quality, prompting]
  frameworks: [diataxis, docs-for-developers]
  last-updated: 2025-11-03
---

# Writing Documentation

*Comprehensive framework combining Diataxis + Docs for Developers for technical documentation*

## Quick Start

### When Writing New Documentation

1. **Identify content type** using decision tree below
2. **Choose template** from `docs/reference/documentation-templates.md`
3. **Write first draft** (functional quality over polish)
4. **Apply quality checks** using multi-pass editing workflow

### When Instructing Claude to Write

Use these prompting patterns:
- "Write a [type] about [topic] following Diataxis [type] principles"
- "Use 'good enough for now' approach - functional over perfect"
- "Apply first editing pass only: technical accuracy"
- "Generate [content] incrementally, document what exists"

### When Reviewing Documentation

Apply multi-pass editing workflow:
1. Technical accuracy (does it work?)
2. Completeness (any gaps?)
3. Structure (logical flow?)
4. Clarity (easy to understand?)

See `docs/reference/docs-for-developers-workflow.md` for detailed checklists.

---

## Content Type Decision

### The Four Diataxis Types

**Quick decision guide**:

```
User learning something new?
└─> TUTORIAL
    Purpose: Learning-oriented, step-by-step guidance
    Example: "Installing Lauti on Uberspace"

User solving specific problem?
└─> HOW-TO GUIDE
    Purpose: Task-oriented, assume knowledge
    Example: "How to Create a Custom Manifest"

User looking up facts?
└─> REFERENCE
    Purpose: Information-oriented, quick lookup
    Example: "CLI Commands Reference", "Manifest Schema"

User understanding concepts?
└─> EXPLANATION
    Purpose: Understanding-oriented, design rationale
    Example: "Why Clean Architecture for Uberman"
```

### Uberman Distribution

Personal tool for experienced user:
- **Reference: 60%** - CLI syntax, manifest schema, error codes (primary need)
- **Apps: 25%** - Tutorial-style installation guides (preserve knowledge)
- **Explanation: 15%** - ADRs, design rationale (future-you context)
- **How-to: 10%** - Infrequent tasks (troubleshooting, config)
- **Tool tutorials: 0%** - You don't need teaching about Uberman itself

### Special Content Types

Beyond the four Diataxis types:

- **App Installation Guides**: Tutorial-style, co-located with manifest (`apps/<app>/README.md`)
- **ADRs**: Architectural decisions, separate from user docs (`decisions/`)
- **Getting Started**: Trust-building first impressions, landing page

**See `docs/reference/diataxis-framework.md` for comprehensive type characteristics**

### Separation Principle

**CRITICAL**: Each document = ONE type only

❌ Don't mix:
- Tutorials in reference docs
- Explanations in how-to guides
- How-to steps in tutorials
- Reference tables in explanations

✅ Do link:
- Tutorial → Reference (for details)
- How-to → Explanation (for rationale)
- Reference → How-to (for usage)
- Explanation → Tutorial (for practice)

---

## Prompting Patterns for Claude

### Pattern 1: Type-Aware Creation

**When**: Starting new documentation
**Goal**: Generate content following Diataxis principles

**Prompt structure**:
```
Write a [tutorial|howto|reference|explanation] about [topic].

Requirements:
- Follow Diataxis [type] principles (see characteristics below)
- Target audience: [Experienced sysadmin/SRE | Beginners | specific persona]
- Tone: [Direct and concise | Friendly and explanatory | Formal and precise]
- Approach: "Good enough for now, safe enough to try"
- [Type-specific requirements]

[Type characteristics from Diataxis framework]
```

**Example (Reference)**:
```
Write a reference document for the `uberman install` command.

Requirements:
- Diataxis reference type: information-oriented, theoretical
- Include: Syntax, options table with types/defaults, exit codes, 3-5 examples
- Exclude: Tutorials, how-to steps (just facts)
- Target: Experienced user who forgot syntax
- Tone: Direct, neutral description
- Format: Easy to scan quickly
```

**Example (How-to)**:
```
Write a how-to guide for creating a custom app manifest.

Requirements:
- Diataxis how-to type: task-oriented, practical
- Assume: User knows TOML, understands Uberspace, familiar with Uberman
- Include: Prerequisites, clear steps, verification, troubleshooting table
- Exclude: TOML basics, Uberspace platform explanation
- Goal: User can create working manifest in 15 minutes
```

### Pattern 2: Incremental Documentation

**When**: Documenting as features ship (preferred approach for Uberman)
**Goal**: Document what exists, defer comprehensive coverage

**Prompt structure**:
```
Document [feature] incrementally:

1. What exists NOW (not what's planned)
2. Minimum viable documentation (defer comprehensive)
3. Focus on: [quick reference | design rationale | usage examples]
4. Add frontmatter: type, title, date
5. Stop when "good enough for now, safe enough to try"
```

**Example**:
```
Document the manifest schema incrementally:

1. What exists: TOML structure in internal/appinstallation/domain/manifest/
2. Minimum viable: Required fields only, 2-3 examples
3. Focus on: Quick reference for when I forget field names
4. Frontmatter: type=reference, title="Manifest Schema", date=2025-11-03
5. Defer: Optional fields, advanced patterns (add when needed)
```

### Pattern 3: Multi-Pass Editing

**When**: Reviewing or improving existing docs
**Goal**: Apply systematic quality checks

**Prompt structure**:
```
Review [document] using [pass name] only:

Technical accuracy pass:
- Follow own instructions, verify all commands work
- Test outputs match examples
- Check terminology consistency

Completeness pass:
- Identify missing prerequisites, TODOs, next steps
- Note version/platform limitations
- Document common errors

Structure pass:
- Evaluate header hierarchy and logic
- Check template adherence
- Verify navigation (links, prerequisites, next steps)

Clarity pass:
- Line-by-line brevity check
- Remove jargon or unnecessary words
- Use active voice, present tense
```

**Example**:
```
Review docs/cli-reference.md using technical accuracy pass only:

- Test every command example (copy-paste and run)
- Verify outputs match documented examples
- Check flag names against actual CLI (--help)
- Confirm exit codes are accurate
- Note any terminology inconsistencies

Do NOT comment on completeness, structure, or clarity yet.
```

### Pattern 4: Quality Gate Check

**When**: Before committing documentation
**Goal**: Ensure minimum quality standards

**Prompt structure**:
```
Apply quality gate checklist to [document]:

Functional quality:
- Accessible: Reading level appropriate, clear formatting
- Purposeful: Goal stated in title + first paragraph
- Findable: Keywords present, clear navigation
- Accurate: All procedures tested and work
- Complete: Prerequisites, steps, next steps defined

Structural quality:
- Clear: Unambiguous steps, errors defined
- Concise: No unnecessary words, brief paragraphs
- Consistent: Same terms = same meaning, style followed

Technical:
- Frontmatter complete: type, title, date
- Examples tested: All code samples work
- Links valid: No broken references
```

**Example**:
```
Apply quality gate checklist to docs/manifest-schema.md:

Functional:
- Is reading level appropriate for experienced user? (10th grade or lower)
- Does title + first paragraph state purpose clearly?
- Can user find this via search keywords?
- Are all TOML examples syntactically correct?
- Are prerequisites (TOML knowledge) and next steps (creating manifest) defined?

Structural:
- Are field descriptions unambiguous?
- Are there unnecessary explanations to remove?
- Do same terms mean same things throughout?

Technical:
- Check frontmatter: type=reference, title present, date present
- Test all TOML examples parse correctly
- Verify all links work (templates, how-to guides)
```

### Pattern 5: Scope Hammering

**When**: Documentation becoming too comprehensive or bloated
**Goal**: Cut scope to "good enough for now"

**Prompt structure**:
```
Apply scope hammer to [document]:

Remove:
- Nice-to-know information (not must-know)
- Detailed explanations (link to explanation docs instead)
- Exhaustive option lists (link to reference instead)

Defer:
- Future features not yet implemented
- Advanced patterns (add when users need them)
- Comprehensive troubleshooting (start with top 3 issues)

Simplify:
- Reduce examples from N to 3-5 most common
- Replace paragraphs with bullet lists
- Collapse detailed steps into single commands

Keep:
- Information that serves immediate purpose
- Minimum to accomplish the stated goal
- What you'd want when returning after 3 months
```

**Example**:
```
Apply scope hammer to apps/wordpress/README.md:

Remove:
- History of WordPress (nice-to-know)
- Detailed Uberspace platform explanation (link to UBERSPACE_INTEGRATION.md)
- Complete list of configuration options (link to WordPress docs)

Defer:
- Advanced caching configuration (add when you implement it)
- Custom theme installation (not needed for basic install)
- Multisite setup (future enhancement)

Simplify:
- Reduce manual installation steps from 15 to 8 key steps
- Replace troubleshooting paragraphs with 3-row issue/solution table
- Collapse post-install configuration into 3 bullet points

Keep:
- `uberman install wordpress` command
- 8 key manual steps (preserves knowledge of what's automated)
- Manifest design rationale (why these database/port choices)
- Top 3 troubleshooting issues
```

---

## Quality Workflow Summary

### Multi-Pass Editing (Brief)

Apply passes in order:

**Pass 1: Technical Accuracy**
- Follow your own instructions
- Test all commands and code samples
- Verify outputs match examples

**Pass 2: Completeness**
- Fill [TODO] markers
- Add prerequisites and next steps
- Document version applicability

**Pass 3: Structure**
- Check header hierarchy
- Verify template adherence
- Ensure navigation works

**Pass 4: Clarity & Brevity**
- Remove unnecessary words
- Use active voice
- Break long paragraphs

**See `docs/reference/docs-for-developers-workflow.md` for complete checklists**

### Plussing Rule (Code Review)

From Pixar's creative process:

> Only criticize if you add constructive suggestion

**Application**:
- ❌ "This section is unclear"
- ✅ "This section is unclear because X. Suggest: Rewrite as '...'"

---

## Templates

Quick links to copy-paste templates:

- **Reference**: CLI commands, manifest schema, error codes
- **How-to**: Task guides with prerequisites, steps, verification
- **Tutorial**: Step-by-step learning with examples
- **Explanation/ADR**: Design decisions with rationale
- **App Install**: What Uberman automates + manual process
- **README**: Project overview with navigation

**See `docs/reference/documentation-templates.md` for all templates**

---

## Uberman Project Context

### Documentation Approach

- **Strategy**: Incremental (document as features ship)
- **Quality bar**: "Good enough for now, safe enough to try"
- **Style**: Direct, concise, no handholding
- **Audience**: You (15-year veteran) + experienced sysadmins/SREs

### Organization

**App docs**: Co-located with manifests
```
apps/<app>/
├── manifest.toml
└── README.md (tutorial-style)
```

**ADRs**: Separate from user docs
```
decisions/
├── template.md
└── NNNN-title.md
```

**User docs**: By need, not Diataxis folders
```
docs/
├── cli-reference.md (type: reference)
├── manifest-schema.md (type: reference)
├── creating-manifests.md (type: howto)
└── troubleshooting.md (type: howto)
```

### Frontmatter (Minimal)

```yaml
---
title: [Title]
type: tutorial|howto|reference|explanation
date: YYYY-MM-DD
---
```

No owner/freshness tracking needed (personal tool).

---

## Anti-Patterns to Avoid

**When instructing Claude**:
- ❌ "Write comprehensive documentation for..."
- ❌ "Explain everything about..."
- ❌ "Create complete guide covering all aspects..."
- ❌ "Write tutorial for beginners..." (you're experienced)

**Content mixing**:
- ❌ Tutorial steps in reference docs
- ❌ Reference tables in tutorials
- ❌ How-to steps in explanations
- ❌ Design rationale in how-to guides

**Scope creep**:
- ❌ Documenting unimplemented features
- ❌ Comprehensive coverage (defer advanced topics)
- ❌ Explaining basics (link to external resources)
- ❌ Tutorial bloat (you don't need hand-holding)

---

## Quick Reference

### Frameworks

- **Diataxis**: docs/reference/diataxis-framework.md
- **Docs for Developers**: docs/reference/docs-for-developers-workflow.md
- **Templates**: docs/reference/documentation-templates.md

### Common Questions

**Q: What content type should this be?**
A: Use decision tree above or read `docs/reference/diataxis-framework.md`

**Q: How do I instruct Claude to write good docs?**
A: Use Pattern 1 (Type-Aware Creation) with specific requirements

**Q: How do I review documentation quality?**
A: Use Pattern 3 (Multi-Pass Editing) with 4 passes

**Q: How do I avoid over-documenting?**
A: Use Pattern 5 (Scope Hammering) and "good enough for now" principle

**Q: What template should I use?**
A: See `docs/reference/documentation-templates.md` for all templates

---

## Success Metrics

Documentation is "good enough" when:

1. ✅ Can you use it after 3 months away?
2. ✅ Does it answer "How do I...?" questions?
3. ✅ Does it explain "Why did I...?" decisions?
4. ✅ Can you instruct Claude to create similar docs?

---

**Last updated**: 2025-11-03
**Frameworks**: Diataxis (diataxis.fr) + Docs for Developers (Bhatti et al., 2021)
