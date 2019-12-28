package config

import (
	"errors"
	"fmt"
	"strings"
)

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
