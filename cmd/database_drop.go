package cmd

import (
	"context"
	"fmt"

	psql "github.com/hall-arranger/versioner/api/postgresql"
	"github.com/hall-arranger/versioner/log"
	"github.com/urfave/cli"
)

// NewDatabaseDropCommand returns a command
// used for dropping databases on remote endpoint.
func NewDatabaseDropCommand() cli.Command {
	return cli.Command{
		Name:   "drop",
		Usage:  "drops database on remote endpoint",
		Action: runDatabaseDropCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "name, n",
				Usage:    "the name of the database to drop",
				Required: true,
			},
		},
	}
}

func runDatabaseDropCommand(ctx *cli.Context) error {
	cmdName := ctx.Command.Name
	log.Info("Start '%s' command", cmdName)

	name := ctx.String("name")
	if name == "" {
		msg := fmt.Sprintf("No new database name specified. See '%s --help'", cmdName)
		return cli.NewExitError(msg, 1)
	}

	runContext := context.Background()
	clientConfig, err := loadClientConfig(ctx)
	if err != nil {
		msg := fmt.Sprintf("Failed to client configuration: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Initializing database client...")
	client, err := psql.NewClient(runContext, clientConfig)
	if err != nil {
		msg := fmt.Sprintf("Failed to initialize database client: %s", err)
		return cli.NewExitError(msg, 1)
	}

	if _, err := client.DropDatabase(runContext, name); err != nil {
		msg := fmt.Sprintf("Failed to drop database '%s': %s", name, err)
		return cli.NewExitError(msg, 1)
	}
	log.Info("Database '%s' was dropped", name)

	log.Info("Closing database client...")
	if err := client.Close(); err != nil {
		msg := fmt.Sprintf("Failed to close database client connection: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Finish '%s' command", cmdName)
	return nil
}
