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
dhanu config --username your_username
dhanu config --password your_app_password
dhanu config --default-recipient my-email@example.com
dhanu config --username my_username --password my_password --host smtp.gmail.com --port 465 --default-recipient my-email@example.com`,

	Run: func(cmd *cobra.Command, args []string) {
		updateConfig(cmd) // Pass the cmd object here
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Define flags for updating configuration
	configCmd.Flags().IntP("port", "P", 0, "SMTP port")
	configCmd.Flags().StringP("host", "H", "", "SMTP host")
	configCmd.Flags().StringP("username", "u", "", "SMTP username")
	configCmd.Flags().StringP("password", "p", "", "SMTP password")
	configCmd.Flags().StringP("default-recipient", "d", "", "Default recipient email")

}

// Function to update the configuration
func updateConfig(cmd *cobra.Command) {
	// Load existing config
	config, configPath, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	// Update fields based on flags
	if username, _ := cmd.Flags().GetString("username"); username != "" {
		config.SMTP.Username = username
	}
	if password, _ := cmd.Flags().GetString("password"); password != "" {
		config.SMTP.Password = password
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
