package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// SaveConfig saves the configuration to the specified file path
func SaveConfig(config Config, configPath string) error {
	viper.Set("smtp.host", config.SMTP.Host)
	viper.Set("smtp.port", config.SMTP.Port)
	viper.Set("smtp.from_email", config.SMTP.FromEmail)    // Updated field name
	viper.Set("smtp.credentials", config.SMTP.Credentials) // Updated field name
	viper.Set("default_recipient", config.DefaultRecipient)
	viper.Set("setup_completed", config.SetupCompleted) // Track setup completion

	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		return err
	}

	// Write the config file
	viper.SetConfigFile(configPath)
	if err := viper.WriteConfigAs(configPath); err != nil {
		return err
	}

	return nil
}
