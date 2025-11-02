---
name: requirements-batch
description: Batch analysis of multiple requirements with quality reports
version: 1.0.0
tags: [requirements, ears, batch, analysis]
---

Analyze multiple requirements in batch mode. Identify common issues, check cross-requirement consistency, and generate a comprehensive quality report.

**Input**: Paste multiple requirements (one per line or numbered list)

**Process**:
1. Parse all requirements
2. Analyze each individually
3. Check for duplicates and contradictions
4. Identify patterns and common issues
5. Assess overall quality metrics
6. Generate prioritized improvement recommendations

**Output**:
```
**Batch Analysis Summary**
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total Requirements: {count}
EARS-Compliant: {count} ({percentage}%)

Quality Distribution:
- High Quality (avg ≥4.0): {count}
- Medium Quality (avg 2.5-3.9): {count}
- Low Quality (avg <2.5): {count}

**Pattern Distribution**:
- Ubiquitous (no keyword): {count}
- State-Driven (while): {count}
- Event-Driven (when): {count}
- Optional Feature (where): {count}
- Unwanted Behavior (if/then): {count}
- Non-EARS: {count}

**Common Issues** (top 5):
1. {Issue type} - {count} occurrences
2. {Issue type} - {count} occurrences
3. {Issue type} - {count} occurrences

**Cross-Requirement Issues**:
- Duplicates: {list of duplicate pairs}
- Contradictions: {list of conflicting requirements}
- Gaps: {missing scenarios identified}

**Detailed Analysis**:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

[REQ-001] {original text}
├─ Status: ✅ PASS / ⚠️ NEEDS IMPROVEMENT
├─ Pattern: {pattern type}
├─ Quality: Clarity {X}/5, Testability {Y}/5, Completeness {Z}/5
├─ Issues: {list if any}
└─ Improved: {EARS version if improvements needed}

[REQ-002] {original text}
├─ Status: ✅ PASS / ⚠️ NEEDS IMPROVEMENT
...

**Priority Improvements** (ranked by impact):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. 🔴 HIGH: {Critical issue affecting multiple requirements}
2. 🟡 MEDIUM: {Important issue affecting quality}
3. 🟢 LOW: {Minor improvement opportunity}
```

**Example**:

Input:
```
1. The system should respond quickly
2. Users must be able to login
3. The app shall validate email addresses
4. System should handle errors gracefully
```

Output:
```
**Batch Analysis Summary**
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total Requirements: 4
EARS-Compliant: 1 (25%)

Quality Distribution:
- High Quality (avg ≥4.0): 1 (REQ-003)
- Medium Quality (avg 2.5-3.9): 0
- Low Quality (avg <2.5): 3 (REQ-001, REQ-002, REQ-004)

**Pattern Distribution**:
- Event-Driven (when): 0
- State-Driven (while): 0
- Ubiquitous (shall): 1
- Non-EARS: 3

**Common Issues** (top 5):
1. Weak modal verbs (should/must) - 3 occurrences
2. Vague terms (quickly, gracefully) - 2 occurrences
3. Missing triggers/conditions - 3 occurrences
4. No measurable criteria - 3 occurrences

**Cross-Requirement Issues**:
- Gaps: No error-specific requirements for login failures
- Gaps: No performance requirements beyond vague "quickly"
- Missing: Security requirements for authentication

**Detailed Analysis**:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

[REQ-001] "The system should respond quickly"
├─ Status: ⚠️ NEEDS IMPROVEMENT
├─ Pattern: None
├─ Quality: Clarity 1/5, Testability 1/5, Completeness 1/5
├─ Issues: Uses "should", vague "quickly", missing trigger, no system name
└─ Improved: [REQ-001] When a user submits a request, the web application shall return a response within 2 seconds.

[REQ-002] "Users must be able to login"
├─ Status: ⚠️ NEEDS IMPROVEMENT
├─ Pattern: None
├─ Quality: Clarity 2/5, Testability 2/5, Completeness 2/5
├─ Issues: Uses "must" not "shall", missing system name, no success criteria
└─ Improved: [REQ-002] When a user enters valid credentials and clicks Login, the authentication system shall grant access to the user's account within 1 second.

[REQ-003] "The app shall validate email addresses"
├─ Status: ✅ PASS (with minor improvements recommended)
├─ Pattern: Ubiquitous
├─ Quality: Clarity 4/5, Testability 4/5, Completeness 3/5
├─ Issues: Missing trigger (when does validation occur?), what happens if invalid?
└─ Improved: [REQ-003] When a user enters an email address in the registration form, the application shall validate the format according to RFC 5322 and shall display an error message if the format is invalid.

[REQ-004] "System should handle errors gracefully"
├─ Status: ⚠️ NEEDS IMPROVEMENT
├─ Pattern: None
├─ Quality: Clarity 1/5, Testability 1/5, Completeness 1/5
├─ Issues: Uses "should", "gracefully" is subjective, no specific errors mentioned
└─ Improved: Split into specific error scenarios, e.g.:
   [REQ-004a] If a database connection fails, then the application shall display "Service temporarily unavailable" and shall log the error.
   [REQ-004b] If a user session expires, then the application shall redirect to the login page with the message "Your session has expired. Please log in again."

**Priority Improvements** (ranked by impact):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. 🔴 HIGH: Replace weak modal verbs (should/must) with "shall" in all requirements
2. 🔴 HIGH: Add specific, measurable criteria for all requirements
3. 🟡 MEDIUM: Add explicit error handling requirements for login failures
4. 🟡 MEDIUM: Break down vague requirement REQ-004 into specific error scenarios
5. 🟢 LOW: Specify system names consistently
```

This provides a comprehensive quality assessment and actionable improvement recommendations for requirement sets.
