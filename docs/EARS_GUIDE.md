# EARS - Easy Approach to Requirements Syntax

**Complete Reference Guide**

## Table of Contents

- [Introduction](#introduction)
- [The 5 EARS Patterns](#the-5-ears-patterns)
- [Syntax Rules](#syntax-rules)
- [Best Practices](#best-practices)
- [Common Mistakes](#common-mistakes)
- [Complex Requirements](#complex-requirements)
- [When NOT to Use EARS](#when-not-to-use-ears)
- [Examples by Domain](#examples-by-domain)

---

## Introduction

### What is EARS?

**EARS (Easy Approach to Requirements Syntax)** is a constrained natural language notation developed by **Alistair Mavin** and colleagues at **Rolls-Royce PLC** in 2009. It emerged from analysis of airworthiness regulations for jet engine control systems and has since been adopted globally by organizations including Airbus, Bosch, NASA, and Siemens.

### Why Use EARS?

EARS addresses common problems in natural language requirements:

**Problems EARS Solves**:
- ❌ **Ambiguity**: Multiple possible interpretations
- ❌ **Inconsistency**: Different authors use different styles
- ❌ **Incompleteness**: Missing triggers, conditions, or responses
- ❌ **Weak Language**: Vague terms like "fast", "easy", "should"
- ❌ **Difficult to Test**: Unclear what constitutes success

**Benefits of EARS**:
- ✅ **Reduces ambiguity** through structured templates
- ✅ **Improves consistency** across different authors
- ✅ **Simplifies understanding** especially for non-native English speakers
- ✅ **Facilitates testability** with clear trigger/response pairs
- ✅ **Enables automation** through structured format
- ✅ **Speeds up review** by making issues more visible

### Core Principle

EARS constrains natural language into **5 specific patterns**, each identified by keywords. Every requirement fits into one (or a combination) of these patterns.

---

## The 5 EARS Patterns

### Pattern 1: Ubiquitous Requirements

**Characteristics**:
- **Always active** with no preconditions
- Continuous properties or behaviors
- No triggering events or conditions
- **No keyword** (just uses "shall")

**Template**:
```
The <system name> shall <system response>
```

**Examples**:

✅ **Good Examples**:
```
The mobile phone shall have a mass of less than 150 grams.
The database shall store data in UTF-8 encoding.
The API shall support HTTPS connections on port 443.
The password field shall accept strings between 8 and 128 characters.
```

❌ **Bad Examples** (and why):
```
❌ "The system should be secure"
   Problem: Uses "should", vague "secure", no specific property

❌ "The app must be fast"
   Problem: Uses "must" not "shall", "fast" is unmeasurable

✅ IMPROVED: "The web application shall encrypt all data in transit using TLS 1.3 or higher."
```

**When to Use**:
- Constant system properties (weight, size, capacity)
- Architectural constraints (protocols, formats, standards)
- Continuous invariants (always-true conditions)
- Static configurations

---

### Pattern 2: State-Driven Requirements

**Characteristics**:
- Active **as long as a condition remains true**
- Behavior depends on system state
- State persists over time
- **Keyword**: **"While"**

**Template**:
```
While <precondition(s)>, the <system name> shall <system response>
```

**Examples**:

✅ **Good Examples**:
```
While there is no card in the ATM, the ATM shall display "Insert card to begin".

While the aircraft is on the ground, the landing gear indicator shall display "LOCKED".

While the user is logged out, the application shall hide all user-specific menu options.

While the battery level is below 20%, the device shall enable power-saving mode.

While the printer is out of paper, the printer shall display the message "Please refill paper" and shall not accept new print jobs.
```

❌ **Bad Examples** (and why):
```
❌ "While offline, handle requests appropriately"
   Problem: "appropriately" is vague, missing system name, weak verb

✅ IMPROVED: "While the network connection is unavailable, the mobile app shall queue user requests locally and shall display 'Working offline' in the status bar."
```

**When to Use**:
- System modes (offline mode, maintenance mode, safe mode)
- Resource states (battery low, memory full, disk space low)
- User states (logged in, authenticated, admin role)
- Environmental conditions (temperature range, network availability)

**Key Distinction**:
- **State-Driven ("While")**: Condition persists - behavior continues while true
- **Event-Driven ("When")**: One-time trigger - behavior happens in response to event

---

### Pattern 3: Event-Driven Requirements

**Characteristics**:
- Specify response to a **specific triggering event**
- One-time or repeated actions
- Event marks a point in time
- **Keyword**: **"When"**

**Template**:
```
When <trigger>, the <system name> shall <system response>
```

**Examples**:

✅ **Good Examples**:
```
When the user clicks the Submit button, the form shall validate all fields within 500 milliseconds.

When 'mute' is selected, the laptop shall suppress all audio output.

When the sensor detects motion, the alarm system shall activate the siren within 1 second.

When a file upload completes, the web application shall display a success notification and shall refresh the file list.

When the system receives a shutdown command, the database shall flush all pending writes to disk before terminating.
```

❌ **Bad Examples** (and why):
```
❌ "When the user logs in, show their dashboard"
   Problem: Informal verb "show", missing system name, no "shall"

✅ IMPROVED: "When the user successfully authenticates, the web application shall display the user's dashboard within 1 second."

❌ "Process orders when they arrive"
   Problem: Missing system name, no "shall", vague "process"

✅ IMPROVED: "When an order is received, the order management system shall validate the order details and shall update the order status to 'Processing' within 5 seconds."
```

**When to Use**:
- User actions (clicks, inputs, selections)
- System events (startup, shutdown, timer expiration)
- External triggers (messages received, files detected)
- State transitions (mode changes, status updates)

**Timing Considerations**:
Add performance criteria when relevant:
```
When the user submits a search query, the system shall return results within 2 seconds.
```

---

### Pattern 4: Optional Feature Requirements

**Characteristics**:
- Apply **only when a specific feature is present**
- Variant or configurable system features
- Product line variations
- **Keyword**: **"Where"**

**Template**:
```
Where <feature is included>, the <system name> shall <system response>
```

**Examples**:

✅ **Good Examples**:
```
Where the car has a sunroof, the car shall have a sunroof control panel on the driver's door.

Where GPS is enabled, the mobile app shall display the user's current location on the map.

Where the premium subscription is active, the streaming service shall allow downloads for offline viewing.

Where biometric authentication is supported, the device shall offer fingerprint or face recognition as login options.

Where the API version is 2.0 or higher, the server shall include HATEOAS links in all responses.
```

❌ **Bad Examples** (and why):
```
❌ "Where available, use dark mode"
   Problem: Missing system name, no "shall", vague "use"

✅ IMPROVED: "Where dark mode is enabled in system settings, the application shall use the dark color theme for all UI elements."
```

**When to Use**:
- Optional hardware components
- Feature toggles or flags
- Subscription tiers or license levels
- Platform-specific capabilities
- Configuration variants

**Combining with Other Patterns**:
```
Where the video conferencing feature is enabled, when a user starts a call, the application shall activate the camera and microphone.
```

---

### Pattern 5: Unwanted Behavior Requirements

**Characteristics**:
- Specify response to **undesired or exceptional situations**
- Error handling and fault tolerance
- Safety requirements
- **Keywords**: **"If...Then"**

**Template**:
```
If <unwanted condition or event>, then the <system name> shall <system response>
```

**Examples**:

✅ **Good Examples**:
```
If an invalid credit card number is entered, then the payment form shall display "Please re-enter credit card details" below the card number field.

If the database connection fails, then the application shall log the error and shall display "Service temporarily unavailable. Please try again later."

If the temperature sensor reads above 85°C, then the cooling system shall activate the emergency shutdown procedure within 2 seconds.

If a user enters an incorrect password 3 times consecutively, then the system shall lock the account for 15 minutes and shall send a security alert email.

If the file size exceeds 10 MB, then the upload system shall reject the file and shall display "File too large. Maximum size is 10 MB."
```

❌ **Bad Examples** (and why):
```
❌ "If there's an error, handle it"
   Problem: Vague "error", weak verb "handle", no system name

✅ IMPROVED: "If a network timeout occurs during data transmission, then the synchronization service shall retry the operation up to 3 times with exponential backoff before reporting failure to the user."
```

**When to Use**:
- Input validation errors
- System failures and exceptions
- Security violations
- Resource exhaustion
- Safety hazards
- Boundary condition violations

**Safety-Critical Systems**:
Be especially precise with unwanted behavior requirements:
```
If the brake pressure sensor detects pressure below 50 psi, then the brake control system shall activate the warning light, shall log the fault code, and shall engage the emergency braking assist within 100 milliseconds.
```

---

## Syntax Rules

### Mandatory Rules

1. **Always use "shall"** for requirements
   - ✅ "The system **shall** validate input"
   - ❌ "The system **should** validate input"
   - ❌ "The system **must** validate input"
   - ❌ "The system **will** validate input"
   - ❌ "The system **may** validate input"

2. **One requirement per sentence**
   - ❌ "The system shall validate input and sanitize data and log errors"
   - ✅ Split into three requirements:
     - REQ-001: "When a user submits a form, the system shall validate all input fields"
     - REQ-002: "When a user submits a form, the system shall sanitize all input data"
     - REQ-003: "If validation fails, then the system shall log the validation error"

3. **Specify the system name explicitly**
   - ❌ "It shall display an error"
   - ✅ "The registration form shall display an error"

4. **Use precise, measurable terms**
   - ❌ "The system shall respond quickly"
   - ✅ "When a user submits a query, the search engine shall return results within 2 seconds"

5. **Include units of measure**
   - ❌ "The device shall weigh less than 500"
   - ✅ "The mobile device shall have a mass of less than 500 grams"

### Recommended Practices

**Active Voice**:
- ✅ "The system shall send a notification"
- ⚠️ "A notification shall be sent by the system" (passive - harder to read)

**Avoid Pronouns**:
- ❌ "When a user logs in, it shall verify their credentials"
- ✅ "When a user logs in, the authentication service shall verify the user's credentials"

**Be Specific About "The System"**:
- Instead of generic "the system", name the specific component:
  - "the web application"
  - "the authentication service"
  - "the payment gateway"
  - "the user interface"

**Quantify Everything Measurable**:
- Response times (seconds, milliseconds)
- Capacities (MB, GB, number of items)
- Durations (minutes, hours, days)
- Frequencies (per second, per minute)
- Accuracies (%, decimal precision)

---

## Best Practices

### Start with the Right Pattern

**Decision Tree**:
```
Is there a feature/variant condition?
├─ YES → Use "Where" (Optional Feature)
└─ NO  → Continue...

Is this about error/exception handling?
├─ YES → Use "If...Then" (Unwanted Behavior)
└─ NO  → Continue...

Is there a triggering event (point in time)?
├─ YES → Use "When" (Event-Driven)
└─ NO  → Continue...

Is there a persisting condition (state over time)?
├─ YES → Use "While" (State-Driven)
└─ NO  → Use basic "shall" (Ubiquitous)
```

### Be Explicit About Context

Don't assume the reader knows context. Always include:
- **Who/what triggers** the behavior
- **What conditions** must be true
- **What the system does** in response
- **When/how to measure** success

**Example Progression**:
```
❌ Level 0: "Validate email"
❌ Level 1: "The system should validate email addresses"
⚠️ Level 2: "The registration form shall validate email addresses"
✅ Level 3: "When a user enters an email address in the registration form, the application shall validate the format according to RFC 5322"
✅✅ Level 4: "When a user enters an email address in the registration form, the application shall validate the format according to RFC 5322 and shall display an error message 'Please enter a valid email address' below the email field if the format is invalid"
```

### Think Testability First

For every requirement, ask: **"How do I test this?"**

If you can't write a clear test case, the requirement isn't complete.

**Testable**:
```
"When a user clicks the Save button, the document editor shall save the file to local storage within 1 second."

Test:
1. Open document editor
2. Make changes
3. Click Save button
4. Measure time to completion (must be ≤1 second)
5. Verify file exists in local storage
6. Verify file contents match document
```

**Not Testable**:
```
"The application shall be user-friendly."

How do you objectively test "user-friendly"?
```

### Use Consistent Terminology

Create and maintain a **glossary** of terms:
- Define technical terms precisely
- Use the same term for the same concept
- Avoid synonyms (don't alternate between "user", "customer", "client")

---

## Common Mistakes

### 1. Weak Modal Verbs

❌ **WRONG**:
- "The system **should** encrypt data"
- "The API **must** return a response"
- "Users **may** change their password"
- "The app **will** display a message"

✅ **CORRECT**:
- "The system **shall** encrypt all data at rest using AES-256"
- "When the API receives a request, the API **shall** return a response within 500 milliseconds"
- "The user account page **shall** provide a 'Change Password' button"
- "If login fails, then the login page **shall** display the message 'Invalid credentials'"

### 2. Vague Quantifiers

❌ **WRONG** (Unmeasurable):
- "fast", "slow", "quick", "rapidly"
- "large", "small", "big", "tiny"
- "many", "few", "several", "some"
- "often", "rarely", "occasionally"
- "approximately", "around", "about"

✅ **CORRECT** (Measurable):
- "within 2 seconds", "under 500 milliseconds"
- "greater than 10 MB", "between 1 and 100 items"
- "at least 50", "no more than 1000"
- "every 5 minutes", "once per hour"
- "±5%", "exactly", "precisely"

### 3. Compound Requirements

❌ **WRONG**:
"The system shall authenticate users, log all access attempts, and send email notifications when suspicious activity is detected."

This is actually 3 requirements:

✅ **CORRECT**:
```
REQ-001: When a user submits login credentials, the authentication service shall verify the credentials against the user database.

REQ-002: When a user attempts to log in, the system shall log the attempt with timestamp, username, IP address, and result (success/failure).

REQ-003: If the system detects more than 5 failed login attempts from the same IP address within 10 minutes, then the security service shall send an email notification to the system administrator.
```

### 4. Missing Triggers or Conditions

❌ **WRONG**:
"The application shall display a progress bar."

✅ **CORRECT**:
"While a file upload is in progress, the application shall display a progress bar showing the percentage completed."

❌ **WRONG**:
"The system shall send a notification."

✅ **CORRECT**:
"When an order status changes to 'Shipped', the order management system shall send a notification email to the customer."

### 5. Subjective or Ambiguous Terms

❌ **AVOID**:
- "user-friendly", "intuitive", "easy to use"
- "robust", "stable", "reliable" (without quantification)
- "adequate", "sufficient", "appropriate"
- "as needed", "if necessary", "when appropriate"

✅ **REPLACE WITH**:
- Specific usability requirements (e.g., "task completion in ≤3 clicks")
- Measurable reliability (e.g., "99.9% uptime", "MTBF of 10,000 hours")
- Concrete thresholds (e.g., "support 1000 concurrent users")
- Explicit conditions (e.g., "when battery level drops below 10%")

### 6. Negative Requirements

⚠️ **AVOID** (when possible):
"The system shall not allow unauthorized access."

✅ **PREFER** (positive statement):
"The system shall grant access only to authenticated and authorized users."

However, "If...Then" patterns naturally express negatives for unwanted behavior:
```
If a user is not authenticated, then the application shall redirect to the login page.
```

---

## Complex Requirements

### Combining Patterns

Real-world requirements often need multiple patterns combined:

**Example 1**: State + Event
```
While the aircraft is on the ground, when the pilot commands reverse thrust, the engine control system shall engage reverse thrust within 500 milliseconds.
```
Combines:
- State-driven: "While the aircraft is on the ground"
- Event-driven: "when the pilot commands reverse thrust"

**Example 2**: Optional + Event + Unwanted
```
Where biometric authentication is enabled, when a user attempts fingerprint login, if the fingerprint does not match after 3 attempts, then the device shall fall back to password authentication.
```
Combines:
- Optional: "Where biometric authentication is enabled"
- Event: "when a user attempts fingerprint login"
- Unwanted: "if the fingerprint does not match after 3 attempts"

**Example 3**: State + Ubiquitous
```
While operating in high-security mode, the system shall encrypt all data transmissions using TLS 1.3 with certificate pinning.
```
Combines:
- State-driven: "While operating in high-security mode"
- Ubiquitous property: encryption requirement

### Guidelines for Complex Requirements

1. **Don't force it**: If combining patterns makes the requirement harder to understand, split it into multiple requirements

2. **Maximum 2-3 patterns**: Beyond that, consider decomposition

3. **Maintain clarity**: The combined requirement should still be clear and testable

4. **Logical order**: Present conditions/triggers in logical sequence (state → event → response)

---

## When NOT to Use EARS

EARS is powerful but not universal. Recognize when other approaches are better:

### 1. Mathematical or Formal Specifications

❌ **DON'T force into EARS**:
Complex algorithms, mathematical formulas, state machines

✅ **USE instead**:
- Mathematical notation
- Formal specification languages (Z, VDM, TLA+)
- State diagrams
- Sequence diagrams

**Example**: For a cryptographic algorithm, reference the standard:
```
"The encryption module shall implement AES-256 as specified in FIPS 197."
```
Rather than trying to describe the algorithm in EARS.

### 2. Highly Technical Domain-Specific Requirements

Some domains have established notations:
- **Hardware**: Circuit diagrams, timing diagrams
- **Protocols**: State machines, packet formats
- **Real-time systems**: Timing analysis, schedulability
- **Safety**: Fault trees, FMEA tables

Use EARS for **high-level** requirements, domain notation for **low-level** details.

### 3. Pure Constraints Without System Response

Some constraints don't fit the trigger→response pattern:

**Example**:
"All code shall comply with MISRA-C coding standards."

This is a constraint, not a behavior. It's acceptable as-is.

### 4. Information or Glossary Items

Definitions and explanations aren't requirements:

**Not a requirement**:
"A 'user session' is the period between login and logout."

This is a definition for the glossary, not a requirement.

---

## Examples by Domain

### Web Applications

```
REQ-WEB-001: When a user submits a registration form, the web application shall validate all required fields and shall display error messages inline within 500 milliseconds.

REQ-WEB-002: While a user is logged in, the application shall display the user's name in the header.

REQ-WEB-003: If a user's session expires, then the application shall redirect the user to the login page with the message "Your session has expired. Please log in again."

REQ-WEB-004: Where a user has admin privileges, the application shall display the Admin menu option in the navigation bar.

REQ-WEB-005: The web application shall use HTTPS for all communications.
```

### Mobile Applications

```
REQ-MOB-001: When the app is launched, the mobile application shall display the home screen within 2 seconds.

REQ-MOB-002: While the device is offline, the app shall queue all user actions locally and shall display "Working offline" in the status bar.

REQ-MOB-003: If the GPS signal is lost, then the navigation app shall display "GPS signal lost" and shall continue displaying the last known location.

REQ-MOB-004: Where location services are enabled, the app shall request the user's location permission on first launch.

REQ-MOB-005: The mobile app shall support iOS 15 or higher and Android 11 or higher.
```

### Safety-Critical Systems

```
REQ-SAFE-001: When the emergency stop button is pressed, the robotic system shall halt all movement within 100 milliseconds.

REQ-SAFE-002: While the safety door is open, the machine shall not allow the cutting blade to operate.

REQ-SAFE-003: If the temperature sensor detects a reading above 90°C, then the control system shall activate the cooling system and shall trigger an alarm within 1 second.

REQ-SAFE-004: Where redundant sensors are installed, the control system shall compare readings and shall use the median value for decision-making.

REQ-SAFE-005: The safety interlock shall fail to a safe state (system shutdown) if power is interrupted.
```

### APIs and Services

```
REQ-API-001: When the API receives a GET request to /users/{id}, the service shall return the user data in JSON format within 200 milliseconds.

REQ-API-002: The REST API shall use API keys for authentication on all endpoints except /health.

REQ-API-003: If an API request contains invalid JSON, then the service shall return HTTP 400 with an error message describing the JSON parsing error.

REQ-API-004: While the rate limit is exceeded, the API shall return HTTP 429 and shall include the Retry-After header.

REQ-API-005: Where the client requests API version 2.0, the service shall include HATEOAS links in all responses.
```

---

## Quick Reference Card

### The 5 Patterns

| Pattern | Keyword | Template | Use Case |
|---------|---------|----------|----------|
| **Ubiquitous** | *(none)* | The \<system\> shall \<response\> | Continuous properties |
| **State-Driven** | **While** | While \<state\>, the \<system\> shall \<response\> | State-dependent behavior |
| **Event-Driven** | **When** | When \<event\>, the \<system\> shall \<response\> | Event triggers |
| **Optional** | **Where** | Where \<feature\>, the \<system\> shall \<response\> | Feature variants |
| **Unwanted** | **If...Then** | If \<condition\>, then \<system\> shall \<response\> | Error/fault handling |

### Quality Checklist

- [ ] Uses "shall" (not should, must, will, may)
- [ ] Specifies system name explicitly
- [ ] Includes only one requirement (atomic)
- [ ] Has clear trigger/condition (if applicable)
- [ ] Uses measurable criteria (with units)
- [ ] Avoids weak/vague terms
- [ ] Is testable (can write test case)
- [ ] Matches one of the 5 EARS patterns

---

## Resources

**Official EARS Resources**:
- Alistair Mavin's EARS site: https://alistairmavin.com/ears/
- Original 2009 paper: "Easy Approach to Requirements Syntax (EARS)"
- INCOSE Requirements Working Group materials

**Standards**:
- ISO/IEC/IEEE 29148:2018 - Requirements Engineering
- INCOSE Guide to Writing Requirements

**Tools**:
- ARM (Automated Requirements Measurement)
- QuARS (Quality Analyzer for Requirements)
- RUBRIC (Boilerplate conformance checking)

---

*This guide is part of the Uberman Requirements Engineering toolkit.*
