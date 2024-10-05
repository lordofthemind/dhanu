package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lordofthemind/dhanu/internals/utils"
	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update configuration settings",
	Long: `Use these commands to update your configuration settings, for example:

dhanu config -S, --show saved configs
dhanu config -P, --port 465
dhanu config -H, --host smtp.yahoo.com
dhanu config -F, --from-email your_email@example.com
dhanu config -C, --credentials your_app_password
dhanu config -D, --default-recipient my-email@example.com
dhanu config -F my_email@example.com -C my_password -H smtp.gmail.com -P 465 -D my-email@example.com`,

	Run: func(cmd *cobra.Command, args []string) {
		// Load existing config to check if setup is completed
		config, configPath, err := configs.LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}

		// Check if the user passed the -S flag
		showConfig, _ := cmd.Flags().GetBool("show")

		if showConfig {
			// Display the saved configuration
			displayConfig(&config)
			return
		}

		// If setup is completed, show help message
		if config.SetupCompleted {
			cmd.Help()
		} else {
			// If setup is not completed, initiate first-time setup
			initiateSetup(&config, configPath)
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
	configCmd.Flags().BoolP("show", "S", false, "Show saved configuration") // Ensure the -S flag is properly registered
}

// Function to display the saved configuration
func displayConfig(config *configs.Config) {
	fmt.Println("Saved Configuration:")
	fmt.Printf("From Email: %s\n", config.SMTP.FromEmail)
	fmt.Printf("Credential: %s\n", config.SMTP.Credentials)
	fmt.Printf("Port: %d\n", config.SMTP.Port)
	fmt.Printf("Host: %s\n", config.SMTP.Host)
	fmt.Printf("Default Recipient: %s\n", config.DefaultRecipient)
	fmt.Printf("Setup Completed: %v\n", config.SetupCompleted)
}

// Function to initiate first-time setup
func initiateSetup(config *configs.Config, configPath string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Initial setup required. Please provide the following details:")

	// Prompt for From Email with validation
	for {
		fmt.Print("From Email (the email from which the emails will be sent): ")
		config.SMTP.FromEmail, _ = reader.ReadString('\n')
		config.SMTP.FromEmail = strings.TrimSpace(config.SMTP.FromEmail)

		if utils.IsValidEmail(config.SMTP.FromEmail) {
			break
		} else {
			fmt.Println("Invalid email format. Please enter a valid email address.")
		}
	}

	fmt.Print("Credential (the app password of the above-provided email): ")
	config.SMTP.Credentials, _ = reader.ReadString('\n')
	config.SMTP.Credentials = strings.TrimSpace(config.SMTP.Credentials)

	// Handle port input with validation for integer
	for {
		fmt.Print("Port (the port number to be used by SMTP for the above email): ")
		portInput, _ := reader.ReadString('\n')
		portInput = strings.TrimSpace(portInput)

		port, err := strconv.Atoi(portInput)
		if err != nil {
			fmt.Println("Invalid input. Please enter a numeric value for the port.")
		} else {
			config.SMTP.Port = port
			break
		}
	}

	fmt.Print("Host (the SMTP host for the above-provided email): ")
	config.SMTP.Host, _ = reader.ReadString('\n')
	config.SMTP.Host = strings.TrimSpace(config.SMTP.Host)

	// Prompt for Default Recipient with validation
	for {
		fmt.Print("Default Recipient (this email address will be used as the default recipient if none is provided): ")
		config.DefaultRecipient, _ = reader.ReadString('\n')
		config.DefaultRecipient = strings.TrimSpace(config.DefaultRecipient)

		if utils.IsValidEmail(config.DefaultRecipient) {
			break
		} else {
			fmt.Println("Invalid email format. Please enter a valid email address.")
		}
	}

	// Confirmation prompt
	for {
		fmt.Println("\nPlease confirm the details you entered:")
		fmt.Printf("From Email: %s\n", config.SMTP.FromEmail)
		fmt.Printf("Credential: %s\n", config.SMTP.Credentials)
		fmt.Printf("Port: %d\n", config.SMTP.Port)
		fmt.Printf("Host: %s\n", config.SMTP.Host)
		fmt.Printf("Default Recipient: %s\n", config.DefaultRecipient)
		fmt.Print("Are these details correct? (y/n): ")

		confirmation, _ := reader.ReadString('\n')
		confirmation = strings.TrimSpace(strings.ToLower(confirmation))

		if confirmation == "y" {
			// Set SetupCompleted to true and save config
			config.SetupCompleted = true
			err := configs.SaveConfig(*config, configPath)
			if err != nil {
				fmt.Println("Error saving configuration:", err)
				return
			}
			fmt.Println("Configuration saved successfully.")
			break
		} else if confirmation == "n" {
			// Re-enter the details if the user chooses 'n'
			fmt.Println("Let's re-enter the details.")
			initiateSetup(config, configPath) // Recursively call initiateSetup to re-enter details
		} else {
			fmt.Println("Invalid option. Please type 'y' for yes or 'n' for no.")
		}
	}
}
