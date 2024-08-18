package utils

// AppConfig represents the application settings
type GlobalAppConfig struct {
	Name     string `toml:"name"`
	Version  string `toml:"version"`
	LogLevel string `toml:"log_level"`
}

// DatabaseConfig represents the database settings
type GlobalDatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// Config represents the overall structure of the config files
type GlobalConfig struct {
	Cms      GlobalAppConfig      `toml:"cms"`
	Database GlobalDatabaseConfig `toml:"database"`
}
