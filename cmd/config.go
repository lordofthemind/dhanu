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
	Long: `Use this command to update your configuration settings, for example:

dhanu config --username your_username`,
	Run: func(cmd *cobra.Command, args []string) {
		updateConfig(cmd) // Pass the cmd object here
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Define flags for updating configuration
	configCmd.Flags().StringP("username", "u", "", "SMTP username")
	configCmd.Flags().StringP("password", "p", "", "SMTP password")
	configCmd.Flags().StringP("host", "H", "", "SMTP host")
	configCmd.Flags().IntP("port", "P", 0, "SMTP port")
}

// Function to update the configuration
func updateConfig(cmd *cobra.Command) {
	// Load existing config
	config, err := configs.LoadConfig()
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

	// Save updated configuration
	err = configs.SaveConfig(config)
	if err != nil {
		fmt.Println("Error saving configuration:", err)
		return
	}

	fmt.Println("Configuration updated successfully.")
}
