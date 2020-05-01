package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vitanovs/versioner/cmd"
	"github.com/vitanovs/versioner/version"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		QuoteEmptyFields:       true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		TimestampFormat:        time.RFC3339,
	})
}

func main() {
	app := cli.NewApp()
	app.Name = "Versioner"
	app.Usage = "PostgreSQL schema migrations versioning tool"
	app.Version = version.BuildInfo

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "endpoint, e",
			Usage:    "remote endpoint address",
			EnvVar:   "VERSIONER_PSQL_ENDPOINT",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "port, p",
			Usage:    "remote endpoint port",
			EnvVar:   "VERSIONER_PSQL_PORT",
			Value:    5432,
			Required: true,
		},
		cli.StringFlag{
			Name:     "database, d",
			Usage:    "remote endpoint database",
			EnvVar:   "VERSIONER_PSQL_DATABASE",
			Required: true,
		},
		cli.StringFlag{
			Name:     "username",
			Usage:    "remote endpoint username",
			EnvVar:   "VERSIONER_PSQL_USERNAME",
			Required: true,
		},
		cli.StringFlag{
			Name:     "password",
			Usage:    "remote endpoint password",
			EnvVar:   "VERSIONER_PSQL_PASSWORD",
			Required: true,
		},
		cli.StringFlag{
			Name:     "sslmode, s",
			Usage:    "remote endpoint ssl mode",
			EnvVar:   "VERSIONER_PSQL_SSLMODE",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		cmd.NewMigrationCommand(),
		cmd.NewDatabaseCommand(),
	}

	app.Run(os.Args)
}
