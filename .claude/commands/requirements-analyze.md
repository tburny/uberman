---
name: requirements-analyze
description: Analyze requirement quality and EARS compliance
version: 1.0.0
tags: [requirements, ears, analysis, quality]
---

Analyze the provided requirement for quality and EARS compliance. Identify issues and suggest improvements.

**Process**:
1. Assess the requirement against INCOSE quality criteria
2. Check EARS pattern compliance
3. Detect weak terms, ambiguities, and missing information
4. Score quality dimensions (Clarity, Testability, Completeness)
5. Generate an improved EARS-compliant version
6. Explain all changes

**Output**:
```
**Original**: {requirement text}

**Quality Scores** (1-5 scale):
- Clarity: X/5 - {specific issues if <5}
- Testability: Y/5 - {specific issues if <5}
- Completeness: Z/5 - {specific issues if <5}

**EARS Pattern**: {current pattern or "None/Incorrect"}

**Issues Detected**:
1. {Issue category}: {specific problem}
2. {Issue category}: {specific problem}

**Weak Terms**: {list any vague/subjective terms}

**Improved Version:**
[REQ-XXX] {EARS-formatted requirement}

**Improvements Made**:
- {Specific change 1}
- {Specific change 2}

**Explanation**: {Why these changes improve quality}
```

**Example**:

Input: "The app should be fast and user-friendly"

Output:
```
**Original**: "The app should be fast and user-friendly"

**Quality Scores** (1-5 scale):
- Clarity: 1/5 - Extremely vague, no specific behavior defined
- Testability: 1/5 - Cannot be objectively verified
- Completeness: 1/5 - Missing trigger, system name, measurable criteria

**EARS Pattern**: None

**Issues Detected**:
1. Weak modal: Uses "should" instead of "shall"
2. Vague terms: "fast" (unmeasurable), "user-friendly" (subjective)
3. Compound requirement: Combines performance and usability
4. Missing system name: Uses "the app" without specification
5. No trigger or condition: Unclear when this applies

**Weak Terms**: fast, user-friendly

**Improved Version:**
Split into two separate requirements:

[REQ-001] When a user initiates a search, the mobile application shall return results within 2 seconds.

[REQ-002] The mobile application shall display feedback within 200 milliseconds for every user action to maintain perceived responsiveness.

**Improvements Made**:
- Replaced "should" with "shall"
- Removed vague term "fast", replaced with specific timing criteria
- Removed subjective "user-friendly", replaced with measurable usability requirement
- Split compound requirement into atomic requirements
- Added specific triggers (user initiates search, user action)
- Named the system (mobile application)

**Explanation**: The original requirement was untestable and unclear. By splitting it into two atomic requirements with measurable criteria, we can now verify both performance (search speed) and usability (responsiveness) objectively.
```

This command provides detailed analysis to help improve any requirement.
