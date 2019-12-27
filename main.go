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
			Name:     "config, c",
			Usage:    "path to configuration file",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		cmd.NewMigrationCommand(),
	}

	app.Run(os.Args)
}
