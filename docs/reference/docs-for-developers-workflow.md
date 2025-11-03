---
title: Docs for Developers Workflow Reference
type: reference
date: 2025-11-03
---

# Docs for Developers Workflow Reference

*Quality framework and editing workflow for technical documentation*

**Source**: "Docs for Developers: An Engineer's Field Guide to Technical Writing"
**Authors**: Jared Bhatti, Zachary Sarah Corleissen, Jen Lambourne, David Nunez, Heidi Waterhouse
**Publisher**: Apress, 2021

---

## Core Principle

> "A document is good when it fulfills its purpose."

### Two Quality Dimensions

**Functional Quality** (Does it work? Does it achieve its goal?)
- Accessible
- Purposeful
- Findable
- Accurate
- Complete

**Structural Quality** (Is it well-written?)
- Clear
- Concise
- Consistent

### Priority Rule

**Functional quality > Structural quality**

- Low functional + high structural = POOR overall quality
- High functional + okay structural = GOOD overall quality

A poorly written doc that accomplishes its goal beats a beautifully written doc that doesn't help users.

---

## Functional Quality Framework

### 1. Accessible

**Definition**: Can users access and understand the content?

**Dimensions**:
- **Language/Reading level**: Target 10th grade or lower
- **Screen reader compatibility**: Alt text for images, captions for videos
- **Clear formatting**: Headers, lists, visual hierarchy

**How to check**:
- Flesch-Kincaid Grade Level (most tools target: 10th grade)
- Automated Readability Index
- Coleman-Liau Index
- Hemingway Editor (hemingwayapp.com)
- WCAG compliance check

**Checklist**:
- [ ] Reading level ≤ 10th grade
- [ ] Images have alt text
- [ ] Videos have captions/subtitles
- [ ] Headers create clear hierarchy
- [ ] Lists used for scanability

### 2. Purposeful

**Definition**: Does the document clearly state and achieve its goal?

**Metrics**:
- **Time to Hello World (TTHW)**: Time from start to first success
- **Goal accomplishment**: Can users do what doc promises?

**How to check**:
- Title clearly indicates what document does
- First paragraph explicitly states purpose
- User can accomplish stated goal by following doc

**Checklist**:
- [ ] Title describes what document helps accomplish
- [ ] First paragraph states purpose clearly
- [ ] Document delivers on promise
- [ ] Success criteria defined
- [ ] TTHW measured (if applicable)

### 3. Findable

**Definition**: Can users locate the information they need?

> "The real findability problem is how to get readers from the wrong place deep within your content to the right place deep within your content." — Mark Baker

**Techniques**:
- Standardized search keywords
- Every page = potential entry point ("Every page is page one")
- Clear navigation (breadcrumbs, sidebars, prerequisites, next steps)
- Escape hatches (callouts for users in wrong place)

**Checklist**:
- [ ] Keywords consistent and predictable
- [ ] Page works as standalone entry point
- [ ] Prerequisites link to prior knowledge
- [ ] Next steps link to follow-up content
- [ ] "Not in right place?" escape hatches provided

### 4. Accurate

**Definition**: Is the information correct and current?

**How to check**:
- Test all procedures personally
- Run all code samples/commands
- Verify outputs match examples
- Check API calls return documented responses
- Confirm version applicability

**Checklist**:
- [ ] All procedures tested and work
- [ ] Code samples run without errors
- [ ] Outputs match documented examples
- [ ] API responses accurate
- [ ] Version/platform limitations noted
- [ ] No outdated information

### 5. Complete

**Definition**: Does documentation cover what users need?

> "Completeness is not the same as telling people everything." — Ch 4, Par 34

**What to include**:
- Prerequisites (what users need first)
- All steps for task
- Next steps (where to go after)
- Version applicability
- Common errors/troubleshooting

**Checklist**:
- [ ] Prerequisites explicitly listed
- [ ] No missing steps in procedures
- [ ] Next steps defined
- [ ] Version/platform applicability stated
- [ ] Common errors documented
- [ ] No [TODO] or [TBD] markers remain

---

## Structural Quality Framework

### 1. Clear

**Definition**: Is the content unambiguous and easy to understand?

**Techniques**:
- Well-ordered headers (logical flow)
- Unambiguous steps (one action per step)
- Errors explicitly defined
- Technical terms explained or linked

**Checklist**:
- [ ] Headers create logical flow
- [ ] Each step has single, clear action
- [ ] Error messages explained
- [ ] Technical jargon defined
- [ ] Examples illustrate concepts

### 2. Concise

**Definition**: Is the content brief but comprehensive?

> "A sentence should contain no unnecessary words, a paragraph no unnecessary sentences, for the same reason that a drawing should have no unnecessary lines and a machine no unnecessary parts." — William Strunk Jr.

**Techniques**:
- Remove unnecessary words
- Paragraphs < 5 sentences
- Use lists instead of long paragraphs
- Active voice ("Run the command" not "The command should be run")
- Present tense ("The system returns" not "The system will return")

**Checklist**:
- [ ] No unnecessary words
- [ ] Paragraphs < 5 sentences
- [ ] Lists used for scanability
- [ ] Active voice throughout
- [ ] Present tense used

### 3. Consistent

**Definition**: Do terms and formatting remain uniform?

**Techniques**:
- Same terms = same meaning throughout
- Follow style guide (Google, Microsoft, etc.)
- Consistent formatting (code blocks, emphasis, links)
- Uniform voice and tone

**Checklist**:
- [ ] Terms used consistently
- [ ] Style guide followed
- [ ] Formatting uniform
- [ ] Voice/tone consistent throughout

---

## Multi-Pass Editing Workflow

**Strategy**: Separate passes for different quality dimensions

### Why Multi-Pass?

Trying to check everything at once is overwhelming. Focus on one dimension per pass.

### The Four Passes

Apply in order for best results.

---

### Pass 1: Technical Accuracy

**Focus**: Does it work?

**Checklist**:
- [ ] Follow your own instructions (copy-paste commands and run them)
- [ ] Test all code samples (verify they execute without errors)
- [ ] Check outputs match examples (exact output comparison)
- [ ] Verify API calls work (test actual endpoints)
- [ ] Confirm terminology consistent (same terms = same meaning)
- [ ] Test across platforms/versions (if applicable)
- [ ] Flag data loss risks with warnings
- [ ] Note any steps that don't work

**Time commitment**: Usually longest pass (actually testing)

**If errors found**: Fix them before proceeding to next pass

---

### Pass 2: Completeness

**Focus**: Any gaps?

**Checklist**:
- [ ] Fill all [TODO]/[TBD] markers
- [ ] Add missing prerequisites
- [ ] Define next steps
- [ ] Document version/platform applicability
- [ ] Add common errors to troubleshooting
- [ ] Check for missing steps (did you skip "obvious" ones?)
- [ ] Verify examples cover main use cases
- [ ] Ensure success criteria defined

**Remember**: Complete ≠ telling everything
- Include what's necessary for purpose
- Link to other docs for detail
- Defer advanced topics

**If gaps found**: Fill them or explicitly defer ("Advanced configuration: [link]")

---

### Pass 3: Structure

**Focus**: Logical flow?

**Checklist**:
- [ ] Title short and specific (describes what doc does)
- [ ] Headers logically ordered (progressive disclosure)
- [ ] Follows template structure (if applicable)
- [ ] Prerequisites at top
- [ ] Next steps at bottom
- [ ] All links work (no 404s)
- [ ] Navigation clear (table of contents, breadcrumbs)
- [ ] Escape hatches for wrong-page visitors
- [ ] Images/diagrams add value (not decoration)

**If structure issues found**: Reorganize before clarity pass

---

### Pass 4: Clarity & Brevity

**Focus**: Easy to understand?

**Checklist**:
- [ ] Line-by-line review (read every sentence)
- [ ] Remove unnecessary words ("In order to" → "To")
- [ ] Break long paragraphs (< 5 sentences)
- [ ] Use lists instead of prose (when listing items)
- [ ] Active voice ("Run command" not "Command should be run")
- [ ] Present tense ("Returns value" not "Will return value")
- [ ] Remove jargon or explain it
- [ ] Check for biased/exclusionary language
- [ ] Run grammar/spelling check

**Common brevity improvements**:
- "In order to" → "To"
- "It should be noted that" → [delete]
- "At this point in time" → "Now"
- "Due to the fact that" → "Because"

**If clarity issues found**: Rewrite and re-check

---

## Plussing Rule (Code Review for Docs)

From Pixar's creative process:

> "You may only criticize an idea if you also add a constructive suggestion."

### Application to Documentation Review

**Bad feedback** (criticism without suggestion):
- ❌ "This section is unclear"
- ❌ "This doesn't make sense"
- ❌ "Needs improvement"

**Good feedback** (criticism + suggestion):
- ✅ "This section is unclear because it uses 'instance' and 'installation' interchangeably. Suggest: Use 'instance' consistently and define it in prerequisites."
- ✅ "Step 3 doesn't make sense because prerequisites aren't listed. Suggest: Add prerequisites section before steps showing required dependencies."
- ✅ "Needs improvement: Examples don't show expected output. Suggest: Add '**Output**:' section after each command showing what users should see."

### Benefits

- **Actionable**: Reviewer provides solution, not just problem
- **Collaborative**: Focuses on improvement, not criticism
- **Efficient**: Author has clear path forward

---

## Planning Questions

Before writing ANY documentation, answer these:

### Audience & Purpose
- [ ] Who is the target audience? (Beginners, experienced users, specific role?)
- [ ] What should they take away? (Main learning/accomplishment?)
- [ ] What prerequisites do they need? (Prior knowledge, tools, access?)

### Content & Scope
- [ ] What are the top 3-5 use cases? (Most common scenarios?)
- [ ] What are known friction points? (Where do users struggle?)
- [ ] What version/platform does this apply to? (Limitations?)

### Success Metrics
- [ ] How will I measure success? (TTHW, support reduction, user feedback?)
- [ ] When is this documentation "good enough"? (Quality bar?)

### Example (Uberman Project)

**Documenting**: `uberman install` command

**Answers**:
- Audience: You (15-year veteran), experienced sysadmins
- Takeaway: Install app with single command, understand what's automated
- Prerequisites: Uberspace account, basic Linux CLI knowledge
- Top use cases: (1) Install WordPress, (2) Install Nextcloud, (3) Install custom app
- Friction: Database naming, port conflicts, supervisord service management
- Applies to: Uberspace v7+, uberman v1.0+
- Success: Can use command after 3 months without rereading docs
- Good enough: Syntax, options, 3 examples, error codes

---

## Quality Metrics

### For Personal Tools (Uberman Context)

Simpler metrics for personal use:

1. **Can you use it after 3 months away?** (primary test)
2. **Does it answer "How do I...?" questions?** (functional)
3. **Does it explain "Why did I...?" decisions?** (rationale)

### For Public Documentation

If sharing more broadly:

**Documentation-Specific Metrics**:
- **Time to Hello World (TTHW)**: Time to first success
- **Page views**: Which docs are accessed most?
- **Time on page**: Long = detailed reading or confusion?
- **Bounce rate**: Do users leave immediately?
- **Search keywords**: What are users looking for?
- **Support issues**: How many docs-related issues?
- **Link validation**: Any broken links?

**Using Metrics Effectively** (Tips from book):
1. **Make a plan**: What will you measure and why?
2. **Establish baseline**: Before/after comparison
3. **Consider context**: High views on error docs = problem, not success
4. **Use clusters**: Multiple metrics together, not single metrics
5. **Mix qualitative + quantitative**: Numbers + user feedback

---

## Maintenance Strategies

### Alignment Techniques

**1. Align with Release Process**
- Documentation required before feature approval
- Track in same system as code (issues/PRs)
- Block merges if docs missing

**Example (Kubernetes)**:
- Enhancement proposals require docs assessment
- Docs tracked in same spreadsheet as features
- All approved together on release date

**2. Assign Document Owners**
- Use CODEOWNERS file
- Metadata in docs: `<!-- Owner: alice@example.com -->`
- Owners review and approve changes

**3. Reward Maintenance**
- Recognition for doc contributions
- Built into performance expectations
- Not "extra" work

### Automation to Eliminate Toil

**Toil Definition** (from SRE book):
> Manual, repetitive, automatable, tactical, devoid of enduring value, and scales linearly as a service grows.

**Automation Tools**:

**1. Content Freshness Checks**
```html
<!-- Freshness: {owner: "alice" reviewed: 2025-01-15} -->
```
- Automated reminders after 6 months
- Owner reviews and updates date
- Or re-writes if outdated

**2. Link Checkers**
- Run in CI/CD before publication
- Catch 404s before users do
- Example tools: linkchecker, htmltest

**3. Prose Linters**
- Vale: Style guide enforcement
- Custom dictionaries (project terms)
- Inclusive language checking
- Grammar/spelling

**4. Reference Doc Generators**
- OpenAPI → API reference
- Javadoc → Code documentation
- `--help` → CLI reference
- Always in sync with code

### Deprecation Process

**Steps**:
1. **Mark as deprecated** - Add callout to doc
2. **Announce** - Release notes, user communications
3. **Provide migration guide** - Before announcement
4. **Delete when safe** - After migration period
5. **Set up redirects** - Prevent 404s

**Deprecation Callout Format**:
```markdown
> ⚠️ **DEPRECATED**
>
> The Audio API was deprecated on 2025-08-20.
> It has been replaced by the Multimedia API.
>
> **Migration guide**: [link]
```

---

## Document Templates

Quick links to all templates:

See `docs/reference/documentation-templates.md` for complete, copy-paste templates:

- **Reference** template
- **How-to Guide** template
- **Tutorial** template
- **Explanation/ADR** template
- **App Installation** template
- **README** template
- **Release Notes** template

---

## Further Reading

### Book

"Docs for Developers: An Engineer's Field Guide to Technical Writing"
- Jared Bhatti, Zachary Sarah Corleissen, Jen Lambourne, David Nunez, Heidi Waterhouse
- Apress, 2021

### Related Frameworks

- **Diataxis**: See `docs/reference/diataxis-framework.md`
- **Shape Up**: See `plans/README.md` (cooldown for reflection)
- **EARS**: See `docs/EARS_GUIDE.md` (requirements format)

### Style Guides

- Google Developer Style Guide
- Microsoft Style Guide
- Red Hat Style Guide

### Tools

- **Readability**: Hemingway Editor (hemingwayapp.com)
- **Prose linting**: Vale (vale.sh)
- **Link checking**: linkchecker, htmltest
- **Accessibility**: WAVE, axe DevTools
