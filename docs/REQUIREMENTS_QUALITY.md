# Requirements Quality Framework

**INCOSE and IEEE 29148 Standards-Based Assessment**

## Table of Contents

- [Quality Criteria Overview](#quality-criteria-overview)
- [INCOSE 42-Rule Framework](#incose-42-rule-framework)
- [Quality Scoring System](#quality-scoring-system)
- [Ambiguity Detection](#ambiguity-detection)
- [Testability Guidelines](#testability-guidelines)
- [Review Checklists](#review-checklists)

---

## Quality Criteria Overview

### The 10 Essential Characteristics

High-quality requirements must exhibit these fundamental properties:

#### 1. **Clarity**
- **Definition**: Requirement has one and only one interpretation
- **Test**: Can multiple people read it and understand the same thing?
- **Bad**: "The system should be user-friendly"
- **Good**: "The application shall display feedback within 200ms for every user action"

#### 2. **Completeness**
- **Definition**: All necessary information is included
- **Test**: Can you implement and test without asking questions?
- **Bad**: "Validate email addresses"
- **Good**: "When a user enters an email in the registration form, the system shall validate format per RFC 5322 and shall display 'Invalid email format' if validation fails"

#### 3. **Consistency**
- **Definition**: No conflicts with other requirements
- **Test**: Can all requirements be satisfied simultaneously?
- **Bad**: REQ-001 "Response time shall be under 1 second" + REQ-002 "System shall perform comprehensive virus scan on all uploads"
- **Good**: Ensure performance requirements account for security scanning overhead

#### 4. **Correctness**
- **Definition**: Accurately represents stakeholder needs
- **Test**: Do stakeholders confirm this is what they need?
- **Method**: Validation sessions, prototyping, acceptance criteria

#### 5. **Testability**
- **Definition**: Can be verified objectively through testing or inspection
- **Test**: Can you write a pass/fail test case?
- **Bad**: "System shall be reliable"
- **Good**: "System shall maintain 99.9% uptime measured monthly"

#### 6. **Feasibility**
- **Definition**: Technically and economically achievable
- **Test**: Can this be built with available technology and budget?
- **Bad**: "System shall read users' minds"
- **Good**: "System shall predict user preferences based on historical behavior with 80% accuracy"

#### 7. **Unambiguity**
- **Definition**: Admits only one interpretation
- **Test**: No vague terms, pronouns without clear referents, or ambiguous quantifiers
- **Bad**: "Fast response time"
- **Good**: "Response time under 2 seconds"

#### 8. **Atomic**
- **Definition**: Addresses exactly one thing
- **Test**: Can this be satisfied with a single solution?
- **Bad**: "System shall validate and sanitize input and log errors"
- **Good**: Split into three separate requirements

#### 9. **Necessity**
- **Definition**: Requirement adds value and is needed
- **Test**: What happens if we don't implement this?
- **Avoid**: Gold-plating, nice-to-haves without justification

#### 10. **Conciseness**
- **Definition**: Brief without sacrificing clarity
- **Test**: Can you remove words without losing meaning?
- **Balance**: Completeness vs. brevity

---

## INCOSE 42-Rule Framework

### Essential Rules (Core Quality)

**Rule 1**: Use "shall" for mandatory requirements
**Rule 2**: One requirement per statement
**Rule 3**: Avoid weak words (adequate, easy, fast, sufficient, user-friendly)
**Rule 4**: Use active voice
**Rule 5**: Avoid negative requirements when possible
**Rule 6**: Specify measurable quantities with units
**Rule 7**: Avoid pronouns; use specific names
**Rule 8**: Define acronyms and abbreviations
**Rule 9**: Use consistent terminology
**Rule 10**: Avoid ambiguous terms and phrases

### Structural Rules

**Rule 11**: Atomic - one requirement per statement
**Rule 12**: Complete - all necessary information included
**Rule 13**: Consistent - no contradictions
**Rule 14**: Traceable - linked to sources
**Rule 15**: Numbered uniquely
**Rule 16**: Organized logically
**Rule 17**: Prioritized
**Rule 18**: Labeled with type (functional, performance, etc.)

### Content Rules

**Rule 19**: State **what**, not **how** (for functional requirements)
**Rule 20**: Include success criteria
**Rule 21**: Include failure handling
**Rule 22**: Specify timing where relevant
**Rule 23**: Specify accuracy/precision where relevant
**Rule 24**: Specify capacity/volume where relevant
**Rule 25**: Include error tolerances

### Language Rules

**Rule 26**: Use simple sentences
**Rule 27**: Use present tense
**Rule 28**: Avoid combining requirements with "and" or "or"
**Rule 29**: Avoid escape clauses (TBD, TBC, etc.)
**Rule 30**: Define all variables

### Testability Rules

**Rule 31**: Requirement must be verifiable
**Rule 32**: Include acceptance criteria
**Rule 33**: Identify verification method (test, inspection, analysis, demonstration)
**Rule 34**: Ensure objectivity (no subjective judgments)

### Additional Quality Rules

**Rule 35**: Requirements must be feasible
**Rule 36**: Requirements must be unambiguous
**Rule 37**: Requirements must be necessary
**Rule 38**: Requirements must be implementation-independent (functional reqs)
**Rule 39**: Requirements must be design-independent (functional reqs)
**Rule 40**: Avoid over-specification
**Rule 41**: Include all stakeholder needs
**Rule 42**: Maintain requirement dependencies

---

## Quality Scoring System

### 1-5 Scale for Each Dimension

Use this rubric to score requirements:

#### Clarity Score

| Score | Description | Criteria |
|-------|-------------|----------|
| **5** | Crystal clear | No possibility of misinterpretation; all terms defined |
| **4** | Clear | Minor ambiguities in non-critical areas only |
| **3** | Mostly clear | Some ambiguous terms but generally understandable |
| **2** | Unclear | Multiple interpretations possible in critical areas |
| **1** | Very unclear | Fundamentally unclear, contradictory, or meaningless |

**Example 5/5**: "When the user clicks Submit, the web form shall validate all fields and shall display results within 500 milliseconds"
**Example 1/5**: "The system should work properly"

#### Testability Score

| Score | Description | Criteria |
|-------|-------------|----------|
| **5** | Fully testable | Objective test criteria explicitly stated with measurable values |
| **4** | Testable | Test approach clear; minor details need definition |
| **3** | Somewhat testable | Can derive test cases but requires interpretation |
| **2** | Hard to test | Mainly subjective criteria; difficult to verify objectively |
| **1** | Untestable | Impossible to verify or test |

**Example 5/5**: "The mobile app shall start within 3 seconds on iPhone 12 or newer"
**Example 1/5**: "The app shall be performant"

#### Completeness Score

| Score | Description | Criteria |
|-------|-------------|----------|
| **5** | Complete | All triggers, conditions, responses, and criteria specified |
| **4** | Nearly complete | Minor details missing but intent is complete |
| **3** | Mostly complete | Some important information implicit or assumed |
| **2** | Incomplete | Critical information missing; significant assumptions required |
| **1** | Very incomplete | Core elements missing; cannot implement |

**Example 5/5**: "If password validation fails 3 times, then the system shall lock the account for 15 minutes and shall email the account owner"
**Example 1/5**: "Handle login errors"

### Overall Quality Score

**Calculate**: Average of Clarity, Testability, and Completeness scores

**Quality Bands**:
- **High Quality** (4.0-5.0): Ready for implementation
- **Medium Quality** (2.5-3.9): Needs refinement
- **Low Quality** (1.0-2.4): Requires major rework

---

## Ambiguity Detection

### Types of Ambiguity

#### 1. **Lexical Ambiguity**
Words with multiple meanings

**Examples**:
- "bank" (financial institution vs. river bank)
- "light" (illumination vs. not heavy)
- "can" (ability vs. container)

**Solution**: Use context-specific terms or define in glossary

#### 2. **Syntactic Ambiguity**
Sentence structure allows multiple interpretations

**Example**: "The system shall notify users and administrators of errors"

Does this mean:
- A) Notify (users and administrators) of errors, OR
- B) Notify users of (administrators of errors)?

**Solution**: Restructure for clarity or use bullet points

#### 3. **Semantic Ambiguity**
Vague or subjective terms

**Examples**:
- "fast", "slow", "quick", "efficient"
- "user-friendly", "intuitive", "simple"
- "high quality", "robust", "reliable"

**Solution**: Replace with measurable criteria

#### 4. **Anaphoric Ambiguity**
Pronouns with unclear referents

**Example**: "When the user logs in, it shall verify credentials"

What is "it"? The system? The login process?

**Solution**: Use specific nouns, avoid pronouns

### Weak Terms to Avoid

**Vague Quantifiers**:
- ❌ fast, quick, slow, rapid
- ❌ large, small, big, tiny
- ❌ many, few, several, some
- ❌ often, rarely, occasionally

**Vague Qualifiers**:
- ❌ adequate, sufficient, appropriate
- ❌ easy, simple, straightforward
- ❌ user-friendly, intuitive
- ❌ robust, reliable, stable (without quantification)

**Escape Clauses**:
- ❌ TBD (to be determined)
- ❌ TBC (to be confirmed)
- ❌ etc., and so on
- ❌ as appropriate, as needed

**Weak Modal Verbs**:
- ❌ should, could, would, might, may

**Replace With**:
- ✅ Specific numbers with units
- ✅ Measurable criteria
- ✅ Objective thresholds
- ✅ "shall" for mandatory requirements

---

## Testability Guidelines

### Can You Write a Test Case?

For every requirement, you should be able to write:

**1. Test Setup**: Preconditions and test data
**2. Test Steps**: Actions to perform
**3. Expected Result**: What constitutes pass/fail
**4. Measurement Method**: How to verify

### Example: Testable vs. Untestable

**❌ Untestable**:
"The system shall be easy to use"

Can't write objective test case - "easy" is subjective

**✅ Testable**:
"New users shall be able to create an account and complete their first purchase within 5 minutes without consulting help documentation"

**Test Case**:
1. **Setup**: Recruit 10 users with no prior system experience
2. **Steps**: Ask each user to create account and purchase item
3. **Measurement**: Time each user; success = account created + purchase completed
4. **Pass Criteria**: ≥ 8 out of 10 users succeed within 5 minutes

### Verification Methods

| Method | Description | Example Requirements |
|--------|-------------|---------------------|
| **Test** | Exercise system with inputs, observe outputs | Functional behavior, performance |
| **Inspection** | Visual examination | Color, size, label text |
| **Analysis** | Mathematical or logical evaluation | Capacity calculations, algorithm correctness |
| **Demonstration** | Show capability in operational scenario | User workflow, integration |

Specify verification method in requirement or traceability matrix.

---

## Review Checklists

### Individual Requirement Review

Use this checklist for each requirement:

**Structure**:
- [ ] Has unique identifier
- [ ] Uses "shall" (not should/must/will/may)
- [ ] Follows EARS pattern (or justified exception)
- [ ] Is atomic (one requirement)
- [ ] Is complete sentence

**Content**:
- [ ] Specifies system name explicitly
- [ ] Includes trigger/condition (if applicable)
- [ ] Describes system response clearly
- [ ] Includes measurable criteria
- [ ] Uses units for quantities
- [ ] Defines all acronyms/abbreviations

**Quality**:
- [ ] Is unambiguous (only one interpretation)
- [ ] Is testable (can write test case)
- [ ] Is complete (all necessary info included)
- [ ] Is necessary (adds value)
- [ ] Is feasible (can be implemented)
- [ ] Is consistent (no contradictions)

**Language**:
- [ ] Uses active voice
- [ ] Avoids weak/vague terms
- [ ] Avoids pronouns (uses specific nouns)
- [ ] Uses present tense
- [ ] Is concise (no unnecessary words)

### Requirements Set Review

For the complete set of requirements:

**Coverage**:
- [ ] All stakeholder needs addressed
- [ ] All system functions covered
- [ ] All quality attributes specified
- [ ] All interfaces defined
- [ ] All error scenarios handled

**Consistency**:
- [ ] No contradictions between requirements
- [ ] Terminology used consistently
- [ ] Numbering scheme logical
- [ ] Abstraction level appropriate

**Traceability**:
- [ ] Requirements linked to sources
- [ ] Requirements linked to design
- [ ] Requirements linked to tests
- [ ] Dependencies identified

**Organization**:
- [ ] Logical grouping/categorization
- [ ] Appropriate hierarchy
- [ ] Priorities assigned
- [ ] Owner/responsible party identified

---

## Quality Metrics

### Requirement Set Metrics

Track these metrics to assess overall quality:

**Compliance Metrics**:
- % of requirements using "shall"
- % of requirements following EARS patterns
- % of requirements with measurable criteria

**Quality Metrics**:
- Average clarity score
- Average testability score
- Average completeness score
- % high quality (avg ≥4.0)
- % low quality (avg <2.5)

**Issue Metrics**:
- Number of weak terms detected
- Number of compound requirements
- Number of ambiguities found
- Number of untestable requirements

**Process Metrics**:
- Requirements churn rate (% changed)
- Review defect density (issues per requirement)
- Time to resolve issues

---

## Tools and Automation

### Automated Analysis

**Available Tools**:
- **ARM**: Automated Requirements Measurement
- **QuARS**: Quality Analyzer for Requirements Specifications
- **RETA**: Requirements Template Analyzer
- **RUBRIC**: Boilerplate conformance checking

**What They Check**:
- Weak term detection
- Passive voice identification
- Ambiguous terminology
- Structural patterns
- Complexity metrics

### Manual Review

**When to use human review**:
- Domain expertise needed
- Context-dependent interpretation
- Stakeholder validation
- Design trade-offs
- Critical safety requirements

**Best Practice**: Combine automated checks with expert review

---

## Standards Reference

### IEEE 29148:2018

Key requirements per this standard:

**Requirement Characteristics**:
- Necessary
- Implementation-free (functional)
- Unambiguous
- Consistent
- Complete
- Singular
- Feasible
- Traceable
- Verifiable

### ISO/IEC 25010

Quality model for systems and software:

**Product Quality**:
- Functional suitability
- Performance efficiency
- Compatibility
- Usability
- Reliability
- Security
- Maintainability
- Portability

Map non-functional requirements to these quality characteristics.

---

*This quality framework is part of the Uberman Requirements Engineering toolkit.*
