package config

import (
	"github.com/spf13/viper"
	"os"
)

// LoadConfig loads the configuration based on the APP_ENV environment variable.
func LoadConfig(configPaths []string) (Config, error) {
	var config Config

	// Determine the environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // Default to 'local'
	}

	// Set the configuration file based on the environment
	viper.SetConfigName("./application_" + env)
	viper.SetConfigType("yaml") // Let's stop making more of these...
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
