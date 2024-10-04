package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	SMTP struct {
		Host        string `mapstructure:"host"`
		Port        int    `mapstructure:"port"`
		FromEmail   string `mapstructure:"from_email"`  // Updated from Username to FromEmail
		Credentials string `mapstructure:"credentials"` // Updated from Password to Credentials
	} `mapstructure:"smtp"`
	DefaultRecipient string `mapstructure:"default_recipient"`
	SetupCompleted   bool   `mapstructure:"setup_completed"` // New field to track if setup is completed
}

// LoadConfig loads the configuration and returns the config and the path
func LoadConfig() (Config, string, error) {
	var config Config

	// Determine the OS and set the default config path
	var configPath string
	if os.Getenv("DHANU_CONFIG") != "" {
		configPath = os.Getenv("DHANU_CONFIG")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return config, "", err
		}

		// Set config path based on the OS
		if runtime.GOOS == "windows" {
			configPath = filepath.Join(homeDir, "AppData", "Roaming", "dhanu", "dhanu.yaml") // Windows
		} else {
			configPath = filepath.Join(homeDir, ".config", "dhanu", "dhanu.yaml") // Linux
		}
	}

	// Check if config file exists, if not, create a new one
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Configuration file not found, creating a new one...")

		// Set default values
		config.SMTP.Host = ""
		config.SMTP.Port = 0
		config.SMTP.FromEmail = ""
		config.SMTP.Credentials = ""
		config.DefaultRecipient = ""
		config.SetupCompleted = false // Mark setup as incomplete

		// Save the default config to file
		if err := SaveConfig(config, configPath); err != nil {
			return config, "", err
		}

		fmt.Println("Default configuration file created at:", configPath)
	}

	// Load the config file using Viper
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return config, "", err
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, "", err
	}

	return config, configPath, nil
}
