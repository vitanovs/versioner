package cmd

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/client"
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
			cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug logging",
			},
		},
	}
}

func runDatabaseDropCommand(ctx *cli.Context) error {
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

	if _, err := psqlClient.DropDatabase(runContext, name); err != nil {
		msg := fmt.Sprintf("failed to drop database '%s': %s", name, err)
		return cli.NewExitError(msg, 1)
	}
	log.Infof("Database '%s' was dropped", name)

	log.Debugf("Closing database client...")
	if err := psqlClient.Close(); err != nil {
		msg := fmt.Sprintf("failed to close database client connection: %s", err)
		return cli.NewExitError(msg, 1)
	}

	return nil
}
