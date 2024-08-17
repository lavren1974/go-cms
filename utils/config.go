// utils/config.go
package utils

import (
	"github.com/BurntSushi/toml"
)

// AppConfig represents the application settings
type AppConfig struct {
	Name     string `toml:"name"`
	LogLevel string `toml:"log_level"`
}

// DatabaseConfig represents the database settings
type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// Config represents the overall structure of the config files
type Config struct {
	App      AppConfig      `toml:"app"`
	Database DatabaseConfig `toml:"database"`
}

// LoadConfig reads and parses a TOML file
func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
