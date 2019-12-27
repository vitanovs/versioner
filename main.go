package main

import (
	"os"

	"github.com/hall-arranger/versioner/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Versioner"
	app.Usage = "The Hall Arranger database schema versioning tool"
	app.Version = version.BuildInfo

	app.Commands = []cli.Command{}

	app.Run(os.Args)
}
