package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/client"
	"github.com/vitanovs/versioner/log"
)

// NewDatabaseCreateCommand returns a command
// used for creating new databases on remote endpoint.
func NewDatabaseCreateCommand() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "creates new database on remote endpoint",
		Action: runDatabaseCreateCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "name, n",
				Usage:    "the name of the database to create",
				Required: true,
			},
		},
	}
}

func runDatabaseCreateCommand(ctx *cli.Context) error {
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
	psqlClient, err := client.NewClient(runContext, clientConfig)
	if err != nil {
		msg := fmt.Sprintf("Failed to initialize database client: %s", err)
		return cli.NewExitError(msg, 1)
	}

	if _, err := psqlClient.CreateDatabase(runContext, name); err != nil {
		msg := fmt.Sprintf("Failed to create database '%s': %s", name, err)
		return cli.NewExitError(msg, 1)
	}
	log.Info("Database '%s' was created", name)

	log.Info("Closing database client...")
	if err := psqlClient.Close(); err != nil {
		msg := fmt.Sprintf("Failed to close database client connection: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Finish '%s' command", cmdName)
	return nil
}
