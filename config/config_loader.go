package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// LoadConfig loads the configuration based on the APP_ENV environment variable.
func LoadConfig(configPaths []string) (Config, error) {
	var config Config

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, fmt.Errorf("failed to read primary configuration: %v", err)
		}
	}

	// Set the configuration file based on the environment (ex: application_local.yaml)
	viper.SetConfigName("application_" + env)

	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, fmt.Errorf("failed to read environment-specific configuration: %v", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal configuration: %v", err)
	}

	return config, nil
}
