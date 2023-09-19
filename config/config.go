// config/config.go

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	// Add other configuration fields as needed
}

var AppConfig Config

func LoadConfig() {
	// Initialize Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config") // You can specify additional config paths

	// Optionally, you can set environment variables with a specific prefix
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()

	// Read the configuration file (e.g., config.yaml)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config file: %s", err))
	}

	// Unmarshal the configuration into the AppConfig struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %s", err))
	}
}

// GetDatabaseURL returns the database URL from the configuration
func (c *Config) GetDatabaseURL() string {
	return c.DatabaseURL
}
