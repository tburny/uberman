---
name: requirements-convert
description: Convert informal requirements to EARS format
version: 1.0.0
tags: [requirements, ears, conversion]
---

Convert the user's informal requirement or description into a properly formatted EARS (Easy Approach to Requirements Syntax) requirement.

**Process**:
1. Parse the user's input to understand the intent
2. Identify the appropriate EARS pattern (Ubiquitous, State-Driven, Event-Driven, Optional Feature, or Unwanted Behavior)
3. Transform into the correct EARS template
4. Ensure "shall" is used
5. Add specific, measurable criteria where needed
6. Remove vague or weak terms

**Output**:
```
**EARS Requirement:**
[REQ-XXX] {EARS-formatted requirement}

**Pattern**: {pattern type}
**Changes Made**:
- {List specific transformations}

**Testability**: {How to verify this requirement}
```

**Examples**:

Input: "The system should respond quickly"
Output:
```
**EARS Requirement:**
[REQ-001] When a user submits a query, the search system shall return results within 2 seconds.

**Pattern**: Event-Driven (When)
**Changes Made**:
- Changed "should" to "shall"
- Added specific trigger ("user submits a query")
- Replaced "quickly" with measurable criterion ("within 2 seconds")
- Specified system name ("search system")

**Testability**: Submit test queries and measure response time with performance monitoring tools
```

If the user's input is unclear or missing critical information, ask targeted clarifying questions before generating the requirement.
