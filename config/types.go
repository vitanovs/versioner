package config

import (
	"errors"
	"fmt"
	"strings"
)

// PostgresConfig defines the tool database configuration.
type PostgresConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	SslMode  string `toml:"ssl_mode"`
}

func (c *PostgresConfig) validate() error {
	if c.Host == "" {
		return errors.New("invalid host ''")
	}

	if c.Database == "" {
		return errors.New("invalid database ''")
	}

	if c.Username == "" {
		return errors.New("invalid username ''")
	}

	return nil
}

// SchemaConfig defines the schema migration resources
// configurations.
type SchemaConfig struct {
	Directory string   `toml:"directory"`
	Sequence  []string `toml:"sequence"`
}

func (c *SchemaConfig) validate() error {
	if c.Directory == "" {
		return errors.New("invalid directory ''")
	}

	for id, migration := range c.Sequence {
		if strings.Contains(migration, " ") {
			return fmt.Errorf("invalid migration %d: name '%s' contains spaces",
				id,
				migration)
		}
	}

	return nil
}
