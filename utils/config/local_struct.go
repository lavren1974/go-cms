package utils

// AppConfig represents the application settings
type LocalAppConfig struct {
	Name     string `toml:"name"`
	Port     string `toml:"port"`
	Theme    string `toml:"theme"`
	LogLevel string `toml:"log_level"`
}

// DatabaseConfig represents the database settings
type LocalDatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// Config represents the overall structure of the config files
type LocalConfig struct {
	App      LocalAppConfig      `toml:"app"`
	Database LocalDatabaseConfig `toml:"database"`
}
