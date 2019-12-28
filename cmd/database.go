package cmd

import (
	"github.com/urfave/cli"
)

// NewDatabaseCommand returns new command
// that is used as a parent for all database
// related commands supported in the tool.
func NewDatabaseCommand() cli.Command {
	return cli.Command{
		Name:  "database",
		Usage: "database operations command",
		Subcommands: []cli.Command{
			NewDatabaseCreateCommand(),
			NewDatabaseDropCommand(),
		},
	}
}
