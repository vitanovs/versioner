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
	app.Usage = "The Hall Arranger database schema versioning tool"
	app.Version = version.BuildInfo

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "endpoint, e",
			Usage:    "remote endpoint address",
			Required: true,
		},
		cli.Int64Flag{
			Name:     "port, p",
			Usage:    "remote endpoint port",
			Required: true,
		},
		cli.StringFlag{
			Name:     "database, d",
			Usage:    "remote endpoint database",
			Required: true,
		},
		cli.StringFlag{
			Name:     "username",
			Usage:    "remote endpoint username",
			Required: true,
		},
		cli.StringFlag{
			Name:     "password",
			Usage:    "remote endpoint password",
			Required: true,
		},
		cli.StringFlag{
			Name:     "sslmode, s",
			Usage:    "remote endpoint ssl mode",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		cmd.NewMigrationCommand(),
		cmd.NewDatabaseCommand(),
	}

	app.Run(os.Args)
}
