package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/client"
)

// NewDatabaseCreateCommand returns a command
// used for creating new databases on remote endpoint.
func NewDatabaseCreateCommand() cli.Command {
	cmd := cli.Command{
		Name:   "create",
		Usage:  "creates new database on remote endpoint",
		Action: runDatabaseCreateCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "name, n",
				Usage:    "the name of the database to create",
				Required: true,
			},
			cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug logging",
			},
		},
	}

	cmd.BashComplete = cli.DefaultCompleteWithFlags(&cmd)
	return cmd
}

func runDatabaseCreateCommand(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	cmdName := ctx.Command.Name

	name := ctx.String("name")
	if name == "" {
		msg := fmt.Sprintf("No new database name specified. See '%s --help'", cmdName)
		return cli.NewExitError(msg, 1)
	}

	runContext := context.Background()

	log.Debug("Loading PostgreSQL client configuration...")
	clientConfig, err := loadClientConfig(ctx)
	if err != nil {
		msg := fmt.Sprintf("failed to client configuration: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Debug("Initializing database client...")
	psqlClient, err := client.NewClient(runContext, clientConfig)
	if err != nil {
		msg := fmt.Sprintf("failed to initialize database client: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Debugf("Creating database '%s'...", name)
	if _, err := psqlClient.CreateDatabase(runContext, name); err != nil {
		msg := fmt.Sprintf("failed to create database '%s': %s", name, err)
		return cli.NewExitError(msg, 1)
	}
	log.Infof("Database '%s' was created", name)

	log.Debug("Closing database client...")
	if err := psqlClient.Close(); err != nil {
		msg := fmt.Sprintf("failed to close database client connection: %s", err)
		return cli.NewExitError(msg, 1)
	}

	return nil
}
