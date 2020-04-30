package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/client"
	"github.com/vitanovs/versioner/config"
)

// NewMigrationApplyCommand returns a command
// used for applying migration definitions
// on a remote database endpoint.
func NewMigrationApplyCommand() cli.Command {
	return cli.Command{
		Name:   "apply",
		Usage:  "applies migrations on remote PostgreSQL endpoint",
		Action: runMigrationApplyCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "config, c",
				Usage:    "path to configuration file",
				Required: true,
			},
			cli.BoolFlag{
				Name:  "debug",
				Usage: "enables debug logging",
			},
		},
	}
}

func runMigrationApplyCommand(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	cmdName := ctx.Command.Name

	configPath := ctx.String("config")
	if configPath == "" {
		msg := fmt.Sprintf("No config path specified. See '%s --help'", cmdName)
		log.Error(msg)
		return cli.NewExitError(msg, 1)
	}

	log.Infof("Loading configuration '%s'...", configPath)
	configuration, err := config.LoadConfig(configPath)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	log.Debug("Loading PostgreSQL client configuration...")
	clientConfig, err := loadClientConfig(ctx)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	cmdCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Debug("Initializing new PostgreSQL client...")
	psqlClient, err := client.NewClient(cmdCtx, clientConfig)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	if err := applyMigrations(cmdCtx, psqlClient, configuration, clientConfig.Username); err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	log.Debug("Closing PostgreSQL client...")
	if err := psqlClient.Close(); err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	return nil
}

func applyMigrations(ctx context.Context, psqlClient *client.Client, configuration *config.Config, user string) error {
	log.Debug("Retrieving last applied migration...")
	lastMigration, err := psqlClient.LastMigration(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve last applied migration: %s", err)
	}

	index := -1
	if lastMigration != nil {

		log.WithFields(map[string]interface{}{
			"created": lastMigration.Created,
		}).Infof("Last applied migration '%s'", lastMigration.Name)

		for i, migration := range configuration.Schema.Sequence {
			if migration == lastMigration.Name {
				index = i
				break
			}
		}
	}

	for i, migration := range configuration.Schema.Sequence {
		path := fmt.Sprintf("%s%c%s", configuration.Schema.Directory, os.PathSeparator, migration)

		logMetadata := map[string]interface{}{
			"path": path,
		}

		if i <= index {
			log.WithFields(logMetadata).Infof("Skipping '%s'. Already applied", migration)
			continue
		}

		log.WithFields(logMetadata).Infof("Applying migration '%s'...", migration)

		log.WithFields(logMetadata).Debugf("Loading migration '%s' content", migration)
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration '%s' content: %s", migration, err)
		}

		log.WithFields(logMetadata).Debugf("Executing migration '%s' content", migration)
		_, err = psqlClient.Execute(ctx, string(bytes))
		if err != nil {
			return fmt.Errorf("failed to apply migration '%s': %s", migration, err)
		}

		newMigration := client.Migration{
			Name:      migration,
			Path:      path,
			AppliedBy: user,
		}

		log.WithFields(logMetadata).Debugf("Registering migration '%s' as applied", migration)
		err = psqlClient.RegisterMigration(ctx, &newMigration)
		if err != nil {
			return fmt.Errorf("failed to register migration '%s': %s", migration, err)
		}

		log.WithFields(logMetadata).Infof("Migration '%s' was applied successfully", migration)
	}

	return nil
}
