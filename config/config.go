package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

const (
	// supportedConfigVersion defines the currently supported
	// configuration version.
	supportedConfigVersion = 1
)

// Config defines the tool configuration.
type Config struct {
	// Version defines the configuration version.
	Version int `toml:"version"`
	// Postgres defines the PostgreSQL database configuration.
	Postgres PostgresConfig `toml:"postgres"`
	// Schema defines the schema resources configuration.
	Schema SchemaConfig `toml:"schema"`
}

// LoadConfig loads configuration from the provided path.
func LoadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, fmt.Errorf("failed decoding %s: %s", path, err)
	}

	if config.Version != supportedConfigVersion {
		return nil, fmt.Errorf("unsupported configuration version %d", config.Version)
	}

	if err = config.Postgres.validate(); err != nil {
		return nil, fmt.Errorf("invalid Postgres configuration: %s", err)
	}

	if err = config.Schema.validate(); err != nil {
		return nil, fmt.Errorf("invalid schema resources configuration: %s", err)
	}

	return &config, nil
}
