# Dhanu Project Documentation

## Overview

Dhanu is a command-line tool designed to manage and send emails with support for configuration, attachments, and more. This project aims to provide a modular, easy-to-use email service with a CLI interface. The project is built in **Go** using **Cobra** for command-line operations and other custom packages for services and utilities.

---

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Commands](#commands)
    - [Config Command](#config-command)
    - [Send Command](#send-command)
- [Makefile](#makefile)
- [Build Process](#build-process)
- [Contributing](#contributing)
- [License](#license)

---

## Project Structure

```
.
├── Changelog.md
├── LICENSE
├── Makefile
├── Plan.md
├── Readme2.md
├── Roadmap.md
├── Structure.md
├── build
│   ├── linux
│   ├── mac
│   └── win
├── cmd
│   ├── config.go
│   ├── root.go
│   └── send.go
├── config.yaml
├── docs
├── examples
├── go.mod
├── go.sum
├── internals
│   ├── services
│   │   ├── EmailService.go
│   │   └── EmailServiceInterface.go
│   └── utils
│       ├── HandleAttachments.go
│       └── Validator.go
├── main.go
├── pkgs
│   ├── configs
│   │   ├── LoadSendConfigs.go
│   │   └── SaveConfigs.go
│   └── logger
└── tree.txt
```

---

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/lordofthemind/dhanu.git
    cd dhanu
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

3. Build the project using Makefile:
    ```bash
    make build
    ```

4. Run the project:
    ```bash
    ./dhanu
    ```

---

## Configuration

To configure the project for sending emails, you need to either pass configurations directly via CLI or edit the configuration file located at `$HOME/.config/dhanu/dhanu.yaml`.

The `config.yaml` file contains details like:
- SMTP Host
- SMTP Port
- From Email (the email that sends the email)
- Credentials (app password)
- Default Recipient

These can be set up either through the CLI or by manually editing the YAML file.

---

## Usage

### Commands

Dhanu provides several commands to manage configuration and send emails:

#### Config Command

The `config` command is used to update and manage the email configuration. You can display, modify, or initiate the configuration setup.

Usage:
```bash
dhanu config
```

Flags:
- `-S`, `--show`: Display the saved configuration.
- `-H`, `--host`: Specify the SMTP host.
- `-P`, `--port`: Specify the SMTP port.
- `-F`, `--from-email`: Set the email address from which emails will be sent.
- `-D`, `--default-recipient`: Set a default recipient email address.
- `-C`, `--credentials`: Set the credentials (app password) for the email.

Example:
```bash
dhanu config --show
dhanu config -F my_email@example.com -C my_password -H smtp.gmail.com -P 465 -D default_recipient@example.com
```

#### Send Command

The `send` command is used to send emails with optional attachments. You can specify the recipient, subject, body, and attachments.

Usage:
```bash
dhanu send
```

Flags:
- `-t`, `--to`: Recipient email address.
- `-s`, `--subject`: Email subject.
- `-b`, `--body`: Email body content.
- `-f`, `--body-file`: Path to a file containing the email body.
- `-a`, `--attachments`: List of file paths or directories to attach to the email.

Example:
```bash
dhanu send -t recipient@example.com -s "Test Subject" -b "This is a test email."
```

Attachments:
```bash
dhanu send -t recipient@example.com -s "Email with Attachment" -b "Please find the attachment." -a /path/to/file.pdf
```

---

## Makefile

The project includes a `Makefile` to simplify building the project for different operating systems. Available commands include:

- **build**: Build the project for all platforms.
  ```bash
  make build
  ```

- **clean**: Remove all build artifacts.
  ```bash
  make clean
  ```

---

## Build Process

To build the project for different platforms (Linux, macOS, Windows):

```bash
make build
```

Build artifacts will be generated in the `build` directory under respective OS folders:
- `build/linux`
- `build/mac`
- `build/win`

---

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Create a new Pull Request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.