# Dhanu CLI Documentation

**Dhanu** is a CLI tool built using Golang and Cobra for sending emails with files and handling script outputs. The name "Dhanu" is derived from the Sanskrit word for bow, symbolizing strength and focus.

## Features

### 1. **Send Script Output via Email**
   - Capture the output of any script or command and send it via email to the configured recipient.
   - Set a default email address for sending outputs, so users don't need to specify it every time.
   - Option to send email notifications only on script success or failure.
   
   **Command Example:**
   ```bash
   dhanu run "script.sh" --to recipient@example.com --subject "Script Output" --on-success
   ```

### 2. **File Attachments**
   - Send one or more files as attachments in an email.
   - Automatically compress large files before sending.

   **Command Example:**
   ```bash
   dhanu send --to recipient@example.com --subject "File Report" --body "Attached are the files" --file ./output.txt
   ```

### 3. **Scheduled/Delayed Emails**
   - Schedule emails to be sent at a specific time or with a delay after script completion.

   **Command Example:**
   ```bash
   dhanu send --to recipient@example.com --subject "Report" --file ./output.txt --delay 10m
   ```

### 4. **Configurable Email Templates**
   - Create and use predefined email templates for common emails such as log reports, error notifications, etc.
   - Customize subject lines and body text with dynamic variables like date, time, and script name.

   **Command Example:**
   ```bash
   dhanu send --template error-report --to recipient@example.com --script-name "data_backup.sh"
   ```

### 5. **Built-in Logging**
   - Maintain a log of all emails sent, including timestamp, recipients, and attachments.
   - Option to attach logs or error reports to the email.

   **Command Example:**
   ```bash
   dhanu logs --attach
   ```

### 6. **Custom Email Recipients and CC/BCC Support**
   - Send emails to multiple recipients and optionally use CC/BCC fields.

   **Command Example:**
   ```bash
   dhanu send --to primary@example.com --cc manager@example.com --bcc audit@example.com --file ./report.txt
   ```

### 7. **Email Notifications for Long-running Scripts**
   - Notify the user when a script starts and completes execution.
   - Send interim progress updates for long-running tasks.

   **Command Example:**
   ```bash
   dhanu run "long-script.sh" --to recipient@example.com --notify-progress 10m
   ```

### 8. **File Watcher Mode**
   - Watch a file or directory for changes and send an email when updates occur (e.g., log file changes).

   **Command Example:**
   ```bash
   dhanu watch --file ./log.txt --to admin@example.com
   ```

### 9. **Environment Variable Integration**
   - Securely read email credentials (SMTP server, email, password) from environment variables or a configuration file.
   - Use CLI flags to override these defaults.

   **Command Example:**
   ```bash
   export DHANU_EMAIL_USER="your-email@example.com"
   export DHANU_EMAIL_PASS="password"
   dhanu send --to recipient@example.com --file ./report.txt
   ```

### 10. **Retry Mechanism for Failed Emails**
   - Automatically retry sending emails in case of failure due to network issues or server downtime.

   **Command Example:**
   ```bash
   dhanu send --to recipient@example.com --file ./output.txt --retry 3 --interval 5m
   ```

### 11. **Interactive Mode**
   - Step through email creation interactively, ideal for beginners.

   **Command Example:**
   ```bash
   dhanu send --interactive
   ```

### 12. **Encryption (Optional)**
   - Support for sending encrypted emails using standards like PGP or S/MIME.

   **Command Example:**
   ```bash
   dhanu send --to recipient@example.com --file ./secret.txt --encrypt --pgp-key ./recipient_key.asc
   ```

### 13. **Verbose or Silent Mode**
   - Control the verbosity level to show detailed execution logs or run silently.

   **Command Example:**
   ```bash
   dhanu run "script.sh" --silent
   dhanu run "script.sh" --verbose
   ```

---

## Commands

### `dhanu send`
Send an email with optional file attachments and custom message content.

```bash
dhanu send --to recipient@example.com --subject "Subject" --body "Message body" --file ./path/to/file
```

| Flag            | Description                                      |
| --------------- | ------------------------------------------------ |
| `--to`          | Recipient email address.                         |
| `--cc`          | CC recipients.                                   |
| `--bcc`         | BCC recipients.                                  |
| `--subject`     | Email subject.                                   |
| `--body`        | Email body.                                      |
| `--file`        | Path to file attachment(s).                      |
| `--template`    | Use a predefined email template.                 |
| `--retry`       | Number of retries if the email fails.            |
| `--delay`       | Delay sending email by the specified duration.   |
| `--encrypt`     | Encrypt email before sending (PGP/S-MIME).       |
| `--silent`      | Run silently without displaying output.          |
| `--verbose`     | Display detailed logs during execution.          |

---

### `dhanu run`
Run a script and send its output as an email.

```bash
dhanu run "script.sh" --to recipient@example.com --on-success
```

| Flag            | Description                                      |
| --------------- | ------------------------------------------------ |
| `--to`          | Recipient email address.                         |
| `--on-success`  | Send email only if the script succeeds.          |
| `--on-failure`  | Send email only if the script fails.             |
| `--notify-progress` | Send progress updates every specified duration. |

---

### `dhanu watch`
Watch a file or directory for changes and notify via email.

```bash
dhanu watch --file ./log.txt --to recipient@example.com
```

| Flag            | Description                                      |
| --------------- | ------------------------------------------------ |
| `--file`        | File or directory to watch.                      |
| `--to`          | Recipient email address.                         |

---

### `dhanu logs`
View or send logs related to email transactions.

```bash
dhanu logs --attach
```

| Flag            | Description                                      |
| --------------- | ------------------------------------------------ |
| `--attach`      | Attach log file to email.                        |

---

## Configuration

### Environment Variables
Set the following environment variables for security and ease of use:

```bash
export DHANU_EMAIL_USER="your-email@example.com"
export DHANU_EMAIL_PASS="password"
export DHANU_SMTP_SERVER="smtp.example.com"
```

### Configuration File (Optional)
You can also use a configuration file (`dhanurc.yml`) to store defaults:

```yaml
email:
  user: "your-email@example.com"
  password: "password"
  smtp_server: "smtp.example.com"
default_recipient: "default@example.com"
```

---

## Installation

To install **Dhanu**, clone the repository and build the project:

```bash
git clone https://github.com/yourusername/dhanu.git
cd dhanu
go build
```

---

## License

Dhanu is licensed under the [MIT License](LICENSE).

---

## Contribution

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for more information.
```

This Markdown file provides a detailed and well-structured documentation for Dhanu, covering all key features, commands, and configuration options. Let me know if you'd like any further customization!