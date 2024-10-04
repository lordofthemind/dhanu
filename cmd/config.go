package cmd

import (
	"fmt"

	"github.com/lordofthemind/dhanu/pkgs/configs"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update configuration settings",
	Long: `Use these commands to update your configuration settings, for example:

dhanu config --port 465
dhanu config --host smtp.yahoo.com
dhanu config --from-email your_email@example.com
dhanu config --credentials your_app_password
dhanu config --default-recipient my-email@example.com
dhanu config --from-email my_email@example.com --credentials my_password --host smtp.gmail.com --port 465 --default-recipient my-email@example.com`,

	Run: func(cmd *cobra.Command, args []string) {
		updateConfig(cmd)
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

	// First time setup for credentials
	if config.SetupCompleted {
		fmt.Println("Initial setup required. Please provide following details.")
		fmt.Print("From Email (the email from which the emails will be sent): ")
		fmt.Scanln(&config.SMTP.FromEmail)
		fmt.Print("Credential (app password of the above provided email): ")
		fmt.Scanln(&config.SMTP.Credentials)
		fmt.Print("Port (the port number which will be use by smtp of the above provided email): ")
		fmt.Scanln(&config.SMTP.Port)
		fmt.Print("Host (the host of the above provided email): ")
		fmt.Scanln(&config.SMTP.Host)
		fmt.Print("Default recipient (this email address will be used as default reciepient in case no reciepient provided): ")
		fmt.Scanln(&config.SMTP.Credentials)
		fmt.Print("Confirm the details above or you want to try again, y or no: ")
		// put condition here
	}

	// Update fields based on flags
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
