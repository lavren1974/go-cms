// utils/config.go
package utils

import (
	"github.com/BurntSushi/toml"
)

// LoadConfig reads and parses a TOML file
func LoadConfigGlobal(path string) (*GlobalConfig, error) {
	var config GlobalConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadConfigLocal(path string) (*LocalConfig, error) {
	var config LocalConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
