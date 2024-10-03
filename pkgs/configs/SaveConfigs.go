package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Function to save the configuration to the file
func SaveConfig(config Config, configPath string) error {
	viper.Set("smtp.host", config.SMTP.Host)
	viper.Set("smtp.port", config.SMTP.Port)
	viper.Set("smtp.username", config.SMTP.Username)
	viper.Set("smtp.password", config.SMTP.Password)
	viper.Set("default_recipient", config.DefaultRecipient)

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
