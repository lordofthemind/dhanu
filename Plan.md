Getting started with building **Dhanu** can be an exciting and rewarding process. Here's a step-by-step guide to help you begin developing your CLI tool:

---

## **1. Set Up Your Development Environment**

### **Install Go (Golang)**

- **Download and Install:**
  - Visit the [official Go website](https://golang.org/dl/) and download the appropriate installer for your operating system.
  - Follow the installation instructions provided.

- **Verify Installation:**
  ```bash
  go version
  ```
  Ensure that the command outputs the installed Go version, confirming a successful installation.

### **Set Up Your Workspace**

- **Configure `GOPATH`:**
  - Although Go modules make `GOPATH` less critical, it's good practice to set it.
  - For Unix/Linux:
    ```bash
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
  - Add these lines to your shell profile (`~/.bashrc`, `~/.zshrc`) to make them persistent.

---

## **2. Initialize Your Project**

### **Create Project Directory**

- **Create and Navigate:**
  ```bash
  mkdir dhanu
  cd dhanu
  ```

### **Initialize Go Modules**

- **Initialize Module:**
  ```bash
  go mod init github.com/yourusername/dhanu
  ```
  Replace `yourusername` with your GitHub username or desired repository name.

---

## **3. Add Cobra for CLI Functionality**

### **Install Cobra CLI Tool**

- **Install Globally:**
  ```bash
  go install github.com/spf13/cobra-cli@latest
  ```
  Ensure `$GOPATH/bin` is in your system `PATH`.

### **Initialize Cobra Application**

- **Create Base Application:**
  ```bash
  cobra-cli init
  ```
  This generates the main application files (`main.go`, `cmd/root.go`).

### **Verify the Structure**

- **Check Generated Files:**
  - `main.go`: The entry point of your application.
  - `cmd/root.go`: The root command definition.

---

## **4. Plan Your Command Structure**

### **Identify Commands and Subcommands**

- **Commands to Implement:**
  - `send`: Send emails with optional attachments.
  - `run`: Execute scripts and send outputs.
  - `watch`: Monitor files or directories for changes.
  - `logs`: View and manage logs.

### **Create Subcommands**

- **Generate Commands:**
  ```bash
  cobra-cli add send
  cobra-cli add run
  cobra-cli add watch
  cobra-cli add logs
  ```
- **Review Created Files:**
  - Each command will have its own file in the `cmd/` directory (e.g., `cmd/send.go`).

---

## **5. Implement Core Features Step by Step**

### **a. Email Sending Capability**

- **Choose an Email Package:**
  - **Option 1: `gomail`**
    ```bash
    go get gopkg.in/gomail.v2
    ```
  - **Option 2: Standard Library `net/smtp`**
    - No installation needed; part of Go's standard library.

- **Create Email Utility Functions:**
  - Abstract email sending logic into reusable functions.
  - Handle authentication, email formatting, and error checking.

### **b. File Attachments**

- **Implement Attachment Handling:**
  - Modify email functions to accept file paths and attach them.
  - Ensure multiple files can be attached in one email.

- **Auto-Compress Large Files:**
  - Use Go's `archive/zip` or `compress/gzip` packages.
  - Check file sizes and compress if they exceed a threshold.

### **c. Script Execution and Output Capture**

- **Execute Scripts or Commands:**
  - Use the `os/exec` package.
    ```go
    import "os/exec"

    func executeScript(scriptPath string) (string, error) {
        cmd := exec.Command("bash", scriptPath)
        output, err := cmd.CombinedOutput()
        return string(output), err
    }
    ```

- **Capture Output:**
  - Collect both `stdout` and `stderr`.
  - Include output in the email body or as an attachment.

---

## **6. Configuration Management with Viper**

### **Install Viper**

```bash
go get github.com/spf13/viper
```

### **Set Up Configuration**

- **Create a Configuration File:**
  - Name it `config.yaml` or `dhanurc.yaml`.
  - Define default settings and credentials.

- **Load Configuration in Code:**
  ```go
  import "github.com/spf13/viper"

  func initConfig() {
      viper.SetConfigName("config")
      viper.AddConfigPath(".")
      err := viper.ReadInConfig()
      if err != nil {
          fmt.Println("Error reading config file", err)
      }
  }
  ```

- **Use Environment Variables:**
  - Viper can automatically read environment variables.
    ```go
    viper.AutomaticEnv()
    ```

---

## **7. Implement Additional Features**

### **Scheduled/Delayed Emails**

- **Use Go's `time` Package:**
  ```go
  time.Sleep(10 * time.Minute) // Delay execution by 10 minutes
  ```

- **Implement Scheduling Logic:**
  - Allow users to specify delay durations.
  - Parse duration strings (e.g., "10m", "2h").

### **Retry Mechanism for Failed Emails**

- **Implement Retry Logic:**
  ```go
  for i := 0; i < retryCount; i++ {
      err := sendEmail()
      if err == nil {
          break
      }
      time.Sleep(interval)
  }
  ```

- **Use Exponential Backoff (Optional):**
  - Increase wait time between retries progressively.

### **File Watcher Mode**

- **Install `fsnotify`:**
  ```bash
  go get github.com/fsnotify/fsnotify
  ```

- **Set Up File Watching:**
  ```go
  import "github.com/fsnotify/fsnotify"

  func watchFile(filePath string) {
      watcher, err := fsnotify.NewWatcher()
      // Handle errors and events
  }
  ```

### **Encryption (Optional)**

- **Integrate PGP Encryption:**
  - Use packages like `golang.org/x/crypto/openpgp`.

- **Handle Key Management:**
  - Allow users to specify paths to public keys.

---

## **8. Logging and Error Handling**

### **Implement Logging**

- **Use Standard Library `log`:**
  ```go
  import "log"

  log.Println("This is a log message")
  ```

- **Or Install `logrus` for Advanced Features:**
  ```bash
  go get github.com/sirupsen/logrus
  ```

### **Error Handling Best Practices**

- **Provide Clear Messages:**
  - Return informative errors to the user.
  - Use error wrapping for context.

- **Centralize Error Handling:**
  - Create helper functions to handle common error scenarios.

---

## **9. Testing**

### **Write Unit Tests**

- **Use Go's Testing Package:**
  ```go
  import "testing"

  func TestSendEmail(t *testing.T) {
      // Test cases here
  }
  ```

- **Run Tests:**
  ```bash
  go test ./...
  ```

### **Test Email Functionality**

- **Set Up a Test SMTP Server:**
  - Use tools like [MailHog](https://github.com/mailhog/MailHog) or [Mailtrap](https://mailtrap.io/).

- **Mock SMTP Server in Tests:**
  - Avoid sending real emails during testing.

---

## **10. Documentation and Examples**

### **Keep Documentation Updated**

- **Update `README.md`:**
  - Provide an overview, installation instructions, and examples.

- **Create Additional Docs:**
  - Use `/docs` directory for detailed guides.

### **Provide Usage Examples**

- **Add Examples in Documentation:**
  - Show how to use each command and flag.

- **Consider a `examples/` Directory:**
  - Include example scripts and config files.

---

## **11. Version Control with Git**

### **Initialize Git Repository**

```bash
git init
git add .
git commit -m "Initial commit"
```

### **Create `.gitignore` File**

- **Exclude Unnecessary Files:**
  ```gitignore
  /bin/
  /vendor/
  *.exe
  *.log
  ```

### **Push to Remote Repository**

```bash
git remote add origin https://github.com/yourusername/dhanu.git
git push -u origin master
```

---

## **12. Incremental Development**

### **Start with Core Features**

- **Implement Basic Email Sending:**
  - Ensure you can send a simple email before adding complexity.

- **Add Features Gradually:**
  - Tackle one feature at a time, such as attachments, then script execution.

### **Use Feature Branches**

- **Create New Branches for Features:**
  ```bash
  git checkout -b feature/attachments
  ```
- **Merge Back into `master` After Testing:**
  ```bash
  git checkout master
  git merge feature/attachments
  ```

---

## **13. Community and Contributions**

### **Set Up Contribution Guidelines**

- **Create `CONTRIBUTING.md`:**
  - Outline how others can contribute.

### **Issue Tracking**

- **Use GitHub Issues:**
  - Track bugs, feature requests, and enhancements.

- **Label and Organize Issues:**
  - Use labels like `bug`, `enhancement`, `question`.

---

## **14. Licensing**

### **Choose a License**

- **Add an Open-Source License:**
  - MIT License is a good starting point.
  - Create a `LICENSE` file with the license text.

---

## **15. Build and Release**

### **Build Executables**

- **Build for Your OS:**
  ```bash
  go build
  ```
  - The executable will be named `dhanu` (or `dhanu.exe` on Windows).

- **Cross-Compile for Other OSes:**
  ```bash
  # For Windows
  GOOS=windows GOARCH=amd64 go build -o dhanu.exe

  # For Linux
  GOOS=linux GOARCH=amd64 go build -o dhanu-linux
  ```

### **Create Releases**

- **Use GitHub Releases:**
  - Tag versions in Git.
    ```bash
    git tag -a v1.0.0 -m "First release"
    git push origin v1.0.0
    ```
  - Upload compiled binaries to the release.

---

## **16. Feedback and Iteration**

### **Collect Feedback**

- **Share with Peers:**
  - Get feedback from friends or colleagues.

- **Monitor Issues and Pull Requests:**
  - Engage with the community.

### **Iterate Based on Feedback**

- **Prioritize Feature Requests:**
  - Focus on the most requested features or bug fixes.

- **Improve Documentation:**
  - Clarify any confusing areas identified by users.

---

## **Common Issues and Solutions**

- **Email Sending Failures:**
  - **Cause:** Incorrect SMTP settings or authentication errors.
  - **Solution:** Double-check SMTP server details and credentials. Ensure that less secure app access is enabled if necessary.

- **Permission Denied When Executing Scripts:**
  - **Cause:** Script lacks execute permissions.
  - **Solution:** Modify permissions.
    ```bash
    chmod +x script.sh
    ```

- **Environment Variables Not Loaded:**
  - **Cause:** Variables not exported or available in the execution context.
  - **Solution:** Export variables in the shell or include them in your config file.

- **Dependencies Not Found:**
  - **Cause:** Missing or outdated Go packages.
  - **Solution:** Run `go get -u ./...` to update dependencies.

---

## **Additional Tips**

- **Code Organization:**
  - Keep your code modular by separating concerns into different packages or files.

- **Consistent Coding Style:**
  - Use tools like `go fmt` to format your code.

- **Error Logging:**
  - Log errors to both the console and a log file for easier debugging.

- **Security Considerations:**
  - Avoid hardcoding sensitive information.
  - Consider using a secrets manager or secure storage.

- **Stay Informed:**
  - Keep up with updates to Go, Cobra, and other dependencies.

---

## **Useful Resources**

- **Go Documentation:** [https://golang.org/doc/](https://golang.org/doc/)
- **Cobra Documentation:** [https://github.com/spf13/cobra](https://github.com/spf13/cobra)
- **Viper Documentation:** [https://github.com/spf13/viper](https://github.com/spf13/viper)
- **Golang Example Projects:** Explore projects on GitHub for inspiration.

---

## **Next Steps**

1. **Set Milestones:**
   - Break down your project into manageable tasks with deadlines.

2. **Engage with the Community:**
   - Join Go and open-source communities for support and collaboration.

3. **Consider Continuous Integration:**
   - Set up CI/CD pipelines with tools like GitHub Actions.

4. **Plan for Future Features:**
   - Think about scalability, additional protocols (like SMS notifications), or integrations.

