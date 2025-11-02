# EARS Patterns - Quick Reference Templates

## Pattern 1: Ubiquitous (No Keyword)

**Use when**: Continuous properties, always-true conditions, constant behaviors

**Template**:
```
The <system name> shall <system response>
```

**Fill-in Example**:
```
The _____________ shall _________________________________
   (system name)         (what the system does/property)

Example:
The mobile application shall support iOS 15 or higher.
```

---

## Pattern 2: State-Driven (While)

**Use when**: Behavior depends on a persisting system state or condition

**Template**:
```
While <precondition(s)>, the <system name> shall <system response>
```

**Fill-in Example**:
```
While ________________________________, the _____________ shall _________________________________
      (what condition must be true)        (system name)         (what the system does)

Example:
While the user is logged out, the application shall hide all user-specific menu options.
```

---

## Pattern 3: Event-Driven (When)

**Use when**: System responds to a specific trigger or event (point in time)

**Template**:
```
When <trigger>, the <system name> shall <system response>
```

**Fill-in Example**:
```
When ________________________________, the _____________ shall _________________________________
     (what triggers the behavior)         (system name)         (what the system does)

Example:
When the user clicks the Submit button, the form validator shall check all required fields within 500 milliseconds.
```

---

## Pattern 4: Optional Feature (Where)

**Use when**: Requirement applies only when a specific feature/variant is present

**Template**:
```
Where <feature is included>, the <system name> shall <system response>
```

**Fill-in Example**:
```
Where ________________________________, the _____________ shall _________________________________
      (what feature must be present)       (system name)         (what the system does)

Example:
Where biometric authentication is enabled, the device shall offer fingerprint or face recognition as login options.
```

---

## Pattern 5: Unwanted Behavior (If...Then)

**Use when**: Specifying response to errors, faults, or exceptional conditions

**Template**:
```
If <unwanted condition or event>, then the <system name> shall <system response>
```

**Fill-in Example**:
```
If ________________________________, then the _____________ shall _________________________________
   (what error/exception occurs)            (system name)         (what the system does)

Example:
If the database connection fails, then the application shall display "Service temporarily unavailable" and shall log the error.
```

---

## Complex Requirements (Combined Patterns)

### State + Event

**Template**:
```
While <state>, when <event>, the <system> shall <response>
```

**Example**:
```
While the aircraft is on the ground, when the pilot commands reverse thrust, the engine control system shall engage reverse thrust within 500 milliseconds.
```

### Optional + Event

**Template**:
```
Where <feature>, when <event>, the <system> shall <response>
```

**Example**:
```
Where dark mode is enabled, when the app launches, the UI shall use the dark color theme for all elements.
```

### Event + Unwanted

**Template**:
```
When <event>, if <error condition>, then the <system> shall <response>
```

**Example**:
```
When a user submits a form, if validation fails, then the system shall display field-specific error messages and shall not submit the form.
```

---

## Pattern Selection Flowchart

```
START
  |
  ├─ Is there a feature/variant condition?
  │   └─ YES → Use "Where" (Optional Feature)
  |
  ├─ Is this about error/exception handling?
  │   └─ YES → Use "If...Then" (Unwanted Behavior)
  |
  ├─ Is there a triggering event (point in time)?
  │   └─ YES → Use "When" (Event-Driven)
  |
  ├─ Is there a persisting condition (state over time)?
  │   └─ YES → Use "While" (State-Driven)
  |
  └─ None of the above
      └─ Use basic "shall" (Ubiquitous)
```

---

## Checklist for Any Requirement

- [ ] Uses "shall" (not should, must, will, may)
- [ ] Specifies system name explicitly
- [ ] Is atomic (one requirement only)
- [ ] Has clear trigger/condition (if applicable)
- [ ] Includes measurable criteria (with units where applicable)
- [ ] Avoids weak/vague terms (fast, easy, user-friendly, etc.)
- [ ] Is testable (can write a test case)
- [ ] Matches one of the 5 EARS patterns

---

*Use these templates as starting points for writing high-quality EARS requirements.*
