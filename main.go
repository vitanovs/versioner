package main

import (
	"os"

	"github.com/hall-arranger/versioner/cmd"
	"github.com/hall-arranger/versioner/version"
	"github.com/urfave/cli"
)

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
	}

	app.Run(os.Args)
}
