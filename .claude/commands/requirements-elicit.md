---
name: requirements-elicit
description: Interactive requirements elicitation through guided conversation
version: 1.0.0
tags: [requirements, ears, elicitation, interactive]
---

Start an interactive requirements elicitation session to gather requirements for a system or feature through guided conversation.

**Process**:
1. Understand the system context (ask if not provided)
2. Ask targeted questions to extract requirements:
   - What triggers the behavior? (events)
   - What conditions must be true? (states)
   - What should the system do? (responses)
   - What are the success criteria? (measurable outcomes)
   - What could go wrong? (error scenarios)
3. Draft EARS requirements from responses
4. Validate understanding with the user
5. Iterate until comprehensive coverage

**Initial Questions** (adapt based on context):
```
**System Understanding**:
1. What system/feature are we defining requirements for?
2. Who are the primary users/stakeholders?
3. What is the main purpose or goal?

**Behavioral Questions**:
4. What events or user actions should trigger the system?
5. Are there specific states or conditions that affect behavior?
6. What should happen when everything works correctly?
7. What should happen when things go wrong?
8. Are there performance, security, or other quality expectations?
```

**Output Format** (iterative):
```
**Current Context**:
System: {system name}
Feature: {feature being specified}
Stakeholders: {key users/stakeholders}

**Requirements Drafted**:
[REQ-001] {EARS requirement from conversation}
[REQ-002] {EARS requirement from conversation}
...

**Next Questions**:
1. {Targeted question about gaps}
2. {Targeted question about edge cases}
3. {Targeted question about non-functional aspects}

**Coverage Assessment**:
✅ Happy path scenarios covered
⚠️  Need more detail on error handling
⚠️  Missing performance requirements
❓ Should we consider {potential scenario}?
```

**Example Session**:

User: "I need requirements for a password reset feature"