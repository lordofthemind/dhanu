### **Development Roadmap for Dhanu (Modular Breakdown)**

Here's a systematic breakdown of **Dhanu** into modules and a roadmap to guide its development. Each module is designed to focus on one key aspect, ensuring that the project is manageable, and tasks are divided logically.

---

## **Phase 1: Core Foundation and CLI Structure**
### **Module 1: Project Initialization**
1. **Set up Go environment** and initialize Go modules (`go mod init`).
2. **Install Cobra** for CLI functionality and initialize the Cobra project (`cobra-cli init`).
3. **Create the root command** and set up basic CLI structure.

**Deliverables:**
- Basic project structure.
- Root CLI command with version and help information.

---

## **Phase 2: Basic Email Sending Functionality**
### **Module 2: Email Functionality**
1. **Install and configure email packages** (e.g., `gomail` or `net/smtp`).
2. **Implement email sending functionality** in a new command (`cobra-cli add send`).
3. Support **basic attachments** (single file).

**Deliverables:**
- Ability to send a simple email from the CLI.
- CLI command `send` with the ability to send attachments.

---

## **Phase 3: Script Execution & Output Handling**
### **Module 3: Script Execution**
1. Create a new command for **running scripts** (`cobra-cli add run`).
2. Implement functionality to **execute shell scripts** and capture output using `os/exec`.
3. **Send script output via email** using the `send` command.

**Deliverables:**
- CLI command `run` to execute a script and email its output.

---

## **Phase 4: Advanced Email Features**
### **Module 4: Advanced Email Options**
1. **Add CC, BCC**, and multiple recipient support.
2. Implement **custom email templates** with variables for dynamic email content (e.g., script name, date).
3. Add the option for **on-success/on-failure email triggers** for script execution.

**Deliverables:**
- Support for multiple recipients and email templates.
- Enhanced email options for conditional notifications.

---

## **Phase 5: File Watching and Notifications**
### **Module 5: File Watcher & Notifications**
1. Install and configure the `fsnotify` package.
2. Create a new command to **watch a file or directory** for changes (`cobra-cli add watch`).
3. Implement logic to **send email notifications** when changes are detected.

**Deliverables:**
- CLI command `watch` that sends email notifications when files are modified.

---

## **Phase 6: Configuration & Environment Setup**
### **Module 6: Configuration Management**
1. Install and configure **Viper** to manage configuration (`go get github.com/spf13/viper`).
2. Enable reading configuration from **environment variables** and config files (`config.yaml`).
3. Implement a configuration command to **set and manage email credentials** and other settings.

**Deliverables:**
- Configuration support for SMTP credentials, default recipients, and other settings.
- CLI command to manage configurations.

---

## **Phase 7: File Compression and Email Attachments**
### **Module 7: File Attachment and Compression**
1. Add support for **multiple file attachments**.
2. Implement **file compression** for large attachments using `archive/zip` or `compress/gzip`.
3. Auto-compress files that exceed a predefined size before sending.

**Deliverables:**
- Ability to attach multiple files to emails.
- Auto-compression for large files.

---

## **Phase 8: Logging & Error Handling**
### **Module 8: Logging and Error Handling**
1. Implement a logging mechanism to **log all sent emails and commands** executed.
2. Create a new command for **viewing logs** (`cobra-cli add logs`).
3. Improve **error handling** and return meaningful error messages for failed operations.

**Deliverables:**
- Logging support for email and command activities.
- Error handling mechanism with retries.

---

## **Phase 9: Email Scheduling & Delayed Execution**
### **Module 9: Scheduled and Delayed Emails**
1. Use Go’s `time` package to add functionality for **delayed or scheduled email sending**.
2. Create a scheduler that allows users to specify email sending after a certain duration or at a specific time.

**Deliverables:**
- Support for scheduling and delaying email notifications.

---

## **Phase 10: Security and Encryption**
### **Module 10: Security & Email Encryption**
1. Add **PGP encryption** support for emails, ensuring that sensitive data is securely sent.
2. Implement key management for **encryption and decryption**.

**Deliverables:**
- Optional email encryption support for sensitive information.

---

## **Phase 11: Testing & Optimization**
### **Module 11: Testing and Refinements**
1. Write **unit tests** for each module (use `testing` package).
2. **Optimize the code** for performance, especially around script execution and file watching.
3. Implement **retry mechanisms** for failed email sends.

**Deliverables:**
- Test coverage for major features.
- Optimized and stable CLI tool.

---

## **Phase 12: Documentation and Release**
### **Module 12: Documentation and Final Release**
1. Finalize and update **documentation** with examples for all commands.
2. Prepare **GitHub repository** for public release, including versioning and licensing.
3. Create a **build and release process** using GitHub Releases or a similar tool.

**Deliverables:**
- Comprehensive documentation and examples.
- Public release of Dhanu CLI tool.

---

### **Development Roadmap (Timeline)**

| **Phase**           | **Module**                            | **Estimated Duration** |
|---------------------|---------------------------------------|------------------------|
| **Phase 1**         | Project Initialization                | 1 day                  |
| **Phase 2**         | Email Functionality                   | 3-4 days               |
| **Phase 3**         | Script Execution                      | 3-4 days               |
| **Phase 4**         | Advanced Email Options                | 5-7 days               |
| **Phase 5**         | File Watcher & Notifications          | 4-5 days               |
| **Phase 6**         | Configuration Management              | 3 days                 |
| **Phase 7**         | File Attachment & Compression         | 4 days                 |
| **Phase 8**         | Logging and Error Handling            | 3-4 days               |
| **Phase 9**         | Scheduled & Delayed Emails            | 3 days                 |
| **Phase 10**        | Security & Email Encryption           | 5-7 days               |
| **Phase 11**        | Testing and Optimizations             | 5-7 days               |
| **Phase 12**        | Documentation & Final Release         | 3-5 days               |
