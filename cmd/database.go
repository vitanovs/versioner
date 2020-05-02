package cmd

import (
	"github.com/urfave/cli"
)

// NewDatabaseCommand returns new command
// that is used as a parent for all database
// related commands supported in the tool.
func NewDatabaseCommand() cli.Command {
	cmd := cli.Command{
		Name:  "database",
		Usage: "database operations command",
		Subcommands: []cli.Command{
			NewDatabaseCreateCommand(),
			NewDatabaseDropCommand(),
		},
	}

	cmd.BashComplete = cli.DefaultCompleteWithFlags(&cmd)
	return cmd
}
