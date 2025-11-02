# Requirements Analysis Checklist

Use this checklist to assess requirement quality systematically.

## Requirement Under Review

**ID**: [REQ-___]
**Text**: ________________________________________________
**Author**: ________________
**Date**: ________________

---

## EARS Pattern Compliance

- [ ] **Pattern Identified**: ☐ Ubiquitous ☐ State-Driven (While) ☐ Event-Driven (When) ☐ Optional (Where) ☐ Unwanted (If/Then)
- [ ] **Uses "shall"** (not should, must, will, may, can)
- [ ] **System name specified** (not "the system" or pronoun)
- [ ] **Trigger/condition clear** (if applicable to pattern)

---

## Quality Scores (1-5 scale)

### Clarity
- [ ] **5** - Crystal clear, no ambiguity
- [ ] **4** - Clear with minor non-critical ambiguities
- [ ] **3** - Generally clear but some ambiguous terms
- [ ] **2** - Multiple interpretations possible
- [ ] **1** - Unclear or contradictory

**Issues noted**: ________________________________________________

### Testability
- [ ] **5** - Fully testable with explicit criteria
- [ ] **4** - Testable with minor details to define
- [ ] **3** - Can derive tests but requires interpretation
- [ ] **2** - Mainly subjective, hard to verify
- [ ] **1** - Untestable

**Test approach**: ________________________________________________

### Completeness
- [ ] **5** - All triggers, conditions, responses, criteria included
- [ ] **4** - Nearly complete, minor details missing
- [ ] **3** - Some important info implicit
- [ ] **2** - Critical information missing
- [ ] **1** - Fundamentally incomplete

**Missing information**: ________________________________________________

---

## Language Check

### Weak Terms (flag if present)
- [ ] Vague quantifiers: fast, slow, quick, large, small, many, few
- [ ] Subjective terms: user-friendly, intuitive, easy, simple
- [ ] Unclear terms: adequate, sufficient, appropriate, as needed
- [ ] Open-ended: etc., and so on, for example
- [ ] Weak modals: should, could, would, might, may

**Terms to replace**: ________________________________________________

### Structure
- [ ] **Active voice** (not passive)
- [ ] **Present tense**
- [ ] **Simple sentence** (not compound)
- [ ] **Specific nouns** (not pronouns)
- [ ] **Measurable quantities** (with units)

---

## Content Assessment

### Atomic
- [ ] **Addresses only one thing**
- [ ] **Can be satisfied with single solution**
- [ ] **Not compound** (no multiple "and"/"or")

**If compound, split into**: ________________________________________________

### Necessary
- [ ] **Adds value**
- [ ] **Stakeholder need identified**
- [ ] **Not gold-plating**

**Justification**: ________________________________________________

### Feasible
- [ ] **Technically possible**
- [ ] **Within constraints** (time, budget, technology)
- [ ] **No known blockers**

**Concerns**: ________________________________________________

### Consistent
- [ ] **No contradictions** with other requirements
- [ ] **Terminology consistent** with glossary
- [ ] **Abstraction level appropriate**

**Conflicts with**: ________________________________________________

---

## Specific Checks

### Triggers and Conditions
- [ ] **Trigger clearly defined** (for event-driven)
- [ ] **Condition clearly stated** (for state-driven)
- [ ] **Optional feature specified** (for optional)
- [ ] **Error condition clear** (for unwanted behavior)

### System Response
- [ ] **Action verb clear** (display, send, validate, etc.)
- [ ] **Response complete** (what, where, how)
- [ ] **Success criteria defined**
- [ ] **Timing specified** (if performance-sensitive)

### Measurability
- [ ] **Numbers have units** (seconds, MB, items, etc.)
- [ ] **Ranges bounded** (min/max values)
- [ ] **Percentages specified** (if applicable)
- [ ] **Thresholds defined**

---

## INCOSE Quality Criteria

- [ ] **Clarity** - Single interpretation
- [ ] **Completeness** - All info included
- [ ] **Consistency** - No conflicts
- [ ] **Correctness** - Represents need accurately
- [ ] **Testability** - Can be verified
- [ ] **Feasibility** - Can be implemented
- [ ] **Unambiguity** - Clear meaning
- [ ] **Atomic** - One thing only
- [ ] **Necessity** - Adds value
- [ ] **Conciseness** - No unnecessary words

---

## Overall Assessment

**Status**:
- [ ] ✅ **APPROVED** - Meets quality standards
- [ ] ⚠️ **APPROVED WITH MINOR CHANGES** - Minor improvements needed
- [ ] ❌ **REJECTED** - Major rework required

**Priority of Issues**:
- [ ] 🔴 **HIGH** - Critical issues blocking approval
- [ ] 🟡 **MEDIUM** - Important improvements recommended
- [ ] 🟢 **LOW** - Nice-to-have improvements

---

## Improvement Recommendations

**Suggested Rewrite** (if applicable):
________________________________________________
________________________________________________
________________________________________________

**Rationale for Changes**:
________________________________________________
________________________________________________

**Additional Requirements Needed** (if split):
________________________________________________
________________________________________________

---

## Reviewer Information

**Reviewer Name**: ________________
**Review Date**: ________________
**Next Review**: ________________

---

*This checklist aligns with INCOSE and IEEE 29148 standards for requirements quality.*
