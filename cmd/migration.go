package cmd

import (
	"github.com/urfave/cli"
)

// NewMigrationCommand returns new command
// that is used as a parent for all migration
// related commands supported in the tool.
func NewMigrationCommand() cli.Command {
	return cli.Command{
		Name:  "migration",
		Usage: "schema migration commands",
		Subcommands: []cli.Command{
			NewMigrationInitCommand(),
			NewMigrationApplyCommand(),
		},
	}
}
