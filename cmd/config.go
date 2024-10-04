package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update configuration settings",
	Long: `Use these commands to update your configuration settings, for example:

dhanu config -P, --port 465
dhanu config -H, --host smtp.yahoo.com
dhanu config -F, --from-email your_email@example.com
dhanu config -C, --credentials your_app_password
dhanu config -D, --default-recipient my-email@example.com
dhanu config -F my_email@example.com -C my_password -H smtp.gmail.com -P 465 -D my-email@example.com`,

	Run: func(cmd *cobra.Command, args []string) {
		// If no flags provided, show help information
		if !cmd.Flags().Changed("port") &&
			!cmd.Flags().Changed("host") &&
			!cmd.Flags().Changed("from-email") &&
			!cmd.Flags().Changed("credentials") &&
			!cmd.Flags().Changed("default-recipient") {
			cmd.Help()
		} else {
			updateConfig(cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Define flags for updating configuration
	configCmd.Flags().IntP("port", "P", 0, "SMTP port")
	configCmd.Flags().StringP("host", "H", "", "SMTP host")
	configCmd.Flags().StringP("from-email", "F", "", "SMTP from_email")
	configCmd.Flags().StringP("credentials", "C", "", "SMTP credentials (password)")
	configCmd.Flags().StringP("default-recipient", "D", "", "Default recipient email")
}

// Function to update the configuration
func updateConfig(cmd *cobra.Command) {
	// Load existing config
	config, configPath, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// First time setup for credentials
	if !config.SetupCompleted {
		fmt.Println("Initial setup required. Please provide the following details:")
		setupInitialConfig(&config, reader)
	} else {
		fmt.Println("You are updating an existing configuration.")
	}

	// Update fields based on flags if provided
	if fromEmail, _ := cmd.Flags().GetString("from-email"); fromEmail != "" {
		config.SMTP.FromEmail = fromEmail
	}
	if credentials, _ := cmd.Flags().GetString("credentials"); credentials != "" {
		config.SMTP.Credentials = credentials
	}
	if host, _ := cmd.Flags().GetString("host"); host != "" {
		config.SMTP.Host = host
	}
	if port, _ := cmd.Flags().GetInt("port"); port != 0 {
		config.SMTP.Port = port
	}
	if recipient, _ := cmd.Flags().GetString("default-recipient"); recipient != "" {
		config.DefaultRecipient = recipient
	}

	// Save updated configuration
	err = configs.SaveConfig(config, configPath)
	if err != nil {
		fmt.Println("Error saving configuration:", err)
		return
	}

	fmt.Println("Configuration updated successfully.")
}

// Helper function to handle initial setup
func setupInitialConfig(config *configs.Config, reader *bufio.Reader) {
	for {
		// Prompt for each field and read user input
		fmt.Print("From Email (the email from which the emails will be sent): ")
		config.SMTP.FromEmail, _ = reader.ReadString('\n')
		config.SMTP.FromEmail = strings.TrimSpace(config.SMTP.FromEmail)

		fmt.Print("Credential (the app password of the above-provided email): ")
		config.SMTP.Credentials, _ = reader.ReadString('\n')
		config.SMTP.Credentials = strings.TrimSpace(config.SMTP.Credentials)

		fmt.Print("Port (the port number to be used by SMTP for the above email): ")
		fmt.Scanln(&config.SMTP.Port)

		fmt.Print("Host (the SMTP host for the above-provided email): ")
		config.SMTP.Host, _ = reader.ReadString('\n')
		config.SMTP.Host = strings.TrimSpace(config.SMTP.Host)

		fmt.Print("Default Recipient (this email address will be used as the default recipient if none is provided): ")
		config.DefaultRecipient, _ = reader.ReadString('\n')
		config.DefaultRecipient = strings.TrimSpace(config.DefaultRecipient)

		// Confirmation prompt
		if confirmDetails(config, reader) {
			config.SetupCompleted = true
			break
		}
	}
}

// Helper function to confirm user details
func confirmDetails(config *configs.Config, reader *bufio.Reader) bool {
	fmt.Println("\nPlease confirm the details you entered:")
	fmt.Printf("From Email: %s\n", config.SMTP.FromEmail)
	fmt.Printf("Credential: %s\n", config.SMTP.Credentials)
	fmt.Printf("Port: %d\n", config.SMTP.Port)
	fmt.Printf("Host: %s\n", config.SMTP.Host)
	fmt.Printf("Default Recipient: %s\n", config.DefaultRecipient)
	fmt.Print("Are these details correct? (y/n): ")

	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(strings.ToLower(confirmation))

	return confirmation == "y"
}
