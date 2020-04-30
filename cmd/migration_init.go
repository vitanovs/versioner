package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/api/postgresql"
)

const (
	versionerSchema = `
		BEGIN TRANSACTION;

		CREATE SCHEMA versioner;

		CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;
		CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;

		CREATE OR REPLACE FUNCTION versioner.set_created() RETURNS trigger AS $set_created$
		     BEGIN
		         NEW.created := now();
		         RETURN NEW;
		     END;
		 $set_created$ LANGUAGE plpgsql;

		CREATE OR REPLACE FUNCTION versioner.migration_create_pk() RETURNS trigger AS $make_pk$
		 DECLARE
		     pk char(64);
		 BEGIN
		     pk = LOWER(ENCODE(public.DIGEST(CONCAT(NEW.name), 'sha256'), 'hex'));
		     IF NEW.id IS DISTINCT FROM pk THEN
		         NEW.id := pk;
		     END IF;
		     RETURN NEW;
		 END;
		 $make_pk$ LANGUAGE plpgsql;

		CREATE TABLE versioner.migration (
		    id VARCHAR(64) NOT NULL,
			name VARCHAR(64) NOT NULL,
			path VARCHAR(255) NOT NULL,
			applied_by VARCHAR(127) NOT NULL,
		    created TIMESTAMP WITH TIME ZONE NOT NULL,
		    CONSTRAINT migration_pk PRIMARY KEY (id)
		);
		
		CREATE TRIGGER migration_create_pk_tr
		     BEFORE INSERT ON versioner.migration
		     FOR EACH ROW
		     EXECUTE PROCEDURE versioner.migration_create_pk();
		
		CREATE TRIGGER migration_set_created_tr
		     BEFORE INSERT ON versioner.migration
		     FOR EACH ROW
		     EXECUTE PROCEDURE versioner.set_created();
		
		COMMIT;
	`
)

// NewMigrationInitCommand returns a command
// used for initializing the versioner on the
// remote PostgreSQL endpoint.
func NewMigrationInitCommand() cli.Command {
	return cli.Command{
		Name:   "init",
		Usage:  "initializes versioner tool schema",
		Action: runMigrationInitCommand,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "debug",
				Usage: "enables debug logging",
			},
		},
	}
}

func runMigrationInitCommand(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("Loading PostgreSQL client configuration...")
	config, err := loadClientConfig(ctx)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	cmdCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Debug("Initializing new PostgreSQL client...")
	client, err := postgresql.NewClient(cmdCtx, config)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	log.Debug("Applying utility schema...")
	_, err = client.Execute(cmdCtx, versionerSchema)
	if err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}
	log.Info("`versioner` utility schema applied successfully")

	log.Debug("Closing PostgreSQL client...")
	if err := client.Close(); err != nil {
		log.Error(err)
		return cli.NewExitError(err, 1)
	}

	return nil
}
