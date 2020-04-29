package cmd

import (
	"fmt"

	"github.com/vitanovs/versioner/api/postgresql"
	"github.com/urfave/cli"
)

func loadClientConfig(ctx *cli.Context) (*postgresql.ClientConfig, error) {
	cmdName := ctx.Command.FullName()

	endpoint := ctx.GlobalString("endpoint")
	if endpoint == "" {
		return nil, fmt.Errorf("No endpoint specified. See '%s --help'", cmdName)
	}

	port := ctx.GlobalInt("port")
	if port == 0 {
		return nil, fmt.Errorf("No port specified. See '%s --help'", cmdName)
	}

	database := ctx.GlobalString("database")
	if database == "" {
		return nil, fmt.Errorf("No database specified. See '%s --help'", cmdName)
	}

	username := ctx.GlobalString("username")
	if username == "" {
		return nil, fmt.Errorf("No username specified. See '%s --help'", cmdName)
	}

	password := ctx.GlobalString("password")
	if password == "" {
		return nil, fmt.Errorf("No password specified. See '%s --help'", cmdName)
	}

	sslMode := ctx.GlobalString("sslmode")
	if sslMode == "" {
		return nil, fmt.Errorf("No SSL mode specified. See '%s --help'", cmdName)
	}

	config := postgresql.ClientConfig{
		Endpoint: endpoint,
		Port:     port,
		Database: database,
		Username: username,
		Password: password,
		SslMode:  sslMode,
	}

	return &config, nil
}
