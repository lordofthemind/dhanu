// pkgs/configs/load.go
package configs

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	SMTP struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"smtp"`
}

func LoadConfig() (Config, error) {
	var config Config

	// Determine the OS and set the default config path
	var configPath string
	if os.Getenv("DHANU_CONFIG") != "" {
		configPath = os.Getenv("DHANU_CONFIG")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return config, err
		}

		// Set config path based on the OS
		if runtime.GOOS == "windows" {
			configPath = filepath.Join(homeDir, "AppData", "Roaming", "dhanu", "dhanu.yaml") // Windows
		} else {
			configPath = filepath.Join(homeDir, ".config", "dhanu", "dhanu.yaml") // Linux
		}
	}

	viper.SetConfigFile(configPath) // Use the config file path

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
