package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	psql "github.com/hall-arranger/versioner/api/postgresql"
	"github.com/hall-arranger/versioner/config"
	"github.com/hall-arranger/versioner/log"
	"github.com/urfave/cli"
)

// NewMigrationRunCommand returns a command
// used for running migration definitions
// on a remote database endpoint.
func NewMigrationRunCommand() cli.Command {
	return cli.Command{
		Name:   "run",
		Usage:  "runs migration definitions",
		Action: runMigrationRunCommand,
	}
}

func runMigrationRunCommand(ctx *cli.Context) error {
	cmdName := ctx.Command.Name
	log.Info("Start '%s' command", cmdName)

	configPath := ctx.GlobalString("config")
	if configPath == "" {
		msg := fmt.Sprintf("No config path specified. See '%s --help'", cmdName)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Loading configuration '%s'", configPath)
	configuration, err := config.LoadConfig(configPath)
	if err != nil {
		msg := fmt.Sprintf("Failed to load configuration: %s", err)
		return cli.NewExitError(msg, 1)
	}

	runContext := context.Background()

	log.Info("Initializing database client...")
	client, err := psql.NewClient(runContext, configuration)
	if err != nil {
		msg := fmt.Sprintf("Failed to initialize database client: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Applying migrations...")
	if err := applyMigrations(runContext, configuration, client); err != nil {
		msg := fmt.Sprintf("Failed to apply migrations: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Closing database client...")
	if err := client.Close(); err != nil {
		msg := fmt.Sprintf("Failed to close database client connection: %s", err)
		return cli.NewExitError(msg, 1)
	}

	log.Info("Finish '%s' command", cmdName)
	return nil
}

func applyMigrations(ctx context.Context, configuration *config.Config, client *psql.Client) error {
	log.Info("Retrieving database version...")
	version, err := client.DatabaseVersion()
	if err != nil {
		msg := fmt.Sprintf("Failed to retrieve database version: %s", err)
		return cli.NewExitError(msg, 1)
	}
	log.Info("Current Database version '%s'", version)

	log.Info("Calculating current version index...")
	lastMigrationIndex := getLastAppliedMigrationIndex(configuration, version)
	log.Info("Current version index %d", lastMigrationIndex)

	for index := 0; index <= lastMigrationIndex; index++ {
		migrationName := configuration.Schema.Sequence[index]
		log.Info("Skipping migration '%s' with index %d, already applied", migrationName, index)
	}

	for index := lastMigrationIndex + 1; index < len(configuration.Schema.Sequence); index++ {
		migrationName := configuration.Schema.Sequence[index]
		log.Info("Applying migration '%s' with index %d...", migrationName, index)

		migrationPath := fmt.Sprintf("%s%c%s",
			configuration.Schema.Directory,
			os.PathSeparator,
			migrationName,
		)

		bytes, err := ioutil.ReadFile(migrationPath)
		if err != nil {
			msg := fmt.Sprintf("Failed to read migration value: %s", err)
			return cli.NewExitError(msg, 1)
		}

		query := string(bytes)
		if _, err := client.Execute(ctx, query); err != nil {
			msg := fmt.Sprintf("Failed to execute migration '%s': %s", migrationName, err)
			return cli.NewExitError(msg, 1)
		}
	}

	return nil
}

func getLastAppliedMigrationIndex(configuration *config.Config, version string) int {
	resultIndex := -1
	for index, migration := range configuration.Schema.Sequence {
		migrationName := strings.Split(migration, ".")[0]
		if migrationName == version {
			resultIndex = index
			break
		}
	}

	return resultIndex
}
