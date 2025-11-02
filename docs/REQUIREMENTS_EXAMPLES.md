# Requirements Examples

**Real-world EARS requirements across different domains**

## Software Applications

### Web Applications

```toml
# Authentication and Authorization
[REQ-AUTH-001]
When a user enters valid credentials and clicks Login, the web application
shall authenticate the user and shall grant access to their account within 1 second.

[REQ-AUTH-002]
While a user is logged in, the application shall display the user's name and
avatar in the header on all pages.

[REQ-AUTH-003]
If a user enters incorrect credentials 3 times consecutively, then the system
shall lock the account for 15 minutes and shall send a security alert email to
the account owner.

[REQ-AUTH-004]
If a user's session is inactive for 30 minutes, then the application shall log
out the user and shall redirect to the login page with the message "Your session
has expired for security reasons."

# Data Management
[REQ-DATA-001]
When a user clicks the Export button, the application shall generate a CSV file
containing all user data and shall initiate download within 3 seconds.

[REQ-DATA-002]
Where a user has admin privileges, the application shall provide access to the
admin dashboard and user management features.

[REQ-DATA-003]
The web application shall encrypt all data at rest using AES-256 encryption.

[REQ-DATA-004]
When a user deletes their account, the application shall permanently remove all
personal data within 30 days in compliance with GDPR requirements.
```

### Mobile Applications

```toml
# Performance and Responsiveness
[REQ-MOB-001]
When the app is launched, the mobile application shall display the home screen
within 2 seconds on devices released in the last 3 years.

[REQ-MOB-002]
When a user taps any interactive element, the application shall provide visual
feedback (ripple, highlight, or animation) within 100 milliseconds.

[REQ-MOB-003]
While the device is offline, the app shall queue all user actions locally and
shall display "Working offline - changes will sync when online" in the status bar.

# Resource Management
[REQ-MOB-004]
Where location services are enabled, the app shall request location permission
on first launch with clear explanation of usage.

[REQ-MOB-005]
While the app is running in the background, the application shall limit location
updates to once every 5 minutes to conserve battery.

[REQ-MOB-006]
If available storage drops below 100 MB, then the app shall delete cached data
older than 7 days and shall notify the user "Clearing old cache to free space."
```

## APIs and Web Services

```toml
# REST API Requirements
[REQ-API-001]
When the API receives a GET request to /users/{id}, the service shall return
the user data in JSON format within 200 milliseconds for 95% of requests.

[REQ-API-002]
The REST API shall use API keys in the Authorization header for authentication
on all endpoints except /health and /version.

[REQ-API-003]
If an API request contains malformed JSON, then the service shall return HTTP 400
with a response body containing {"error": "Invalid JSON", "details": "<specific parse error>"}.

[REQ-API-004]
While an API consumer exceeds the rate limit of 1000 requests per hour, the service
shall return HTTP 429 with the Retry-After header indicating seconds until reset.

[REQ-API-005]
Where the client requests API version 2.0 or higher (via Accept-Version header),
the service shall include HATEOAS links in all responses.

# Data Validation
[REQ-API-006]
When the API receives a POST request to create a resource, the service shall
validate all required fields and shall return HTTP 422 with field-specific error
messages if validation fails.

[REQ-API-007]
If an API request references a non-existent resource, then the service shall
return HTTP 404 with {"error": "Resource not found", "resource": "<type>", "id": "<id>"}.
```

## Safety-Critical Systems

```toml
# Industrial Control
[REQ-SAFE-001]
When the emergency stop button is pressed, the robotic system shall halt all
movement within 100 milliseconds and shall maintain the stopped state until
manually reset by an operator.

[REQ-SAFE-002]
While the safety door is open, the machine control system shall not allow the
cutting blade to operate and shall display "SAFETY DOOR OPEN" on the operator panel.

[REQ-SAFE-003]
If the temperature sensor detects a reading above 90°C, then the thermal
management system shall activate the cooling system and shall trigger an audible
alarm within 1 second.

[REQ-SAFE-004]
Where redundant sensors are installed, the control system shall compare all sensor
readings and shall use the median value for decision-making, and shall log a fault
if readings diverge by more than 10%.

[REQ-SAFE-005]
The safety interlock shall fail to a safe state (complete system shutdown) if
power is interrupted or if any safety circuit is compromised.

# Medical Devices
[REQ-MED-001]
When a critical alarm condition is detected, the patient monitor shall emit an
audible alarm at 65-85 dB(A) at 1 meter distance within 250 milliseconds.

[REQ-MED-002]
If the blood pressure reading exceeds 180/120 mmHg, then the monitoring system
shall display a red alert, shall sound a continuous alarm, and shall log the
event with timestamp.

[REQ-MED-003]
The infusion pump shall maintain flow rate accuracy within ±5% of the set rate
under all operating conditions specified in the user manual.
```

## E-Commerce

```toml
# Shopping Cart
[REQ-ECOM-001]
When a user adds an item to their cart, the shopping cart system shall update
the cart count badge and shall save the cart state to the database within 500
milliseconds.

[REQ-ECOM-002]
While items remain in the cart, the application shall preserve the cart contents
for 30 days or until the user completes checkout, whichever occurs first.

[REQ-ECOM-003]
If an item in the cart becomes out of stock, then the system shall remove the
item, shall notify the user with message "Item X is no longer available and has
been removed from your cart", and shall update the total price.

# Checkout and Payment
[REQ-ECOM-004]
When a user submits payment information, the payment gateway shall process the
transaction and shall return a success or failure response within 10 seconds.

[REQ-ECOM-005]
If payment processing fails, then the checkout system shall display the specific
error message from the payment gateway, shall not clear the shipping/billing
information, and shall allow the user to retry or choose a different payment method.

[REQ-ECOM-006]
The payment system shall comply with PCI DSS requirements and shall not store
full credit card numbers in any system logs or databases.

# Order Management
[REQ-ECOM-007]
When an order is successfully placed, the order management system shall send a
confirmation email to the customer within 1 minute containing order number, items,
prices, and estimated delivery date.

[REQ-ECOM-008]
When an order status changes, the notification service shall send a status update
email to the customer and shall update the order tracking page in real-time.
```

## IoT and Embedded Systems

```toml
# Smart Home
[REQ-IOT-001]
When motion is detected by the sensor, the smart lighting system shall turn on
the lights within 500 milliseconds.

[REQ-IOT-002]
While no motion is detected for 5 consecutive minutes, the lighting system shall
gradually dim the lights over 10 seconds and then turn them off.

[REQ-IOT-003]
Where voice control is enabled, the smart home hub shall respond to voice commands
within 1 second and shall provide audible confirmation of actions taken.

[REQ-IOT-004]
If the internet connection is lost, then the smart home devices shall continue to
operate using locally cached settings and shall log all state changes for sync
when connectivity is restored.

# Wearable Devices
[REQ-WEAR-001]
The fitness tracker shall measure heart rate with accuracy of ±5 BPM compared to
a medical-grade reference monitor under normal operating conditions.

[REQ-WEAR-002]
When the device battery level drops below 10%, the wearable shall display a low
battery notification and shall reduce sensor polling frequency to extend battery
life by at least 2 hours.

[REQ-WEAR-003]
The wearable device shall sync collected data with the mobile app within 30 seconds
when a Bluetooth connection is established.
```

## Data Processing and Analytics

```toml
# Batch Processing
[REQ-DATA-001]
When a batch job is triggered, the data processing system shall process up to
1 million records within 10 minutes using the allocated compute resources.

[REQ-DATA-002]
If a batch job fails, then the system shall log the error with stack trace, shall
send an alert email to the operations team, and shall preserve the input data for
manual review.

[REQ-DATA-003]
The ETL pipeline shall maintain data accuracy with zero data loss during transformation
and loading stages.

# Real-time Analytics
[REQ-ANALYTICS-001]
When new event data arrives, the analytics engine shall process and update dashboards
within 2 seconds to provide near-real-time visibility.

[REQ-ANALYTICS-002]
The analytics system shall handle peak loads of 10,000 events per second without
data loss or performance degradation below SLA thresholds.

[REQ-ANALYTICS-003]
If the data stream processing falls behind by more than 60 seconds, then the system
shall trigger an alert and shall automatically scale up processing capacity.
```

## Before/After Transformations

### Example 1: Login Feature

**❌ Before (informal, low quality)**:
"Users should be able to login quickly. Handle errors appropriately."

**✅ After (EARS, high quality)**:
```
[REQ-LOGIN-001] When a user enters valid credentials and clicks the Login button,
the authentication service shall verify the credentials against the user database
and shall grant access to the user's account within 1 second.

[REQ-LOGIN-002] If a user enters invalid credentials, then the login page shall
display "Invalid email or password. Please try again." and shall not reveal
whether the email or password was incorrect for security reasons.

[REQ-LOGIN-003] If a user enters incorrect credentials 3 times consecutively within
15 minutes, then the system shall lock the account for 15 minutes and shall send
a security alert email to the registered email address.

[REQ-LOGIN-004] If a user's session is inactive for 30 minutes, then the application
shall log out the user automatically and shall redirect to the login page with the
message "Your session has expired for security reasons. Please log in again."
```

### Example 2: File Upload

**❌ Before (vague, untestable)**:
"Allow users to upload files. Files should be validated and stored securely."

**✅ After (EARS, testable)**:
```
[REQ-UPLOAD-001] When a user selects a file and clicks Upload, the web application
shall validate that the file size does not exceed 10 MB and that the file type is
one of: PDF, DOC, DOCX, JPG, PNG.

[REQ-UPLOAD-002] If the selected file exceeds 10 MB, then the application shall
display "File too large. Maximum size is 10 MB." and shall not upload the file.

[REQ-UPLOAD-003] If the selected file type is not in the allowed list, then the
application shall display "Invalid file type. Allowed types: PDF, DOC, DOCX, JPG,
PNG." and shall not upload the file.

[REQ-UPLOAD-004] When a valid file is uploaded, the storage service shall scan
the file for malware and shall store the file only if the scan returns clean.

[REQ-UPLOAD-005] The file storage service shall encrypt all uploaded files at rest
using AES-256 encryption.

[REQ-UPLOAD-006] When a file upload completes successfully, the application shall
display "File uploaded successfully: [filename]" and shall refresh the file list
to include the new file.
```

---

*These examples demonstrate EARS best practices across various domains. Use them as templates for your own requirements.*
